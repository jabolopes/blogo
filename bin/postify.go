package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

func parsePost(postFilename string) (Post, error) {
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

	if len(post.PostTitle) == 0 {
		return Post{}, fmt.Errorf("post %q is missing the title", postFilename)
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

	return post, nil
}

func postify(ctx context.Context, postFilename string) error {
	post, err := parsePost(postFilename)
	if err != nil {
		return err
	}

	htmlContent := &strings.Builder{}
	if err := markdown(ctx, strings.NewReader(post.PostContent), htmlContent); err != nil {
		return err
	}
	post.HTMLContent = htmlContent.String()

	return storePost(postifiedFilename(postFilename), post)
}
