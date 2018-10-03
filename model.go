package main

import (
	"database/sql"
	"errors"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) getUser(db *sql.DB) error {
	return db.QueryRow("Select name,age from users where id = $1", u.Id).Scan(&u.Name, &u.Age)
}

func (u *User) updateuser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) deleteUser(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *User) createUser(db *sql.DB) error {
	err := db.QueryRow("Insert into users(name,age) values($1,$2) returning id", u.Name, u.Age).Scan(&u.Id)
	if err != nil {
		return err
	}
	return nil
}

func getUsers(db *sql.DB, start int64, count int64) ([]User, error) {
	return nil, errors.New("Not implemented")
}
