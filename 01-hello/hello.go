package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	const englishHelloPrefix = "Hello, "
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
