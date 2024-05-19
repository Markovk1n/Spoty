package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/internal/service/spotify"
)

type ArtistPageResult struct {
	ArtistInfo *spotify.Artist
	TopTracks  *spotify.TracksResult
}

func (h *Handler) getArtists(c *gin.Context) {
	token1, _ := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)
	client := spotify.NewClient(token1)

	artistResult, _ := client.Artist.List("2CIMQHirSU0MQqyYHq0eOx,57dN52uHvrHOxijzpIgu3E,1vCWHaC5f2uS3yhpwWbIA6")
	c.HTML(http.StatusOK, "artists.html", artistResult)
}

func (h *Handler) getArtistById(c *gin.Context) {
	artistId := c.Param("id")
	token1, err := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)

	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	client := spotify.NewClient(token1)
	topTracks, _ := client.Artist.GetTopTracks(artistId, "EU")

	client = spotify.NewClient(token1)
	artistInfo, _ := client.Artist.Get(artistId)

	// result2, _ := client.Artist.GetAlbums("4tZwfgrHOc3mvqYlEYSvVi", "single,appears_on", "US", 10, 0)

	result := ArtistPageResult{
		ArtistInfo: artistInfo,
		TopTracks:  topTracks,
	}

	c.HTML(http.StatusOK, "artist.html", result)
}
