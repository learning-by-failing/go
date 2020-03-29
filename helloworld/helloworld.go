package main

import (
	"fmt"
	"LearnGo/helloworld/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	// moreStrings has to be built before importing, only functions starting in upper case are exported
	fmt.Println(morestrings.ReverseRunes("ciao"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}