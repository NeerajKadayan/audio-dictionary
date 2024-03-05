package main

import (
	"github.com/NeerajKadayan/audio-dictionary/internal"
	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	internal.RouterMain(app)

	err := app.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

}
