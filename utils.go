package meta

import (
	"bytes"
	"encoding/json"
)

// JSONStringfyMeta changes Metas to json string
func JSONStringfyMeta(metas Metas) (string, error) {
	jsonBytes, err := json.Marshal(metas)
	if err != nil {
		return "", err
	}
	out := new(bytes.Buffer)
	json.Indent(out, jsonBytes, "", "    ")
	return out.String(), nil
}
