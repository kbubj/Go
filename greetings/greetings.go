package greetings

import "fmt"

func Hello(name string) string {
	massage := fmt.Sprintf("Hi %v. Welcome!", name)
	return massage
}
