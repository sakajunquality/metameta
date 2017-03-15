package meta

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSONStringfyMeta changes Metas to json string
func JSONStringfyMeta(metas Metas) string {
	jsonBytes, err := json.Marshal(metas)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
	}
	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	return out.String()
}
