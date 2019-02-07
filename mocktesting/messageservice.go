package main

type messageService interface {
	getMessage(name string) string
}

type greetingMessageService struct {
	gg greetingGenerator
}

func (g greetingMessageService) getMessage(name string) string {
	return g.gg.getGreeting() + " " + name
}
