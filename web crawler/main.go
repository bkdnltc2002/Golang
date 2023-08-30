package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main(){

	baseURL := "tiki.vn"
	startingURL := "https://" +baseURL
	allowedURL := []string{baseURL}



	c := colly.NewCollector(
		colly.AllowedDomains(allowedURL...),
		colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Limit(&colly.LimitRule{DomainGlob: "*",Parallelism: 5})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Getting error",err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited",r.Request.URL)
	})

	c.OnHTML("a[href]",func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnHTML("title",func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		
	})


	fmt.Println("Starting crawl at: ",startingURL)

	if err:= c.Visit(startingURL); err!=nil{
		fmt.Println("Error on start of crawl:",err)
	}
	c.Wait()

}