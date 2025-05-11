package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func (v *MultiSet) String() string {
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

func (v *MultiSet) Contains(value int) bool{
	if BinarySearch(v.data, value) != -1{
		return true
	}
	return false
}
func NewMultiSet(capacity int) *MultiSet{
	if capacity < 0{
		capacity = 0
	}

	return &MultiSet{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}
func (v *MultiSet) Reserve(capacity int) {
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
func (v *MultiSet) Sort() {
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

func (v *MultiSet) Insert(value int){
	if v.size == v.capacity {
		v.Reserve(max(1, 2 * v.capacity))
	}

	v.data[v.size] = value
	v.size++
	v.Sort()

}

func (v *MultiSet) Erase(value int) {
	if BinarySearch(v.data, value) == -1{
		fmt.Println("value not found")
		return
	}
	var indexQueE int
	for i := 0; i < v.size-1; i++{
		if v.data[i] == value{
			indexQueE = i
		}
	}
	
	for i := indexQueE; i < v.size-1; i++{
		v.data[i] = v.data[i+1]
	}
	v.size -= 1
}


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	 ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "magic":
			// value, _ := strconv.Atoi(args[1])
		case "insert":
			for _, part := range args[1:] {
				value, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("fail: invalid value", part)
					continue
				}
			ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			for _, part := range args[1:] {
				value, err := strconv.Atoi(part)
				if err != nil {
					fmt.Println("fail: invalid value", part)
					continue
				}
			ms.Erase(value)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
