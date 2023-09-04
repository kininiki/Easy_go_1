package main

import "fmt"

func main() {

	var num1, num2, num3 int
	var number int
	//var number int = 189
	//fmt.Scanf("%s\n", &number)
	fmt.Printf("Введите любое трёхзначное число: ")
	fmt.Scanf("%d", &number)

	num1 = number / 100
	num3 = number - (number/10)*10
	num2 = (number - (number/100)*100) / 10

	fmt.Printf("Первое число: %d\n", num1)
	fmt.Printf("Последнее число: %d\n", num3)
	fmt.Printf("Среднее число: %d\n", num2)

}
