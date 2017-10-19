package main

import (
	"errors"
	"fmt"
	"os"
)

//!+ errorOn
type errorOn struct {
	open    error
	close   error
	process error
}

func (e errorOn) Open() error {
	fmt.Println("Open()")
	return e.open
}

func (e errorOn) Close() error {
	fmt.Println("Close()")
	return e.close
}

func (e errorOn) Process() error {
	fmt.Println("Process()")
	return e.process
}

//!- errorOn

func main() {
	ocp := errorOn{}

	for _, arg := range os.Args {
		switch arg {
		case "open":
			ocp.open = errors.New("open")
		case "close":
			ocp.close = errors.New("close")
		case "process":
			ocp.process = errors.New("process")
		}
	}

	err := process(ocp)
	fmt.Printf("Error %T: %v\n", err, err)
}
