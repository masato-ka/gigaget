package list

import (
	"context"
	"flag"
	"fmt"
	"gigaget/util"
	"github.com/google/subcommands"
	"net/http"
	"os"
)

type ListCmd struct {
	category string
	num      int
}

func (*ListCmd) Name() string     { return "list" }
func (*ListCmd) Synopsis() string { return "Show article list of GigaZine to stdout." }
func (*ListCmd) Usage() string {
	return `list [-category <category> -num <number>]:
  Show article list to stdout.
`
}

func (p *ListCmd) SetFlags(f *flag.FlagSet) {

	f.StringVar(&p.category, "category", "", "specify category")
	f.IntVar(&p.num, "num", 40, "list size")
	//f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *ListCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	id := util.GetCategory(p.category)
	url := "https://gigazine.net/news/" + id
	if p.category == "" {
		url = "https://gigazine.net/"
	}
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		os.Exit(1)
	}

	defer resp.Body.Close()
	articles := util.ArticleParse(resp.Body, p.num)

	//return articles

	for i, a := range articles {
		fmt.Printf("%d : %s - %s  (%s)\n", i+1, a.Data, a.Title, a.Category)
	}
	fmt.Println()
	return subcommands.ExitSuccess
}
