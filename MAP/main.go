package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)

	// another way
	// var otherColors map[string]string
	// otherColors["white"] = "#ffffff"

	// yet another way
	moreColors := make(map[string]string)
	moreColors["black"] = "#000000"

	// fmt.Println(otherColors)
	fmt.Println(moreColors)
	delete(moreColors, "black")
	fmt.Println(moreColors)
}
