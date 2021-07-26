package main

import "fmt"

func main() {

	var score = 87

	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("B")
	case 8:
		fmt.Println("C")
	case 7:
		fmt.Println("D")
	case 6:
		fmt.Println("E")
	default:
		fmt.Println("F")
	}
}
