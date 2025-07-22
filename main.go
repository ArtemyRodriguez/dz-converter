package main

import "fmt"

func main() {
	fmt.Println(">>> Конвертер валют <<<")
	result := getResult()
	fmt.Print(result)
}
func getUserInput() (int, float64) {
	var currencyOperationCode int
	var currencyLevel float64
	fmt.Print("Выберите тип операции: >>> 1 - перевести Рубли в Доллары <<< / >>> 2 - перевести Доллары в Рубли <<<")
	fmt.Scan(&currencyOperationCode)
	switch {
	case currencyOperationCode == 1:
		fmt.Print("Какое кол-во рублей вы хотите обменять? - ")
		fmt.Scan(&currencyLevel)

	case currencyOperationCode == 2:
		fmt.Print("Какое кол-во долларов вы хотите обменять? - ")
		fmt.Scan(&currencyOperationCode)
	}

	return currencyOperationCode, currencyLevel
}
func getResult() float64 {
	var exchangeRate float64
	var result float64
	exchangeRate = 78.0
	currencyOperationCode, currencyLevel := getUserInput()

	switch {
	case currencyOperationCode == 1:
		result = currencyLevel / exchangeRate

	case currencyOperationCode == 2:
		result = currencyLevel * exchangeRate

	}
	return result
}
