package main

import (
	"fmt"
	"LearnGo/helloworld/morestrings"
)

func main() {
	// moreStrings has to be built before importing, only functions starting in upper case are exported
	fmt.Println(morestrings.ReverseRunes("ciao"))
}