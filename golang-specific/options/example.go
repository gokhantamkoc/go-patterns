package options

import "fmt"

type Option func(greeting *Greeting)

// NewGreeting creates a Greeting
func NewGreeting(options ...Option) *Greeting {
	greeting := &Greeting{
		name: "Stranger",
	}

	for _, option := range options {
		option(greeting)
	}

	return greeting
}

func Name(name string) Option {
	return func(greeting *Greeting) {
		greeting.name = name
	}
}

type Greeting struct {
	name string
}

func (g *Greeting) get() string {
	return fmt.Sprintf("Hello %s", g.name)
}
