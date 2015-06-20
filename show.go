package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/codegangsta/cli"
)

type Problem struct {
	title      string
	conditions string
	content    string
}

func fetchProblem(url string) (p Problem) {
	doc, _ := goquery.NewDocument(url)
	description := doc.Find(".description").Each(func(_ int, s *goquery.Selection) {
		s.Find("h1").Remove()
		s.Find(".source").Remove()
		s.Find(".dat").Remove()
		s.Find(".spacer60").Remove()
	})

	p.title = doc.Find(".title").Text()
	p.conditions = doc.Find(".text-red3").Text()
	p.content = description.Text()

	return
}

func (p *Problem) showProblem() {
	fmt.Println("--- Title ---\n", p.title)
	fmt.Println("--- Condtions ---\n", p.conditions)
	fmt.Print("--- Content ---\n", p.content)
}

func doShow(c *cli.Context) {
	originURL := "http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id="
	problemNum := c.Args()[0]
	p := fetchProblem(originURL + problemNum)
	p.showProblem()
}
