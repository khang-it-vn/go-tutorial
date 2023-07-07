package main
import (
    _"fmt"
     "go-tutorial/model"
     "go-tutorial/service"
)
func main(){
    var list = []model.Student{}
    var student model.Student
    
    student = model.InitStudent("Khang", 22, "191106", "address", 3.8, 162.2)
    
    service.AddStudent(student, &list)
    student.Name= "New Name"
    service.EditByMSSV(student,"191106", &list)
    service.PrintSpliceStudent(list)
    var student2 model.Student
    student2 = model.InitStudent("Khang", 22, "191107", "address", 2.8, 162.2)
    
    service.AddStudent(student2, &list)
    service.PrintSpliceStudent(list)
    service.SortWithSelectionSort(&list)
    // service.DeleteByMSSV(&list, "191106")
    service.PrintSpliceStudent(list)


}