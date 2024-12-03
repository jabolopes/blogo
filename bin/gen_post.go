package main

import (
	"os"
	"text/template"
)

func genPost(postFilename string) error {
	tmpl, err := template.ParseFiles(pageTemplateName, contentTemplateName)
	if err != nil {
		return err
	}

	content, err := renderPost(postFilename)
	if err != nil {
		return err
	}

	blogConfig["Content"] = string(content)
	return tmpl.Execute(os.Stdout, blogConfig)
}
