package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
	z = 0

}

func main() {

	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5

}
