package main

import "fmt"

func main() {
	var num1, num2 float32

	fmt.Print("Number 1 :")
	fmt.Scanln(&num1)
	fmt.Print("Number 2 :")
	fmt.Scanln(&num2)

	penjumlahan := num1 + num2
	pengurangan := num1 + num2
	perkalian := num1 + num2
	pembagian := num1 + num2

	fmt.Printf("\nHasil Penjumlahan: %.2f + %.2f = %.2f\n", num1, num2, penjumlahan)
	fmt.Printf("\nHasil Pengurangan: %.2f - %.2f = %.2f\n", num1, num2, pengurangan)
	fmt.Printf("\nHasil Perkalian: %.2f * %.2f = %.2f\n", num1, num2, perkalian)
	if num2 != 0 {
		fmt.Printf("\nHasil Pembagian: %.2f / %.2f = %.2f\n", num1, num2, pembagian)
	} else {
		fmt.Printf("Pembagian tidak bisa dilakukan dengan angka 0")
	}
}
