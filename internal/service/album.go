package service

import (
	"github.com/markovk1n/spoty/internal/repository"
	"github.com/markovk1n/spoty/models"
)

type AlbumService struct {
	repository repository.Album
}

func NewAlbumService(repo repository.Album) *AlbumService {
	return &AlbumService{repository: repo}
}

func (t *AlbumService) CreateCommentForAlbum(userID int, input models.AlbumComment) error {
	if err := t.repository.CreateCommentForAlbum(userID, input); err != nil {
		return err
	}
	return nil
}

func (t *AlbumService) GetAlbumsComm(albumId string) ([]models.AlbumCommentResp, error) {
	return t.repository.GetAlbumsComm(albumId)
}
