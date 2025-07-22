package main

import "fmt"

func main() {
	fmt.Println(">>> Конвертер валют (USD, EUR, RUB) <<<")

	// Ввод данных
	amount := getInput("Введите сумму: ")
	from := getInput("Выберите исходную валюту (USD/EUR/RUB): ")
	to := getInput("Выберите целевую валюту (USD/EUR/RUB): ")

	// Конвертация (заглушка)
	result := convert(amount, from, to)

	// Вывод результата
	fmt.Printf("Результат: %.2f %s = %.2f %s\n", amount, from, result, to)
}

// Функция для ввода данных
func getInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

// Функция конвертации (заглушка)
func convert(amount float64, from, to string) float64 {
	// Здесь должна быть логика конвертации
	// Пока просто возвращаем ту же сумму
	return amount
}
