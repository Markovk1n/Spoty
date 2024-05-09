package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/internal/service/spotify"
)

func (h *Handler) getArtists(c *gin.Context) {
	token1, _ := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)

	client := spotify.NewClient(token1)
	artistResult, _ := client.Artist.List("2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6")
	c.HTML(http.StatusOK, "artists.html", artistResult)
}

func (h *Handler) getArtistById(c *gin.Context) {

	client := spotify.NewClient(token)

	result, _ := client.Artist.Get("4tZwfgrHOc3mvqYlEYSvVi")
	// result1, _ := client.Artist.GetTopTracks("4tZwfgrHOc3mvqYlEYSvVi", "US")
	// result2, _ := client.Artist.GetAlbums("4tZwfgrHOc3mvqYlEYSvVi", "single,appears_on", "US", 10, 0)

	// AllInfo := models.AllArtistInfo{
	// 	result,
	// 	result1,
	// 	result2,
	// }
	c.HTML(http.StatusOK, "Artist.html", result)
}
