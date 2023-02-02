package printing

import (
	"fmt"

	"github.com/ToRnaDo1337/bonustask/math"
)

// PrintCalculatorResult prints the result of a calculation based on the operation
func PrintCalculatorResult(c *math.Calculator, operation string) {
	var result int
	switch operation {
	case "add":
		result = math.Add(c.Num1, c.Num2)
	case "subtract":
		result = math.Subtract(c.Num1, c.Num2)
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
