package page

import "testing"

func TestGetMainPage(t *testing.T) {
	t.Log(GetMainPage().Html())
}
