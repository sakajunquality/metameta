package meta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const HTMLRawString = `
	<!doctype html>
	<html>
	<head>
		<title>title</title>
		<meta charset="utf-8" />
		<meta http-equiv="Content-type" content="text/html; charset=utf-8" />
		<meta property="og:site_name" content="title" />
		<meta name="twitter:card" content="summary">
	</head>

	<body>
	<div>
		<h1>h1</h1>
		<p>this is a paragraph</p>
	</div>
	</body>
	</html>
`

func TestParseMetaData(t *testing.T) {
	expected := Metas{
		Meta{Index: "charset", Name: "utf-8", Content: ""},
		Meta{Index: "http-equiv", Name: "Content-type", Content: "text/html; charset=utf-8"},
		Meta{Index: "property", Name: "og:site_name", Content: "title"},
		Meta{Index: "name", Name: "twitter:card", Content: "summary"},
	}

	actual := ParseMetaData(HTMLRawString)
	assert.Equal(t, expected, actual)
}

func TestGetMetaIndexes(t *testing.T) {
	expected := [4]string{
		MetaIndexCharset,
		MetaIndexName,
		MetaIndexProperty,
		MetaIndexHttpEquiv,
	}

	actual := getMetaIndexes()
	assert.Equal(t, expected, actual)
}
