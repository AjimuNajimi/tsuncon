package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		printUsage()
	} else {
		fmt.Print("ξ ﾟ⊿ ﾟ)ξ < args is ")
		fmt.Println(args)
		fmt.Println(args[0])
	}
}

func printUsage() {
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Hello, I'm Tsunko!")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < I will tell you how many your contributions on GitHub.")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Please input and run 'tsuncon <your GitHub name>'")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Bye!")
}
