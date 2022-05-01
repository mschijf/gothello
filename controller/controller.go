package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gothello/docs"
	"net/http"
	"os"
)

const ADDRESS = "localhost"
const PORT = "8080"

var router *gin.Engine

func init() {
	router = gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		println("ERROR: setting trusted proxies not succeeding" + err.Error())
		return
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @Summary      Get Sizes of the Board
// @Tags         size
// @Router       /sizes [get]
func getSizes(c *gin.Context) {
	sizeRecord := size{MaxRow: 8, MaxCol: 8}
	c.IndentedJSON(http.StatusOK, sizeRecord)
}

func getHtml(c *gin.Context) {
	html, err := os.ReadFile("./view/othello.html")
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(string(html)))
}

func TestController() {
	router.GET("/api/v1/sizes", getSizes)
	router.GET("/", getHtml)
	err := router.Run(ADDRESS + ":" + PORT)
	if err != nil {
		println("ERROR: running the router/handler not working")
		return
	}
}
