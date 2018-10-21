package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	t := time.Now()
	layout := "2006-01-02"
	ymd := t.Format(layout)

	// Set URL
	url := "https://github.com/users/" + username + "/contributions?to=" + ymd

	// Get document
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	// Get the number of contributions
	selection := doc.Find("rect").Last()
	attr, exists := selection.Attr("data-count")
	// Convert string to int
	attrint, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}

	if exists == true {
		if 0 <= attrint && attrint <= 1 {
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < The number of your contribution is " + attr + ".")
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Let's contribute!")
		} else if 2 <= attrint && attrint <= 5 {
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < The number of your contributions is " + attr + ".")
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Let's contribute!")
		} else if 6 <= attrint && attrint <= 10 {
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < The number of your contributions is " + attr + ".")
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Nice!")
		} else {
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < The number of your contributions is " + attr + ".")
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < Greeeeat!")
		}
	} else {
		fmt.Println("ξ ﾟ⊿ ﾟ)ξ < The user you inputted doesn't exist.")
	}

	return nil
}
