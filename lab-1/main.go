package main
import (
    "fmt"

    // import from package
     "go-tutorial/model"
)
func main(){
    fmt.Println("Hello World")
    var student model.Student
    student.Name = "Khang"
    fmt.Println(student)
}