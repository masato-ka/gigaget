package util

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"os"
)

type Article struct {
	Title    string
	Category string
	Link     string
	Data     string
	//time
}

type Articles []*Article

func createArticles(title string, category string, link string, date string) (a *Article) {
	a = new(Article)
	a.Title = title
	a.Category = category
	a.Link = link
	a.Data = date
	return a
}

func ArticleParse(reader io.Reader, num int) Articles {

	var arts Articles

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println("No url root")
		os.Exit(1)
	}
	// main section content div card
	count := 0
	doc.Find("div#main .content .card").Each(func(i int, s *goquery.Selection) {
		if count > num {
			return
		}
		text := s.Find("a span").Text()
		category := s.Find(".ctlink .catab").Text()
		time, _ := s.Find(".Data time").Attr("datetime")

		arts = append(arts, createArticles(text, category, "", time))
		count++
	})
	return arts
}
