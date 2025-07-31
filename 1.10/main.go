package main

import "fmt"

func main() {

	groupedTemps := make(map[int][]float32) // карта для хранения температур

	temps := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, t := range temps {

		rangeKey := int(t) / 10 // Определяем диапазон в который входит температура

		tempList, exists := groupedTemps[rangeKey]
		if exists {
			tempList = append(tempList, t)
			groupedTemps[rangeKey] = tempList
		} else {
			groupedTemps[rangeKey] = []float32{t}
		}
	}

	fmt.Println(groupedTemps)
}
