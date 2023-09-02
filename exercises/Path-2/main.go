package main

import (
	resolve "exercise-path-2-tutorial/resolve"
	"fmt"
)

func main() {
	path := "debai.txt"
	for {
		resolve.ReadFile(path)
		var choose int8
		fmt.Printf("\nEnter the number of exercise: ")
		fmt.Scan(&choose)
		switch choose {
		case 77:
			resolve.Bai_77_Method_1()
		case 78:
		case 79:
		case 80:
		}
	}
}
