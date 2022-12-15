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

func TestSwitch(t *testing.T) {
	scr := NewScript("test/switch")
	scr.Run()
}
