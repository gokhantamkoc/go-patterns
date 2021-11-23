package function_type

import "errors"

// functiontype shows that functions are first class citizen in Golang.
// They can be used as a type whenever we want to easily implement a Strategy pattern or similar.
// This pattern is heavily used in the golang http package https://golang.org/pkg/net/http/#Handler

type Greeting func(name string) (string, error)

func GreetingService(requestBody string, greeting Greeting) (string, error) {
	if requestBody != "" {
		return greeting(requestBody)
	}
	return "", errors.New("FATAL ERROR")
}
