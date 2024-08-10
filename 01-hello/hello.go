package main

import (
	"fmt"
)

const (
	indonesian = "Indonesian"
	french     = "French"

	englishHelloPrefix    = "Hello, "
	indonesianHelloPrefix = "Halo, "
	frenchHelloPrefix     = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case indonesian:
		prefix = indonesianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
