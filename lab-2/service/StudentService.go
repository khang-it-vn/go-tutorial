package service

import (
    "github.com/nguyenhoangkhangithutech/lab-2/model"
)

func AddStudent(student model.Student, listStudent *[]model.Student){

    *listStudent = append(*listStudent, student)
}