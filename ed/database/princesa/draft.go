package main
import (
    "fmt"
    "container/list"
)


func main() {
    var elementos int
    var espada int
    fmt.Scan(&elementos, &espada)

    elementosVet := list.New()

    for i := 1; i <= elementos; i++{
        elementosVet.PushBack(i)
    }
    
    for elementosVet.Len() != 0 {
        fmt.Print("[ ")
        var e *list.Element
        var nextElement *list.Element
        var nextElement2 *list.Element
        var achou bool = false

        for e = elementosVet.Front(); e != nil; e = e.Next(){
            if e.Value == espada {
                achou = true
                if e.Next() != nil{
                    nextElement = e.Next()
                } else {
                    nextElement = elementosVet.Front()
                }
                
                fmt.Printf("%v> ", e.Value)
            } else {
                fmt.Printf("%v ", e.Value)
            }
        }

        if achou {
            nextElement2 = nextElement.Next()
            if nextElement2 != nil {
                espada = nextElement2.Value.(int)
                elementosVet.Remove(nextElement)
            } else {
                nextElement2 = elementosVet.Front()
                espada = nextElement2.Value.(int)
                elementosVet.Remove(nextElement)
            }
            achou = false
        }
        
       
        fmt.Print("]\n")
    }
    
}
