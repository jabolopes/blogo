package main

import (
	"os"
	"text/template"
)

func renderPostsForIndex(filenames []string) ([]byte, error) {
	posts, err := getPostsSortedDescending(filenames)
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

func genIndex(postFilenames []string) error {
	tmpl, err := template.ParseFiles(templateName, contentTemplateName)
	if err != nil {
		return err
	}

	content, err := renderPostsForIndex(postFilenames)
	if err != nil {
		return err
	}

	blogConfig["Content"] = string(content)
	return tmpl.Execute(os.Stdout, blogConfig)
}
