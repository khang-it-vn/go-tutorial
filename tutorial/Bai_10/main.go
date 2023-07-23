package main

import "fmt"

// hàm khống có giá trị trả về
func sum(a int, b int) {
	fmt.Println(a + b)
}

// xuất dữ liệu, thông báo một cái gì đó

// hàm có giá trị trả về
func sum_2(a int, b int) int {
	return a + b
}

// cho các cái phép tính toán
// hãy cộng 2 số sau đó in ra màn hình, và sau đó lại lấy tổng đó + 100
func main() {

	var a, b int = 1, 2

	sum(a, b)

	gt_sum := sum_2(a, b)

	fmt.Println("Giá trị sau khi sum: ", gt_sum)

}
