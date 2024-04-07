package user

import "time"

type User struct {
	ID                int       `json:"id`
	Nama_Pengguna     string    `json:"nama_pengguna`
	Email_Pengguna    string    `json:"email_pengguna`
	Password          string    `json:"password`
	Tanggal_Pembuatan time.Time `json:"tanggal_pembuatan`
}
