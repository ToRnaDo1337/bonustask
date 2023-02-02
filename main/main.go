package main

import (
	"github.com/ToRnaDo1337/bonustask/math"
)

func main() {
	c := &math.Calculator{Num1: 10, Num2: 5}
	printing.PrintCalculatorResult(c, "add")
}
