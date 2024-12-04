package main

import (
	"os"
	"text/template"
)

func loadIndexPosts() ([]Post, error) {
	posts, err := loadAllPostsSortedDescending()
	if err != nil {
		return nil, err
	}

	posts = posts[:min(len(posts), indexPostsNum)]
	return posts, nil
}

func genIndex() error {
	tmpl, err := template.ParseFiles(pageTemplateName, indexTemplateName)
	if err != nil {
		return err
	}

	posts, err := loadIndexPosts()
	if err != nil {
		return err
	}

	blogConfig["Posts"] = posts
	return tmpl.Execute(os.Stdout, blogConfig)
}
