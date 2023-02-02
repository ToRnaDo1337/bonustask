package math

// Add returns the sum of two numbers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference of two numbers
func Subtract(a, b int) int {
	return a - b
}

// Calculator is a struct that has methods for arithmetic operations
type Calculator struct {
	num1, num2 int
}

// Multiply returns the product of two numbers stored in the Calculator struct
func (c *Calculator) Multiply() int {
	return c.num1 * c.num2
}

// Divide returns the result of dividing two numbers stored in the Calculator struct
func (c *Calculator) Divide() int {
	return c.num1 / c.num2
}
