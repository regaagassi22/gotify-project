package repository

import (
	"database/sql"
	"gotify-project/structs"
)

func GetAllMusic(DB *sql.DB) (results []structs.Music, err error) {
	sql := "SELECT * FROM musics"

	rows, err := DB.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var music = structs.Music{}

		err = rows.Scan(&music.ID, &music.Title, &music.Artist, &music.Duration, &music.GenreId)
		if err != nil {
			panic(err)
		}
		results = append(results, music)
	}
	return
}

func InsertMusic(DB *sql.DB, music structs.Music) (err error) {
	sql := "INSERT INTO musics (title, artist, duration, genre_id) VALUES ($1, $2, $3, $4)"

	errs := DB.QueryRow(sql, music.Title, music.Artist, music.Duration, music.GenreId)

	return errs.Err()
}

func UpdateMusic(DB *sql.DB, music structs.Music) (err error) {
	sql := "UPDATE musics SET title = $1, artist = $2, duration = $3, genre_id = $4 WHERE id = $5"

	errs := DB.QueryRow(sql, music.Title, music.Artist, music.Duration, music.GenreId, music.ID)

	return errs.Err()
}

func DeleteMusic(DB *sql.DB, music structs.Music) (err error) {
	sql := "DELETE FROM musics WHERE id = $1"

	errs := DB.QueryRow(sql, music.ID)

	return errs.Err()

}
