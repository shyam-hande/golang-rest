package models

import (
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `insert into users (email, password) values (?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	_, err = res.LastInsertId()

	return err
}

func (u *User) ValidateCredentails() error {
	query := "select id, password from users where email = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	var retrivedPassword string

	row := stmt.QueryRow(u.Email)

	err = row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return err
	}

	isPasswordValid, err := utils.CheckPassword(retrivedPassword, u.Password)

	if !isPasswordValid {
		return err
	}

	return nil
}
