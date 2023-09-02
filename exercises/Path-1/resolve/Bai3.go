package resolve

import "fmt"

// Bài 3: Tính S(n) = 1 + ½ + 1/3 + … + 1/n
func Bai_3_Method_1() {
	fmt.Println("\n Nhap n method 1: ")
	var n int

	fmt.Scan(&n)

	result := bai_3_method_1(float64(n))

	fmt.Printf("\nResult of Bài 3: Tính S(n) = 1 + ½ + 1/3 + … + 1/%v = %v\n", n, result)
}

func bai_3_method_1(n float64) float64 {
	if n == 1 {
		return 1
	}
	return 1/n + bai_3_method_1(n-1)
}

func Bai_3_Method_2() {
	fmt.Println("\nNhap n method 2:")
	var n int
	fmt.Scan(&n)
	result := bai_3_method_2(float64(n))

	fmt.Printf("\nResult Of Bài 3: Tính S(n) = 1 + ½ + 1/3 + … + 1/%v = %v", n, result)
}

func bai_3_method_2(n float64) float64 {
	var result float64
	for i := n; i >= 1; i-- {
		result += 1 / i
	}

	return result
}
