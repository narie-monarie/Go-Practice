package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Apple"] = 40
	m["Orange"] = 50

	for key, val := range m {
		fmt.Println("[", key, " -> ", val, "]")
	}
}
