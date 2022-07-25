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
