package main

import (
	"github.com/ANPez/whoapi/ctrl"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/episodes", ctrl.GetEpisodes)

	r.GET("/doctors/:id/actors", ctrl.GetActors)

	r.GET("/doctors", ctrl.GetDoctors)

	r.Run()

}
