package service

import (
	"github.com/markovk1n/spoty/internal/repository"
	"github.com/markovk1n/spoty/models"
)

type Authorization interface {
	CreateUser(user models.User) error
	GenerateToken(username, password string) (string, error)
}

type Track interface {
	CreateCommentForTrack(userID int, input models.TrackComment) error
	GetTrackComm(trackId string) ([]models.TrackCommentResp, error)
}
type Album interface {
	CreateCommentForAlbum(userID int, input models.AlbumComment) error
	GetAlbumsComm(albumId string) ([]models.AlbumCommentResp, error)
}

type Service struct {
	Authorization
	Track
	Album
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Track:         NewTrackService(repos.Track),
		Album:         NewAlbumService(repos.Album),
	}
}
