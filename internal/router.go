package internal

import (
	"github.com/NeerajKadayan/audio-dictionary/internal/api"
	"github.com/gin-gonic/gin"
)

// RouterMain --
func RouterMain(route *gin.Engine) {

	route.GET("/ping", api.Ping)

	route.GET("/track_search", api.TrackSearch)

	route.GET("/get_lyrics", api.GetLyrics)

	route.GET("/top_track", api.TopTrack)

	route.GET("/artist_info", api.ArtistInfo)
}
