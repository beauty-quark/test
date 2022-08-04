// Что нужно: нужно сделать программу, которая на вход получает список еды, которое я съел за выходные (будет указано название и количество калорий),
// а на выходе показывает:
// - сколько я суммарно съел калорий Total
// - какое блюдо было самым калорийным TopKkal
// - какое блюдо я ел чаще всегo TopNameFood
// - средняя калорийность MiddleKkal

package main

import "fmt"

type FoodList struct {
	List []Food
}

var ImputData = FoodList{
	List: []Food{
		{
			NameFood: "bread",
			Kkal:     350,
		},
		{
			NameFood: "omichka",
			Kkal:     380,
		},
		{
			NameFood: "bread",
			Kkal:     350,
		},
		{
			NameFood: "peach",
			Kkal:     200,
		},
		{
			NameFood: "milk",
			Kkal:     200,
		},
		{
			NameFood: "chips",
			Kkal:     400,
		},
		{
			NameFood: "peanut paste",
			Kkal:     400,
		},
	},
}

type Food struct {
	NameFood string
	Kkal     int
}
type Output struct {
	Total       int
	TopKkal     Food
	TopNameFood string
	MidKkal     int
}

func main() {

	result := count(ImputData)

	fmt.Printf("%+v", result)
}
