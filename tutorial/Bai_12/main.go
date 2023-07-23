package main

import "fmt"

type Car struct {
	Name  string
	Price float32
}

func main() {

	slice_car := []Car{}

	var car1 Car
	car1.Name = "BMW"
	car1.Price = 2.3
	var car2 Car
	car2.Name = "Volvo"
	car2.Price = 7.2

	slice_car = append(slice_car, car1, car2)

	// Hayx in ra giá tiền của BMW

	for _, value := range slice_car {
		if value.Name == "Volvo" {
			fmt.Println(value.Price)
		}
	}

}
