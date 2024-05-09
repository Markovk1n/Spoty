package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/markovk1n/spoty/models"
)

const (
	usersTable = "users"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(username, password string) (models.User, error)
}

type Repository struct {
	Authorization
	Track
	Album
}
type Track interface {
	CreateCommentForTrack(userID int, input models.TrackComment) error
	GetTrackComm(trackId string) ([]models.TrackCommentResp, error)
}
type Album interface {
	CreateCommentForAlbum(userID int, input models.AlbumComment) error
	GetAlbumsComm(albumId string) ([]models.AlbumCommentResp, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: newAuthPostgres(db),
		Track:         NewTrackRepository(db),
		Album:         NewAlbumRepository(db),
	}
}
