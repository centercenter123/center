package stack
import (
	"container/list"
	"fmt"
)

type Stack struct {
	list.List
}

//初始化栈
func New() *Stack{
	stack := new(Stack)
	return stack
}

//查看栈长度
func (stack Stack) Len() int {
	return stack.List.Len()
}

//查看栈顶元素
func (stack Stack) Front() interface{}	{
	element := stack.List.Front()
	if element != nil {
		return element.Value
	}
	return nil
}

//出栈
func (stack *Stack) Pop() interface{}{
	element := stack.List.Front()
	v := stack.List.Remove(element)
	return v
}
//入栈
func (stack *Stack) Push(v interface{}) {
	stack.List.PushFront(v)
}

//打印栈
func (stack *Stack) Print() {
	fmt.Print("root -> ")
	element := stack.List.Front()
	for element != nil {
		fmt.Print(element.Value, " -> ")
		element = element.Next()
	}
	fmt.Println("nil")
}



