package main

import "fmt"

type Car struct {
	Name  string
	Price float32
}

func main() {

	// cach 1 define map
	// var mapNumber = map[string]int{"one": 1, "two": 2, "three": 3}

	// cach 2

	// mapCar := map[string]float32{"BMW": 2.3, "Volvo": 7.2, "MG": 10.5}
	// fmt.Println(mapNumber["one"])

	// fmt.Println(mapCar["Volvo"])
	var bmw Car
	bmw.Name = "BMW"
	bmw.Price = 3.2

	var volvo Car
	volvo.Name = "Volvo"
	volvo.Price = 7.2
	mapCar2 := map[string]Car{"BMW": bmw, "Volvo": volvo}

	fmt.Println("Truoc khi xoa ", mapCar2)

	delete(mapCar2, "Volvo")

	fmt.Println("Sau khi xoa", mapCar2)

}
