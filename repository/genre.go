package repository

import (
	"database/sql"
	"gotify-project/structs"
)

func GetAllGenre(DB *sql.DB) (results []structs.Genre, err error) {
	sql := "SELECT * FROM genres"

	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var genre = structs.Genre{}

		err = rows.Scan(&genre.ID, &genre.NamaGenre)
		if err != nil {
			panic(err)
		}
		results = append(results, genre)
	}
	return
}

func InsertGenre(DB *sql.DB, genre structs.Genre) (err error) {
	sql := "INSERT INTO genres (nama_genre) VALUES ($1)"

	errs := DB.QueryRow(sql, genre.NamaGenre)

	return errs.Err()
}

func UpdateGenre(DB *sql.DB, genre structs.Genre) (err error) {
	sql := "UPDATE genres SET nama_genre = $1 WHERE id = $2"

	errs := DB.QueryRow(sql, genre.NamaGenre, genre.ID)

	return errs.Err()
}

func DeleteGenre(DB *sql.DB, genre structs.Genre) (err error) {
	sql := "DELETE FROM genres WHERE id = $1"

	errs := DB.QueryRow(sql, genre.ID)

	return errs.Err()

}
