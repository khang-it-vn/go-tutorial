/*
Chương trình này sẽ đọc vào các loại điểm của sinh viên (điểm chuyên cần, điểm giữa kỳ, và điểm cuối kỳ)
	và xếp loại điểm theo quy luật sau:
– if điểm trung bình >=9 =>loại=A
– if điểm trung bình >= 7 và <9 => loại=B
– điểm còn lại là  =>loại=C

*/

package main

import "fmt"

func main() {
	var (
		diem_chuyen_can int = 9
		diem_giua_ky    int = 9
		diem_cuoi_ky    int = 0
	)

	diem_tb := (float32)(diem_chuyen_can+diem_giua_ky+diem_cuoi_ky) / 3

	fmt.Printf("Diem trung binh la %v\n", diem_tb)

	if diem_tb >= 9 {
		fmt.Println("Loại = A")
	} else if diem_tb >= 7 && diem_tb < 9 { // Diemtb >= 7 và < 9 7 ..... 8.999999 <9
		fmt.Println("Loại = B")
	} else {
		fmt.Println("Loại = C")
	}

}
