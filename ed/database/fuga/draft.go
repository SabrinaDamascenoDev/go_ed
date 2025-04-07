package main
import "fmt"
func main(){
	var helicoptero, policial, fugitivo, direcao int
    fmt.Scan(&helicoptero, &policial, &fugitivo, &direcao)
    var conseguiu int = 0
	var conseguiuOuNao bool
	var i int = fugitivo

	if direcao == -1{
		for conseguiu != 1{
			if i == helicoptero{
				conseguiu = 1
				conseguiuOuNao = true
				break
			} else if i == policial{
				conseguiu = 1
				conseguiuOuNao = false
				break
			}
			i--
			if(i == 0){
				i=15
			}
		}

    } else if direcao == 1 {
        for conseguiu != 1{
            if i == helicoptero{
                conseguiu = 1
				conseguiuOuNao = true
				break
            }
            if i == policial{
                conseguiu = 1
                conseguiuOuNao = false
                break
            }
            i++
            if(i==15){
                i = 0
            }
        }
    }

    if conseguiuOuNao == true{
        fmt.Println("S")
    } else { 
        fmt.Println("N")
    }   
}
	