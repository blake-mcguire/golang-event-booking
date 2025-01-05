package models

import (
	"github.com/blake-mcguire/golang-event-booking/main/db"
	"github.com/blake-mcguire/golang-event-booking/main/utils"
)

type User struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
    stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err 

	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId() 


	u.ID = userId
	return err
} 