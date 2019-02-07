package main

import "fmt"

func main() {

	messageService := greetingMessageService{timeOfDayGreetingGenerator{}}

	fmt.Println(messageService.getMessage("Sam"))
}
