package main

import (
	"github.com/gin-gonic/gin"
	"pbf/package/animal"
)

func main() {
	r := gin.Default()

	animal.ApplyRoutes(r)

	err := r.Run()
	if err != nil {
		return
	}
}
