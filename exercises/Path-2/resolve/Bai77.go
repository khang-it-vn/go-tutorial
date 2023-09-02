package resolve

import "fmt"

func Bai_77_Method_1() {
	fmt.Println("\nEnter the value for n method 1: ")
	var n int64
	fmt.Scan(&n)
	result := bai_77_method_1(n)
	fmt.Printf("\n Result of the Bài 77: Viết chương trình tính tổng của dãy số sau: S(n) = 1 + 2 + 3 + … + %v = %v\n", n, result)
}

func bai_77_method_1(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n + bai_77_method_1(n-1)
}
