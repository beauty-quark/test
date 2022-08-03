package main

import "fmt"

var InputData = Input{
	FoodList: []FoodData{
		{
			NameFood: "Vok",
			Kkal:     400,
		},
		{
			NameFood: "Salat",
			Kkal:     300,
		},
		{
			NameFood: "IceCreem",
			Kkal:     500,
		},
	},
}

func main() {
	println("mne ploho")

	result := calculete(InputData)

	fmt.Printf("result %+v", result)
}

// Что нужно: нужно сделать программу, которая на вход получает список еды, которое я съел за выходные (будет указано название и количество калорий),
// а на выходе показывает:
// - сколько я суммарно съел калорий
// - какое блюдо было самым калорийным
// - какое блюдо я ел чаще всегo
type Input struct {
	FoodList []FoodData
}

type FoodData struct {
	NameFood string
	Kkal     int
}

// NameFood Vok 400 Kkkal
// NameFood Salat 300 Kkkal
// NameFood IceCreem 500 Kkkal

// sameKkal 5000
// KklaTop 500 NameFood IceCreem
// nil

type Output struct {
	Sum     int
	KklaTop FoodData
	Popular string
}
