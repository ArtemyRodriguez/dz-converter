package main

import "fmt"

func main() {

	fmt.Println(" >>> Конвертер валют! <<<")
	var mainResult string = getResult()
	fmt.Print(mainResult)

}
func getUserChoise() int {
	var userChoise int
	//выбираем валюту
	fmt.Print("Выберите валюту: 1 - рубли, 2 - доллары. Ваш выбор: ")
	fmt.Scan(&userChoise)
	return userChoise
}
func getResult() string {
	var usdToRub float64
	usdToRub = 78.3
	var moneyInWallet float64
	var mainResult string
	var userChoise = getUserChoise()
	//Вводим переменную результата
	var result float64
	//конвертируем валюту через оператор свитч
	switch {
	//конвертируем рубли
	case userChoise == 1:
		fmt.Print("Введите, пожалуйста, кол-во рублей: ")
		fmt.Scan(&moneyInWallet)
		result = moneyInWallet / usdToRub
		fmt.Printf("Если перевести %.f руб. в доллары, то вы получите %.f $.", moneyInWallet, result)
		mainResult := fmt.Sprintf("Если перевести %.f руб. в доллары, то вы получите %.f $.", moneyInWallet, result)
		return mainResult

		//конвертируем доллары
	case userChoise == 2:
		fmt.Print("Введите, пожалуйста, кол-во долларов: ")
		fmt.Scan(&moneyInWallet)
		result = moneyInWallet * usdToRub
		mainResult := fmt.Sprintf("Если перевести %.f долл. в рубли, то вы получите %.f руб.", moneyInWallet, result)
		return mainResult
	}
	return mainResult
}
