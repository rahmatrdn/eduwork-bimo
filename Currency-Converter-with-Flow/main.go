package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	gbpToRupiah = 19939
	jpyToRupiah = 105
	eurToRupiah = 17050
	usdToRupiah = 15651
	maxExchange = 1000
)

var balanceInput string

func main() {
	currencyInput()
}

func currencyInput() {
	var input string

	fmt.Println("Selamat datang, silakan pilih nomor atau ketikkan langsung mata uang yang ingin dikonversi\n(1)USD\n(2)EUR\n(3)GBP\n(4)JPY\n(5)Keluar Program")

	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Mohon isi terlebih dahulu mata uang yang ingin Anda konversi\n")
		currencyInput()
	} else {
		validateCurrencyInput(input)
	}
}

func validateCurrencyInput(input string) {
	var balance float64
	lowerInput := strings.ToLower(input)

	if lowerInput == "1" || lowerInput == "usd" {
		fmt.Print("Silakan masukkan saldo Dollar Anda : ")
		balance = validateExchangeProcess()
		exchangeUsdToRupiah(balance)
	} else if lowerInput == "2" || lowerInput == "eur" {
		fmt.Print("Silakan masukkan saldo Euro Anda : ")
		balance = validateExchangeProcess()
		exchangeEurToRupiah(balance)
	} else if lowerInput == "3" || lowerInput == "gbp" {
		fmt.Print("Silakan masukkan saldo Poundsterling Anda : ")
		balance = validateExchangeProcess()
		exchangeGbpToRupiah(balance)
	} else if lowerInput == "4" || lowerInput == "jpy" {
		fmt.Print("Silakan masukkan saldo Yen Anda : ")
		balance = validateExchangeProcess()
		exchangeJpyToRupiah(balance)
	} else if lowerInput == "5" || lowerInput == "keluar program" {
		fmt.Println("Terima kasih telah menggunakan program ini.\n")
	} else {
		fmt.Println("Maaf, mata uang tersebut belum dapat kami konversi, silakan coba yang tersedia.\n")
		currencyInput()
	}
}

func validateExchangeProcess() float64 {
	fmt.Scan(&balanceInput)
	balance, err := strconv.ParseFloat(balanceInput, 64)

	if err != nil {
		fmt.Print("Mohon isi saldo Anda dengan angka : ")
		return validateExchangeProcess()
	}

	return balance
}

func exchangeUsdToRupiah(balance float64) {
	rupiah := balance * usdToRupiah
	fmt.Println("Kurs USD sampai hari ini mencapai : ", usdToRupiah)
	output(rupiah)
}

func exchangeGbpToRupiah(balance float64) {
	rupiah := balance * gbpToRupiah
	fmt.Println("Kurs GBP sampai hari ini mencapai : ", gbpToRupiah)
	output(rupiah)
}

func exchangeJpyToRupiah(balance float64) {
	rupiah := balance * jpyToRupiah
	fmt.Println("Kurs JPY sampai hari ini mencapai : ", jpyToRupiah)
	output(rupiah)
}

func exchangeEurToRupiah(balance float64) {
	rupiah := balance * eurToRupiah
	fmt.Println("Kurs EUR sampai hari ini mencapai : ", eurToRupiah)
	output(rupiah)
}

func output(rupiah float64) {
	if rupiah > 1000 {
		fmt.Printf("Hasil konversi cukup besar dengan jumlah : %v IDR\n\n", rupiah)
	} else {
		fmt.Printf("Hasil konversi : %v IDR\n\n", rupiah)
	}
	currencyInput()
}
