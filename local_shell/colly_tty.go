package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

// GOOS=linux GOARCH=amd64 go build -o $out/name.go main.go
// GOOS=linux GOARCH=amd64 go build   -o  ./tty.go local_shell/colly_tty.go

func main() {
	for {
		c := colly.NewCollector()
		//time.Sleep(time.Second)
		err := c.Visit("https://www.ttysq.com/thread-3777094-1-1.html")
		fmt.Println(err, 123)
	}
}
