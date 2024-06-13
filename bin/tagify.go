package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func tagify() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Tags: ") {
			l := strings.TrimPrefix(line, "<p>Tags: ")
			l = strings.TrimSuffix(l, "</p>")
			splits := strings.Split(l, ",")
			for i, split := range splits {
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf(`<p>Tags: <a href='tag_%s.html'>%s</a></p>`, split, split)
				fmt.Println()
			}
			continue
		}

		fmt.Println(line)
	}
	return scanner.Err()
}
