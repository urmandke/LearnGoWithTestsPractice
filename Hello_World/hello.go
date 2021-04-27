package main

import "fmt"

//Helllo returns Hello, world string
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
