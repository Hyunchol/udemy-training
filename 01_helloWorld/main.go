package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
