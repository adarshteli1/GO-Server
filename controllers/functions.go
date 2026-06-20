package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Formfunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("name")
		address := c.PostForm("address")
		responseString := fmt.Sprintf("The name is : %v\n and tge address is %v", name, address)

		c.String(http.StatusOK, responseString)
	}
}

func Hellofunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello welcome to my Server")

	}

}
