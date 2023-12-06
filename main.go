package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	isRoman  = false
	isArabic = false
)

var (
	errNumSystems     = errors.New("ошибка: используются одновременно разные системы счисления")
	errOutOfRange     = errors.New("ошибка: число вне диапазона")
	errRomanBelowZero = errors.New("ошибка: в римской системе нет отрицательных чисел")
	errWrongOp        = errors.New("ошибка: неверный символ операции")
	errWrongFormat    = errors.New("ошибка: формат математической операции не удовлетворяет заданию")
)

func main() {

	for {
		fmt.Println("Введите математическую операцию")

		spl := getInput()

		checkNumbers(spl)

		if isRoman && !isArabic {
			spl = convertToArabic(spl)
		} else if isRoman && isArabic {
			log.Fatalln(errNumSystems)
		}

		a, err := strconv.Atoi(spl[0])
		
		if err != nil {
			log.Fatalln(errWrongFormat)
		}
		b, err := strconv.Atoi(spl[2])
		
		if err != nil {
			log.Fatalln(errWrongFormat)
		}

		if (a < 1 || a > 10) || (b < 1 || b > 10) {
			log.Fatalln(errOutOfRange)
		}

		op := spl[1]

		if op == "+" || op == "-" || op == "/" || op == "*" {

			result := calcFunc(a, b, op)

			if isRoman {
				if result < 1 {
					log.Fatalln(errRomanBelowZero)

				}
				fmt.Println(toRoman(result))
				continue
			}
			fmt.Println(result)

		} else {
			log.Fatalln(errWrongOp)
		}
	}
}

func getInput() []string {
	var a, op, b string

	_, err := fmt.Scanln(&a, &op, &b)
	if err != nil {
		log.Fatalln(errWrongFormat)
	}
	input := []string{a, op, b}
	return input
}

func calcFunc(a int, b int, op string) int {

	var result int

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	}
	return result
}

func checkNumbers(spl []string) {
	for i := range spl {
		if spl[i] == "1" || spl[i] == "2" || spl[i] == "3" || spl[i] == "4" || spl[i] == "5" || spl[i] == "6" || spl[i] == "7" || spl[i] == "8" || spl[i] == "9" || spl[i] == "10" {
			isArabic = true
		}
		if spl[i] == "I" || spl[i] == "II" || spl[i] == "III" || spl[i] == "IV" || spl[i] == "V" || spl[i] == "VI" || spl[i] == "VII" || spl[i] == "VIII" || spl[i] == "IX" || spl[i] == "X" {
			isRoman = true
		}
	}
}

func convertToArabic(spl []string) []string {

	for k, v := range spl {

		switch v {
		case "I":
			spl[k] = strings.Replace(spl[k], "I", "1", 1)
			fallthrough
		case "II":
			spl[k] = strings.Replace(spl[k], "II", "2", 1)
			fallthrough
		case "III":
			spl[k] = strings.Replace(spl[k], "III", "3", 1)
			fallthrough
		case "IV":
			spl[k] = strings.Replace(spl[k], "IV", "4", 1)
			fallthrough
		case "V":
			spl[k] = strings.Replace(spl[k], "V", "5", 1)
			fallthrough
		case "VI":
			spl[k] = strings.Replace(spl[k], "VI", "6", 1)
			fallthrough
		case "VII":
			spl[k] = strings.Replace(spl[k], "VII", "7", 1)
			fallthrough
		case "VIII":
			spl[k] = strings.Replace(spl[k], "VIII", "8", 1)
			fallthrough
		case "IX":
			spl[k] = strings.Replace(spl[k], "IX", "9", 1)
			fallthrough
		case "X":
			spl[k] = strings.Replace(spl[k], "X", "10", 1)
		}
	}
	return spl
}

func toRoman(num int) string {
	convList := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"}}

	var roman strings.Builder

	for _, conv := range convList {
		for num >= conv.value {
			roman.WriteString(conv.digit)
			num -= conv.value
		}
	}
	return roman.String()
}
