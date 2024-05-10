package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markovk1n/spoty/internal/service/spotify"
	"github.com/markovk1n/spoty/models"
)

func (h *Handler) home(c *gin.Context) {
	userID := GetUserID(c)
	token1, _ := spotify.GetSpotifyToken(h.ClientID, h.ClientSecret)
	at := "5IS4dQ9lDW01IY1buR7bW7,7dGJo4pcD2V6oG8kP0tJRR,3YQKmKGau1PzlVlkL1iodx,6olE6TJLqED3rqDCT0FyPh"

	client := spotify.NewClient(token1)
	albumsResult, _ := client.Album.List("382ObEPsp2rxGrnsizN5TX,1A2GTWGtFfWp7KSQTwWOyo,2noRn2Aes5aoNVsU6iWThc,2cWBwpqMsDJC1ZUwz813lo,2WT1pbYjLJciAR26yMebkH,3PZmQxxLUZwyyMgXWUpmuw")
	art := GetArtistsForHome(token1, at)

	var flag bool
	_, err := c.Request.Cookie("session_token")
	if err != nil {
		flag = false

	} else {
		flag = true
	}

	// artRes := []spotify.Artist{}
	// ///
	// artists := []string{
	// "5IS4dQ9lDW01IY1buR7bW7",
	// 	"7dGJo4pcD2V6oG8kP0tJRR",
	// 	"3YQKmKGau1PzlVlkL1iodx",
	// 	"6olE6TJLqED3rqDCT0FyPh",
	// }

	// for i := range artists {
	// 	artist1, _ := client.Artist.Get(artists[i])
	// 	artRes = append(artRes, *artist1)
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println(artist1)

	// }

	// ///

	homeResult := &Home{
		User: flag,
		HomeUser: models.User{
			Id: userID,
		},
		HomeAlbums:  albumsResult,
		HomeArtists: art,
	}

	c.HTML(http.StatusOK, "index.html", homeResult)

}

type Home struct {
	User        bool
	HomeUser    models.User
	HomeAlbums  *spotify.AlbumsResult
	HomeArtists *spotify.ArtistsResult
}

func GetArtistsForHome(token, ids string) *spotify.ArtistsResult {
	url := fmt.Sprintf("https://api.spotify.com/v1/artists?ids=%s", ids)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	res := &spotify.ArtistsResult{}
	json.Unmarshal(body, res)
	return res
}
