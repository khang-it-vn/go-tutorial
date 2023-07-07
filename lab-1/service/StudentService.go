package service
import(
    "go-tutorial/model"
    "fmt"
)
/** Add student to list
 * @param Student to add for list
 * @param List Student 
 */
func AddStudent(student model.Student, list *[]model.Student)*[]model.Student{
    *list = append(*list, student)
    return list
}


func PrintSpliceStudent(list []model.Student){
    fmt.Println(list)
}

func EditByMSSV(student model.Student, mssv string, list *[]model.Student) {
    for i  := 0 ; i < len(*list); i++{
      
        if (*list)[i].Mssv == mssv{
            model.UpdateInfo(&(*list)[i], student.Name, student.Age, student.Address, student.Gpa, student.Height)
            break
        }
    }
}

func DeleteByMSSV(list *[]model.Student, mssv string){
    for i:= 0; i < len(*list); i++{
        if (*list)[i].Mssv == mssv{
            *list = append((*list)[:i], (*list)[i+1:]...)
        }
    }
}

func SortWithSelectionSort(list *[]model.Student){
    for i:= 0; i < len(*list) - 1; i++{
        index_min := i 
        for j:=i + 1; j < len(*list); j++{
            if (*list)[index_min].Gpa > (*list)[j].Gpa{
                index_min = j
            }
        }
        swap(&(*list)[index_min], &(*list)[i])
    }
}

func swap(student *model.Student, student2 *model.Student){
    studentTemp := *student
    *student = *student2
    *student2 = studentTemp
}