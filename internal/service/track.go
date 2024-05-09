package service

import (
	"github.com/markovk1n/spoty/internal/repository"
	"github.com/markovk1n/spoty/models"
)

type TrackService struct {
	repository repository.Track
}

func NewTrackService(repo repository.Track) *TrackService {
	return &TrackService{repository: repo}
}

func (t *TrackService) CreateCommentForTrack(userID int, input models.TrackComment) error {
	if err := t.repository.CreateCommentForTrack(userID, input); err != nil {
		return err
	}
	return nil
}

func (t *TrackService) GetTrackComm(trackId string) ([]models.TrackCommentResp, error) {
	return t.repository.GetTrackComm(trackId)
}
