package main

import "fmt"

func genRobotsTXT() error {
	fmt.Printf(`User-agent: *
Sitemap: %s/sitemap.xml
`, blogURL)
	return nil
}
