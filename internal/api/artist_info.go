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

func ArtistInfo(c *gin.Context) {

	// Encode query parameters
	queryParams := url.Values{}
	queryParams.Set("apikey", config.MusixMatchApiKey)
	queryParams.Set("artist_id", c.Query("artist_id")) // 51771

	// Musixmatch API endpoint for getting artist information
	endpoint := "https://api.musixmatch.com/ws/1.1/artist.get?" + queryParams.Encode()

	// Making the API request
	response, err := http.Get(endpoint)
	if err != nil {
		log.Println("failed to make request", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	defer response.Body.Close()

	// Decode JSON response
	var data dtos.ArtistResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println("error decoding json response", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Check if artist information was found
	if data.Message.Body.Artist.ArtistID == 0 {
		log.Println("artist not found")
		c.JSON(http.StatusNotFound, gin.H{"message": "artist not found"})
		return
	}

	c.JSON(http.StatusOK, data)
}
