package main

import (
	"fmt"

	"github.com/PuerkitoBio/purell"
)

func main() {
	if normalized, err := purell.NormalizeURLString(
		"hTTp://someWEBsite.com:80/Amazing%3f/url/",
		purell.FlagLowercaseScheme|purell.FlagLowercaseHost|purell.FlagUppercaseEscapes,
	); err != nil {
		panic(err)
	} else {
		fmt.Print(normalized)
	}
}
