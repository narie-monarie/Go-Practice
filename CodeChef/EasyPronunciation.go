package codechef

import "fmt"

func main() {
	s := "murdyk"
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			count++
		}
	}
	println(count)
	fmt.Print([]rune("Hello"))
}
