package entities

type User struct {
	ID          int    `json:"id"`
	NamaLengkap string `json:"nama_lengkap"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
