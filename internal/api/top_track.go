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

func TopTrack(c *gin.Context) {

	// Encode query parameters
	queryParams := url.Values{}
	queryParams.Set("apikey", config.MusixMatchApiKey)
	queryParams.Set("country", c.Query("country"))

	// Musixmatch API endpoint for getting top tracks
	endpoint := "https://api.musixmatch.com/ws/1.1/chart.tracks.get?" + queryParams.Encode()

	// Making the API request
	response, err := http.Get(endpoint)
	if err != nil {
		log.Println("failed to make request", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer response.Body.Close()

	// Decode JSON response
	var data dtos.MusixmatchResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println("error decoding json response", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Check if any tracks were found
	if len(data.Message.Body.TrackList) == 0 {
		log.Println("no tracks found")
		c.JSON(http.StatusNotFound, gin.H{"message": "no tracks found"})
		return
	}

	// Print information about the top track
	topTrack := data.Message.Body.TrackList[0].Track

	c.JSON(http.StatusOK, topTrack)
}
