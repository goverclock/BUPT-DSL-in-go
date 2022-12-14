package main

import (
	"DSL"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " <script name>")
		os.Exit(-1)
	}
	scr := DSL.NewScript("scripts/" + args[1])
	scr.Run()
}
