package main

import "fmt"

func tong_hieu(a float32, b float32) (float32, float32) {
	sum := a + b
	hieu := a - b
	return sum, hieu
}

// parameter
func xuat(x float32) {
	fmt.Println(x)
}

func main() {

	// cho vào 2 số a và b có kiểu dữ liệu là float32 ,
	// sau đó thực hiện tính + - * / và xuất kết quả ra màn hình
	var x, y float32 = 3, 4

	s, h := tong_hieu(x, y)

	fmt.Println("tong ", s, " Hieu ", h)
}
