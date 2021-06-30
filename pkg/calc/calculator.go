package calc

import "fmt"

type calculator struct{}

func NewCalculator() calculator {
	return calculator{}
}

func (c *calculator) Calculate(firstNum, secondNum float64, operator string) float64 {
	switch {
	case operator == "/":
		return c.division(firstNum, secondNum)
	case operator == "-":
		return c.subtraction(firstNum, secondNum)
	case operator == "+":
		return c.addition(firstNum, secondNum)
	case operator == "*":
		return c.multiplication(firstNum, secondNum)
	default:
		fmt.Println("Введен неизвестный арифметический оператор: ", operator)
		return 0
	}
}

func (c *calculator) division(firstNum, secondNum float64) float64 {
	if secondNum == 0 {
		fmt.Println("Ошибка: попытка деления на 0")
		return 0
	}
	return firstNum / secondNum
}

func (c *calculator) addition(firstNum, secondNum float64) float64 {
	return firstNum + secondNum
}

func (c *calculator) subtraction(firstNum, secondNum float64) float64 {
	return firstNum + secondNum
}

func (c *calculator) multiplication(firstNum, secondNum float64) float64 {
	return firstNum * secondNum
}
