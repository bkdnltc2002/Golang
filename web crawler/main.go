package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
)

type Web struct {
	id                  int
	title, content, url string
}

var webinfo []Web

func main() {
	count := 0
	baseURL := "www.bachhoaxanh.com"
	startingURL := "https://" + baseURL
	allowedURL := []string{baseURL}
	web := Web{}

	c := colly.NewCollector(
		colly.AllowedDomains(allowedURL...),
		colly.Async(true),
	)

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Getting error", err)
	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Visited", r.Request.URL)
		web.url = r.Request.AbsoluteURL("")
		// web.content = string(r.Body)
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",count)
		count = count +1
	})

	c.OnHTML("head", func(e *colly.HTMLElement) {
		web.title = e.ChildText("title")
		webinfo = append(webinfo, web)
		fmt.Println(web.url,web.title)
		fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbb",count)
		time.Sleep(time.Millisecond*300)
		fmt.Println("ccccccccccccccccccccc",count)
		fmt.Println("")
		count = count +1
	})

	c.OnHTML("a[href]",func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	fmt.Println("Starting crawl at: ", startingURL)


	if err := c.Visit(startingURL); err != nil {
		fmt.Println("Error on start of crawl:", err)
	}
	c.Wait()

}
