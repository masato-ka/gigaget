package util

import "flag"

func ArgParser() (string, int, int) {
	category := flag.String("Category", "", "Select Category of GigaZine.(default=none)")
	num := flag.Int("num", 40, "The number of Article. (default=40)")
	date := flag.Int("Data", 7, "Specify how many days ago Articles are displayed. (default=7)")
	flag.Parse()
	return *category, *num, *date
}
