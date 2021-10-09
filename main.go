package main

import "fmt"

func main() {

	var weekday int
	var numberOfGuests int
	var checks int

	discount := 0
	allowance := 0

	fmt.Println("Введите день недели:")
	fmt.Scan(&weekday)
	fmt.Println("Введите число гостей:")
	fmt.Scan(&numberOfGuests)
	fmt.Println("Введите сумму чека:")
	fmt.Scan(&checks)

	if weekday == 1 {
		discount = (checks / 100) * 10
		fmt.Println("Скидка по понедельникам:", discount)
	} else if weekday == 5 {
		if checks > 10000 {
			discount = (checks / 100) * 5
			fmt.Println("Скидка по пятницам:", discount)
		}
	}

	if numberOfGuests > 5 {
		allowance = (checks / 100) * 10
		fmt.Println("Надбавка на обслуживание:", allowance)
	}

	fmt.Println("Сумма к оплате:", checks-discount+allowance)
}
