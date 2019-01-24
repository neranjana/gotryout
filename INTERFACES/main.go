package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// func (eb englishBot) getGreeting() string can be written
// as follow, omiting eb if it is not used
func (englishBot) getGreeting() string {
	// custom logic to generate English greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// custom logic to generate Spanish greeting
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
