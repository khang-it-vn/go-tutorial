package main
import ("fmt")

 var sum2 = 999
func main (){
    // 123, 862175, 2, -67  => int
    // 67 12 => uint
    // Hello => string
    // 4.2 => float32, float64
    // true, false => bool

    // Cach 1:
    var x int
    var y int = 1000
    x = 99
    var number1 uint = 88
    fmt.Println("Gia tri x = ",x, " , Gia tri y = ", y)
    fmt.Println(number1)

    // cach 2
    var so1, so2 = 5,9
    fmt.Println(so1, so2)
    // cach 3
    var (
        so3 int
        so4 string = "Xin Chao bai 3"
        so5 bool = true
    )
    fmt.Println(so3, so4, so5)
    fmt.Println(sum2)

    const PI = 3.14
    fmt.Println(PI)
    fmt.Print(PI)
    fmt.Print(PI)



}