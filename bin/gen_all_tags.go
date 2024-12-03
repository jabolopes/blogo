package main

import (
	"fmt"
	"os"
	"sort"
	"text/template"
)

type Tag struct {
	// e.g., 'tag_poetry.html'.
	Href string
	// e.g., 'poetry'.
	Name string
	// e.g., '1 post' or '2 posts'.
	Count string
}

func indexTags(filenames []string) ([]Tag, error) {
	counts := map[string]int{}

	for _, filename := range filenames {
		post, err := getPost(filename)
		if err != nil {
			return nil, err
		}

		for _, tag := range post.Tags {
			counts[tag] = counts[tag] + 1
		}
	}

	var tags []Tag
	for name, count := range counts {
		var countStr string
		if count == 1 {
			countStr = fmt.Sprintf("%d post", count)
		} else {
			countStr = fmt.Sprintf("%d posts", count)
		}

		tag := Tag{
			fmt.Sprintf("tag_%s.html", name),
			name,
			countStr,
		}
		tags = append(tags, tag)
	}

	sort.Slice(tags, func(i, j int) bool {
		t1 := tags[i]
		t2 := tags[j]
		return t1.Name <= t2.Name
	})

	return tags, nil
}

func genAllTags(postFilenames []string) error {
	tmpl, err := template.ParseFiles(pageTemplateName, allTagsTemplateName)
	if err != nil {
		return err
	}

	tags, err := indexTags(postFilenames)
	if err != nil {
		return err
	}

	blogConfig["Title"] = fmt.Sprintf(allTagsTitleFormat, blogName)
	blogConfig["Tags"] = tags
	return tmpl.Execute(os.Stdout, blogConfig)
}
