package resolve

import "fmt"

// Bài 2: Tính S(n) = 1^2 + 2^2 + … + n^2
func Bai_2_Method_1() {
	fmt.Println("\nNhap n method 1: ")
	var n int
	fmt.Scan(&n)
	result := bai_2_method_1(n)
	fmt.Printf("\nResult of 1^2 + 2^2 + … + %v^2 = %v", n, result)
}

func bai_2_method_1(n int) int {

	if n == 1 {
		return 1
	}
	return n*n + bai_2_method_1(n-1)
}

func Bai_2_Method_2() {
	fmt.Println("\nNhap n method 2: ")
	var n int
	fmt.Scan(&n)
	result := bai_2_method_2(n)

	fmt.Printf("\nResult of 1^2 + 2^2 + … + %v^2 = %v", n, result)

}

func bai_2_method_2(n int) int {
	sum_times := 0

	for i := 1; i <= n; i++ {
		sum_times += i * i
	}

	return sum_times
}
