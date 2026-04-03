package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter number to convert: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			fmt.Println("Invalid input.")
			continue
		}

		if input == "quit" {
			return
		}

		fields := strings.Fields(input)
		if len(fields) != 2 {
			fmt.Println("Invalid format. Use: <number> <base>")
			continue
		}

		value := fields[0]
		base := strings.ToLower(fields[1])

		switch base {
		case "hex":
			convertToDecimal(value, 16)
		case "bin":
			convertToDecimal(value, 2)
		case "dec":
			convertFromDecimal(value)
		default:
			fmt.Println("Unsupported base. Use: hex, bin, or dec.")
		}
	}
}

func convertToDecimal(value string, base int) {
	num, err := strconv.ParseInt(value, base, 64)
	if err != nil {
		fmt.Println("Invalid number for the specified base.")
		return
	}

	fmt.Println("✦ Decimal:", num)
}

func convertFromDecimal(value string) {
	num, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("Invalid decimal number.")
		return
	}

	binary := strconv.FormatInt(num, 2)
	hex := strings.ToUpper(strconv.FormatInt(num, 16))

	fmt.Println("✦ Binary: ", binary)
	fmt.Println("✦ Hex:    ", hex)
}
