package main

import "fmt"

func printYellow(str string) {
	fmt.Printf("\033[0m\033[1;33m%s\033[0m", str)
}

func printPink(str string) {
	fmt.Printf("\033[0m\033[1;35m%s\033[0m", str)
}

func printGreen(str string) {
	fmt.Printf("\033[0m\033[1;32m%s\033[0m", str)
}

func printBlue(str string) {
	fmt.Printf("\033[0m\033[1;34m%s\033[0m", str)
}

func printlnYellow(str string) {
	fmt.Printf("\033[0m\033[1;33m%s\033[0m\n", str)
}

func printlnPink(str string) {
	fmt.Printf("\033[0m\033[1;35m%s\033[0m\n", str)
}

func printlnGreen(str string) {
	fmt.Printf("\033[0m\033[1;32m%s\033[0m\n", str)
}

func printlnBlue(str string) {
	fmt.Printf("\033[0m\033[1;34m%s\033[0m\n", str)
}
