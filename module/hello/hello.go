package hello

import (
	"fmt"

	"rsc.io/quote"
)

// Hello returns hello world
func Hello(str string) {
	fmt.Println(quote.Hello(), str)
}
