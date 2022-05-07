package board

const BoardSize = 8

func init() {
	if BoardSize%2 != 0 {
		panic("BoardSize must be even!!")
	}
}
