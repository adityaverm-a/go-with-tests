package main

import "fmt"

const englishHelloPrefix = "Hello, "

// It is good to separate your "domain" code from the outside world (side-effects).
//The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
func Hello(name string, language string) string {
	if name == "" {
		name = "Sir"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Mr. Wayne", ""))
}
