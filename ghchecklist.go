package main

import (
	"fmt"
	"os"

	"github.com/bfontaine/ghchecklist/githubchecklist"
)

func main() {
	s, err := ghchecklist.Transform(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(s)
}
