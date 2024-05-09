package models

type TrackComment struct {
	ID      int    `json:"id" db:"id"`
	TrackID string `json:"track_id" db:"track_id"`
	Comment string `form:"comment" db:"comment"`
}

type AlbumComment struct {
	ID      int    `json:"id" db:"id"`
	AlbumID string `json:"album_id" db:"album_id"`
	Comment string `form:"comment" db:"comment"`
}

type AlbumCommentResp struct {
	Username string
	Comment  string
}

type TrackCommentResp struct {
	Username string
	Comment  string
}
