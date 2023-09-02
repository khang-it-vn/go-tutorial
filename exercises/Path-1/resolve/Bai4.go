package resolve

import "fmt"

// Bài 4: Tính S(n) = ½ + ¼ + … + 1/2n
func Bai_4_Method_1() {
	fmt.Println("\nEnter the value for n method 1:")

	var n float64
	fmt.Scan(&n)
	result := bai_4_method_1(n)

	fmt.Printf("\nValue of Bài 4: Tính S(n) = ½ + ¼ + … + 1/2.%v = %v", n, result)
}

func bai_4_method_1(n float64) float64 {
	if n == 1 {
		return 1.0 / 2
	}
	return 1/(2*n) + bai_4_method_1(n-1)
}

func Bai_4_Method_2() {
	fmt.Println("\nEnter the value for n method 2:")
	var n float64
	fmt.Scan(&n)

	result := bai_4_method_2(n)

	fmt.Printf("\nValue of Bài 4: Tính S(n) = ½ + ¼ + … + 1/2.%v = %v", n, result)
}

func bai_4_method_2(n float64) float64 {

	var result float64
	for i := 1.0; i <= n; i++ {
		result += (1 / (2 * i))
	}
	return result
}
