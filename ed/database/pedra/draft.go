package main
import "fmt"
func main() {
    var competidores int
    fmt.Scan(&competidores)

    var valor1, valor2 int 
    var pontuacao = make([]int, competidores)

    var pontuacaoDoCompetidor int 
    for i := 0; i < competidores; i++{
        fmt.Scan(&valor1, &valor2)

        if valor1 >= 10 && valor2 >= 10{ 
            pontuacaoDoCompetidor = valor1 - valor2
            if pontuacaoDoCompetidor<0{
                pontuacaoDoCompetidor=pontuacaoDoCompetidor*-1
            }
            pontuacao[i] = pontuacaoDoCompetidor
        } else {
            pontuacao[i] = 10000
        }
    }

    var menorPontuacao int = pontuacao[0] 
    var competidorVencedor int

    for i := 0; i < competidores; i++{
        if(pontuacao[i] < menorPontuacao){
            menorPontuacao = pontuacao[i]
            competidorVencedor = i
        }
    }

    fmt.Println(competidorVencedor)

}
