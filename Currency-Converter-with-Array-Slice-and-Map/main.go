package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	currencyList [5]string
	currencyMap  = make(map[string]map[string]float64)

	opening                   = "\nSelamat datang di program Currency Converter\nBerikut merupakan mata uang yang kami sediakan, \n(1)IDR\n(2)USD\n(3)GBP\n(4)JPY\n(5)EUR\n(0)EXIT"
	fromCurrencyInputPrompt   = "Silakan pilih nomor atau ketikkan langsung (sesuai pilihan di atas) mata uang asal : "
	validateFromCurrencyInput = "Mohon isi terlebih dahulu mata uang yang ingin Anda konversi\n"
	toCurrencyInputPrompt     = "Silakan pilih nomor atau ketikkan langsung (sesuai pilihan di atas) mata uang tujuan : "
	validateToCurrencyinput   = "Mohon isi terlebih dahulu mata uang tujuan konversi Anda\n"
	noCurrencyAvailable       = "Maaf, mata uang tersebut belum dapat kami konversi, silakan coba yang tersedia.\n"
	balanceInputPrompt        = "Silakan masukkan saldo Anda : "
	validateBalancePrompt     = "Mohon isi saldo Anda dengan angka"
	exitPrompt                = "Terima kasih telah menggunakan program kami.\n"

	idrConversionValue = []float64{1, 0.000064136435, 0.000050347904, 0.0095582167, 0.000058865304}
	usdConversionValue = []float64{15591.762, 1, 0.78501252, 149.02944, 0.91781379}
	gbpConversionValue = []float64{19861.80, 1.273865, 1, 189.84339, 1.1691709}
	jpyConversionValue = []float64{104.62203, 0.0067100838, 0.0052674998, 1, 0.0061586074}
	eurConversionValue = []float64{16987.936, 1.0895456, 0.85530696, 162.37437, 1}
)

func main() {

	currencyList = [5]string{"IDR", "USD", "GBP", "JPY", "EUR"}

	go updateConversionValues()

	fmt.Println(opening)
	input()
}

func updateConversionValues() {
	for {
		now := time.Now()

		// if already 15.00, updates the data
		if now.Hour() >= 18 && now.Minute() >= 30 {
			idrConversionValue = []float64{1, 0.000066136435, 0.000050247904, 0.0097582167, 0.000055865304}
			usdConversionValue = []float64{16591.762, 1, 0.80501252, 152.02944, 0.92081379}
			gbpConversionValue = []float64{20861.80, 1.293865, 1, 190.84339, 1.1751709}
			jpyConversionValue = []float64{105.62203, 0.0069100838, 0.0055674998, 1, 0.0063086074}
			eurConversionValue = []float64{17287.936, 1.0905456, 0.85630696, 165.37437, 1}

			fmt.Println("\nNilai tukar telah diperbarui pada jam 18.30")
		}

		// sleep for minimize resource consumption
		time.Sleep(time.Minute)
	}
}

func updateCurrencyMap() {
	idr := make(map[string]float64)
	usd := make(map[string]float64)
	gbp := make(map[string]float64)
	jpy := make(map[string]float64)
	eur := make(map[string]float64)

	initializeEachCurrencyMap(idr, idrConversionValue)
	initializeEachCurrencyMap(usd, usdConversionValue)
	initializeEachCurrencyMap(gbp, gbpConversionValue)
	initializeEachCurrencyMap(jpy, jpyConversionValue)
	initializeEachCurrencyMap(eur, eurConversionValue)

	currencyMap = map[string]map[string]float64{
		"IDR": idr, "USD": usd, "GBP": gbp, "JPY": jpy, "EUR": eur,
		"1": idr, "2": usd, "3": gbp, "4": jpy, "5": eur,
	}
}

func initializeEachCurrencyMap(currencyName map[string]float64, currencyConversionValue []float64) {
	for i := 0; i < len(currencyList); i++ {
		currencyName[currencyList[i]] = currencyConversionValue[i]
	}
}

func readInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

func input() {
	updateCurrencyMap()
	fromCurrencyInput := currencyProcess(fromCurrencyInputPrompt, validateFromCurrencyInput)
	toCurrencyInput := currencyProcess(toCurrencyInputPrompt, validateToCurrencyinput)
	balance := numberInputProcess(balanceInputPrompt, validateBalancePrompt)
	convertValue := currencyConverter(balance, toCurrencyInput, fromCurrencyInput)
	output(convertValue, toCurrencyInput)
}

func currencyProcess(prompt1 string, prompt2 string) string {
	currencyInput := readInput(prompt1)

	if validateChoice(currencyInput) == noCurrencyAvailable {
		fmt.Println(noCurrencyAvailable)
		return currencyProcess(prompt1, prompt2)
	} else if validateChoice(currencyInput) == "exit" {
		fmt.Println(exitPrompt)
		os.Exit(0)
	}
	return validateChoice(currencyInput)
}

func validateChoice(input string) string {
	upperInput := strings.ToUpper(input)

	if _, ok := currencyMap[upperInput]; ok {
		numberInput, err := strconv.ParseInt(upperInput, 10, 64)

		if err != nil {
			return upperInput
		}
		return currencyList[numberInput-1]
	} else if upperInput == "0" || upperInput == "EXIT" {
		return "exit"
	}
	return noCurrencyAvailable
}

func numberInputProcess(prompt1 string, prompt2 string) float64 {
	numberInput := readInput(prompt1)
	validNumber, err := strconv.ParseFloat(numberInput, 64)

	if err != nil {
		fmt.Println(prompt2)
		return numberInputProcess(prompt1, prompt2)
	}
	return validNumber
}

func currencyConverter(balance float64, toCurrenyInput string, fromCurrencyInput string) float64 {
	return balance * currencyMap[fromCurrencyInput][toCurrenyInput]
}

func output(convertValue float64, toCurrencyInput string) {
	fmt.Printf("Berikut merupakan hasil konversi Anda : %.2f %v\n\n", convertValue, strings.ToUpper(toCurrencyInput))
	main()
}
