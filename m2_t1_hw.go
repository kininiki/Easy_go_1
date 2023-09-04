/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/
package main

import (
	"fmt"
)

func main() {
	var num, n1, n2, n3, k1, k2, k3, res int

	fmt.Printf("Введите трёхзначное число:")
	fmt.Scanf("%d\n", &num)

	if num >= 100 && num <= 999 {
		n1 = num / 100
		n2 = (num - num/100*100 - num%10) / 10
		n3 = num % 10

		if n1 > n2 && n1 > n3 {
			k1 = n1
			if n2 > n3 {
				k2 = n2 * 10
				k3 = n3 * 100
			} else {
				k2 = n3 * 10
				k3 = n2 * 100
			}
		} else if n2 > n1 && n2 > n3 {
			k1 = n2
			if n1 > n3 {
				k2 = n1 * 10
				k3 = n3 * 100
			} else {
				k2 = n3 * 10
				k3 = n1 * 100
			}
		} else {
			k1 = n3
			if n2 > n1 {
				k2 = n2 * 10
				k3 = n1 * 100
			} else {
				k2 = n1 * 10
				k3 = n2 * 100
			}
		}
		res = k1 + k2 + k3
		fmt.Printf("Числа в порядке возрастания: %d\n", res)

	} else {
		fmt.Printf("Вы ввели некорректное значение.")
	}

}
