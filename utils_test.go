package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONStringfyMeta(t *testing.T) {
	testMetas := Metas{
		Meta{Index: "testIndex", Name: "testName", Content: "testContent"},
	}

	expected := `[
    {
        "index": "testIndex",
        "name": "testName",
        "content": "testContent"
    }
]`

	actual, err := JSONStringfyMeta(testMetas)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
