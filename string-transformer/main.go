package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func cap(s string) string {
	return strings.Title(s)
}

func Title(s string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true,
		"but": true, "or": true, "for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true,
		"is": true, "it": true,
	}

	words := strings.Fields(s)

	for i, w := range words {
		lower := strings.ToLower(w)

		if i == 0 || !smallWords[lower] {
			runes := []rune(lower)
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
			}
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}

	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Text Formatter Started")
	fmt.Println("Type 'exit' anytime to quit")

	for {
		fmt.Print("\nEnter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Goodbye paul")
			break
		}

		fmt.Print("choose upper, lower, cap, or title: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "exit" {
			fmt.Println("Goodbye paul")
			break
		}

		switch choice {
		case "upper":
			fmt.Println(ToUpper(text))
		case "lower":
			fmt.Println(ToLower(text))
		case "cap":
			fmt.Println(cap(text))
		case "title":
			fmt.Println(Title(text))
		default:
			fmt.Println("Invalid choice")
		}
	}
}
