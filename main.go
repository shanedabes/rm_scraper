package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/hekmon/transmissionrpc"
	"github.com/jessevdk/go-flags"
)

var (
	NoMagnetFoundErr = errors.New("no magnet link found")
)

type Options struct {
	Host     string `short:"H" long:"host" description:"Transmission server address" env:"RMS_HOST" default:"localhost"`
	Port     uint16 `short:"P" long:"port" description:"Transmission server port" env:"RMS_PORT" default:"9091"`
	Secure   bool   `short:"s" long:"secure" description:"Connect to transmission using tls" env:"RMS_SECURE"`
	User     string `short:"u" long:"user" description:"Transmission server user" env:"RMS_USER" required:"true"`
	Password string `short:"p" long:"password" description:"Transmission server password" env:"RMS_PASSWORD" required:"true"`
}

func getMagnet(doc *goquery.Document) (string, error) {
	href, found := doc.Find(".magnet").Attr("href")

	if !found {
		return "", NoMagnetFoundErr
	}

	return href, nil
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)

	if err != nil {
		log.Fatal(err)
	}

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

	fmt.Println("Found magnet link:", magnet)

	bt, err := transmissionrpc.New(opts.Host, opts.User, opts.Password, &transmissionrpc.AdvancedConfig{
		HTTPS: opts.Secure,
		Port:  opts.Port,
	})
	if err != nil {
		log.Fatal(err)
	}

	torrent, err := bt.TorrentAdd(&transmissionrpc.TorrentAddPayload{
		Filename: &magnet,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added file to transmission:", *torrent.Name)
}
