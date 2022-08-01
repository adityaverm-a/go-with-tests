package hello

import "fmt"

const englishSir = "Sir"
const englishHelloPrefix = "Hello, "

const french = "French"
const frenchSir = "Monsieur"
const frenchHelloPrefix = "Bonjour, "

const japanese = "Japanese"
const japaneseSir = "様"
const japaneseHelloPrefix = "こんにちは、"

// It is good to separate your "domain" code from the outside world (side-effects).
//The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.
func Hello(name string, language string) string {
	if name == "" {
		name = getName(language)
	}

	return greetingPrefix(language) + name
}

func getName(language string) (name string) {
	switch language {
	case french:
		name = frenchSir

	case japanese:
		name = japaneseSir

	default:
		name = englishSir
	}

	return
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix

	case japanese:
		prefix = japaneseHelloPrefix

	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("", japanese))
}

// Quick Summary

// Got some understanding of:

// Some of Go's syntax around
// - Writing tests
// - Declaring functions, with arguments and return types
// - if, const and switch
// - Declaring variables and constants

// The TDD process and why the steps are important
// - Write a failing test and see it fail so we know we have written a relevant test for our requirements and seen that it produces an easy to understand description of the failure
// - Writing the smallest amount of code to make it pass so we know we have working software
// - Then refactor, backed with the safety of our tests to ensure we have well-crafted code that is easy to work with
