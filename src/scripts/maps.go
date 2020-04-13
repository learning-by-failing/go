package main

import (
	"fmt"
)

func main() {
	votes := make(map[string]int)
	votes["Good"] = 8
	votes["Bad"] = 4
	fmt.Println(votes)

	cakeIngrediends := map[string]int{
		"sugar":  100,
		"Choco":  50,
		"Banana": 200,
		"Water":  800,
	}
	fmt.Println(cakeIngrediends)
	for key, value := range cakeIngrediends {
		//the output order is not guarantied to be the same of the entered item order
		fmt.Println("Key:", key, "value:", value)
	}
	fmt.Println("Banana weight is", cakeIngrediends["Banana"], "gr.")
	delete(cakeIngrediends, "Water")
	fmt.Println(cakeIngrediends)
}
