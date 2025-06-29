package models

import (
	"github.com/idstev/marketplace/config"
	"log"
)

func CreateUser(user User) error {
	query := `INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4)`
	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password, user.Role)
	return err
}


func GetUserByEmailAndPassword(email, password string) (*User, error) {
	var user User
	query := `SELECT id, name, email, role FROM users WHERE email=$1 AND password=$2`
	row := config.DB.QueryRow(query, email, password)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
	if err != nil {
		log.Println("Error al crear usuario:", err)
	}
	return &user, nil
}
