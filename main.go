package main

import (
	"fmt"
	"gigaget/util"
	"net/http"
	"os"
)

func main() {

	category, num, _ := util.ArgParser() //TODO date
	id := util.GetCategory(category)
	url := "https://gigazine.net/news/" + id
	if id == "" {
		url = "https://gigazine.net/"
	}
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}

	defer resp.Body.Close()
	articles := util.ArticleParse(resp.Body, num)

	for i, a := range articles {
		fmt.Printf("%d : %s - %s  (%s)\n", i+1, a.Data, a.Title, a.Category)
	}

}
