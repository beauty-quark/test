package main

func calculete(inp Input) Output {
	sum := 0
	topFood := FoodData{}
	for _, food := range inp.FoodList {
		sum += food.Kkal

		if food.Kkal > topFood.Kkal {
			topFood = food
		}
	}

	return Output{
		Sum:     sum,
		KklaTop: topFood,
	}
}
