package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Set struct{
	data []int
	size int
	capacity int
}

func NewSet(capacity int) *Set{
	if capacity < 0{
		capacity = 0
	}
	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}
func binarySearch(value int, slice []int, inf int, sup int) int{
	if sup < inf{
		return -1
	}
	meio := (inf+sup)/2
	if slice[meio] == value{
		return meio
	}

	if value < slice[meio]{
		return binarySearch(value, slice, inf, meio-1)
	} 
	return binarySearch(value, slice, meio+1, sup)
}

func BinarySearch(slice []int, value int) int{
	return binarySearch(value, slice, 0, len(slice)-1)
}
func (v *Set) Sort() {
	if v.size == 1{
		return
	}
	new := make([]int, v.size)
	copy(new, v.data[:v.size])

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

func (v *Set) Reserve(capacity int) {
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


func (v *Set) Insert(value int){
	if BinarySearch(v.data, value) != -1{
		return
	}
	if v.size == v.capacity {
		v.Reserve(max(1, 2 * v.capacity))
	}

	v.data[v.size] = value
	v.size++
	v.Sort()


}

func (v *Set) Contains(value int) bool{
	if BinarySearch(v.data, value) != -1{
		return true
	}
	return false
}

func (v *Set) String() string {
	result := "["
	for i := 0; i < v.size; i++ {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprint(v.data[i])
	}
	result += "]"
	return result
}
func (v *Set) Erase(value int) {
	if BinarySearch(v.data, value) == -1{
		fmt.Println("value not found")
		return
	}
	for i := value-1; i < v.size-1; i++{
		v.data[i] = v.data[i+1]
	}
	v.size -= 1
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("fail: invalid value", part)
					continue
				}
			v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			for _, part := range parts[1:] {
				value, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("fail: invalid value", part)
					continue
				}
			v.Erase(value)
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
