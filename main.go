package main

import "fmt"

func fizzbuzz(i int) {
	for num := 1; num <= i; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if num%3 == 0 {
			fmt.Println("fizz")
		} else if num%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(num)
		}
	}
}

func main() {
	var num int
	fmt.Print("Введите число (до скольки считать): ")
	fmt.Scan(&num)
	fizzbuzz(num)
}