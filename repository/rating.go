package repository

import (
	"database/sql"
	"gotify-project/structs"
)

func GetAllRating(DB *sql.DB) (results []structs.Rating, err error) {
	sql := "SELECT * FROM ratings"

	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = structs.Rating{}

		err = rows.Scan(&user.ID, &user.Rating, &user.UserId, &user.MusicId)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}
	return
}

func InsertRating(DB *sql.DB, user structs.Rating) (err error) {
	sql := "INSERT INTO ratings (rating, user_id, music_id) VALUES ($1, $2 ,$3)"

	errs := DB.QueryRow(sql, user.Rating, user.UserId, user.MusicId)

	return errs.Err()
}

func UpdateRating(DB *sql.DB, user structs.Rating) (err error) {
	sql := "UPDATE ratings SET rating = $1, user_id = $2, music_id = $3 WHERE id = $4"

	errs := DB.QueryRow(sql, user.Rating, user.UserId, user.MusicId, user.ID)

	return errs.Err()
}

func DeleteRating(DB *sql.DB, user structs.Rating) (err error) {
	sql := "DELETE FROM ratings WHERE id = $1"

	errs := DB.QueryRow(sql, user.ID)

	return errs.Err()

}
