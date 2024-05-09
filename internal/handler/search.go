package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/internal/service/spotify"
)

func (h *Handler) search(c *gin.Context) {
	search := c.Query("search")
	token1, err := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	client := spotify.NewClient(token1)

	// var input models.Search

	// if err := c.ShouldBind(search); err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	return
	// }

	res, _ := client.Search.Get(search, "track,artist,album", 10, 0)

	c.HTML(http.StatusOK, "search.html", gin.H{
		"results": res,
	})

}
