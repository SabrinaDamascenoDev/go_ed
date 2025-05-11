package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
    scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	linha := scanner.Text()
	campos := strings.Fields(linha)

	var elementos []int
	for _, campo := range campos {
		valor, err := strconv.Atoi(campo)
		if err == nil {
			elementos = append(elementos, valor)
		}
	}

    scanner.Scan()
	elementoAcharStr := scanner.Text()
	elementoAchar, _ := strconv.Atoi(elementoAcharStr)

    var indexColocar int
    if BinarySearch(elementos, elementoAchar) == -1{
        if len(elementos)  == 0 {
            elementos = append(elementos, elementoAchar)
        } else {
            for i, v := range elementos{
                if elementoAchar < v{
                    indexColocar = i
                    break
                } 
            }
            if elementoAchar > elementos[len(elementos)-1]{
                indexColocar = len(elementos)
            }
            elementos = append(elementos, 0)
            for i := len(elementos)-1; i > indexColocar; i--{
                elementos[i] = elementos[i-1]
            } 
            
            elementos[indexColocar] = elementoAchar

        }
        
    } else {
        for i, v := range elementos{
            if elementoAchar == v{
                indexColocar = i
            }
        }
    }

    fmt.Println(indexColocar)

}
