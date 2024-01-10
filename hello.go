package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello(""))
	fmt.Print("The sum of 3+2 is %d", Add(3, 2))
}

func Hello(name string) string {
	if name == "" {
		return "Hello, World!"
	} else {
		return "Hello, " + name + "!"
	}
}
