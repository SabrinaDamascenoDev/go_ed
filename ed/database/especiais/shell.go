package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	for i := 0; i <len(vet); i++{
		if vet[i] < 0{
			vet[i] = vet[i]*-1
 		}
	}
	sort.Ints(vet)
	var pares []Pair
	contador := make(map[int]int)

	for _, v := range vet{
		contador[v]++
	}
	
	for v, qnt := range contador{
		pares = append(pares, Pair{One: v, Two: qnt})

	}
	return pares
}

func teams(vet []int) []Pair {
	var grupos []Pair
	if len(vet) == 0 {
		return grupos
	} else {
		contagem := 1
		elemeneto := vet[0]

		for i := 1; i < len(vet); i++{
			if elemeneto == vet[i]{
				contagem++
			} else {
				grupos = append(grupos, Pair{One: elemeneto, Two: contagem})
				contagem =1
				elemeneto = vet[i]
			}
		}
		grupos = append(grupos, Pair{One: elemeneto, Two: contagem})
		
		return grupos
	}
}

func mnext(vet []int) []int {
	n := len(vet)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		if vet[i] > 0 {
			temMulherAoLado := false

			if i > 0 && vet[i-1] < 0 {
				temMulherAoLado = true
			}
			if i < n-1 && vet[i+1] < 0 {
				temMulherAoLado = true
			}

			if temMulherAoLado {
				result[i] = 1
			}
		}
	}
	return result
}




func alone(vet []int) []int {
	n := len(vet)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		if vet[i] > 0 {
			temMulherAoLado := false
			if i > 0 && vet[i-1] < 0 {
				temMulherAoLado = true
			}
			if i < n-1 && vet[i+1] < 0 {
				temMulherAoLado = true
			}

			if temMulherAoLado {
				result[i] = 0
			} else {
				result[i] = 1
			}
		}
	}
	return result
}

func couple(vet []int) int {
	casais := 0
	for i := 0; i < len(vet); i++ {
		for j := 0; j < len(vet); j++{
			if vet[i] + vet[j] == 0{
				casais++
				vet[i] = 0
				vet[j] = 1000
			}
		}
	}
	return casais
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	for i := range seq{
		if pos+i > len(vet){
			return false
		} 
		if vet[pos+i] != seq[i]{
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	posicao := -1
	for i := 0; i < len(vet) - len(seq) + 1; i++{
		if hasSubseq(vet, seq, i){
			posicao = i
			break
		}
	}
	return posicao
	
}

func erase(vet []int, posList []int) []int {
	newVet := []int{}

	for i := 0; i< len(posList); i++{
		for j := 0; j < len(vet); j++{
			if posList[i] == j{
				vet[j] = 0
			}
		}
	}

	for _, v := range vet{
		if v != 0{
			newVet = append(newVet, v)
		}
	}

	return newVet
}

func clear(vet []int, value int) []int {
	novoArray := []int{}
	for i := 0; i < len(vet); i++{
		if vet[i] == value{
			vet[i] = 0
		}
	}

	for i := 0; i < len(vet); i++{
		if vet[i] != 0{
			novoArray = append(novoArray, vet[i])
		}
	}
	return novoArray

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
