package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root 
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func equals(list1 *LList, list2 *LList) bool {
	if list1.size != list2.size {
		return false
	}

	n1 := list1.Front()
	n2 := list2.Front()

	for i := 0; i < list1.size; i++ {
		if n1.Value != n2.Value {
			return false
		}
		n1 = n1.Next()
		n2 = n2.Next()
	}

	return true
}


func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func (node *Node) Next() *Node{
	if node.next == node.root || node == node.root{
		return nil
	}

	return node.next
}

func (node *Node) Prev() *Node{
	if node.prev == node.root || node == node.root{
		return nil
	}
	return node.prev
}
func NewNode(value int) *Node{
	return &Node{
		Value: value,
	}
}



func NewList() *LList{
	root := NewNode(0)
	root.next = root
	root.prev = root
	root.root = root
	return &LList{root:root}
}

func (l LList) Front() *Node{
	if l.root.next == l.root{
		return nil
	}

	return l.root.next
}

func (l LList) Back() *Node {
	if l.root.prev == l.root {
		return nil
	}
	return l.root.prev
}


func (l *LList) Insert(n *Node, value int) {
	node := NewNode(value)
	node.root = l.root
	node.prev = n.prev
	node.next = n

	n.prev.next = node
	n.prev = node
	l.size++
}

func (l *LList) Search(value int) *Node{
	if l.root.next == l.root{
		return nil
	}

	for n := l.Front(); n != nil; n = n.Next(){
		if n.Value == value{
			return n
		}
	}
	return nil
}

func (l *LList) Remove(node *Node){
	node.next.prev = node.prev
	node.prev.next = node.next
	node = nil
}

func addsorted(list *LList, value int) {
	if list.size == 0{
		list.PushBack(value)
		return
	}
	for node := list.Front(); node != nil; node = node.Next(){
		if value < node.Value{
			list.Insert(node, value)
			return
		}
	}
	list.PushBack(value)

}



func (l *LList) PushFront(value int){
	l.Insert(l.root.next, value)
}

func (l *LList) Clear() {
	current := l.root.next
	for current != l.root {
		next := current.next
		current.prev = nil
		current.next = nil
		current = next
	}
	l.root.next = l.root
	l.root.prev = l.root
}

func (l *LList) String() string {
	var sb strings.Builder
	sb.WriteString("[")

	for node := l.Front(); node != nil; node = node.Next() {
		sb.WriteString(fmt.Sprintf("%d", node.Value))
		if node.Next() != nil {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

func reverse(list *LList) {
	if list.size == 0 {
		return
	}
	newList := NewLList()
	for node := list.Back(); node != nil; node = node.Prev() {
		newList.PushBack(node.Value)
	}
	list.Clear()
	for node := newList.Front(); node != nil; node = node.Next() {
		list.PushBack(node.Value)
	}
}

func merge(list *LList, list2 *LList) *LList{
	newList := NewLList()
	for node := list.Front(); node != nil; node = node.Next() {
		addsorted(newList, node.Value)
	}

	for node := list2.Front(); node != nil; node = node.Next() {
		addsorted(newList, node.Value)
	}
	
	return newList

}


func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			 lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
			addsorted(lla, value)
		 }
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla)
		case "merge":
			 lla := str2list(args[1])
			 llb := str2list(args[2])
			 merged := merge(lla, llb)
			 fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
