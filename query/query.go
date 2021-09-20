package query

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"github.com/kaseiaoki/gozilla/array"
	"net/url"
)

func GetLink(baseUrl string) []string{
    res, _ := http.Get(baseUrl)
    defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)
	
    det := chardet.NewTextDetector()
    detRslt, _ := det.DetectBest(buf)
	
    bReader := bytes.NewReader(buf)
    reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

    doc, _ := goquery.NewDocumentFromReader(reader)

	var arr = []string{baseUrl}
	bu, _ := url.Parse(baseUrl)
    doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		var abu = toAbsUrl(bu, url)
		if(abu.Hostname() == bu.Hostname()) {
			arr = append(arr, abu.String())
		}
  	})
	au := array.Uniq(arr)
	return au
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


