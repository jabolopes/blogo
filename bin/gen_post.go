package main

import (
	"os"
	"text/template"
)

func genPost(postFilename string) error {
	tmpl, err := template.ParseFiles(pageTemplateName, postTemplateName)
	if err != nil {
		return err
	}

	post, err := loadPost(postifiedFilename(postFilename))
	if err != nil {
		return err
	}

	blogConfig["Post"] = post
	return tmpl.Execute(os.Stdout, blogConfig)
}
