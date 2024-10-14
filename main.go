package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var arr []float64
	var num float64

	for i := 0; i <= 12500; i++ {
		fmt.Fscan(os.Stdin, &num)
		arr = append(arr, num)
		var min int
		var max int
		// if len(arr) > 1 {
		// men := Average(arr)
		// standard := calculateStandardDeviation(arr, men)
		sort.Float64s(arr)
		medin := Medin(arr)
		min = int(medin - 40)
		max = int(medin + 40)
		// }
		// if len(arr) == 1 {
		// 	min = int(math.Max(0, num-40))
		// 	max = int(num) + 40
		// }
		fmt.Println(min, max)
	}
}

// func Average(datasep []float64) float64 {
// 	var countr float64
// 	for i := 0; i < len(datasep); i++ {
// 		countr += datasep[i]

// 	}

// 	return (countr / float64(len(datasep)))

// }
//
//	func calculateStandardDeviation(numbers []float64, mean float64) float64 {
//		sumSquaredDiff := 0.0
//		for _, num := range numbers {
//			diff := num - mean
//			sumSquaredDiff += diff * diff
//		}
//		variance := sumSquaredDiff / float64(len(numbers))
//		return math.Sqrt(variance)
//	}
func Medin(data []float64) float64 {
	if len(data)%2 != 0 {
		return data[len(data)/2]
	} else {
		var c float64
		c = (data[len(data)/2] + data[(len(data)/2)-1]) / 2
		return c
	}

}
