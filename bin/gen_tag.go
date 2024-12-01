package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"text/template"
)

func indexPostsByTag(filenames []string) (map[string][]Post, error) {
	index := map[string][]Post{}

	for _, filename := range filenames {
		post, err := getPost(filename)
		if err != nil {
			return nil, err
		}

		for _, tag := range post.Tags {
			index[tag] = append(index[tag], post)
		}
	}

	for _, posts := range index {
		sort.Slice(posts, func(i, j int) bool {
			return comparePostsDescending(posts[i], posts[j])
		})
	}

	return index, nil
}

func genTag(postFilenames []string) error {
	tmpl, err := template.ParseFiles(templateName, contentTemplateName)
	if err != nil {
		return err
	}

	index, err := indexPostsByTag(postFilenames)
	if err != nil {
		return err
	}

	for tagName, posts := range index {
		var content []byte
		for _, post := range posts {
			data, err := renderPost(post.MarkdownFilename)
			if err != nil {
				return err
			}

			content = append(content, data...)
		}

		outputFilename := path.Join(outputDistDirectory, fmt.Sprintf("tag_%s.html", tagName))
		outputFile, err := os.Create(outputFilename)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		config := map[string]interface{}{}
		for key, value := range blogConfig {
			config[key] = value
		}
		config["Title"] = fmt.Sprintf(genTagTitleFormat, blogName, tagName)
		config["Content"] = string(content)

		if err := tmpl.Execute(outputFile, config); err != nil {
			return err
		}
	}

	return nil
}
