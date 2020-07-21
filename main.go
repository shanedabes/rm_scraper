package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	NoMagnetFoundErr = errors.New("no magnet link found")
)

func getMagnet(doc *goquery.Document) (string, error) {
	href, found := doc.Find(".magnet").Attr("href")

	if !found {
		return "", NoMagnetFoundErr
	}

	return href, nil
}

func main() {
	res, err := http.Get("https://myrunningman.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	magnet, err := getMagnet(doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(magnet + "\n")
}
