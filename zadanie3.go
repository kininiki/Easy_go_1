/* На вход целое число
Если двузначное, на выходе - сумма цифр числа.
Если трехзначное, на выходе - произведение положительных цифр.
Для других чисел вывести, что опереция не может быть выполнена.
*/
package main

import "fmt"

func main() {
	var num, z1, z2, z3, proiz int
	fmt.Printf("Введите целое число: ")
	fmt.Scanf("%d", &num)

	if (num/100 == 0) && (num/10 > 0) {
		fmt.Printf("Вы ввели двузначное число. Сумма цифр числа: %d\n", num/10+num%10)

	} else if num/100 > 0 && num/1000 == 0 {
		z1 = num / 100
		z2 = (num - num/100*100) / 10
		z3 = num - num/10*10

		if z2 == 0 && z3 == 0 {
			proiz = z1
		} else if z2 == 0 && z3 > 0 {
			proiz = z1 * z3
		} else if z3 == 0 && z2 > 0 {
			proiz = z1 * z2
		} else {
			proiz = z1 * z2 * z3
		}
		fmt.Printf("Вы ввели трёхзначное число. Произведение положительных цифр числа: %d\n", proiz)

	} else {
		fmt.Printf("Операция не может быть выполнена.")
	}

}
