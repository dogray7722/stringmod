package main

import (
	"fmt"
	"github.com/dogray7722/stringmod/quotes"
	"github.com/dogray7722/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(quotes.GetEmoji("turtle"))

}
