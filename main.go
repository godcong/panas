package main

import (
	"fmt"
	"github.com/godcong/panas/page"
	"os"
	"strconv"
	"sync"
)

func main() {
	times := 1000
	max := true
	limit := 1000
	if len(os.Args) > 0 {
		parseInt, err := strconv.ParseInt(os.Args[0], 10, 32)
		if err == nil {
			//return
			limit = int(parseInt)
		}
	}
	if len(os.Args) > 1 {
		parseInt, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err == nil {
			//return
			times = int(parseInt)
		}
	}
	if times > 0 {
		fmt.Println("loop run: ", max, "times", times)
		max = false
	}
	//loop for handle new message
	for ; times > 0 || max; times-- {
		wg := sync.WaitGroup{}
		for ; limit > 0; limit-- {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, p := range page.List() {
					fmt.Println("pa:", p)
					page.GetPage(p)
				}
			}()
		}

		wg.Wait()
	}
}
