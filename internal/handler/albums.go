package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/internal/service/spotify"
	"github.com/markovk1n/spoty/models"
)

func (h *Handler) getAllAlbums(c *gin.Context) {
	token1, _ := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)

	client := spotify.NewClient(token1)
	albumsResult, _ := client.Album.List("382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc,2cWBwpqMsDJC1ZUwz813lo,2WT1pbYjLJciAR26yMebkH,3PZmQxxLUZwyyMgXWUpmuw")

	c.HTML(http.StatusOK, "albums.html", albumsResult)

}

func (h *Handler) getAlbumById(c *gin.Context) {
	albumId := c.Param("id")
	token1, err := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	client := spotify.NewClient(token1)
	albumResult, _ := client.Album.Get(albumId)

	comm, err := h.services.Album.GetAlbumsComm(albumId)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}

	res := AlbumResponse{
		Comments:  comm,
		AlbumInfo: albumResult,
	}

	// resultTarcks, _ := client.Album.GetTracks("382ObEPsp2rxGrnsizN5TX", 10, 0)
	// Ress := AlbumResponse{

	// }
	fmt.Println(comm, albumId)
	c.HTML(http.StatusOK, "album.html", res)

}

type AlbumResponse struct {
	Comments  []models.AlbumCommentResp
	AlbumInfo *spotify.Album
}

func (h *Handler) CreateCommentForAlbum(c *gin.Context) {
	userID := GetUserID(c)

	var input models.AlbumComment
	input.AlbumID = c.Param("id")

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.CreateCommentForAlbum(userID, input); err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, "created")
}
