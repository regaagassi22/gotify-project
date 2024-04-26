package structs

type Role struct {
	ID       int    `json:"id"`
	NamaRole string `json:"nama_role"`
}

type Genre struct {
	ID        int    `json:"id"`
	NamaGenre string `json:"nama_genre"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
}

type Music struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Duration int    `json:"duration"`
	GenreId  int    `json:"genre_id"`
}

type Rating struct {
	ID      int `json:"id"`
	Rating  int `json:"rating"`
	UserId  int `json:"user_id"`
	MusicId int `json:"music_id"`
}
