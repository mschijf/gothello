package board

type Move struct {
	discsFlipped, discPlayed uint64
}

func (mv *Move) isPass() bool {
	return mv.discPlayed == 0
}
