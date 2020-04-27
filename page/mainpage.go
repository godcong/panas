package page

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

var mainUrl = "https://www.nas2x.com/"

func GetMainPage() *goquery.Document {
	get, err := http.Get(mainUrl)
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
