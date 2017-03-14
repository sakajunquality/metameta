package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sakajunquality/metameta/meta"
)

func main() {
	html := meta.GetHTMLContent("http://example.com/")
	metas := meta.ParseMetaData(html)
	printMetaJSON(metas)
}

func printMetaJSON(m meta.Metas) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
	}
	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	fmt.Println(out.String())
}
