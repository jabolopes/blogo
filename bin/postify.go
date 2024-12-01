package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func postify(outputDirectory, titleHref string) error {
	postTitle := ""

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if len(postTitle) == 0 && len(line) > 0 {
			postTitle = line
			fmt.Printf(`### <a href="%s">%s</a>`, titleHref, line)
			fmt.Println()
			continue
		}

		if strings.HasPrefix(line, "Date: ") {
			line = strings.TrimPrefix(line, "Date: ")
			fmt.Printf(`<div class="postDate">%s</div>`, line)
			fmt.Println()
			continue
		}

		if strings.HasPrefix(line, "Tags: ") {
			line = strings.TrimPrefix(line, "Tags: ")
			tags := strings.Split(line, ",")

			fmt.Printf(`Tags: `)
			for i, tag := range tags {
				tag = strings.TrimSpace(tag)

				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf(`<a href='tag_%s.html'>%s</a>`, tag, tag)
			}
			fmt.Println()

			continue
		}

		fmt.Println(line)
	}

	return scanner.Err()
}
