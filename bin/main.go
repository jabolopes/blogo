package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	templateName         = "templates/index.template"
	allPostsTemplateName = "templates/all_posts.template"
	allTagsTemplateName  = "templates/all_tags.template"
	contentTemplateName  = "templates/content.template"
	feedTemplateName     = "templates/feed.template"

	indexPostsNum = 10

	indexFilename = "index.html"
	feedFilename  = "feed.rss"

	genTagTitleFormat   = `%s &mdash; Posts tagged "%s"`
	allPostsTitleFormat = `%s &mdash; All posts`
	allTagsTitleFormat  = `%s &mdash; All tags`

	feedDateFormat = time.RFC1123Z
	postDateFormat = "January 02, 2006"

	authorEmail     = "jadesmith@email.com"
	authorName      = "Jade Smith"
	authorURL       = "https://github.com/jadesmith"
	blogDescription = "Jade Smith's cool blogo"
	blogLanguage    = "en"
	blogName        = "Cool Blogo"
	blogURL         = "http://jadesmith.blogo"
	license         = "&copy;"
)

var blogConfig = map[string]interface{}{
	"Title":           blogName,
	"BlogDescription": blogDescription,
	"BlogLanguage":    blogLanguage,
	"BlogName":        blogName,
	"License":         license,
	"AuthorURL":       authorURL,
	"AuthorName":      authorName,
	"AuthorEmail":     authorEmail,
}

func main() {
	genAllPostsCmd := flag.NewFlagSet("gen-all-posts", flag.ExitOnError)
	genAllTagsCmd := flag.NewFlagSet("gen-all-tags", flag.ExitOnError)
	genFeedCmd := flag.NewFlagSet("gen-feed", flag.ExitOnError)
	genIndexCmd := flag.NewFlagSet("gen-index", flag.ExitOnError)
	genPostCmd := flag.NewFlagSet("gen-post", flag.ExitOnError)

	genTagCmd := flag.NewFlagSet("gen-tag", flag.ExitOnError)
	out := genTagCmd.String("out", "", "Output directory, e.g., 'out/'")

	postifyCmd := flag.NewFlagSet("postify", flag.ExitOnError)
	titleHref := postifyCmd.String("titleHref", "", "Hyperlink for the title, e.g., 'mypost.html'")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommand, e.g., 'gen-all-posts', 'gen-all-tags', etc")
		os.Exit(1)
	}

	command := os.Args[1]
	var err error
	switch command {
	case "gen-all-posts":
		genAllPostsCmd.Parse(os.Args[2:])
		err = genAllPosts(genAllPostsCmd.Args())
	case "gen-all-tags":
		genAllTagsCmd.Parse(os.Args[2:])
		err = genAllTags(genAllTagsCmd.Args())
	case "gen-feed":
		genFeedCmd.Parse(os.Args[2:])
		err = genFeed(genFeedCmd.Args())
	case "gen-index":
		genIndexCmd.Parse(os.Args[2:])
		err = genIndex(genIndexCmd.Args())
	case "gen-post":
		genPostCmd.Parse(os.Args[2:])

		args := genPostCmd.Args()
		if len(args) <= 0 {
			fmt.Fprintf(os.Stderr, "command %q expects 1 argument\n", command)
			os.Exit(1)
		}

		err = genPost(args[0])
	case "gen-tag":
		genTagCmd.Parse(os.Args[2:])

		if len(*out) == 0 {
			fmt.Fprintf(os.Stderr, "flag --out must be given (non-empty)\n")
			os.Exit(1)
		}

		err = genTag(*out, genTagCmd.Args())
	case "postify":
		postifyCmd.Parse(os.Args[2:])

		if len(*titleHref) == 0 {
			fmt.Fprintf(os.Stderr, "flag --titleHref must be given (non-empty)\n")
			os.Exit(1)
		}

		err = postify(*titleHref)
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", command)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute command %q: %v\n", command, err)
		os.Exit(1)
	}
}
