package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/internal/service/spotify"
)

func (h *Handler) search(c *gin.Context) {
	client := spotify.NewClient(token)
	r, _ := client.Search.Get("Muse", "track,artist", 10, 0)
	result, _ := json.Marshal(r)
	c.JSON(http.StatusOK, result)
}
