package main
import "fmt"
func main() {
    var tamFila int
    fmt.Scan(&tamFila)
    elementosFila := make([]int, tamFila)

    for i := range elementosFila{
        fmt.Scan(&elementosFila[i])
    }
    
    var tamElementosApagar int
    fmt.Scan(&tamElementosApagar)
    var val int
    elementosApagar := make(map[int]struct{})
    for i := 0; i<tamElementosApagar; i++{
        fmt.Scan(&val)
        elementosApagar[val] = struct{}{}
    }

    listaNova := []int{}
    for _,i := range elementosFila{
        _, ok := elementosApagar[i]
        if !ok {
            listaNova = append(listaNova, i)
        } 
    }

    for v :=0; v<len(listaNova); v++{
        fmt.Printf("%v ", listaNova[v])  
    }
    fmt.Print("\n")
}
