package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	nim     int
	nama    string
	jurusan string
}

var studentMap = make(map[int]Student)

func main() {
	for {
		fmt.Println("\nSelamat datang, apa yang akan Anda lakukan hari ini?\n(1) Daftar Mahasiswa\n(2) Hapus Data Mahasiswa\n(3) Tampilkan Semua Data\n(0) Keluar Program")
		input := numberInputProcess("Fitur (ketikkan dengan angka) : ")

		switch input {
		case 1:
			tambahMahasiswa()
		case 2:
			hapusMahasiswa()
		case 3:
			tampilkanData()
		case 0:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Maaf, fitur belum tersedia, silakan coba yang lain.")
		}
	}
}

func numberInputProcess(prompt string) int {
	numberInput := readInput(prompt)

	validNumber, err := strconv.Atoi(numberInput)

	if err != nil {
		fmt.Println("Isilah dengan angka yang sesuai, silakan coba lagi.")
		return numberInputProcess(prompt)
	}

	return validNumber
}

func readInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}

func tambahMahasiswa() {
	fmt.Println("\nPendaftaran Mahasiswa\nSilakan isi data mahasiswa di bawah ini,")

	nim := numberInputProcess("NIM : ")
	nama := readInput("Nama : ")
	jurusan := readInput("Jurusan : ")

	newStudent := Student{nim: nim, nama: nama, jurusan: jurusan}

	studentMap[nim] = newStudent

	fmt.Println("Mahasiswa berhasil ditambahkan!!")
}

func hapusMahasiswa() {
	nim := numberInputProcess("\nHapus Data Mahasiswa\nNIM Mahasiswa : ")

	if validateNimInMap(nim) {
		delete(studentMap, nim)
		fmt.Println("Data mahasiswa berhasil dihapus.")
	} else {
		fmt.Println("Tidak terdapat data mahasiswa yang tersimpan.")
	}
}

func validateNimInMap(nim int) bool {
	key := nim
	_, ok := studentMap[key]
	return ok
}

func tampilkanData() {
	studentCounter := 1
	for nim, student := range studentMap {
		fmt.Printf("\nMahasiswa %d\nNIM : %d\nNama: %s\nJurusan : %s\n", studentCounter, nim, student.nama, student.jurusan)
		studentCounter++
	}
}
