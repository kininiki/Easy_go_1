/*
Посчитать размер вклада при определенном количестве лет и проценте.
Использовать ежегодную капитализацию.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var vklad, proz, razmer float64
	var years int
	fmt.Printf("Введите сумму вклада (руб): ")
	fmt.Scanf("%f", &vklad)

	fmt.Printf("Введите количество лет: ")
	fmt.Scanf("%d", &years)

	fmt.Printf("Введите процент: ")
	fmt.Scanf("%f", &proz)

	razmer = vklad * math.Pow((1+proz/100), float64(years))
	fmt.Printf("Через %d лет размер вклада будет: %.2f\n", years, razmer)

}
