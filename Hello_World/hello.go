package main

import "fmt"

const englishHelloPrefix = "Hello, "

//Helllo returns Hello, world string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
