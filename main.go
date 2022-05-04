// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

package main

import "gothello/controller"

func main() {
	//var myBoard = board.InitStartBoard()
	//fmt.Printf("hallo: %d\n", myBoard.Perft(11))
	controller.RunController()
}
