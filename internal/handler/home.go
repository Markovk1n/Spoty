package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/internal/service/spotify"
)

func (h *Handler) home(c *gin.Context) {

	token1, _ := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)

	client := spotify.NewClient(token1)
	albumsResult, _ := client.Album.List("382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc,2cWBwpqMsDJC1ZUwz813lo,2WT1pbYjLJciAR26yMebkH,3PZmQxxLUZwyyMgXWUpmuw")
	artisResult, _ := client.Artist.Get("5IS4dQ9lDW01IY1buR7bW7")
	homeResult := &Home{
		HomeAlbums:  albumsResult,
		HomeArtists: artisResult,
	}
	fmt.Println(token1)

	c.HTML(http.StatusOK, "index.html", homeResult)

}

type Home struct {
	HomeAlbums  *spotify.AlbumsResult
	HomeArtists *spotify.Artist
}
