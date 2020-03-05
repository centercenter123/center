package trie 

import (
	"fmt"
	"container/list"
)

type Element struct {
	value rune
	child []*Element
}


//查找某个字符是否存在
func (e *Element) search(ch rune) *Element {
	child := e.child
	if child == nil {
		return nil
	}
	for _, value := range child {
		if value.value == ch {
			return value
		}
	}
	return nil
}

//插入一个字符
func (e *Element) insert(ch rune) *Element{
	next := e.search(ch)
	if next == nil {
		next = &Element{
			value: ch,
		}
		child := e.child
		if child == nil {
			child = []*Element{}
		}
		child = append(child, next)
		e.child = child
	}
	return next
}


type Trie struct {
	root *Element
}

//初始化trie
func Init() Trie {
	trie := Trie {
		root: &Element{
			value: 0,
		},
	}
	fmt.Printf("%+v \n", trie.root)
	str := []string {
		"ab",
		"abc",
		"abd",
		"12445",
	}
	trie.InitByArr(str)
	trie.Print()
	return trie
}

//出入字符串数组
func (trie Trie) InitByArr(arr []string) Trie {
	if trie.root == nil {
		trie = Init()
	}
	for _, str := range arr {
		e := trie.root
		for _, ch := range str {
			e = e.insert(ch)
		}
		e = e.insert(0)
	}

	return trie
}

func (trie Trie) Print() {
	root := trie.root
	list1, list2 := list.New(), list.New()

	list1.PushBack(root)
	for {
		e := list1.Front()
		if e == nil {
			if list2.Front() == nil {
				break
			}
			list1, list2 = list2, list1
			fmt.Println()
			continue
		} 
		element := (e.Value).(*Element)
		element.print()
		child := element.child
		if child != nil {
			for _, value := range child {
				list2.PushBack(value)
			}
		}
		list1.Remove(e)	
	}

}

//打印一个节点包含元素
func (e *Element) print() {
	child := e.child 
	if child == nil {
		fmt.Print("(nil)")
	} else {
		for _, value := range child {
			if value == nil || value.value == 0{
				fmt.Print("(nil)")
			} else {
				fmt.Printf("%c", value.value)
			}
		}
	}
	fmt.Print(" ")
}

func (trie Trie) startWith(str string) []string{
	e := trie.root
	result := []string{}
	if str == "" {
		return result
	}

	for _, value := range str {
		e = e.search(value)
		if e == nil {
			return result
		}
	}
	return result
	
}
