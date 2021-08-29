package main

import "fmt"

const helloEn = "Hello, "
const helloId = "Hai, "

func main() {
	fmt.Println(Hello("world!", ""))
}

func Hello(n string, lang string) string {
	if n == "" {
		n = "World"
	}

	return Greeting(lang) + n
}

func Greeting(lang string) string {
	switch lang {
	case "id":
		return helloId
	default:
		return helloEn
	}
}
