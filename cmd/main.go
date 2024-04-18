package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rushg171/go-link-parser"
)

func main() {
	file, err := os.Open("ex1.html")
	if err != nil {
		log.Fatal(err)
	}
	allLinks, err := link.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(allLinks)
}
