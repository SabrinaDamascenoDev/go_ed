package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings" 
)

type Node struct{
	value int
	next *Node
	prev *Node
	root *Node
}

func NewNode(value int) *Node{
	return &Node{
		value: value,
	}
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

type Lista struct{
	root *Node
}

func NewList() *Lista{
	root := NewNode(0)
	root.next = root
	root.prev = root
	root.root = root
	return &Lista{root:root}
}

func (l Lista) Front() *Node{
	if l.root.next == l.root{
		return nil
	}

	return l.root.next
}

func (l Lista) Back() *Node {
	if l.root.prev == l.root {
		return nil
	}
	return l.root.prev
}


func (l *Lista) Insert(n *Node, value int) {
	node := NewNode(value)
	node.root = l.root
	node.prev = n.prev
	node.next = n

	n.prev.next = node
	n.prev = node
}

func (l *Lista) Search(value int) *Node{
	if l.root.next == l.root{
		return nil
	}

	for n := l.Front(); n != nil; n = n.Next(){
		if n.value == value{
			return n
		}
	}
	return nil
}

func (l *Lista) Remove(node *Node){
	node.next.prev = node.prev
	node.prev.next = node.next
	node = nil
}



func (l *Lista) PushBack(value int){
	l.Insert(l.root, value)
}

func (l *Lista) PushFront(value int){
	l.Insert(l.root.next, value)
}

func (l *Lista) Clear() {
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


func (l *Lista) String() string {
	if l.root.next == l.root {
		return "[]"
	}

	elementos := []string{}
	for node := l.Front(); node != nil; node = node.Next() {
		elementos = append(elementos, fmt.Sprint(node.value))
	}

	return "[" + strings.Join(elementos, ", ") + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewList()

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
		case "show":
			fmt.Println(ll.String())
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			if ll.root.next == ll.root {
				fmt.Println("[ ]\n[ ]")
				break
			}
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")

		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
			   fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
			 	ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
