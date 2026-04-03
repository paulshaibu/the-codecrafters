A collection of simple terminal-based tools written in Golang.

This project contains three interactive command-line programs:

Text Formatter
Calculator
Number Base Converter

Each program runs in a loop and behaves like a small terminal application.

Requirements
Go installed (go version)

Download Go: https://go.dev/dl/

How to Run

Clone or download the project.

go run main.go

Or run each tool separately depending on the file name:

go run formatter.go
go run calculator.go
go run converter.go
Text Formatter

An interactive text processing tool.

Features
Convert text to UPPERCASE
Convert text to lowercase
Capitalize words
Smart title formatting
Supported Commands
Command	Description
upper	Converts text to uppercase
lower	Converts text to lowercase
cap	Capitalizes each word
title	Proper title capitalization
exit	Quit program
Example
Enter text: my head is spinning
choose upper, lower, cap, or title: title

My Head Is Spinning
Paul Calculator

A simple terminal calculator supporting basic arithmetic operations.

Features
Addition
Subtraction
Multiplication
Division
Continuous calculation loop
Supported Operators
Operator	Action
+	Add
-	Subtract
*	Multiply
/	Divide
help	Show commands
quit	Exit calculator
Example
Enter the first number: 5
Enter the Second Number: 3
Enter the Operator: +

Output = 5 + 3 = 8
Number Base Converter

Convert numbers between Binary, Hexadecimal, and Decimal.

Supported Conversions
Input	Result
bin	Binary → Decimal
hex	Hex → Decimal
dec	Decimal → Binary & Hex
Format
<number> <base>
Example
Enter number to convert:
1010 bin

✦ Decimal: 10
Enter number to convert:
15 dec

✦ Binary: 1111
✦ Hex:    F
Commands
Command	Description
quit	Exit converter
Concepts Practiced
Go functions
Loops (for)
Switch statements
Maps
String manipulation
Rune handling
Error handling
CLI interaction
User input processing
