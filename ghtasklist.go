package main

import (
	"fmt"
	"os"

	"github.com/bfontaine/ghtasklist/ghtasklists"
)

func main() {
	s, err := ghtasklists.Transform(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(s)
}
