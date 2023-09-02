package main

import (
	resolve "exercise-path-1-tutorial/resolve"
	"fmt"
)

func main() {

	path := "debai.txt"
	for {
		resolve.ReadFile(path)
		var choose int8
		fmt.Println("Enter the number of exercise: ")
		fmt.Scan(&choose)
		switch choose {
		case 1:
			resolve.Bai_1_Method_1()
			resolve.Bai_1_Method_2()
		case 2:
			resolve.Bai_2_Method_1()
			resolve.Bai_2_Method_2()
		case 3:
			resolve.Bai_3_Method_1()
			resolve.Bai_3_Method_2()
		case 4:
			resolve.Bai_4_Method_1()
			resolve.Bai_4_Method_2()
		case 5:
			resolve.Bai_5_Method_1()
		}

	}
}
