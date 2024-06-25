package animal

import "fmt"

func Move(meters int) {
	fmt.Printf("Moved %d meters\n", meters)
}

func Eat(kg int) {
	fmt.Printf("Spent %d minutes eating\n", kg*5)
}
