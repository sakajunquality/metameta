package meta

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	MetaIndexCharset   = "charset"
	MetaIndexName      = "name"
	MetaIndexProperty  = "property"
	MetaIndexHttpEquiv = "http-equiv"
)

// Meta is the unit for parsed meta tag
type Meta struct {
	Index   string `json:"index"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Metas contains all the meta
type Metas []Meta

// GetHTMLContent gets HTML DOM and returns it as string
func GetHTMLContent(targetURL string) string {
	client := &http.Client{}
	res, _ := client.Get(targetURL)
	html, _ := ioutil.ReadAll(res.Body)

	return string(html)
}

// ParseMetaData retrieves meta tags from html string
func ParseMetaData(html string) Metas {
	metas := Metas{}

	r := strings.NewReader(html)
	doc, _ := goquery.NewDocumentFromReader(r)

	doc.Find("head > meta").Each(func(_ int, m *goquery.Selection) {
		var index string
		var name string
		var content string

		indexes := getMetaIndexes()
		for _, i := range indexes {
			indexName, ok := m.Attr(i)
			if ok {
				index = i
				name = indexName
				break
			}
		}

		contentValue, ok := m.Attr("content")
		if ok {
			content = contentValue
		}

		if index != "" {
			metas = append(metas, Meta{
				Index:   index,
				Name:    name,
				Content: content,
			})
		}
	})

	return metas
}

func getMetaIndexes() [4]string {
	return [4]string{
		MetaIndexCharset,
		MetaIndexName,
		MetaIndexProperty,
		MetaIndexHttpEquiv,
	}
}
