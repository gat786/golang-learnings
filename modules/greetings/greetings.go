package greetings

import "fmt"

// In golang a function whos name starts with a capital letter
// can be used in another module
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
