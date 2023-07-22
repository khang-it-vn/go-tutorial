// Slice
package main

import (
	"fmt"
)

func main() {

	// 1 cú pháp khai báo
	// slice1 := []int{1, 2, 3}

	// fmt.Println(len(slice1)) // len tính chiều dài
	// fmt.Println(cap(slice1)) // capacity khả năng chứa

	// slice1 = append(slice1, 4, 5, 6, 7, 8, 9, 10)

	// fmt.Println(len(slice1)) // len tính chiều dài
	// fmt.Println(cap(slice1)) // capacity khả năng chứa

	// cach 2: make ()

	// slice2 := make([]string, 4, 5)
	// slice2[0] = "1"
	// slice2[1] = "2"
	// slice2[2] = "3"
	// slice2[3] = "4"

	// fmt.Println(len(slice2)) //
	// fmt.Println(cap(slice2))

	// slice2 = append(slice2, "5", "6", "7", "8")

	// fmt.Println(len(slice2)) //
	// fmt.Println(cap(slice2))

	// slice2 = append(slice2, "9", "10", "11", "12")

	// fmt.Println(len(slice2)) //
	// fmt.Println(cap(slice2))

	// cách 3: tạo một slice bằng một array

	// arr := [4]float32{2.0, 1.0, 3.0, 4.0}

	// slice3 := arr[1:3] // tham số đầu tiên là vị trí chỉ số bắt đầu , lấy tới < hơn tham số vị trí thứ 2

	// fmt.Println(slice3)

	// slice4 := []int{1, 2, 3, 4}

	// slice4[3] = 10

	// fmt.Println(slice4)

	// danh sách sinh viên
	// 1. sinh viên 1
	// 2. sinh viên 2 .....

	slice_ds_sv := []string{"sv1", "sv2", "sv3", "sv4", "sv5", "sv6", "sv7"}
	slice_sv_345 := make([]string, len(slice_ds_sv))
	//
	fmt.Println(cap(slice_sv_345))
	copy(slice_sv_345, slice_ds_sv)
	fmt.Println(slice_sv_345)

	// cách khỏi tạo một slice: make , array , []{}, append(slice,...), copy() len cap()
}
