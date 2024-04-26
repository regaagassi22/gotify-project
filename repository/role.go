package repository

import (
	"database/sql"
	"gotify-project/structs"
)

func GetAllRole(DB *sql.DB) (results []structs.Role, err error) {
	sql := "SELECT * FROM roles"

	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var role = structs.Role{}

		err = rows.Scan(&role.ID, &role.NamaRole)
		if err != nil {
			panic(err)
		}
		results = append(results, role)
	}
	return
}

func InsertRole(DB *sql.DB, role structs.Role) (err error) {
	sql := "INSERT INTO roles (nama_role) VALUES ($1)"

	errs := DB.QueryRow(sql, role.NamaRole)

	return errs.Err()
}

func UpdateRole(DB *sql.DB, role structs.Role) (err error) {
	sql := "UPDATE roles SET nama_role = $1 WHERE id = $2"

	errs := DB.QueryRow(sql, role.NamaRole, role.ID)

	return errs.Err()
}

func DeleteRole(DB *sql.DB, role structs.Role) (err error) {
	sql := "DELETE FROM roles WHERE id = $1"

	errs := DB.QueryRow(sql, role.ID)

	return errs.Err()

}
