package main

import (
	"fmt"
	"gothello/board"
	"time"
)

func main() {
	//var bbWhite uint64
	//
	//bbWhite = 7
	//x := bits.OnesCount64(bbWhite)
	//fmt.Printf("%064b heeft lengte %d\n", bbWhite, x)

	var bb = board.InitStartBoard()

	for i := 0; i < 14; i++ {
		currentTime := time.Now()
		var result = bb.Perft(i)
		diff := time.Now().Sub(currentTime)
		fmt.Printf("depth %3d  : %12.6f ms --> %14d\n", i, diff.Seconds(), result)
	}
}
