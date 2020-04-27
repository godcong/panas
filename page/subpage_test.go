package page

import "testing"

func TestGetPage(t *testing.T) {
	for {
		for _, s := range pageList {
			GetPage(s).Html()
		}
	}
}
