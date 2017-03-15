package main

import (
	"github.com/sakajunquality/metameta"
)

func main() {
	html := meta.GetHTMLContent("http://example.com/")
	metas := meta.ParseMetaData(html)
	meta.JSONStringfyMeta(metas)
}
