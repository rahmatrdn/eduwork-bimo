package main

import "fmt"

const (
	dollarToRupiah    = 15000
	maxDollarExchange = 100
)

func main() {
	saldoDollar := 1000.0

	fmt.Println("Saldo dollar awal : ", saldoDollar)

	saldoRupiah := exchangeDollarToRupiah(saldoDollar)

	fmt.Println("Saldo rupiah : ", saldoRupiah)
}

func exchangeDollarToRupiah(dollar float64) float64 {
	if dollar > maxDollarExchange {
		fmt.Println("Maaf tidak bisa menukar lebih dari 100")
		return dollar
	}

	rupiah := dollar * dollarToRupiah
	return rupiah
}
