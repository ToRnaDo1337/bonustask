package main

import (
	"github.com/ToRnaDo1337/bonustask/math"
	"github.com/ToRnaDo1337/bonustask/test"
)

func main() {
	c := &math.Calculator{Num1: 10, Num2: 5}
	test.PrintCalculatorResult(c, "add")
}
