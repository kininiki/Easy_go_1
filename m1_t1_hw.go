/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/
package main

import "fmt"

func main() {
	var s, ras, result float64
	const rub float64 = 48

	fmt.Printf("Введите расстояние (50-10000 км):")
	fmt.Scanf("%f\n", &s)
	fmt.Printf("Введите расход в литрах (5-25 литров) на 100 км:")
	fmt.Scanf("%f\n", &ras)

	if s >= 50 && s <= 10000 && ras >= 5 && ras <= 25 {
		result = ras / 100 * s * rub
		fmt.Printf("Стоимость поездки в рублях: %.1f", result)
	} else {
		fmt.Printf("Вы ввели некорректные значения.")
	}

}
