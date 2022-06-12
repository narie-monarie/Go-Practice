package main

import "fmt"

func IsUnique(s string) bool {
	m := make(map[rune]bool)
	for _, i := range s {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = true
	}
	return true

}

func main() {
	fmt.Println(IsUnique("hello"))
}
