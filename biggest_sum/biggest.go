package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 对整数数组进行从小到大排序
func sortArray(nums []int) []int {
	// 创建一个新的切片来存储排序结果，避免修改原始数组
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)

	// 使用Go标准库中的sort包进行排序
	sort.Ints(sortedNums)

	return sortedNums
}

func total_sum(nums []int) int {
	var result int = 0
	for i := 0; i < len(nums); i += 2 { // Fixed increment operator from "i = +2" to "i += 2"
		result = result + nums[i]
	}
	return result
}

func generateRandomArray(length int, maxValue int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue) // Generate random numbers between 0 and maxValue-1
	}
	return arr
}

func main() {
	// 示例数组
	// Generate a random array with specified length
	// Example: Generate an array with 10 elements, max value 100
	arr := generateRandomArray(6, 100)
	// arr := []int{6, 2, 6, 5, 1, 2}

	// 打印原始数组
	fmt.Println("原始数组:", arr)

	// 对数组进行排序并打印结果
	sortedArr := sortArray(arr)
	fmt.Println("排序后数组:", sortedArr)
	fmt.Println("最大和:", total_sum(sortedArr))
}
