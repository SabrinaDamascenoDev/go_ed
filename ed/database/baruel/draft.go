package main
import "fmt"

func main() {
    var totalFigurinhas, totalBaruel int

    fmt.Scan(&totalFigurinhas, &totalBaruel)

    var figurinhasTodas = make([]int, totalBaruel)
     // Povoando o Array de Todas as Figs
    for i := 0; i < totalBaruel; i++{
        fmt.Scan(&figurinhasTodas[i])
    }

    figurinhasRepetidas := []int{}

    // Povoando Array de Figs Repetidas
    for i := 0; i < totalBaruel-1; i++{
        if(figurinhasTodas[i] == figurinhasTodas[i+1]){
          figurinhasRepetidas = append(figurinhasRepetidas, figurinhasTodas[i])
        }
    }
    
    // Se nao tiver nenhuma fig repetida print N
    if len(figurinhasRepetidas) == 0{
        fmt.Print("N")
    }

    //Imprime os valores que estão repetidos na formatação pedida
    var primeiro bool = true
    for i:= 0; i < len(figurinhasRepetidas); i++{
        if(primeiro == true){
            fmt.Printf("%v", figurinhasRepetidas[i])
            primeiro = false
        } else {
            fmt.Printf(" %v", figurinhasRepetidas[i])
        }
        
    }

    figurinhasNaoAchadas := []int{}
    var achou bool
    fmt.Printf("\n")
    figurinhas := []int{}

    for i := 0; i<len(figurinhasRepetidas); i++{
        for j := 0; j < totalBaruel; j++{
            if figurinhasRepetidas[i] == figurinhasTodas[j]{
                figurinhasTodas[j] = 900 + i
                break
            }
        }
    }
    
    for i := 1; i <= totalFigurinhas; i++{
        for j := 0; j < totalBaruel; j++{
            if figurinhasTodas[j] == i{
                figurinhas = append(figurinhas, i)
            }
        }
    }

    for i := 1; i <= totalFigurinhas; i++{
        achou = false
        for j := 0; j < len(figurinhas); j++{
            if figurinhas[j] == i{
                achou = true
                break
            }
        }
        if achou == false{
            figurinhasNaoAchadas = append(figurinhasNaoAchadas, i)
        }
    }
    if len(figurinhasNaoAchadas) == 0{
        fmt.Print("N")
    }
    primeiro = true
    for i:=0; i < len(figurinhasNaoAchadas); i++{
        if(primeiro == true){
            fmt.Printf("%v", figurinhasNaoAchadas[i])
            primeiro = false
        } else {
            fmt.Printf(" %v", figurinhasNaoAchadas[i])
        }
        
    }
    fmt.Println()

}
