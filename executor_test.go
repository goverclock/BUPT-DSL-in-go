package DSL

import (
	"os"
	"os/exec"
	"testing"
)

func TestDispatchFunc(t *testing.T) {
	scr := NewScript("test/dispatch")
	scr.Run()
}

func TestDispatchFuncFail(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		scr := NewScript("test/badfunc")
		scr.Run()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestDispatchFuncFail")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

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

func TestSwitch(t *testing.T) {
	scr := NewScript("test/switch")
	scr.Run()
}
