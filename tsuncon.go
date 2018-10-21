package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// parse args
	flag.Parse()
	args := flag.Args()

	// If args is nothing, print usage and exit.
	if len(args) == 0 {
		printUsage()
		os.Exit(0)
	}

	// Get username from args
	username := args[0]
	getContributions(username)
}

func printUsage() {
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Hello, I'm Tsunko!")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < I will tell you how many your contributions on GitHub.")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Please input and run 'tsuncon <your GitHub name>'")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Bye!")
}

func getContributions(username string) error {
	// processing
	return nil
}
