package main

import "fmt"

type rider int

var x rider

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println(x)
}
