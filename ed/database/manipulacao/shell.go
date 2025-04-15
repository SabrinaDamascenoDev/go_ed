package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func getMen(vet []int) []int {
	vetHomens := []int{}
	for i := 0; i<len(vet); i++{
		if vet[i] > 0{
			vetHomens = append(vetHomens, vet[i])
		}
	}
	return vetHomens
}

func getCalmWomen(vet []int) []int {
	vetCalmWomen := []int{}
	for i := 0; i < len(vet); i++{
		if vet[i] > -10 && vet[i] < 0{
			vetCalmWomen = append(vetCalmWomen, vet[i])
		}
	}
	return vetCalmWomen
}

func sortVet(vet []int) []int {
	vetSort := []int{}
	for i := 0; i < len(vet); i++{
		vetSort = append(vetSort, vet[i])
	}
	sort.Ints(vetSort)
	return vetSort
}

func sortStress(vet []int) []int {
	vetSortStress := []int{}
	for i := 0; i < len(vet); i++{
		if vet[i] > -10 && vet[i] < 0{
			vetSortStress = append(vetSortStress, vet[i])
			vet[i] = 0
		}
	}
	vetExem := []int{}
	for i := 0; i < len(vet); i++{
		if vet[i] > 0 && vet[i] < 13{
			vetExem = append(vetExem, vet[i])
			vet[i] = 0
		}
	}
	sort.Ints(vetExem)

	for _, val := range vetExem {
		vetSortStress = append(vetSortStress, val)
	}

	sort.Ints(vet)
	
	for i := len(vet)-1; i >= 0; i--{
		if vet[i] != 0{
			vetSortStress = append(vetSortStress, vet[i])
		}
		
	}
	return vetSortStress
}

func reverse(vet []int) []int {
	vetReverse := []int{}
	primeiro := len(vet)-1
	for i := primeiro; i >= 0; i--{
		vetReverse = append(vetReverse, vet[i])
	}
	return vetReverse
}

func reverseInplace(vet []int) {
	_ = vet
}
func unique(vet []int) []int {

	repetidos := make(map[int]bool)
	vetUnicos := []int{}
	for _, v := range vet{
		if !repetidos[v]{
			repetidos[v] = true
			vetUnicos = append(vetUnicos, v)
		}
	}
	
	return vetUnicos
}

func repeated(vet []int) []int {
	vetRepeated := []int{}
	sort.Ints(vet)
	for i := 0; i < len(vet)-1; i++{
		if vet[i] == vet[i+1]{
			vetRepeated = append(vetRepeated, vet[i])
		}
	}
	return vetRepeated
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			printVec(reverse(str2vet(args[1])))
		case "reverse_inplace":
			vet := str2vet(args[1])
			reverseInplace(vet)
			printVec(vet)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

