package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	greetingGenerator := timeOfDayGreetingGenerator{}
	assert.Equal(t, "Good Evening", greetingGenerator.getGreeting())
}
