package main

func count(inp FoodList) Output {
	total := 0
	topKkal := Food{}
	topNameFoodMap := make(map[string]int)

	for _, food := range inp.List {
		total += food.Kkal

		if food.Kkal > topKkal.Kkal {
			topKkal = food
		}

		count := topNameFoodMap[food.NameFood]
		count++
		topNameFoodMap[food.NameFood] = count
	}

	topCount := 0
	topName := ""
	for name := range topNameFoodMap {
		if topNameFoodMap[name] > topCount {
			topCount = topNameFoodMap[name]
			topName = name
		}
	}

	return Output{
		Total:       total,
		TopKkal:     topKkal,
		TopNameFood: topName,
		MidKkal:     total / len(inp.List),
	}
}
