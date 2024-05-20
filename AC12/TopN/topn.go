package main

import (
	"fmt"
)

func main() {
	var K int
	for {
		_, err1 := fmt.Scan(&K)
		if err1 != nil {
			return
		}

		switch {
		case K <= 1:
			fmt.Println("Top 1")
		case K <= 3:
			fmt.Println("Top 3")
		case K <= 5:
			fmt.Println("Top 5")
		case K <= 10:
			fmt.Println("Top 10")
		case K <= 25:
			fmt.Println("Top 25")
		case K <= 50:
			fmt.Println("Top 50")
		case K <= 100:
			fmt.Println("Top 100")
		}
	}
}
