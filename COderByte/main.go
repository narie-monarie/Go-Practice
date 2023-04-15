package main

import (
	"fmt"
	"strconv"
)

func intersection(str1, str2 string) string {
	charMap := make(map[rune]bool)
	intersection := ""

	// Loop through each character of str1 and add it to the map
	for _, char := range str1 {
		charMap[char] = true
	}

	// Loop through each character of str2 and if it's present in the map, add it to the intersection string
	for _, char := range str2 {
		if _, ok := charMap[char]; ok {
			intersection += string(char)
		}
	}

	return intersection
}

func main() {
	x := []string{"1, 3, 4, 7, 13", "1, 2, 4, 13, 15"}
	a := ""
	b := ""

	for i := 0; i < len(x[0]); i++ {
		if x[0][i] == ',' {
			continue
		}
		a = string(x[0][i])
	}

	for i := 0; i < len(x[1]); i++ {
		if x[1][i] == ',' {
			continue
		}

		b = string(x[1][i])
		fmt.Println(strconv.Atoi(b))
	}
	fmt.Println(b)

	fmt.Println(intersection(a, b))

	//
	// c := strings.Split(a, ", ")
	// d := strings.Split(b, ", ")
	//
	// // ints1 := make([]int, len(c))
	//
	// for i, v := range c {
	// 	num, _ := strconv.Atoi(v)
	// 	ints1[i] = num
	// }
	//
	// ints2 := make([]int, len(d))
	//
	// for i, v := range d {
	// 	num, _ := strconv.Atoi(v)
	// 	ints2[i] = num
	// }

	// var e []string
	//
	// for i := 0; i < len(c); i++ {
	// 	e = append(e, c[i])
	// }
	//
	// for i := 0; i < len(d); i++ {
	// 	e = append(e, d[i])
	// }
	// freq := make(map[string]int)
	// for _, num := range e {
	// 	freq[num]++
	// }
	// // for num, count := range freq {
	// 	if count > 1 {
	// 		// fmt.Printf("%s,", num)
	// 	}
	// }
	//
	//	for i := 0; i < len(d); i++ {
	//	e = append(e, d[i])
	//}
}
