package main

import (
	"fmt"
)

// Complete the riddle function below.
func riddle(arr []int64) []int64 {
	final_result := make([]int64, 0)

	// Iterate through different window sizes from 1 to 5
	for pair_size := 1; pair_size <= len(arr); pair_size++ {
		window := make([][]int64, 0)
		new_result := make([]int64, 0)

		// Create sliding windows of current size
		for i := 0; i <= len(arr)-pair_size; i++ {
			pair := make([]int64, pair_size)
			for j := 0; j < pair_size; j++ {
				pair[j] = arr[i+j]
			}
			window = append(window, pair)
		}
		// fmt.Println("window: ", window)

		// Find minimum in each window
		for i := 0; i < len(window); i++ {
			min := window[i][0]
			for j := 1; j < pair_size; j++ {
				if window[i][j] < min {
					min = window[i][j]
				}
			}
			new_result = append(new_result, min)
		}
		// fmt.Println("new_result: ", new_result)

		// Find maximum of all minimums for current window size
		if len(new_result) > 0 {
			max := new_result[0]
			for i := 1; i < len(new_result); i++ {
				if new_result[i] > max {
					max = new_result[i]
				}
			}
			final_result = append(final_result, max)
		}
		// fmt.Println("final_result: ", final_result)
	}
	return final_result
}

func main() {
	array := []int64{1, 2, 3, 5, 1, 13, 3}
	result := riddle(array)
	fmt.Println("Result_pro:", result)
}
