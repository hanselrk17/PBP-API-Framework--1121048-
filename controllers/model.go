package controllers

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country"`
	UserType int    `json:"user_type"`
}

type Song struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Duration float64 `json:"duration"`
	Singer   string  `json:"singer"`
}

type Playlists struct {
	ID          int    `json:"id"`
	Name        string `json:"nama"`
	DateCreated string `json:"date_created"`
	State       bool   `json:"state"`
	UserId      int    `json:"user_id"`
}

type DetailPlaylistSong struct {
	PlaylistId int `json:"playlist_id"`
	SongId     int `json:"song_id"`
	TimePlayed int `json:"time_played"`
}

type PopularSong struct {
	DataSong   Song `json:"Song"`
	TimePlayed int  `json:"time_played"`
}

type RecommendedSong struct {
	DataSong Song `json:"Song"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseDoang struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
