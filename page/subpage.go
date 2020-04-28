package page

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

var pageList = []string{
	"https://www.nas2x.com/",
	"https://www.nas2x.com/forums/synology-dsm/",
	"https://www.nas2x.com/forums/other-nas/",
	"https://www.nas2x.com/forums/linux/",
	"https://www.nas2x.com/forums/router/",
	"https://www.nas2x.com/forums/windows-server/",
	"https://www.nas2x.com/forums/virtualization/",
	"https://www.nas2x.com/forums/share/",
	"https://www.nas2x.com/forums/11/",
	"https://www.nas2x.com/forums/15/",
	"https://www.nas2x.com/forums/2/",
}

func GetPage(url string) *goquery.Document {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			return
		}
	}()
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

func List() []string {
	return pageList
}
