package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(post.PostTitle) == 0 && len(line) > 0 {
			post.PostTitle = line
			continue
		}

		if strings.HasPrefix(line, `<div class="subtitle">`) {
			line = strings.TrimPrefix(line, `<div class="subtitle">`)
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
	}

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
		post.ParsedDate, err = time.Parse("January 02, 2006", post.PostDate)
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
	filename = strings.ReplaceAll(filename, ".md", ".pre")
	filename = path.Join("out", filename)
	return os.ReadFile(filename)
}
