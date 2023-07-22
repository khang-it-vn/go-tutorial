package main

import "fmt"

func main() {

	// Chúng ta muốn in ra câu xin chào 100 lần

	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	soLan := 1
	for {

		if soLan > 10 {
			break
		}
		if soLan%2 == 0 {
			soLan++
			continue
		}
		fmt.Println("Xin chào", soLan)
		soLan++
	}
}
