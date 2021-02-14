package main

import (
	"flag"
	"fmt"
	"os"
)

func checkDBExists() {
	if _, err := os.Stat("comics.json"); os.IsNotExist(err) {
		fmt.Println("Comics database not found. Run with --full-refresh flag.")
	}
}

func refresh() {
	fmt.Println("Checking for new comics.")
}

func fullRefresh() {
	fmt.Println("Regenerating comics database.")
}

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] keyword\n", os.Args[0])
	flag.PrintDefaults()
}

func search(keyword string) {
	fmt.Printf("Searching for comics with keyword: %s\n", keyword)
}

func main() {

	flag.Usage = usage

	refreshFlag := flag.Bool("refresh", false, "Check for new comics.")
	fullRefreshFlag := flag.Bool("full-refresh", false, "Regenerate comic database fronm scratch.")
	flag.Parse()

	if !*fullRefreshFlag && !*refreshFlag && len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *fullRefreshFlag {
		fullRefresh()
	}

	if *refreshFlag {
		refresh()
	}

	search(flag.Args()[0])

}
