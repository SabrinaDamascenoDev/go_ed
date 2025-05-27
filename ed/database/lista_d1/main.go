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
}
 
type Lista struct{
	root *Node
}

func NewNode(value int) *Node{
	return &Node{
		value: value,
		next: nil,
		prev: nil,
	}
}

func NewLList() *Lista{
	return &Lista{
		root: nil,
	}
}


func (l *Lista) PushFront(value int) {
	if l.root == nil {
		l.root = &Node{}
		l.root.next = l.root
		l.root.prev = l.root
	}

	novo := NewNode(value)
	novo.next = l.root.next
	novo.prev = l.root
	l.root.next.prev = novo
	l.root.next = novo
}

func (l *Lista) PushBack(value int){
	if l.root == nil {
		l.root = &Node{}
		l.root.next = l.root
		l.root.prev = l.root
	}

	novo := NewNode(value)
	ultimo := l.root.prev
	
	novo.next = l.root
	novo.prev = ultimo
	ultimo.next =  novo
	l.root.prev = novo

}


func (l *Lista) Clear() {
	if l.root == nil {
		return
	}

	
	atual := l.root.next
	for atual != l.root{
		next := atual.next
		atual.next = nil
		atual = next
	}

	l.root.next = nil
	l.root = nil
}

func (l *Lista) Size() int{
	if l.root == nil{
		return 0
	}
	count := 0
	atual := l.root.next
	for atual != l.root{
		count+=1
		atual = atual.next
	}
	return count
}

func (l *Lista) PopFront() {
	if l.root == nil{
		return
	}
	if l.root == l.root.next {
		l.root = nil
		return
	}
	depoisPrimeiro := l.root.next.next

	l.root.next = depoisPrimeiro
	depoisPrimeiro.prev = l.root
}





func (l *Lista) PopBack() {
	if l.root == nil{
		return
	}
	if l.root == l.root.next {
		l.root = nil
		return
	}
	penultimo := l.root.prev.prev

	l.root.prev = penultimo
	penultimo.next = l.root
}

func toString(node *Node) string {
    if node == nil || node.next == node {
        return "[]"
    }

    elementos := "["
    current := node.next
    first := true

    for current != node {
        if !first {
            elementos += ", "
        }
        elementos += fmt.Sprint(current.value)
        current = current.next
        first = false
    }

    elementos += "]"
    return elementos
}


func (l * Lista) String() string{
	return toString(l.root)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList() 

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
			fmt.Println(ll.Size())
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
		 	ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			 ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
