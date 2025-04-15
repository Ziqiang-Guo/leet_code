package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		// 检查右边两个位置是否存在
		if flowerbed[i] == 0 {
			emptyPrev := (i == 0) || (flowerbed[i-1] == 0)
			emptyNext := (i == length-1) || (flowerbed[i+1] == 0)
			if emptyPrev && emptyNext {
				count++
				i++ // 跳过下一个位置，因为已经种了花
			}
		}
	}
	return count >= n
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 9))
}
