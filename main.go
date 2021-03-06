package main

import (
	"fmt"
	"github.com/godcong/panas/page"
	"go.uber.org/atomic"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	times := int64(3)
	limit := 1000
	if len(os.Args) > 1 {
		parseInt, err := strconv.ParseInt(os.Args[1], 10, 32)
		if err == nil {
			//return
			limit = int(parseInt)
		}
	}
	if len(os.Args) > 2 {
		parseInt, err := strconv.ParseInt(os.Args[2], 10, 32)
		if err == nil {
			//return
			times = parseInt
		}
	}

	tm := time.Now()
	random := rand.New(rand.NewSource(tm.UnixNano()))
	stop := atomic.NewBool(false)
	//loop for handle new message
	for ; limit > 0; limit-- {
		go func() {
			for {
				for _, p := range page.List() {
					page.GetPage(p)
					if stop.Load() {
						return
					}
					time.Sleep(time.Duration(random.Int63n(1000)) * time.Millisecond)
				}
			}
		}()
	}

	for r, _ := TimeCheck(tm, times); r > 0; r, _ = TimeCheck(tm, times) {
		fmt.Println("time remain", r)
		time.Sleep(time.Second)
	}

	stop.Store(true)
	//waiting for all done
	time.Sleep(3 * time.Second)
	fmt.Println("done")

}
