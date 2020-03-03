package tree

import (
	"fmt"
	"container/list"
)

type Element struct {
	value      interface{}
	leftChild  *Element
	rightChild *Element
}

//取节点值
func (e Element) Value() interface{} {
	return e.value
}

//取左孩子节点
func (e *Element) LeftChild() *Element {
	return e.leftChild
}

//取右孩子节点
func (e *Element) RightChild() *Element {
	return e.rightChild
}

type Tree struct {
	root Element
}


func Init() *Tree {
	tree := new(Tree)
	tree.root = Element{
		value: 1,
		leftChild: &Element{
			value: 8,
			leftChild: &Element{
				value: 2,
			},
			rightChild: &Element{
				value: 6,
				rightChild: &Element{
					value: 7,
				},
			},
		},
		rightChild: &Element{
			value: 9,
			leftChild: &Element{
				value: 10,
			},
			rightChild: &Element{
				value: 15,
			},
		},
	}
	return tree
}

//根据数据初始化树
func InitByArray(arr []interface{}) *Tree {
	tree := new(Tree)
	treeArray := make([]*Element, len(arr))
	childKey := 1
	if len(arr) == 0 || arr[0] == nil {
		return tree
	}
	element := &Element{
		value: arr[0],
	}
	tree.root = *element
	treeArray[0] = &tree.root
	for index,value := range arr {
		if childKey >= len(arr) {
			break
		}
		if value == nil {
			continue
		}
		element = treeArray[index]
		//fmt.Println(element)

		//左孩子
		if childKey < len(arr) && arr[childKey] != nil {
			leftChild := &Element{
				value: arr[childKey],
			}
			//fmt.Printf("%+v \n", leftChild)
			element.leftChild = leftChild
			treeArray[childKey] = leftChild
		}
		childKey ++
		//右孩子
		if childKey < len(arr) && arr[childKey] != nil {
			rightChild := &Element{
				value: arr[childKey],
			}
			//fmt.Printf("%+v \n", rightChild)
			element.rightChild = rightChild
			treeArray[childKey] = rightChild
		}		
		childKey ++
	}

	return tree
}

//广度优先遍历树
func (tree *Tree) BFS() {
	var element *list.Element
	var node *Element
	list := list.New()
	list.PushBack(&tree.root)
	for {
		element = list.Front()
		if element == nil {
			break
		}
		node = (element.Value).(*Element)

		fmt.Print(node.value, " ")
		leftChild := node.leftChild
		if leftChild != nil {
			list.PushBack(leftChild)
		}
		rightChild := node.rightChild
		if rightChild != nil {
			list.PushBack(rightChild)
		}
		list.Remove(element)		
	}
}

//层次遍历输出树
func (tree *Tree) BFS1() {
	var element *list.Element
	var node *Element
	list0 := list.New()
	list1 := list.New()
	//list := list.New()
	list0.PushBack(&tree.root)
	fmt.Print("[ ")
	for {
		element = list0.Front()
		if element == nil {
			if(list1.Front() == nil){
				break
			}
			fmt.Println("]")
			fmt.Print("[ ")
			list0, list1 = list1, list0
			continue
		}

		node = (element.Value).(*Element)

		fmt.Print(node.value, " ")
		leftChild := node.leftChild
		if leftChild != nil {
			list1.PushBack(leftChild)
		}
		rightChild := node.rightChild
		if rightChild != nil {
			list1.PushBack(rightChild)
		}
		list0.Remove(element)	
	}
	fmt.Println("]")
}

//BFS查看数的最小和最大深度
func (tree *Tree) BFS2() (int, int){
	var element *list.Element
	var node *Element
	list0 := list.New()
	list1 := list.New()
	max := 0 
	min := 0
	level := 1
	list0.PushBack(&tree.root)
	for {
		element = list0.Front()
		if element == nil {
			if(list1.Front() == nil){
				return max, min
			}
			list0, list1 = list1, list0
			level ++
			continue
		}

		node = (element.Value).(*Element)
		fmt.Printf("%+v \n", node)
		leftChild := node.leftChild
		if leftChild != nil {
			list1.PushBack(leftChild)
		}
		rightChild := node.rightChild
		if rightChild != nil {
			list1.PushBack(rightChild)
		}
		if leftChild == nil && rightChild == nil {
			//fmt.Printf("%+v", node)
			if max < level {
				max = level
			}
			if min == 0 || min > level {
				min = level
			}
		}
		list0.Remove(element)	
	}
}

//深度优先遍历数
func (tree *Tree) DFS() {
	root := tree.root
	root.DLR()

}
//前序遍历
func (e *Element) DLR(){
	fmt.Print(e.value, " ")
	if e.leftChild != nil {
		e.leftChild.DLR()
	}
	if e.rightChild != nil {
		e.rightChild.DLR()
	}
}

func (tree *Tree) DFS1() (int, int){
	root := tree.root
	return root.DLR1(1, 0, 0)
}

//前序遍历
func (e *Element) DLR1(level int, max int, min int) (int, int) {
	if e.leftChild == nil && e.rightChild == nil {
		if max < level {
			max = level
		}
		if min == 0 || min > level {
			min = level
		}
	}
	if e.leftChild != nil {
		max, min = e.leftChild.DLR1(level + 1, max, min)
	}
	if e.rightChild != nil {
		max, min = e.rightChild.DLR1(level +1, max, min)
	}
	return max, min
}
