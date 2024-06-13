package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"
)

type Month struct {
	// 'November 2022'.
	Date  string
	Posts []Post
	// Parsed 'Date'.
	ParsedDate time.Time
}

func groupByMonth(posts []Post) ([]Month, error) {
	group := map[string]Month{}

	for _, post := range posts {
		splits := strings.Split(post.PostDate, " ")
		if len(splits) != 3 {
			continue
		}

		key := fmt.Sprintf("%s %s", splits[0], splits[2])
		if _, ok := group[key]; !ok {
			parsedDate, err := time.Parse("January 2006", key)
			if err != nil {
				return nil, err
			}

			group[key] = Month{key, nil, parsedDate}
		}

		month := group[key]
		month.Posts = append(month.Posts, post)
		group[key] = month
	}

	var months []Month
	for _, month := range group {
		months = append(months, month)
	}

	sort.Slice(months, func(i, j int) bool {
		t1 := months[i].ParsedDate
		t2 := months[j].ParsedDate
		// Sort in descending order (newest to oldest).
		return t2.Before(t1)
	})

	for _, month := range months {
		sort.Slice(month.Posts, func(i, j int) bool {
			return comparePostsDescending(month.Posts[i], month.Posts[j])
		})
	}

	return months, nil
}

func genAllPosts(postFilenames []string) error {
	tmpl, err := template.ParseFiles(templateName, allPostsTemplateName)
	if err != nil {
		return err
	}

	posts, err := getPosts(postFilenames)
	if err != nil {
		return err
	}

	months, err := groupByMonth(posts)
	if err != nil {
		return err
	}

	blogConfig["Title"] = fmt.Sprintf(allPostsTitleFormat, blogName)
	blogConfig["Months"] = months
	return tmpl.Execute(os.Stdout, blogConfig)
}
