package main

type candys []string

func makeCandy() []string { // different color and taste candy maker function
	customCandys := candys{}

	types := []string{"sweet", "salty", "sour"}
	colors := getColor()
	for _, type1 := range types {
		for _, color := range colors {
			customCandys = append(customCandys, type1+color+"candy")
		}
	}
	return customCandys
}

func getColor() []string {
	colors := candys{"red", "green"}
	return colors
}
