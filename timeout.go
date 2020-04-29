package main

import "time"

func TimeCheck(t time.Time, limit int64) (remain int64, fa bool) {
	now := time.Now().Unix()
	f := t.Unix() + limit
	remain = -(now - f)
	if remain < 0 {
		remain = 0
	}
	return remain, remain <= 0
}
