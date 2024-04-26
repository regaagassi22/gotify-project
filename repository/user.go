package repository

import (
	"database/sql"
	"fmt"
	"gotify-project/structs"
)

func GetAllUser(DB *sql.DB) (results []structs.User, err error) {
	sql := "SELECT * FROM users"

	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RoleId)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}
	return
}

func InsertUser(DB *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (username, email, password, role_id) VALUES ($1, $2 ,$3, $4)"

	errs := DB.QueryRow(sql, user.Username, user.Email, user.Password, user.RoleId)

	return errs.Err()
}

func UpdateUser(DB *sql.DB, user structs.User) (err error) {
	sql := "UPDATE users SET username = $1, email = $2, password = $3, role_id = $4 WHERE id = $5"

	errs := DB.QueryRow(sql, user.Username, user.Email, user.Password, user.RoleId, user.ID)

	return errs.Err()
}

func DeleteUser(DB *sql.DB, user structs.User) (err error) {
	sql := "DELETE FROM users WHERE id = $1"

	errs := DB.QueryRow(sql, user.ID)

	return errs.Err()

}

func Login(DB *sql.DB, username string, password string) (results []structs.User, err error) {
	sql := "SELECT * FROM users WHERE username = $1 and password = $2"
	fmt.Println(username)
	fmt.Println(password)

	rows, err := DB.Query(sql, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RoleId)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}
	return
}
