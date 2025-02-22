package main

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

func loadPostsByTag() (map[string][]Post, error) {
	index := map[string][]Post{}

	posts, err := loadAllPostsSortedDescending()
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		for _, tag := range post.Tags {
			index[tag.Name] = append(index[tag.Name], post)
		}
	}

	return index, nil
}

func genTag() error {
	tmpl, err := template.ParseFiles(pageTemplateName, tagTemplateName)
	if err != nil {
		return err
	}

	index, err := loadPostsByTag()
	if err != nil {
		return err
	}

	for tagName, posts := range index {
		outputFilename := path.Join(outputDistDirectory, Tag{tagName}.Href())
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
		config["Posts"] = posts

		if err := tmpl.Execute(outputFile, config); err != nil {
			return err
		}
	}

	return nil
}
