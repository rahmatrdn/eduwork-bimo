package entities

type User struct {
	ID           int    `json:"id"`
	NamaLengkap  string `json:"nama_lengkap"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	NomorTelepon string `json:"nomor_telepon"`
}
