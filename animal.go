package animal

import "fmt"

func Move(meters int) {
	fmt.Printf("Moved %d meters", meters)
}

func Eat(kg int) {
	fmt.Printf("Spent %d minutes eating", kg*5)
}
