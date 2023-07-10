package main

import (
    f "fmt"
    "github.com/nguyenhoangkhangithutech/lab-2/service"
    "github.com/nguyenhoangkhangithutech/lab-2/model"
)
func main ( ){

    listStudent := []model.Student{}
    var student model.Student

    student.Name = "Nguyen Hoang Khang"
    student.Age = 23
    student.Score = 3.8

    service.AddStudent(student,  &listStudent)

    f.Println(listStudent)

}