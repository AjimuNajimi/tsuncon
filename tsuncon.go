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
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < 私の名前はつん子よ！")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < 面倒だけどあんたの今日のContribution数を教えてあげるわ、感謝しなさい！")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < こういう風に言ってくれれば教えてあげるわ -> 'tsuncon <your GitHub username>'")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < じゃあね！")
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
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < あんたの今日のContribution数は" + attr + "だわ！")
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < もっと頑張りなさいよね！")
		} else if 2 <= attrint && attrint <= 5 {
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < あんたの今日のContribution数は" + attr + "だわ！")
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < もっとやれるんじゃないの？")
		} else if 6 <= attrint && attrint <= 10 {
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < あんたの今日のContribution数は" + attr + "だわ！")
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < まあまあね、褒めてあげる。")
		} else {
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < あんたの今日のContribution数は" + attr + "だわ！")
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < へぇ、なかなかやるわね。ちょっと見直したかも。")
		}
	} else {
		fmt.Println("ξ ﾟ⊿ ﾟ)ξ < そんなユーザー名は存在しないわ！")
	}

	return nil
}
