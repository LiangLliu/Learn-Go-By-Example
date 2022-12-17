package main

import "fmt"

func main() {
	fmt.Println(Hello("World"))
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if len(name) <= 0 {
		name = "World"
	}
	return englishHelloPrefix + name
}
