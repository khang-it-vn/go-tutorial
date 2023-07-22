package main

import "fmt"

func main() {

	var number int
	// Ví dụ: Nếu như chúng ta nhập vào 1 thì in ra Một nếu nhập 2 thì in ra 2, còn lại in ra Không xác định
	fmt.Println("Nhập vào một giá trị bất kỳ tôi sẽ in nó ra chữ ")
	fmt.Scanln(&number)
	switch number {
	case 1:
		fmt.Println("Một")
	case 2:
		fmt.Println("Hai")
	case 3:
		fmt.Println("Ba")
	default:
		fmt.Println("Không xác định")
	}
}
