package main

import "fmt"

type QuyenSach struct {
	TenSach string
	NSX     string
	TacGia  string
	Gia     float32
}

func main() {

	var qs QuyenSach

	qs.TenSach = "Sách Lập Trình"
	qs.NSX = "22/12/1999"
	qs.TacGia = "Coder"
	qs.Gia = 30000.500

	fmt.Println(qs)
	fmt.Println("tác Giả: ", qs.TacGia)

}
