package main

import "fmt"

type Alarm struct {
	Time  string
	Sound string
}

func alarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Duto",
	}
	return a
}

func main() {

	fmt.Println(alarm("11:21"))
}
