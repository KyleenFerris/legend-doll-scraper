package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	searchUrl := ""
	c := colly.NewCollector(
		colly.MaxDepth(2), //First call is the search php query, then the actual search result
		colly.Async(true),
	)
	searchCollector := colly.NewCollector() //We need to make a seperate collector for the search, so we can handle it differently.

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnHTML("form#searchForm", func(e *colly.HTMLElement) {

		actionURL := e.Attr("action")
		formData := make(map[string]string)
		formData["keywords"] = "wings"
		e.Request.Post(actionURL, formData)

	})

	searchCollector.OnHTML(".goodsItem", func(e *colly.HTMLElement) {

		fmt.Println(e.DOM.Find(".height_name").Text())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response received:", r.StatusCode)
		searchUrl = r.Request.URL.String()

	})

	c.Visit("https://legenddoll.net")
	c.Wait()

	searchCollector.Visit(searchUrl)

}
