package printing

import (
	"fmt"

	"github.com/ToRnaDo1337/bonustask/math"
)

func PrintCalculatorResult(c *math.Calculator, operation string) {
	var result int
	switch operation {

	case "multiply":
		result = c.Multiply()
	case "divide":
		result = c.Divide()
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Printf("Result: %d\n", result)
}
