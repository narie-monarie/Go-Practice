package romanToInt

func romanToInt(s string) int {

	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	var intNum int = 0
	var prevNum int = 0
	var sum int = 0

	for _, j := range s {

		intNum = m[string(j)]
		if prevNum < intNum {
			sum = sum - (prevNum * 2)

		}
		sum = sum + intNum
		prevNum = intNum

	}

	return sum

}
