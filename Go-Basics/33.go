package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 1)
	fmt.Println("Im sleeping")
}

func main() {

	go slowFunc()
	fmt.Println("Get Rekt")
	time.Sleep(time.Second * 2)

}
