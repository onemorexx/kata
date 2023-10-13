package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var symbolsRomans = []string{"I", "V", "X", "L", "C", "D", "M"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	readString = strings.ReplaceAll(readString, "\r\n", "")
	parse(readString)
}

func parse(data string) {
	split := strings.Split(data, " ")
	if len(split) != 3 {
		panic("failed parse")
	}

	if isRoman(split[0]) {
		if isRoman(split[2]) {
			firstDigit := romanToInt(split[0])
			secondDigit := romanToInt(split[2])

			if firstDigit <= 0 || secondDigit <= 0 || firstDigit >= 11 || secondDigit >= 11 {
				panic("first or second digit < I or > X ")
			}

			result := process(split[1], firstDigit, secondDigit)
			result1 := integerToRoman(result)

			if result <= 0 {
				panic("no 0 or negative in romanian")
			}
			fmt.Println(result1)
			return
		} else {
			panic("scnd element not roman")
		}
	}
	if isRoman(split[2]) {
		panic("frst element not roman")
	}

	firstDigit, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}

	secondDigit, err := strconv.Atoi(split[2])
	if err != nil {
		panic(err)
	}

	if firstDigit <= 0 || secondDigit <= 0 || firstDigit >= 11 || secondDigit >= 11 {
		panic("first or second digit < 1 or > 10 ")
	}

	result := process(split[1], firstDigit, secondDigit)
	fmt.Println(result)
}

func add(firstDigit int, secondDigit int) int {
	return firstDigit + secondDigit
}

func subtract(firstDigit int, secondDigit int) int {
	return firstDigit - secondDigit
}

func divide(firstDigit int, secondDigit int) int {
	return firstDigit / secondDigit
}

func multiplication(firstDigit int, secondDigit int) int {
	return firstDigit * secondDigit
}

func isRoman(str string) bool {
	for i := range symbolsRomans {
		if strings.Contains(str, symbolsRomans[i]) {
			return true
		}
	}
	return false
}

func process(operation string, firstDigit int, secondDigit int) int {
	result := 0
	switch operation {
	case "+":
		result = add(firstDigit, secondDigit)
	case "-":
		result = subtract(firstDigit, secondDigit)
	case "/":
		result = divide(firstDigit, secondDigit)
	case "*":
		result = multiplication(firstDigit, secondDigit)
	default:
		panic("unknown operation")
	}
	return result
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	s = strings.Replace(s, "CM", "CCCCCCCCC", -1) // 900
	s = strings.Replace(s, "CD", "CCCC", -1)      // 400
	s = strings.Replace(s, "XC", "XXXXXXXXX", -1) // 90
	s = strings.Replace(s, "XL", "XXXX", -1)      // 40
	s = strings.Replace(s, "IX", "IIIIIIIII", -1) // 9
	s = strings.Replace(s, "IV", "IIII", -1)      // 4

	var sum int
	for _, roman := range s {
		sum += romans[roman]
	}
	return sum
}
