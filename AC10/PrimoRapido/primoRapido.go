package main

import "fmt"

func achaPrimo(x int) bool {
	if x <= 1 {
		return false
	}
	if x <= 5 && x != 4 {
		return true
	}
	if x%2 == 0 || x%3 == 0 || x%5 == 0 {
		return false
	}
	for i := 5; i*i <= x; i += 6 {
		if x%i == 0 || x%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	var N, X int
	fmt.Scanln(&N)
	for N != 0 {
		fmt.Scanln(&X)
		if achaPrimo(X) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
		N--
	}
}
