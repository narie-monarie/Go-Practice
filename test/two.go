package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	fmt.Println(sort.IntsAreSorted(s))
	sort.Ints(s)
	fmt.Println(sort.IntsAreSorted(s))
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(sort.IntsAreSorted(s))
	fmt.Println(s)
	w := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(w)
	fmt.Println(w)
}
