package stack

import (
	container "go-stl/container"
	arraylist"go-stl/arraylist"
)
type Stacker interface{
	container.Container
	Push()(value interface{})
	Pop()(value interface{},ok bool)
	Peek()(value interface{},ok bool)
}
type Stack struct{
	list *arraylist.Arraylist
}
func (stack *Stack)Push()(value interface{},ok bool){
	value,ok=stack.list.Add()
}