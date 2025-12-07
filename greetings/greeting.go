package greetings

import "fmt"

func Greeting(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func Goodbye(name string) string {
	return fmt.Sprintf("Goodbye, %s", name)
}
