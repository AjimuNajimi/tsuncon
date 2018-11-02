package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Level1 -> 0 to 1
var Level1 = []string{
	"ξ ﾟ⊿ ﾟ)ξ < もっと頑張りなさい！",
	"ξ ﾟ⊿ ﾟ)ξ < しっかりしなさいよね！",
	"ξ ﾟ⊿ ﾟ)ξ < あんたなら出来るはずよ！",
}

// Level2 -> 2 to 5
var Level2 = []string{
	"ξ ﾟ⊿ ﾟ)ξ < もっとやれるんじゃないの？",
	"ξ ﾟ⊿ ﾟ)ξ < まだまだ開発するのよ！",
	"ξ ﾟ⊿ ﾟ)ξ < こんなものじゃ足りないわ！",
}

// Level3 -> 6 to 10
var Level3 = []string{
	"ξ ﾟ⊿ ﾟ)ξ < まあまあね、褒めてあげる。",
	"ξ ﾟ⊿ ﾟ)ξ < ふーん、やるじゃない。",
	"ξ ﾟ⊿ ﾟ)ξ < ついでに褒めてあげる。ついでによ、つ・い・で・に！",
}

// Level4 -> 11 or more
var Level4 = []string{
	"ξ ﾟ⊿ ﾟ)ξ < へぇ、すごいわね。ちょっと見直したかも。",
	"ξ ﾟ⊿ ﾟ)ξ < す、すご……。ちょっ、今のなし！",
	"ξ ﾟ⊿ ﾟ)ξ < 結構すごいんじゃないの、あんたにしては。",
}

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
	GetContributions(username)
}

func printUsage() {
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < 私の名前はつん子よ！")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < 面倒だけどあんたの今日のContribution数を教えてあげるわ、感謝しなさい！")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < こういう風に言ってくれれば教えてあげるわ -> 'tsuncon <your GitHub username>'")
	fmt.Println("ξ ﾟ⊿ ﾟ)ξ < じゃあね！")
}

func getContributions(username string) error {
	// Deal with time zone
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
		fmt.Println("ξ ﾟ⊿ ﾟ)ξ < そんなユーザー名は存在しないわ！\n")
		return err
	}

	if exists == true {
		// Generate seed
		rand.Seed(time.Now().UnixNano())
		switch {
		case 0 <= attrint && attrint <= 1:
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < あんたの今日のContribution数は" + attr + "だわ！")
			fmt.Println(Level1[rand.Intn(3)])
		case 2 <= attrint && attrint <= 5:
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < あんたの今日のContribution数は" + attr + "だわ！")
			fmt.Println(Level2[rand.Intn(3)])
		case 6 <= attrint && attrint <= 10:
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < あんたの今日のContribution数は" + attr + "だわ！")
			fmt.Println(Level3[rand.Intn(3)])
		default:
			fmt.Println("ξ ﾟ⊿ ﾟ)ξ < あんたの今日のContribution数は" + attr + "だわ！")
			fmt.Println(Level4[rand.Intn(3)])
		}
	}

	return nil
}
