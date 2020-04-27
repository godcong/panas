package main

import "github.com/godcong/panas/page"

func main() {
	//loop for handle new message
	for {
		for _, p := range page.List() {
			page.GetPage(p)
		}
	}
}
