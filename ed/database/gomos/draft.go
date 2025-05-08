package main
import "fmt"

type Gomo struct{
    X int
    Y int
} 

func main() {
    var quantidadeGomos int
    var direcao string

    fmt.Scan(&quantidadeGomos, &direcao)

     var posicaoCobra []Gomo

     // Povoando array com pares x, y dos valores da posição de cada gomo da cobrinha
    for i := 0; i < quantidadeGomos; i++{
        var x, y int
        fmt.Scan(&x, &y)
        posicaoCobra = append(posicaoCobra, Gomo{X:x, Y:y})
    }
    
    var primeiraPosição Gomo = posicaoCobra[0]

    if direcao == "L"{
        posicaoCobra[0].X = posicaoCobra[0].X-1
    } else if direcao == "R"{
        posicaoCobra[0].X = posicaoCobra[0].X+1
    } else if direcao == "U"{
        posicaoCobra[0].Y = posicaoCobra[0].Y-1
    } else {
        posicaoCobra[0].Y = posicaoCobra[0].Y+1
    }

    for i := 1; i < len(posicaoCobra); i++{
        var posicaoSeguinte Gomo = posicaoCobra[i]
        posicaoCobra[i] = primeiraPosição
        primeiraPosição = posicaoSeguinte
    }

    for _, gomo := range posicaoCobra{
        fmt.Printf("%v %v\n", gomo.X, gomo.Y)
    }
}
