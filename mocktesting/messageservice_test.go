package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type greegingGeneratorMock struct {
	mock.Mock
}

func (todggMock *greegingGeneratorMock) getGreeting() string {
	args := todggMock.Called()
	return args.String(0)
}

func TestGetMessage(t *testing.T) {

	ggMock := new(greegingGeneratorMock)
	ggMock.On("getGreeting").Return("Good Morning")
	ms := greetingMessageService{ggMock}
	assert.Equal(t, "Good Morning John", ms.getMessage("John"))
}
