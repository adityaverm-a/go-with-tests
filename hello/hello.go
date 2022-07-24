package main

import "fmt"

const japanese = "Japanese"
const englishHelloPrefix = "Hello, "
const japaneseHelloPrefix = "こんにちは、"

// It is good to separate your "domain" code from the outside world (side-effects).
//The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
func Hello(name string, language string) string {
	if name == "" {
		name = "Sir"
	}

	if language == japanese {
		return japaneseHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", "Japanese"))
}
