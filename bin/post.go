package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Post struct {
	// './mypage.html'
	PostURL string
	// 'Post title'
	PostTitle string
	// 'November 4, 2022'.
	PostDate string
	// Parsed 'PostDate'.
	ParsedDate time.Time
	// Tags, e.g., 'poetry', 'prose'.
	Tags []string
	// e.g., 'mypost.md'
	MarkdownFilename string
	// Post content without the title, date, tags, etc.
	PostContent string
	// Markdown content after it was rendered by the post template.
	MarkdownContent string
	// HTML content after it was rendered by the markdown program.
	HTMLContent string
}

func postifiedFilename(filename string) string {
	filename = path.Base(filename)
	filename = strings.TrimSuffix(filename, path.Ext(filename))
	filename = strings.TrimSuffix(filename, ".")
	filename = filename + ".post"
	return path.Join(outputPostsDirectory, filename)
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

	sort.Slice(posts, func(i, j int) bool {
		return comparePostsDescending(posts[i], posts[j])
	})

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

func getPost(postFilename string) (Post, error) {
	file, err := os.Open(postFilename)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	post := Post{
		MarkdownFilename: postFilename,
	}

	var postContent strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(post.PostTitle) == 0 && len(line) > 0 {
			post.PostTitle = line
			continue
		}

		if strings.HasPrefix(line, `Date: `) {
			line = strings.TrimPrefix(line, `Date: `)
			line = strings.Split(line, ` &mdash;`)[0]
			post.PostDate = line
			continue
		}

		if strings.HasPrefix(line, "Tags: ") {
			line = strings.TrimPrefix(line, "Tags: ")
			splits := strings.Split(line, ",")
			for i := range splits {
				splits[i] = strings.TrimSpace(splits[i])
			}

			post.Tags = splits
			continue
		}

		postContent.WriteString(line)
		postContent.WriteString("\n")
	}

	post.PostContent = postContent.String()

	if err := scanner.Err(); err != nil {
		return Post{}, err
	}

	{
		base := path.Base(strings.TrimSuffix(postFilename, path.Ext(postFilename)))
		post.PostURL = fmt.Sprintf("./%s.html", base)
	}

	if len(post.PostDate) == 0 {
		return Post{}, fmt.Errorf("post %q is missing the date", postFilename)
	}

	{
		var err error
		post.ParsedDate, err = time.Parse(postDateFormat, post.PostDate)
		if err != nil {
			return Post{}, fmt.Errorf("failed to parse post date: %v", err)
		}
	}

	// TODO: Validate that PostTitle is not empty.

	return post, nil
}

func getPosts(filenames []string) ([]Post, error) {
	var posts []Post
	for _, filename := range filenames {
		post, err := getPost(filename)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPostsSortedDescending(filenames []string) ([]Post, error) {
	posts, err := getPosts(filenames)
	if err != nil {
		return nil, err
	}

	sort.Slice(posts, func(i, j int) bool {
		return comparePostsDescending(posts[i], posts[j])
	})

	return posts, nil
}

func comparePostsDescending(p1, p2 Post) bool {
	// Sort in descending order (newest to oldest).
	if p2.ParsedDate.Before(p1.ParsedDate) {
		return true
	}

	if p2.ParsedDate.After(p1.ParsedDate) {
		return false
	}

	return p1.PostTitle <= p2.PostTitle
}

// TODO: Avoid hardcoded "out" directory since it can be changed externally.
func renderPost(filename string) ([]byte, error) {
	file, err := os.Open(postifiedFilename(filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var post Post
	if err := json.Unmarshal(data, &post); err != nil {
		return nil, err
	}

	return []byte(post.HTMLContent), nil
}
