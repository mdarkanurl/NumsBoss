package main

import "fmt"

func decimalToBinary(num int) int {
	return num
}

func binaryToDecimal(num int) int {
	return num
}

func main() {
	var UserChoice string

	fmt.Print("What types of number do you convert: ")
	fmt.Scanln(&UserChoice)

	switch UserChoice {
	case "Decimal":
		fmt.Println("You chose Decimal to Binary")
		decimalToBinary(0)
	case "Binary":
		fmt.Println("You chose Binary to Decimal")
		binaryToDecimal(0)
	default:
		fmt.Println("Invalid choice")
		return
	}

	fmt.Println(UserChoice)
}
