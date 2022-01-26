package main

import (
	"fmt"
	"is105-modules/myquote"
)

func main() {
	fmt.Println(myquote.GetGlassQuote())
	fmt.Println(myquote.GetGoQuote())
	fmt.Println(myquote.GetHelloQuote())
	fmt.Println(myquote.GetOptQuote())
}
