package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeCheck(t *testing.T) {
	tm := time.Now()
	for r, b := TimeCheck(tm, 30); r > 0; r, b = TimeCheck(tm, 30) {
		fmt.Println(r, b)
		time.Sleep(time.Second)
	}
}
