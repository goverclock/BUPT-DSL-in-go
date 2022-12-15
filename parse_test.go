package DSL

import "testing"

func TestParseArgs(t *testing.T) {
	arg := "username,\"G\",100.0"

	ret := parseArgs(arg)
	ans := []string{"username", "\"G\"", "100.0"}

	if len(ans) != len(ret) {
		t.FailNow()
	}
	for i, v := range ans {
		if ret[i] != v {
			t.FailNow()
		}
	}
}

func TestParseStatement(t *testing.T) {
	arg := "funcname(arg1,arg2,arg3)"
	ret1, ret2 := parseStatement(arg)
	if ret1 != "funcname" {
		t.FailNow()
	}
	ans2 := []string{"arg1", "arg2", "arg3"}
	if len(ans2) != len(ret2) {
		t.FailNow()
	}
	for i, v := range ans2 {
		if ret2[i] != v {
			t.FailNow()
		}
	}

	// bad arg
	arg = "funcname"
	ret1, ret2 = parseStatement(arg)
	if ret1 != "" || ret2 != nil {
		t.FailNow()
	}
}
