package main
import "fmt"
func main() {
    var valores int
    fmt.Scan(&valores)

    var numerosCasais = make([]int, valores)

    for i := 0; i < valores; i++{
        fmt.Scan(&numerosCasais[i])
    }
    var casais int
    
    for i := 0; i < valores; i++{
        for j := 0; j < valores; j++{
            if numerosCasais[i]+numerosCasais[j] == 0{
                casais++
                numerosCasais[i] = 0
                numerosCasais[j] = 1000
                break
            }
        }
    }
    fmt.Println(casais)
}
