package model

type Student struct{
	Name string;
	Age int;
	Mssv string;
	Address string;
	Gpa float32;
	Height float32;
}

func InitStudent(name string, age int,mssv string, address string, gpa float32, height float32) Student{
	var student Student
	student.Name = name
	student.Age = age
	student.Mssv = mssv
	student.Address = address
	student.Gpa = gpa
	student.Height = height

	return student
}

func UpdateInfo(student *Student, name string, age int, address string, gpa float32, height float32) {
	student.Name = name
	student.Age = age
	student.Address = address
	student.Gpa = gpa
	student.Height = height
}
