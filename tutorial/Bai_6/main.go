package main

import "fmt"

func main() {

	// + - * / %
	var (
		a int
		b int
		c float32
		d int
	)

	a = 10
	b = 20
	c = (float32)(a) / (float32)(b)

	fmt.Println(c)
	d = a % b
	fmt.Println(d)
}
