package main

import (
	"fmt"
	"time"
)

type greetingGenerator interface {
	getGreeting() string
}

type timeOfDayGreetingGenerator struct{}

func (timeOfDayGreetingGenerator) getGreeting() string {
	// custom logic to generate greeting based on the time of day
	currentTime := time.Now()
	timeStampString := currentTime.Format("2006-01-02 15:04:05")
	layOut := "2006-01-02 15:04:05"
	timeStamp, _ := time.Parse(layOut, timeStampString)
	hr, _, _ := timeStamp.Clock()
	var greeting string
	if hr < 12 {
		greeting = "Good Morning"
	} else {
		greeting = "Good Evening"
	}
	fmt.Println("inside getGreeting")
	return greeting
}
