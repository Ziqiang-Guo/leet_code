package main

import (
	"fmt"
	"sort"
)

func sortArray(nums []int) []int {
	// 创建一个新的切片来存储排序结果，避免修改原始数组
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	// 使用Go标准库中的sort包进行排序
	sort.Ints(sortedNums)
	return sortedNums
}

func findContentChildren(g []int, s []int) int {
	// 对数组进行排序
	g = sortArray(g)
	s = sortArray(s)
	// 定义一个变量来存储满足条件的孩子数量
	count := 0
	// 遍历饼干数组
	for a, v := range s {
		print(a)
		// 如果当前饼干可以满足当前孩子的胃口，则满足条件的孩子数量加1
		if v >= g[count] {
			count = count + 1
		}
	}
	return count
}

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 2, 3}
	fmt.Println("\nThe can feed ", findContentChildren(g, s), "chirdrens")

}
