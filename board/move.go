package board

type tMove struct {
	discsFlipped, discPlayed uint64
}

func (mv *tMove) isPass() bool {
	return mv.discPlayed == 0
}
