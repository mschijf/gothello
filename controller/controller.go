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
const BOARD_COOKIE = "OTHELLOSTATUS"

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

func setBoardStringCookie(c *gin.Context, cookieValue string) {
	c.SetCookie(BOARD_COOKIE, cookieValue, 3600*24*365, "/", ADDRESS, false, true)
}

// @Router       /v1/api/board [post]
func getNewBoard(c *gin.Context) {
	result, cookieValue := service.GetNewBoard()
	setBoardStringCookie(c, cookieValue)
	c.IndentedJSON(http.StatusOK, result)
}

// @Router       /v1/api/board [get]
func getBoard(c *gin.Context) {
	cookie, _ := c.Cookie(BOARD_COOKIE)
	result, cookieValue := service.GetBoard(cookie)
	setBoardStringCookie(c, cookieValue)
	c.IndentedJSON(http.StatusOK, result)
}

// @Router       /v1/api/move/{column}/{row} [post]
func doMove(c *gin.Context) {
	col, _ := strconv.Atoi(c.Param("column"))
	row, _ := strconv.Atoi(c.Param("row"))
	cookie, _ := c.Cookie(BOARD_COOKIE)
	result, cookieValue := service.DoMove(cookie, col, row)
	setBoardStringCookie(c, cookieValue)
	c.IndentedJSON(http.StatusOK, result)
}

// @Router       /v1/api/move/passmove [post]
func doPassMove(c *gin.Context) {
	cookie, _ := c.Cookie(BOARD_COOKIE)
	result, cookieValue := service.DoPassMove(cookie)
	setBoardStringCookie(c, cookieValue)
	c.IndentedJSON(http.StatusOK, result)
}

// @Router       /v1/api/move/takeback/ [post]
func takeBackLastMove(c *gin.Context) {
	cookie, _ := c.Cookie(BOARD_COOKIE)
	result, cookieValue := service.TakeBackLastMove(cookie)
	setBoardStringCookie(c, cookieValue)
	c.IndentedJSON(http.StatusOK, result)
}

// @Router       / [get]
func getHtml(c *gin.Context) {
	html, err := os.ReadFile("./view/othello.html")
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", html)
}

func RunController() {
	router.GET("/", getHtml)

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
