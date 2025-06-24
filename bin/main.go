package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	pageTemplateName     = "templates/page.template"
	allPostsTemplateName = "templates/all-posts.template"
	allTagsTemplateName  = "templates/all-tags.template"
	feedTemplateName     = "templates/feed.template"
	indexTemplateName    = "templates/index.template"
	postTemplateName     = "templates/post.template"
	tagTemplateName      = "templates/tag.template"

	indexPostsNum = 10

	indexFilename = "index.html"
	feedFilename  = "feed.rss"

	genTagTitleFormat   = `%s &mdash; Posts tagged "%s"`
	allPostsTitleFormat = `%s &mdash; All posts`
	allTagsTitleFormat  = `%s &mdash; All tags`

	feedDateFormat = time.RFC1123Z
	// Date format that the author uses in the Markdown posts.
	postParseDateFormat = "2006/01/02"
	// Date format that the final rendered blog uses.
	postDisplayDateFormat = "January 02, 2006"
	// Date format that the final rendered blog uses.
	monthDisplayDateFormat = "January 2006"

	outputDirectory      = "out"
	outputDistDirectory  = "out/dist"
	outputPostsDirectory = "out/posts"

	authorEmail     = "jadesmith@email.com"
	authorName      = "Jade Smith"
	authorURL       = "https://github.com/jadesmith"
	blogDescription = "Jade Smith's cool blogo"
	blogImage       = "images/logo.png"
	blogLanguage    = "en"
	blogName        = "Cool Blogo"
	blogURL         = "http://jadesmith.blogo"
	license         = "&copy;"
)

var blogConfig = map[string]interface{}{
	"Title":                  blogName,
	"BlogDescription":        blogDescription,
	"BlogImage":              blogImage,
	"BlogLanguage":           blogLanguage,
	"BlogName":               blogName,
	"License":                license,
	"AuthorURL":              authorURL,
	"AuthorName":             authorName,
	"AuthorEmail":            authorEmail,
	"PostDisplayDateFormat":  postDisplayDateFormat,
	"MonthDisplayDateFormat": monthDisplayDateFormat,
}

func main() {
	ctx := context.Background()

	genAllPostsCmd := flag.NewFlagSet("gen-all-posts", flag.ExitOnError)
	genAllTagsCmd := flag.NewFlagSet("gen-all-tags", flag.ExitOnError)
	genFeedCmd := flag.NewFlagSet("gen-feed", flag.ExitOnError)
	genIndexCmd := flag.NewFlagSet("gen-index", flag.ExitOnError)
	genPostCmd := flag.NewFlagSet("gen-post", flag.ExitOnError)
	genSitemapCmd := flag.NewFlagSet("gen-sitemap", flag.ExitOnError)
	genTagCmd := flag.NewFlagSet("gen-tag", flag.ExitOnError)
	postifyCmd := flag.NewFlagSet("postify", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected subcommand, e.g., 'gen-all-posts', 'gen-all-tags', etc")
		os.Exit(1)
	}

	command := os.Args[1]
	var err error
	switch command {
	case "gen-all-posts":
		genAllPostsCmd.Parse(os.Args[2:])
		err = genAllPosts()
	case "gen-all-tags":
		genAllTagsCmd.Parse(os.Args[2:])
		err = genAllTags()
	case "gen-feed":
		genFeedCmd.Parse(os.Args[2:])
		err = genFeed()
	case "gen-index":
		genIndexCmd.Parse(os.Args[2:])
		err = genIndex()

	case "gen-post":
		genPostCmd.Parse(os.Args[2:])

		args := genPostCmd.Args()
		if len(args) <= 0 {
			fmt.Fprintf(os.Stderr, "command %q expects 1 argument\n", command)
			os.Exit(1)
		}

		err = genPost(args[0])

	case "gen-sitemap":
		genSitemapCmd.Parse(os.Args[2:])
		err = genSitemap()

	case "gen-tag":
		genTagCmd.Parse(os.Args[2:])
		err = genTag()

	case "postify":
		postifyCmd.Parse(os.Args[2:])

		if len(postifyCmd.Args()) != 1 {
			fmt.Fprintf(os.Stderr, "command %q expects 1 argument\n", command)
			os.Exit(1)
		}

		err = postify(ctx, postifyCmd.Args()[0])

	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", command)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute command %q: %v\n", command, err)
		os.Exit(1)
	}
}
