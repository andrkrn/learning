package main

import "fmt"

func main() {
	fmt.Println(Hello("world!"))
}

func Hello(p string) string {
	return "Hello, " + p
}
