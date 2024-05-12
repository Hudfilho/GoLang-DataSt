package main

import "fmt"

func main() {
	var C, P, F int
	fmt.Scanln(&C, &P, &F)
	if C*F <= P {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
