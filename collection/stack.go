package collection

type Stack[T any] struct {
	stack [100]*T
	size  int
}

func (thisStack *Stack[T]) Push(element *T) {
	thisStack.stack[thisStack.size] = element
	thisStack.size++
}

func (thisStack *Stack[T]) Pop() *T {
	thisStack.size--
	return thisStack.stack[thisStack.size]
}

func (thisStack *Stack[T]) IsEmpty() bool {
	return thisStack.size <= 0
}

func (thisStack *Stack[T]) FromTop(i int) *T {
	return thisStack.stack[thisStack.size-i-1]
}

func (thisStack *Stack[T]) Size() int {
	return thisStack.size
}
