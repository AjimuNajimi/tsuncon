package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Hello, World!")
	flag.Parse()
	args := flag.Args()
	fmt.Print("ξ ﾟ⊿ ﾟ)ξ < args is ")
	fmt.Println(args)
}
