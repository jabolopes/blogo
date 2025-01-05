package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"text/template"
	"time"
)

type Month struct {
	// Date of this month (only year and month).
	Date time.Time
	// Posts whose Post.Date lie within this month.
	Posts []Post
}

func groupByMonth(posts []Post) ([]Month, error) {
	group := map[time.Time]Month{}

	for _, post := range posts {
		date := time.Date(post.Date.Year(), post.Date.Month(), 1, 0, 0, 0, 0, post.Date.Location())
		if _, ok := group[date]; !ok {
			group[date] = Month{date, nil}
		}

		month := group[date]
		month.Posts = append(month.Posts, post)
		group[date] = month
	}

	months := slices.Collect(maps.Values(group))
	slices.SortFunc(months, func(m1, m2 Month) int {
		// Sort in descending order (newest to oldest).
		return m2.Date.Compare(m1.Date)
	})

	return months, nil
}

func genAllPosts() error {
	tmpl, err := template.ParseFiles(pageTemplateName, allPostsTemplateName)
	if err != nil {
		return err
	}

	posts, err := loadAllPostsSortedDescending()
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
