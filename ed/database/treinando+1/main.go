package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func __toStr(vet []int) string{
	if len(vet) == 0{
		return ""
	}
	if len(vet) == 1 {
		return fmt.Sprint(vet[0])
	}

	return fmt.Sprint(vet[0]) + ", " + __toStr(vet[1:])

}

func tostr(vet []int) string {
	texto := __toStr(vet)
	return "[" + texto + "]"
}

func __toStrRev(vet []int) string{
	if len(vet) == 0{
		return ""
	} 
	if len(vet) == 1 {
		return fmt.Sprint(vet[0])
	}

	return __toStrRev(vet[1:]) +  ", " + fmt.Sprint(vet[0])

}

func tostrrev(vet []int) string {
	texto := __toStrRev(vet)
	return "[" + texto + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	for i, j := 0, len(vet)-1; i < j; i, j = i+1, j-1{
		vet[i], vet[j] = vet[j], vet[i]
	}
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0{
		return 0 
	}
	var soma int
	soma = vet[0] + sum(vet[1:])
	return soma
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0{
		return 1
	}
	var multe int
	multe = vet[0] * mult(vet[1:])
	return multe
}

// min: retorna o índice e valor do menor valor
func min(vet []int) (int, int) {
	if len(vet) == 0{
		return -1, -1
	}
	if len(vet) == 1 {
		return vet[0], 0
	}
	menor := vet[0]
	indice := 0
	menorRec, indiceRec := min(vet[1:])
	indiceRec+=1
	if menorRec < menor{
		menor = menorRec
		indice = indiceRec
	}
	
	return menor, indice
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			_, index := min(vet)
			fmt.Println(index)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
