package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

// Crawler struct
type Crawler struct {
	URL string
}

// type doc *Document

// Crawle method
func (c Crawler) Crawle(selector string) {

	response, _ := http.Get(c.URL)
	defer response.Body.Close()

	buf, _ := ioutil.ReadAll(response.Body)
	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(c.detectCharset(buf), bReader)
	doc, _ := goquery.NewDocumentFromReader(reader)
	selection := doc.Find(selector)

	fmt.Println(selection.Text())
}

func (c Crawler) detectCharset(buf []byte) string {
	det := chardet.NewTextDetector()
	detResult, _ := det.DetectBest(buf)
	return detResult.Charset
}
