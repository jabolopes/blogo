package main

import (
	"os"
	"text/template"
)

func loadIndexPosts() ([]byte, error) {
	posts, err := loadAllPostsSortedDescending()
	if err != nil {
		return nil, err
	}

	posts = posts[:min(len(posts), indexPostsNum)]

	var content []byte
	for _, post := range posts {
		data, err := renderPost(post.MarkdownFilename)
		if err != nil {
			return nil, err
		}

		content = append(content, data...)
	}

	return content, nil
}

func genIndex() error {
	tmpl, err := template.ParseFiles(pageTemplateName, contentTemplateName)
	if err != nil {
		return err
	}

	content, err := loadIndexPosts()
	if err != nil {
		return err
	}

	blogConfig["Content"] = string(content)
	return tmpl.Execute(os.Stdout, blogConfig)
}
