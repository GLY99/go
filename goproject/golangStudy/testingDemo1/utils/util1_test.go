package utils

import (
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := AddUpper(3)
	if res != 6 {
		t.Fatalf("utils.AddUpper(3) run fail, real=%v, expect=%v", res, 6)
	}
	t.Logf("utils.AddUpper(3) run succ")
}
