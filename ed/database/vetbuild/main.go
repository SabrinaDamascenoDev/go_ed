package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func (v *Vector) Reserve(capacity int) {
	if capacity < v.size {
		return
	}
	novo := make([]int, capacity)
	for i := range v.size {
		novo[i] = v.data[i]
	}
	v.capacity = capacity
	v.data = novo
}

func (v *Vector) PushBack(value int) {
	if v.size == v.capacity {
		v.Reserve(max(1, 2 * v.capacity))
	}
	v.data[v.size] = value
	v.size += 1
}

func (v *Vector) Clear() {
	v.data = make([]int, v.capacity)
	v.size = 0 
}

func (v *Vector) PopBack() ([]int, error){
	if v.size == 0 {
		return v.data, fmt.Errorf("vector is empty")
	}
	v.data[v.size-1] = 0
	v.size-=1

	return v.data, nil
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%v capacity:%v", v.size, v.capacity)
}

func (v *Vector) Insert(index int, value int) error{

	if v.size == v.capacity {
		v.Reserve(max(1, 2 * v.capacity))
	}
	if index > v.capacity || index < 0{
		return fmt.Errorf("Index não existe")
	}
	

	for i := v.size-1; i >= index; i--{
		v.data[i+1] = v.data[i]
	}
	
	v.data[index] = value
	v.size+=1
	return nil
}

func (v *Vector) Erase(index int) error{
	if index > v.capacity || index < 0{
		return fmt.Errorf("index out of range")
	}
	

	for i := index; i < v.size-1; i++{
		v.data[i] = v.data[i+1]
	}
	
	v.size-=1
	return nil
}

func (v *Vector) String() string {
	saida := "["
	for i := range v.size {
		if i != 0 {
			saida += ", "
		}
		saida += fmt.Sprint(v.data[i])
	}
	return saida + "]"
}

//construtor
func NewVector(capacity int) *Vector {
	if capacity < 0 {
		capacity = 0
	}
	return &Vector{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) GetSize() int {
	return v.size
}

func (v *Vector) Reverse() {
	var new []int
	for i := v.size-1; i >= 0; i--{
		new = append(new, v.data[i])
	} 
	for i := 0; i < len(new); i++{
		v.data[i] = new[i]
	}
	
}

func (v *Vector) At(index int) (int, error){
	if index < 0 || index > v.size{
		return 0, fmt.Errorf("index out of range")
	}

	return v.data[index], nil
}

func (v *Vector) Set(index int, value int) error{
	if index < 0 || index > v.size{
		return fmt.Errorf("index out of range")
	}
	v.Erase(index)
	v.Insert(index, value)

	return nil
}

func (v *Vector) Sort() {
	var new []int
	for i := v.size-1; i >= 0; i--{
		new = append(new, v.data[i])
	} 
	for i := 0; i < len(new); i++{
		for j := 0; j < len(new); j++{
			if new[i] < new[j]{
				element := new[i]
				new[i] = new[j]
				new[j] = element
				
			}
		}
	}
	for i := 0; i < len(new); i++{
		v.data[i] = new[i]
	}	
}

func (v *Vector) IndexOf(value int) int{
	for i := 0; i < v.size; i++{
		if value == v.data[i]{
			return i
		}
	}
	
	return -1
}

func (v *Vector) Contains(value int) bool{
	for i := 0; i < v.size; i++{
		if value == v.data[i]{
			return true
		}
	}
	return false
}

func (v *Vector) Sorted() []int{
	var new []int
	for i := v.size-1; i >= 0; i--{
		new = append(new, v.data[i])
	} 
	for i := 0; i < len(new); i++{
		for j := 0; j < len(new); j++{
			if new[i] < new[j]{
				element := new[i]
				new[i] = new[j]
				new[j] = element
				
			}
		}
	}
	return new
}
// stack
// heap

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
			 	v.PushBack(value)
			}
		case "show":
			fmt.Println(v)
		case "status":
			fmt.Println(v.Status())
		case "pop":
			value, err := v.PopBack()
			_= value
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			 if err != nil {
			 	fmt.Println(err)
			 }
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
			 	fmt.Println(err)
			} else {
			 	fmt.Println(value)
		 	}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
			 	fmt.Println(err)
			}
		case "reserve":
			 newCapacity, _ := strconv.Atoi(parts[1])
			 v.Reserve(newCapacity)
		case "sort":
			v.Sort()
		case "sorted":
			fmt.Println("[" + Join(v.Sorted(), ", ") + "]")
		case "reverse":
			v.Reverse()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
