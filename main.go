package main

import (
	animalAdaptor "packageByFeature/Package/Animal/Get/Adaptor"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/animal/name", animalAdaptor.GetName)

	err := r.Run()
	if err != nil {
		return
	}
}
