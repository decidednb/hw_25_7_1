package main

import (
	"fmt"
	"hw_25_7_1/pkg/calc"
	"strconv"
)

var input string

func main() {

	for {
		fmt.Print("Введите первое число: ")
		firstNum := scanNum()

		fmt.Print("Введите арифметический оператор (/, -, +): ")
		operator := scanOperator()

		fmt.Print("Введите второе число: ")
		secondNum := scanNum()

		calculator := calc.NewCalculator()
		result := calculator.Calculate(firstNum, secondNum, operator)
		fmt.Println("Результат вычисления: ", result)
	}
}

func scanNum() (number float64) {

	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Не удалось считать число: ", err)
			continue
		}

		number, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Print("Не удалось конвертировать число, повторите ввод: ")
			continue
		}

		return number
	}
}

func scanOperator() (operator string) {

	for {
		_, err := fmt.Scanln(&operator)
		if err != nil {
			fmt.Print("Не удалось считать оператор, повторите ввод: ")
			continue
		}
		return operator
	}
}
