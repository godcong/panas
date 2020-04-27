package page

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

var pageList = []string{
	"https://www.nas2x.com/forums/synology-dsm/",
}

func GetPage(url string) *goquery.Document {
	get, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		panic(err)
	}
	reader, err := goquery.NewDocumentFromReader(bytes.NewReader(all))
	if err != nil {
		panic(err)
	}
	return reader
}
