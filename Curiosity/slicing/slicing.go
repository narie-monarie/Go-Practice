package main

import "fmt"

func main() {
	meetings := []string{"12:15AM-02:00PM", "02:30PM-03:10PM"}
	var meetingCopy []string
	for i := 0; i < len(meetings); i++ {
		meetingCopy = append(meetingCopy, meetings[i])
	}
	fmt.Println(meetingCopy)

	str := "hello_"

	if str[len(str)-1] == '_' {
		fmt.Println("False")
	}
}
