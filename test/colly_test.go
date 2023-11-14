package test

import (
	"fmt"
	"github.com/gocolly/colly"
	"testing"
)

func TestColly(t *testing.T) {
	for {
		c := colly.NewCollector()
		//time.Sleep(time.Second)
		err := c.Visit("https://www.ttysq.com/thread-3777094-1-1.html")

		fmt.Println(err, 123)
	}
}
