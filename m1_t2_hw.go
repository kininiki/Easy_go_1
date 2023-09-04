/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, res int
	var s string = "0"
	var str string

	fmt.Printf("Введите трёхзначное число:")
	fmt.Scanf("%d\n", &num)

	if num > 999 || num < 100 {
		fmt.Printf("Вы ввели некорректное значение.")
	} else {

		if num%10 == 0 && (num%100-num%10) == 0 {
			res = num/100 + num%10*100 + (num%100 - num%10)
			str = strconv.Itoa(res)

			fmt.Printf("Результат: %s\n", s+s+str)

		} else if num%10 == 0 {
			res = num/100 + num%10*100 + (num%100 - num%10)
			str = strconv.Itoa(res)

			fmt.Printf("Результат: %s\n", s+str)

		} else {
			res = num/100 + num%10*100 + (num%100 - num%10)
			fmt.Printf("Результат: %d\n", res)
		}
	}

}
