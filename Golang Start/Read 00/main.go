package main

import (
	"fmt"
)

// import "fmt"

func newCard() string {
	return "new Card on the Block!"
}
func main() {
	card := deck{newCard(), "I got it", "The best way", "Ace", "Global Elite", "Supreme Master First Class"}
	hand, remainingCards := deal(card, 2)
	hand.print()
	remainingCards.print()
	// greeting := "Hi There"
	// for _, i := range greeting {
	// 	fmt.Printf("%s", string(i))
	// }
}
