package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

type Item struct {
	PostTitle   string
	PostContent string
	PostURL     string
	PostDate    string
}

func createFeedItems() ([]Item, error) {
	posts, err := loadAllPostsSortedDescending()
	if err != nil {
		return nil, err
	}

	posts = posts[:min(len(posts), indexPostsNum)]

	var items []Item
	for _, post := range posts {
		postUrl := fmt.Sprintf("%s/%s.html", blogURL, strings.TrimSuffix(post.MarkdownFilename, ".md"))

		content := post.HTMLContent
		{
			lines := strings.Split(content, "\n")
			for i := 0; len(lines) > 0 && i < 2; {
				if len(lines[0]) > 0 {
					i++
				}
				lines = lines[1:]
			}
			content = strings.Join(lines, "\n")
		}

		postDate := post.Date.Format(feedDateFormat)

		item := Item{
			post.PostTitle,
			content,
			postUrl,
			postDate,
		}

		items = append(items, item)
	}

	return items, nil
}

func genFeed() error {
	tmpl, err := template.ParseFiles(feedTemplateName)
	if err != nil {
		return err
	}

	items, err := createFeedItems()
	if err != nil {
		return err
	}

	blogConfig["IndexURL"] = fmt.Sprintf("%s/%s", blogURL, indexFilename)
	blogConfig["PubDate"] = time.Now().Format(feedDateFormat)
	blogConfig["FeedURL"] = fmt.Sprintf("%s/%s", blogURL, feedFilename)
	blogConfig["Items"] = items
	return tmpl.Execute(os.Stdout, blogConfig)
}
