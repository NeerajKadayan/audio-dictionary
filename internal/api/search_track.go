package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/NeerajKadayan/audio-dictionary/config"
	"github.com/NeerajKadayan/audio-dictionary/internal/dtos"
	"github.com/gin-gonic/gin"
)

func TrackSearch(c *gin.Context) {

	queryParams := url.Values{}
	queryParams.Set("apikey", config.MusixMatchApiKey)
	queryParams.Set("q_track", c.Query("q_track"))

	endpoint := "https://api.musixmatch.com/ws/1.1/track.search?" + queryParams.Encode()

	response, err := http.Get(endpoint)
	if err != nil {
		log.Println("failed to make request", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	defer response.Body.Close()

	var data dtos.MusixmatchResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println("error decoding json response", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if len(data.Message.Body.TrackList) == 0 {
		log.Println("no tracks found")
		c.JSON(http.StatusNotFound, gin.H{"message": "no tracks found"})
		return
	}

	c.JSON(http.StatusOK, data)
}
