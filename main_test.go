package main

import (
	"log"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

var exampleBody = `
<html>
<body>
  <div class="container">
    <div class="row">
	  <div class="col-xs-12 col-sm-6 col-md-3">
	    <h4>
		  <a class="magnet" href="magnet:?xt=urn:btih:blah">magnet</a>
		</h4>
	  </div>
	</div>
  </div>
</body>
</html>
`

func TestGetMagnet(t *testing.T) {
	b := strings.NewReader(exampleBody)

	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		log.Fatal(err)
	}

	got, err := getMagnet(doc)

	assert.Nil(t, err)
	assert.Equal(t, got, "magnet:?xt=urn:btih:blah")
}

func TestGetMagnetError(t *testing.T) {
	b := strings.NewReader("")

	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		log.Fatal(err)
	}

	_, err = getMagnet(doc)

	assert.Equal(t, err, NoMagnetFoundErr)
}
