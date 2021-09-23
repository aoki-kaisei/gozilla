package query

import (
	"github.com/gocolly/colly"
	"github.com/kaseiaoki/gozilla/array"
	"log"
	"net/url"
)

func GetLink(baseUrl string) []string {
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
		absurl := toAbsUrl(bu, link)

		if absurl != nil && !array.Contains(memo, absurl.String()) {
			memo = append(memo, absurl.String())
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

	if relurl.Host != baseurl.Host {
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
