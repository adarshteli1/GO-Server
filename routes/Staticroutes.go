package routes

import (
	controller "Go-Server/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Staticroutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/hello", controller.Hellofunc())
	incomingRoutes.POST("/form", controller.Formfunc())

	fileServer := http.FileServer(http.Dir("./static"))
	incomingRoutes.NoRoute(gin.WrapH(fileServer))
}
