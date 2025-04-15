package main

import (
	"fmt"
)

func change_for_lemonade(bills []int) bool {
	five_bills := 0
	ten_bills := 0
	twenty_bills := 0
	for _, v := range bills {
		if v == 5 {
			five_bills++
		} else if v == 10 {
			if five_bills > 0 {
				five_bills--
				ten_bills++
			} else {
				fmt.Print("5 bills are not enough")
				return false
			}
		} else if v == 20 {
			if five_bills > 0 && ten_bills > 0 {
				five_bills--
				ten_bills--
				twenty_bills++
			} else {
				return false
			}
		}
	}
	print("5 bills: ", five_bills)
	return five_bills >= 0
}

func main() {
	get_paied := []int{5, 5, 5, 10, 10, 20}
	fmt.Println("Hello, World!")
	fmt.Println(change_for_lemonade(get_paied))
}
