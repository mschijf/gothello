package board

type MoveStack struct {
	stack []*Move
}

func (ms *MoveStack) push(move *Move) {
	ms.stack = append(ms.stack, move)
}

func (ms *MoveStack) pop() *Move {
	n := len(ms.stack) - 1 // Top element
	move := ms.stack[n]
	ms.stack = ms.stack[:n]
	return move
}

func (ms *MoveStack) isEmpty() bool {
	return len(ms.stack) == 0
}

func (ms *MoveStack) fromTop(i int) *Move {
	return ms.stack[len(ms.stack)-i-1]
}

func (ms *MoveStack) size() int {
	return len(ms.stack)
}
