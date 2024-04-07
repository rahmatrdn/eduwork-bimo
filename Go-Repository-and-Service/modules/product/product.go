package product

import "time"

type Product struct {
	ID                int       `json:"id`
	Nama_Produk       string    `json:"nama_produk`
	Deskripsi         string    `json:"deskripsi`
	Harga             float64   `json:"harga`
	Stok              int       `json:"stok`
	Tanggal_Pembuatan time.Time `json:"tanggal_pembuatan`
}
