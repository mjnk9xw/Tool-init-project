package routers

import (
	"GauGau/agen/clean-code/api/controllers"
	"GauGau/agen/clean-code/services"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

func NewGin(s services.IService) {
	g := gin.Default()
	g.Use(gin.Recovery())

	controllers.NewControllers(g)

	g.Run()
}

func NewEcho(s services.IService) {
	// TODO:
}
