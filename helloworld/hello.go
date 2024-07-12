package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjur, "
	spanishHelloPrefix = "Hola, "
)

func hello(name string, language string) string {
	if len(name) == 0 {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
