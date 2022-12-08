package DSL

import (
	"fmt"
	"testing"
)

func TestParseArgs(t *testing.T) {
	arg := "username,\"G\",100.0"

	ret := parseArgs(arg)
	ans := []string{"username", "\"G\"", "100.0"}

	fmt.Println(ret)

	if len(ans) != len(ret) {
		t.FailNow()
	}
	for i, v := range ans {
		if ret[i] != v {
			t.FailNow()
		}
	}
}
