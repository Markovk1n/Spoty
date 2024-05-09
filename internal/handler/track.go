package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/internal/service/spotify"
	"github.com/markovk1n/spoty/models"
)

func (h *Handler) getTrack(c *gin.Context) {
	trackId := c.Param("id")
	token1, err := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	client := spotify.NewClient(token1)
	trackResult, _ := client.Track.Get(trackId)

	comm, err := h.services.Track.GetTrackComm(trackId)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	res := TrackResponse{
		Comments:  comm,
		TrackInfo: trackResult,
	}

	c.HTML(http.StatusOK, "track.html", res)

}

type TrackResponse struct {
	Comments  []models.TrackCommentResp
	TrackInfo *spotify.Track
}

func (h *Handler) CreateCommentForTrack(c *gin.Context) {
	userID := GetUserID(c)
	var input models.TrackComment
	input.TrackID = c.Param("id")
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.CreateCommentForTrack(userID, input); err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, "created")
}
func GetUserID(c *gin.Context) int {
	if obj, ok := c.Get("user_id"); ok {
		if userID, ok := obj.(int); ok {
			return userID
		}
	}
	return 0
}
