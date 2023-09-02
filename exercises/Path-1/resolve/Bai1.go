package resolve

import "fmt"

// Bài 1: Tính S(n) = 1 + 2 + 3 + … + n
func Bai_1_Method_1() {
	var n int
	fmt.Println("\nNhap n method 1: ")
	fmt.Scan(&n)
	result := bai_1_method_1(n)
	fmt.Printf("\nresult sum of 1 - %v = %v ", n, result)
}

func bai_1_method_1(n int) int {
	if n == 1 {
		return 1
	}
	return bai_1_method_1(n-1) + n
}

// Bài 1: Tính S(n) = 1 + 2 + 3 + … + n
func Bai_1_Method_2() {
	var n int
	fmt.Println("\nNhap n method 2: ")
	fmt.Scan(&n)

	result := bai_1_method_2(n)

	fmt.Printf("\nresult 1 to %v = %v", n, result)
}

func bai_1_method_2(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
