package main

import "fmt"

func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""

	for n > 0 {
		remainder := n % 2
		binary = fmt.Sprint(remainder) + binary
		n = n / 2
	}

	return binary
}

func binaryToDecimal(binary string) int {
	decimal := 0
	power := 1

	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			decimal += power
		}

		power *= 2
	}

	return decimal
}

func main() {
	for {
		var UserChoice string

		fmt.Print("What types of numbers do you convert: ")
		fmt.Scanln(&UserChoice)

		switch UserChoice {
		case "Decimal":
			fmt.Print("Give your decimal number: ")
			var decimalNumber int
			fmt.Scanln(&decimalNumber)
			fmt.Println("Your binary number is:", decimalToBinary(decimalNumber))
		case "Binary":
			fmt.Print("Give your binary number: ")
			var bainaryNumber string
			fmt.Scanln(&bainaryNumber)
			fmt.Println("Your binary number is:", binaryToDecimal(bainaryNumber))
		case "Exit":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
