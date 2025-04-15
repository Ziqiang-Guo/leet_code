package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'luckBalance' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. 2D_INTEGER_ARRAY contests
 */

func luckBalance(k int32, contests [][]int32) int32 {
	var n int32 = 0
	fmt.Println(contests)
	fmt.Println("n=", n, "k=", k)
	// Sort slice in descending order based on the first number of each subslice
	sort.Slice(contests, func(i, j int) bool {
		return contests[i][0] > contests[j][0]
	})
	fmt.Println(contests)
	for _, v := range contests {
		if v[1] == 0 {
			n += (v[0])
			fmt.Println("n +", v[0], "n=", n)
		} else if v[1] == 1 && k > 0 {
			n += (v[0])
			fmt.Println("n +", v[0], "n=", n)
			k--
		} else if v[1] == 1 && k == 0 {
			n -= int32(v[0])
			fmt.Println("n -", v[0], "n=", n)
		} else {
			panic("Bad input")
		}
	}
	return int32(n)
}

func main() {
	fmt.Println(luckBalance(3, [][]int32{{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0}}))
}
