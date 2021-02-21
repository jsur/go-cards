// files in the same package can freely call functions from each other
package main

import "fmt"

func main() {
	// := tells the compiler to infer type
	// used only when a new variable is defined
	card := "Ace of Spades"
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
