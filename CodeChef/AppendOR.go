package main

import "fmt"

func x0r() {
	var a, b, target, d int
	fmt.Scanln(&a)
	var c []int

	for a != 0 {

		fmt.Scanln(&b, &target)
		x := 0
		for b != 0 {
			fmt.Scanf("%d", &x)
			c = append(c, x)
			b--
		}

		for i := 0; i < len(c); i++ {
			d |= c[i]
		}

		e := d ^ target

		if e|d != target {
			fmt.Println(-1)
		} else {
			fmt.Println(e)
		}

		c = c[:0]
		d = 0
		a--

	}
}
func main() {
	x0r()
}
