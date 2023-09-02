package resolve

import "fmt"

func Bai_5_Method_1() {
	fmt.Println("\nEnter the value of method 1")
	var n float64
	fmt.Scan(&n)
	result := bai_5_method_1(n)
	fmt.Printf("\nValue of Bài 5: Tính S(n) = 1 + 1/3 + 1/5 + … + 1/(2.%v + 1) = %v\n", n, result)
}

func bai_5_method_1(n float64) float64 {
	if n == 0 {
		return 1
	}
	return 1.0/(2*n+1) + bai_5_method_1(n-1)
}
