package test

import "github.com/ToRnaDo1337/bonustask/math"

// PrintResult prints the result of an arithmetic operation
func PrintResult(result int) {
	println("Result:", result)
}

// PrintCalculatorResult prints the result of a Calculator struct operation
func PrintCalculatorResult(c *math.Calculator, operation string) {
	switch operation {
	case "add":
		result := math.Add(c.num1, c.num2)
		println("Result:", result)
	case "subtract":
		result := math.Subtract(c.num1, c.num2)
		println("Result:", result)
	case "multiply":
		result := c.Multiply()
		println("Result:", result)
	case "divide":
		result := c.Divide()
		println("Result:", result)
	default:
		println("Invalid operation")
	}
}
