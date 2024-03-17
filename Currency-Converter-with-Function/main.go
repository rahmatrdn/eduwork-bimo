package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var balanceInput string
var fromCurrencyInput string
var toCurrencyInput string
var percentageInput string
var balance float64
var percentage float64
var convertValue float64

func main() {
	fmt.Println("\nSelamat datang pada program Currency Converter with Function")
	fmt.Println("Berikut merupakan mata uang yang kami sediakan, \n(1)USD\n(2)EUR\n(3)GBP\n(4)JPY\n(5)IDR\n(6)Keluar Program")
	input()
	output(convertValue, toCurrencyInput)
}

func input() {
	fromCurrencyProcess()
	toCurrencyProcess()
	balance = balanceInputProcess()
	percentage = percentageInputProcess()
	convertValue = currencyConverter(percentage, balance, toCurrencyInput, fromCurrencyInput)
}

func fromCurrencyProcess() {
	fmt.Print("Silakan pilih nomor atau ketikkan langsung (sesuai pilihan di atas) mata uang asal : ")

	_, err := fmt.Scan(&fromCurrencyInput)

	if err != nil {
		fmt.Println("Mohon isi terlebih dahulu mata uang yang ingin Anda konversi\n")
		fromCurrencyProcess()
	} else if validateChoice(fromCurrencyInput) == "failed" {
		fromCurrencyProcess()
	} else if validateChoice(fromCurrencyInput) == "quit" {
		fmt.Println("Terima kasih telah menggunakan program ini.\n")
		os.Exit(0)
	} else {
		fromCurrencyInput = validateChoice(fromCurrencyInput)
	}
}

func toCurrencyProcess() {
	fmt.Print("Silakan pilih nomor atau ketikkan langsung (sesuai pilihan di atas) mata uang tujuan : ")

	_, err := fmt.Scan(&toCurrencyInput)

	if err != nil {
		fmt.Println("Mohon isi terlebih dahulu mata uang tujuan konversi Anda\n")
		toCurrencyProcess()
	} else if validateChoice(toCurrencyInput) == "failed" {
		toCurrencyProcess()
	} else if validateChoice(toCurrencyInput) == "quit" {
		fmt.Println("Terima kasih telah menggunakan program ini.\n")
		os.Exit(0)
	} else {
		toCurrencyInput = validateChoice(toCurrencyInput)
	}
}

func validateChoice(input string) string {
	lowerInput := strings.ToLower(input)

	if lowerInput == "1" || lowerInput == "usd" {
		lowerInput = "usd"
		return lowerInput
	} else if lowerInput == "2" || lowerInput == "eur" {
		lowerInput = "eur"
		return lowerInput
	} else if lowerInput == "3" || lowerInput == "gbp" {
		lowerInput = "gbp"
		return lowerInput
	} else if lowerInput == "4" || lowerInput == "jpy" {
		lowerInput = "jpy"
		return lowerInput
	} else if lowerInput == "5" || lowerInput == "idr" {
		lowerInput = "idr"
		return lowerInput
	} else if lowerInput == "6" || lowerInput == "keluar program" {
		return "quit"
	} else {
		fmt.Println("Maaf, mata uang tersebut belum dapat kami konversi, silakan coba yang tersedia.\n")
	}

	return "failed"
}

func balanceInputProcess() float64 {
	fmt.Print("Silakan masukkan saldo Anda : ")

	fmt.Scan(&balanceInput)
	balance, err := strconv.ParseFloat(balanceInput, 64)

	if err != nil {
		fmt.Println("Mohon isi saldo Anda dengan angka")
		return balanceInputProcess()
	}

	return balance
}

func percentageInputProcess() float64 {
	fmt.Print("Silakan masukkan persentase perubahan nilai tukar : ")

	fmt.Scan(&percentageInput)
	percentage, err := strconv.ParseFloat(percentageInput, 64)

	if err != nil {
		fmt.Println("Mohon isi persentase perubahan dengan angka saja")
		return percentageInputProcess()
	}

	return percentage
}

// Penyesuaian nilai kurs mata uang dengan persentase perubahan
func currencyConverter(percentage float64, balance float64, toCurrency string, fromCurrency string) float64 {
	switch fromCurrency {
	case "idr":
		oldValue := getIdrValue(toCurrency)
		return balance * (oldValue + (oldValue * percentage / 100))
	case "usd":
		oldValue := getUsdValue(toCurrency)
		return balance * (oldValue + (oldValue * percentage / 100))
	case "eur":
		oldValue := getEurValue(toCurrency)
		return balance * (oldValue + (oldValue * percentage / 100))
	case "gbp":
		oldValue := getGbpValue(toCurrency)
		return balance * (oldValue + (oldValue * percentage / 100))
	case "jpy":
		oldValue := getJpyValue(toCurrency)
		return balance * (oldValue + (oldValue * percentage / 100))
	}

	return balance
}

func getIdrValue(toCurrency string) float64 {
	switch toCurrency {
	case "usd":
		return 0.000064136435

	case "eur":
		return 0.000058865304

	case "gbp":
		return 0.000050347904

	case "jpy":
		return 0.0095582167
	}

	return 1
}

func getUsdValue(toCurrency string) float64 {
	switch toCurrency {
	case "idr":
		return 15591.762

	case "eur":
		return 0.91781379

	case "gbp":
		return 0.78501252

	case "jpy":
		return 149.02944
	}

	return 1
}

func getEurValue(toCurrency string) float64 {
	switch toCurrency {
	case "usd":
		return 1.0895456

	case "idr":
		return 16987.936

	case "gbp":
		return 0.85530696

	case "jpy":
		return 162.37437
	}

	return 1
}

func getGbpValue(toCurrency string) float64 {
	switch toCurrency {
	case "usd":
		return 1.273865

	case "eur":
		return 1.1691709

	case "idr":
		return 19861.80

	case "jpy":
		return 189.84339
	}

	return 1
}

func getJpyValue(toCurrency string) float64 {
	switch toCurrency {
	case "usd":
		return 0.0067100838

	case "eur":
		return 0.0061586074

	case "gbp":
		return 0.0052674998

	case "idr":
		return 104.62203
	}

	return 1
}

func output(convertValue float64, toCurrency string) {
	fmt.Printf("Berikut merupakan hasil konversi Anda : %v %v\n\n", convertValue, strings.ToUpper(toCurrency))
	main()
}
