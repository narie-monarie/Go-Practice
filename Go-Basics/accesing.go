package main

import "fmt"

func main() {
	zombies := []string{"Paul", "wahala", "Deh", "George"}
	zombies = append(zombies, "elena")
	fmt.Println(zombies)
}
