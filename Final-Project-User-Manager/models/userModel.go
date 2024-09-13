package models

import (
	"database/sql"

	"eduwork-bimo/Final-Project-User-Manager/entities"
	helpers "eduwork-bimo/Final-Project-User-Manager/helpers"

	_ "github.com/go-sql-driver/mysql"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db: db}
}

func (u UserModel) GetByUsername(username string) (*entities.User, error) {
	row := u.db.QueryRow("SELECT user_id, nama_lengkap, username, email, password, nomor_telepon FROM users WHERE username = ?", username)
	var user entities.User
	err := row.Scan(&user.ID, &user.NamaLengkap, &user.Username, &user.Email, &user.Password, &user.NomorTelepon)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m UserModel) Create(user *entities.User) error {
	_, err := m.db.Exec("INSERT INTO users (nama_lengkap, username, email, password, nomor_telepon) VALUES (?, ?, ?, ?, ?)", user.NamaLengkap, user.Username, user.Email, user.Password, user.NomorTelepon)
	return err
}

func (m UserModel) Update(user *entities.User) error {
	_, err := m.db.Exec("UPDATE users SET nama_lengkap = ?, username = ?, email = ?, nomor_telepon = ? WHERE user_id = ?", user.NamaLengkap, user.Username, user.Email, user.NomorTelepon, user.ID)
	return err
}

func (m UserModel) ResetPassword(user *entities.User) error {
	hashedPass, err := helpers.EncryptPass(user.Password)
	_, err = m.db.Exec("UPDATE users SET password = ? WHERE nomor_telepon = ?", hashedPass, user.NomorTelepon)
	return err
}

func (m UserModel) DeleteUser(userID int) error {
	_, err := m.db.Exec("DELETE FROM users WHERE user_id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}
