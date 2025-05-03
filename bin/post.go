package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func removeDiacritics(text string) (string, error) {
	chain := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	newText, _, err := transform.String(chain, text)
	return newText, err
}

type Tag struct {
	// Tag name, e.g., 'poetry', 'type theory', etc.
	Name string
}

// Href returns the relative URL for a tag.
func (t Tag) Href() string {
	name, err := removeDiacritics(t.Name)
	if err != nil {
		name = t.Name
	}

	return fmt.Sprintf("tag_%s.html", url.QueryEscape(name))
}

type Post struct {
	// './mypage.html'
	PostURL string
	// 'Post title'
	PostTitle string
	// Date of the post.
	Date time.Time
	// Tags, e.g., 'poetry', 'prose'.
	Tags []Tag
	// e.g., 'mypost.md'
	MarkdownFilename string
	// Post content without the title, date, tags, etc.
	PostContent string
	// HTML content after it was rendered by the markdown program,
	// without the title, date, tags, etc.
	HTMLContent string
}

func postifiedFilename(filename string) string {
	filename = path.Base(filename)
	filename = strings.TrimSuffix(filename, path.Ext(filename))
	filename = strings.TrimSuffix(filename, ".")
	filename = filename + ".post"
	return path.Join(outputPostsDirectory, filename)
}

func comparePostsDescending(p1, p2 Post) int {
	// Sort in descending order (newest to oldest).
	if n := p2.Date.Compare(p1.Date); n != 0 {
		return n
	}
	return cmp.Compare(p1.PostTitle, p2.PostTitle)
}

func loadPost(filename string) (Post, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Post{}, err
	}

	var post Post
	if err := json.Unmarshal(data, &post); err != nil {
		return Post{}, err
	}

	return post, nil
}

func loadAllPosts() ([]Post, error) {
	filenames, err := filepath.Glob(fmt.Sprintf("%s/*.post", outputPostsDirectory))
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(filenames))
	for _, filename := range filenames {
		post, err := loadPost(filename)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func loadAllPostsSortedDescending() ([]Post, error) {
	posts, err := loadAllPosts()
	if err != nil {
		return nil, err
	}

	slices.SortFunc(posts, comparePostsDescending)

	return posts, nil
}

func storePost(filename string, post Post) error {
	postJson, err := json.Marshal(post)
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if _, err := outputFile.Write(postJson); err != nil {
		return err
	}

	return outputFile.Close()
}
