package board

import (
	"fmt"
	"testing"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//    depth   0  :     0.000000 ms -->              1
//    depth   1  :     0.000001 ms -->              4
//    depth   2  :     0.000001 ms -->             12
//    depth   3  :     0.000005 ms -->             56
//    depth   4  :     0.000017 ms -->            244
//    depth   5  :     0.000100 ms -->           1396
//    depth   6  :     0.000515 ms -->           8200
//    depth   7  :     0.003094 ms -->          55092
//    depth   8  :     0.021635 ms -->         390216
//    depth   9  :     0.137550 ms -->        3005288
//    depth  10  :     1.109636 ms -->       24571284
//    depth  11  :     9.388059 ms -->      212258800
//    depth  12  :    83.742958 ms -->     1939886636
//    depth  13  :   782.551742 ms -->    18429641748
//
//    speed: 23.550.700 per second
//
//    see also http://www.aartbik.com/strategy.php
//
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Test_tBitBoard_perft(t *testing.T) {
	hb := InitStartBoard()

	tables := []struct {
		x int
		n int64
	}{
		{9, 3005288},
		{10, 24571284},
	}

	for _, table := range tables {
		nodeCount := hb.bitBoard.perft(table.x, black, false)
		if nodeCount != table.n {
			t.Errorf("Perft of %d was incorrect, got: %d, want: %d.", table.x, nodeCount, table.n)
		}
	}
}

func Test_bitBoard_perft_print(t *testing.T) {
	var hb = InitStartBoard()

	for i := 0; i < 14; i++ {
		currentTime := time.Now()
		result := hb.bitBoard.perft(i, black, false)
		diff := time.Now().Sub(currentTime)
		fmt.Printf("depth %3d  : %12.6f ms --> %14d\n", i, diff.Seconds(), result)
	}
}
