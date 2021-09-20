package query

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/PuerkitoBio/goquery"
    "github.com/saintfish/chardet"
    "golang.org/x/net/html/charset"
)

func GetTitle(url string) {

    res, _ := http.Get(url)
    defer res.Body.Close()

	buf, _ := ioutil.ReadAll(res.Body)
	
    det := chardet.NewTextDetector()
    detRslt, _ := det.DetectBest(buf)
	
    bReader := bytes.NewReader(buf)
    reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

    doc, _ := goquery.NewDocumentFromReader(reader)

    rslt := doc.Find("title").Text()
    fmt.Println(rslt)
}