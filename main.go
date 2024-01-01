package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func handleErr(err error) {
	fmt.Println("Error occured : ", err)
}

func main() {
	fmt.Println("Welcome to golang webscrapper")
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	fmt.Println("Enter URL of site to visit : ")
	reader := bufio.NewReader(os.Stdin)
	site, _ := reader.ReadString('\n')
	site = strings.TrimSpace(site)

	c.OnHTML(".mw-parser-output", func(h *colly.HTMLElement) {
		fmt.Println("Hello")
		links := h.ChildAttrs("a", "href")
		for _, link := range links {
			fmt.Println(link)
		}
	})

	c.OnError(func(r *colly.Response, e error) {
		handleErr(e)
	})

	err := c.Visit(site)
	handleErr(err)
}
