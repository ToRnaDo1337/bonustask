package math

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

type Calculator struct {
	Num1, Num2 int
}

// Multiply returns the product of two numbers stored in the Calculator struct
func (c *Calculator) Multiply() int {
	return c.Num1 * c.Num2
}

// Divide returns the result of dividing two numbers stored in the Calculator struct
func (c *Calculator) Divide() int {
	return c.Num1 / c.Num2
}
