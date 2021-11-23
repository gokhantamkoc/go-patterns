package function_type

import (
	"errors"
	"testing"
)

func TestGreetingServiceTests(t *testing.T) {
	t.Run("success when response body is not empty", func(t *testing.T) {
		result, err := GreetingService("world", func(name string) (string, error) {
			return "Hello " + name, errors.New("")
		})
		if result != "Hello world" && err != nil {
			t.Errorf("Expected: Hello world, actual: %s", result)
		}
	})
	t.Run("fail when responseBody is empty", func(t *testing.T) {
		result, err := GreetingService("", func(name string) (string, error) {
			return "Hello " + name, errors.New("")
		})
		if err == nil {
			t.Errorf("Expected: Hello world, actual: %s", result)
		}
	})
}
