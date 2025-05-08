package main
import "fmt"

type Par struct{
    divisao  int
    resto int
}
func main(){
    var val int
    fmt.Scan(&val)

    resul := RestoeValor(val)
    for i := len(resul)-1; i >=0; i--{
        fmt.Printf("%v %v\n", resul[i].divisao, resul[i].resto)
    }
}

func RestoeValor (val int) []Par{
    if val == 0{
        return []Par{}
    }

    newValor := val/2
    resto := val%2
    restoValor := []Par{}

    restoValor = append(restoValor, Par{divisao: newValor, resto: resto})
    restoValor = append(restoValor, RestoeValor(newValor)...)

    return restoValor
    
}


