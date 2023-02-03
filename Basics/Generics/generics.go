package main

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	println(getSumGen(7, 8))
}
