package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sc = fmt.Println

func main() {
	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	iPrimeArr := []int{2, 3, 5, 7, 11}

	var sPrimeArray []string

	for _, i := range iPrimeArr {
		sPrimeArray = append(sPrimeArray, strconv.Itoa(i))
	}

	for _, num := range sPrimeArray {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Open("data")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scan1 := bufio.NewScanner(f)

	for scan1.Scan() {
		sc("Prime: ", scan1.Text())
	}

	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}
