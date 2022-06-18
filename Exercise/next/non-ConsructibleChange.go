package main

import (
	"fmt"
	"sort"
)

func change(coins []int) int {
	changeCreated := 0
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})
	for _, coin := range coins {
		if coin > changeCreated+1 {
			return changeCreated + 1
		}
		changeCreated += coin
	}
	return changeCreated + 1
}

func main() {
	cash := []int{5, 7, 1, 1, 2, 3, 22}
	fmt.Println(change(cash))
}
