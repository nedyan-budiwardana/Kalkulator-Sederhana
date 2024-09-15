package main

import (
	"fmt"
)

func main() {

	var bilangan1, bilangan2 float64
	var operasi string

	fmt.Print("Input Operasi : ")
	fmt.Scanln(&operasi)

	fmt.Print("Input Bilangan 1 : ")
	fmt.Scanln(&bilangan1)

	fmt.Print("Input Bilangan 2 : ")
	fmt.Scanln(&bilangan2)

	switch operasi {

	case "+":
		fmt.Printf("%.3f %s %.3f = %.3f", bilangan1, operasi, bilangan2, bilangan1+bilangan2)

	case "-":
		fmt.Printf("%.3f %s %.3f = %.3f", bilangan1, operasi, bilangan2, bilangan1-bilangan2)

	case "*":
		fmt.Printf("%.3f %s %.3f = %.3f", bilangan1, operasi, bilangan2, bilangan1*bilangan2)

	case "/":
		if bilangan2 == 0.0 {
			fmt.Println("Tidak terdefinisi")
		} else {
			fmt.Printf("%.3f %s %.3f = %.3f", bilangan1, operasi, bilangan2, bilangan1/bilangan2)
		}

	default:
		fmt.Println("Kalkulator tidak mengerti input user")
	}
}
