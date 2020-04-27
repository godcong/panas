package main

import (
	"fmt"
	"github.com/godcong/panas/page"
)

func main() {
	//loop for handle new message
	for {
		for _, p := range page.List() {
			fmt.Println("pa:", p)
			page.GetPage(p)
		}
	}
}
