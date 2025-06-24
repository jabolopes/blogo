package main

import (
	"os"

	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
)

func createSitemap(posts []Post) []byte {
	sm := stm.NewSitemap(1)
	sm.SetVerbose(false)
	sm.SetDefaultHost(blogURL)

	sm.Create()

	sm.Add(stm.URL{
		{"loc", "index.html"},
		{"changefreq", "always"},
		{"mobile", true},
	})
	sm.Add(stm.URL{
		{"loc", "all-posts.html"},
		{"changefreq", "always"},
		{"mobile", true},
	})
	sm.Add(stm.URL{
		{"loc", "all-tags.html"},
		{"changefreq", "always"},
		{"mobile", true},
	})

	tagHrefs := map[string]struct{}{}
	for _, post := range posts {
		sm.Add(stm.URL{{"loc", post.PostURL}})

		for _, tag := range post.Tags {
			tagHrefs[tag.Href()] = struct{}{}
		}
	}

	for tag, _ := range tagHrefs {
		sm.Add(stm.URL{{"loc", tag}})
	}

	return sm.XMLContent()
}

func genSitemap() error {
	posts, err := loadAllPostsSortedDescending()
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(createSitemap(posts))
	return err
}
