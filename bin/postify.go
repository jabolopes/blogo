package main

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"text/template"
)

func postify(ctx context.Context, postFilename string) error {
	post, err := getPost(postFilename)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(markdownTemplateName)
	if err != nil {
		return err
	}

	blogConfig["PostURL"] = post.PostURL
	blogConfig["PostTitle"] = post.PostTitle
	blogConfig["PostDate"] = post.PostDate
	blogConfig["PostContent"] = post.PostContent
	blogConfig["Tags"] = post.Tags

	output := &bytes.Buffer{}

	if err := tmpl.Execute(output, blogConfig); err != nil {
		return err
	}

	post.MarkdownContent = output.String()

	htmlContent := &bytes.Buffer{}
	if err := markdown(ctx, output, htmlContent); err != nil {
		return err
	}

	post.HTMLContent = htmlContent.String()

	postJson, err := json.Marshal(post)
	if err != nil {
		return err
	}

	outputFile, err := os.OpenFile(postifiedFilename(postFilename), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	if _, err := outputFile.Write(postJson); err != nil {
		return err
	}

	if err := outputFile.Close(); err != nil {
		return err
	}

	return nil
}
