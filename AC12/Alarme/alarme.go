package main

import (
	"fmt"
)

func main() {
	for {
		var H1, M1, H2, M2 int
		fmt.Scan(&H1, &M1, &H2, &M2)

		if H1 == 0 && M1 == 0 && H2 == 0 && M2 == 0 {
			break
		}

		hora_atual := H1*60 + M1
		alarme := H2*60 + M2

		var sono int
		if alarme >= hora_atual {
			sono = alarme - hora_atual
		} else {
			sono = (1440 - hora_atual) + alarme
		}

		fmt.Println(sono)
	}
}
