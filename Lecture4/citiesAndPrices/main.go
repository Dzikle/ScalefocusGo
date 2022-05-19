package main

//pull
import (
	"fmt"
	"math/rand"
	"time"
)

//For Review
func citiesAndPrices() ([]string, []int) {
	rand.Seed(time.Now().UnixMilli())
	cityChoices := []string{"Berlin", "Moscow", "Chicago", "Tokyo", "London"}
	dataPointCount := 100

	// randomly choose cities
	cities := make([]string, dataPointCount)
	for i := range cities {
		cities[i] = cityChoices[rand.Intn(len(cityChoices))]
	}

	prices := make([]int, dataPointCount)
	for i := range prices {
		prices[i] = rand.Intn(100)

	}
	return cities, prices
}

func groupSlices(keySlice []string, valueSlice []int) map[string][]int {

	result := map[string][]int{}

	for idx, v := range keySlice {
		prices := []int{}

		if val, ok := result[v]; ok {
			//fmt.Println("OK")
			val = append(val, valueSlice[idx])
			result[v] = val
		} else {
			prices = append(prices, valueSlice[idx])
			result[v] = prices
		}
	}

	for k, v := range result {
		fmt.Println(k, v, "\n")
	}

	// fmt.Println(result)
	return result

}

func main() {

	groupSlices(citiesAndPrices())
}
