package controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gothello/docs"
	"gothello/service"
	"net/http"
	"os"
	"strconv"
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

func getNewBoard(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetNewBoard())
}

func getBoard(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetBoard())
}

func doMove(c *gin.Context) {
	col, _ := strconv.Atoi(c.Param("column"))
	row, _ := strconv.Atoi(c.Param("row"))
	c.IndentedJSON(http.StatusOK, service.DoMove(col, row))
}

func doPassMove(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.DoPassMove())
}

func takeBackLastMove(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.TakeBackLastMove())
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

func RunController() {
	router.GET("/", getHtml)

	router.GET("/api/v1/sizes", getSizes)

	router.GET("/api/v1/board", getBoard)
	router.POST("/api/v1/board", getNewBoard)
	router.POST("/api/v1/move/:column/:row/", doMove)
	router.POST("/api/v1/move/passmove/", doPassMove)
	router.POST("/api/v1/move/takeback/", takeBackLastMove)

	err := router.Run(ADDRESS + ":" + PORT)
	if err != nil {
		println("ERROR: running the router/handler not working")
		return
	}
}
