package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("..\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []string

	for scanner.Scan() {

		line := scanner.Text()

		numbers = append(numbers, line)
	}

	sum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		sum = addTwoSnafus(sum, numbers[i])
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func addTwoSnafus(snafu1 string, snafu2 string) string {

	if len(snafu1) > len(snafu2) {
		addedzeros := len(snafu1) - len(snafu2)

		prefix := ""
		for i := 0; i < addedzeros; i++ {
			prefix += "0"
		}
		snafu2 = prefix + snafu2
	} else if len(snafu2) > len(snafu1) {
		addedzeros := len(snafu2) - len(snafu1)

		prefix := ""
		for i := 0; i < addedzeros; i++ {
			prefix += "0"
		}
		snafu1 = prefix + snafu1
	}

	snafulen := len(snafu1)
	res := ""
	transfer := 0
	for i := snafulen - 1; i >= 0; i-- {

		sum := SnafuToDecimal(snafu1[i:i+1]) + SnafuToDecimal(snafu2[i:i+1]) + transfer
		digit := DecimalToSnafu(sum)
		res = digit[(len(digit)-1):] + res
		if len(digit) > 1 {
			transfer = SnafuToDecimal(digit[:(len(digit) - 1)])
		} else {
			transfer = 0
		}
	}

	if transfer != 0 {
		res = DecimalToSnafu(transfer) + res
	}

	return res
}

func DecimalToSnafu(number int) string {

	length := 1
	res := "2"

	if number > 0 {
		for SnafuToDecimal(res) < number {
			res += "2"
			length += 1
		}
	} else if number < 0 {
		res := "="
		for SnafuToDecimal(res) > number {
			res += "="
			length += 1
		}
	} else {
		return "0"
	}

	res = constructSnafu(length, length, number, "")
	return res
}

func constructSnafu(length int, fullLength int, goal int, current string) string {

	if length == 0 {
		return current
	}

	if length > 0 {

		for i := 0; i < 5; i++ {
			var digit string
			if i == 0 {
				digit = "="
			} else if i == 1 {
				digit = "-"
			} else if i == 2 {
				digit = "0"
			} else if i == 3 {
				digit = "1"
			} else if i == 4 {
				digit = "2"
			}

			tmp := constructSnafu(length-1, fullLength, goal, current+digit)
			if SnafuToDecimal(tmp) == goal {
				return tmp
			}
		}
	}

	return ""
}

func SnafuToDecimal(number string) int {

	weight := 1
	res := 0

	for i := len(number) - 1; i >= 0; i-- {
		var coef int
		if number[i] == '-' {
			coef = -1
		} else if number[i] == '=' {
			coef = -2
		} else {
			coef = int(number[i]) - 48
		}

		res += coef * weight
		weight *= 5
	}

	return res
}
