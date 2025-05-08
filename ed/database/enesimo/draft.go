package main
import "fmt"

func enesimo(x int, div int) int{
    if div == 1 {
        return x
    }

    return enesimo(x, div-1)
}
func main() {
    var x int
    fmt.Scan(&x)
    
}
