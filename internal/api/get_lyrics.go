package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/NeerajKadayan/audio-dictionary/config"
	"github.com/gin-gonic/gin"
)

func GetLyrics(c *gin.Context) {

	// Encode query parameters
	queryParams := url.Values{}
	queryParams.Set("apikey", config.MusixMatchApiKey)
	queryParams.Set("track_id", c.Query("track_id"))

	// Musixmatch API endpoint for getting track lyrics
	endpoint := "https://api.musixmatch.com/ws/1.1/track.lyrics.get?" + queryParams.Encode()

	// Making the API request
	response, err := http.Get(endpoint)
	if err != nil {
		log.Println("failed to make request", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer response.Body.Close()

	// debug logs
	log.Println("resp: \n", response.Body)

	// Decode JSON response
	var data interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println("error decoding json response", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
