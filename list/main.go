package main
import (
	"fmt"
	//importa
	"container/list"
)
func main(){
	// Como inicaliza
	newList := list.New()
	 // como povoar
	for i := range 3{
		newList.PushBack(i)
	}
	
	//Como imprimir pq a lista não tem indece, tem só o lugar namemória que o elemento ta guardado
	for e := newList.Front(); e != nil; e = e.Next(){
		fmt.Print(e.Value)
	}

	//Exemplo de como guardar um elemento de uma lista
	//Cria uma var do tipo *list.Element, pq nessa variável vai ficar armazenado o ponteiro do elemento específico da lista
	var proximoElementoDaList *list.Element
	for e := newList.Front(); e != nil; e = e.Next(){
		if e.Next() != nil {
			proximoElementoDaList = e.Next()
		}
	}

	// tem que colocar o método value para que imprima o valor do elemento e não o ponteiro em que o elemento está 
	fmt.Print(proximoElementoDaList.Value)

	//Como remover 
	if proximoElementoDaList.Value == 2{
		newList.Remove(proximoElementoDaList)
	}

	for e := newList.Front(); e != nil; e = e.Next(){
		fmt.Print(e.Value)
	}

}