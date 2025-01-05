package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"text/template"
)

type TagCount struct {
	// e.g., 'poetry'.
	Tag Tag
	// e.g., '1 post' or '2 posts'.
	Count string
}

func loadAllTags() ([]TagCount, error) {
	counts := map[string]int{}

	posts, err := loadAllPosts()
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		for _, tag := range post.Tags {
			counts[tag.Name] = counts[tag.Name] + 1
		}
	}

	var tags []TagCount
	for name, count := range counts {
		var countStr string
		if count == 1 {
			countStr = fmt.Sprintf("%d post", count)
		} else {
			countStr = fmt.Sprintf("%d posts", count)
		}

		tag := TagCount{
			Tag{name},
			countStr,
		}
		tags = append(tags, tag)
	}

	slices.SortFunc(tags, func(t1, t2 TagCount) int {
		return cmp.Compare(t1.Tag.Name, t2.Tag.Name)
	})

	return tags, nil
}

func genAllTags() error {
	tmpl, err := template.ParseFiles(pageTemplateName, allTagsTemplateName)
	if err != nil {
		return err
	}

	tags, err := loadAllTags()
	if err != nil {
		return err
	}

	blogConfig["Title"] = fmt.Sprintf(allTagsTitleFormat, blogName)
	blogConfig["Tags"] = tags
	return tmpl.Execute(os.Stdout, blogConfig)
}
