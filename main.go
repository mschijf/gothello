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
	//var bb = board.InitStartBoard()
	//
	//for i := 0; i < 12; i++ {
	//	currentTime := time.Now()
	//	var result = bb.Perft(i)
	//	diff := time.Now().Sub(currentTime)
	//	fmt.Printf("depth %3d  : %12.6f ms --> %14d\n", i, diff.Seconds(), result)
	//}
	controller.RunController()
}
