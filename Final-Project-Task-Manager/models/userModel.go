package models

import (
	"database/sql"
	"eduwork-bimo/Final-Project-Task-Manager/entities"
	_ "github.com/go-sql-driver/mysql"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{db: db}
}

func (u UserModel) GetByUsername(username string) (*entities.User, error) {
	row := u.db.QueryRow("SELECT id, nama_lengkap, email, username, password FROM users WHERE username = ?", username)
	var user entities.User
	err := row.Scan(&user.ID, &user.NamaLengkap, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m UserModel) Create(user *entities.User) error {
	_, err := m.db.Exec("INSERT INTO users (nama_lengkap, email, username, password) VALUES (?, ?, ?, ?)", user.NamaLengkap, user.Email, user.Username, user.Password)
	return err
}
