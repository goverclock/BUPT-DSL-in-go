package DSL

import (
	"bufio"
	"log"
	"os"
)

type Script struct {
	lines []string
}

// load script from a file
func NewScript(fname string) *Script {
	// create object
	obj := new(Script)
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		obj.lines = append(obj.lines, scan.Text())
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	// analysis the script
	
	return obj
}
