package board

type move struct {
	discsFlipped, discPlayed uint64
}

func (mv *move) isPass() bool {
	return mv.discPlayed == 0
}
