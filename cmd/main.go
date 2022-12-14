package main

import (
	"DSL"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("bad")
		os.Exit(-1)
	}
	scr := DSL.NewScript("../scripts/" + args[1])
	scr.Run()
}
