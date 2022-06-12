package main

import "fmt"

func main() {
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	value, doesItExist := colors["Red"]
	fmt.Println(value, doesItExist)

	color := map[string]string{
		"Blue":   "#f0f8ff",
		"Red":    "#f0f5kk",
		"Grey":   "#fffff0",
		"Yellow": "#f06f80",
	}

	delete(color, "Grey")
	for key, value := range color {
		fmt.Println(key, value)
	}
}
