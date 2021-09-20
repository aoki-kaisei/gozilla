package query

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"log"
	"github.com/kaseiaoki/gozilla/array"
)

func GetLink(baseUrl string) []string{
		memo := []string{}
		c := colly.NewCollector()
		url, err := url.Parse(baseUrl)
		if err != nil {
			log.Fatal(err)
		  }
		host := url.Hostname()
		bu, _ := url.Parse(baseUrl) 
		
		c.AllowedDomains = []string{host}

		c.OnHTML("a[href]", func(element *colly.HTMLElement) {
			link := element.Attr("href")
			if(!array.Contains(memo, toAbsUrl(bu, link).String())) {
				memo = append(memo, toAbsUrl(bu, link).String()) 
				fmt.Print(memo)
				c.Visit(element.Request.AbsoluteURL(link))
			}
		})
		
	
		c.OnRequest(func(request *colly.Request) {
		})
	
		c.Visit(baseUrl)

		return memo
}

func toAbsUrl(baseurl *url.URL, weburl string) *url.URL {
	relurl, err := url.Parse(weburl)
	if err != nil {
		return nil
	}

	absurl := baseurl.ResolveReference(relurl)
	absurlParced := absurl.Scheme + "://" + absurl.Host + absurl.Path

	rel, err := url.Parse(absurlParced)
	if err != nil {
		return nil
	}
	return rel
}

