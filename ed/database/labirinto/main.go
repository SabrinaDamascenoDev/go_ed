package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

// compara duas posições
func (p Pos) equals(other Pos) bool {
	return p.l == other.l && p.c == other.c
}

func (p Pos) getNeighbors() []Pos {
	return []Pos{
		{p.l, p.c - 1},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l + 1, p.c},
	}
}

// Função recursiva para marcar o caminho
func procurarSaida(mat [][]rune, atual, fim Pos) bool {
	nl := len(mat)
	nc := len(mat[0])
	if atual.l < 0 || atual.c < 0 || atual.l >= nl || atual.c >= nc {
		return false
	}
	if atual == fim {
		mat[atual.l][atual.c] = '.'
		return true
	}
	if mat[atual.l][atual.c] != '_' {
		return false
	}

	mat[atual.l][atual.c] = '.'

	for _, viz := range atual.getNeighbors() {
		if procurarSaida(mat, viz, fim) {
			return true
		}
	}

	// Se não for caminho válido, volta a ser '_'
	mat[atual.l][atual.c] = '_'
	return false
}

func removerErros(mat [][]rune) {
	for l := range mat {
		for c := range mat[0] {
			if mat[l][c] == 'x' {
				mat[l][c] = '_'
			}
		}
	}
}

func main() {
	var nl, nc int
	fmt.Scan(&nl, &nc)

	scanner := bufio.NewScanner(os.Stdin)
	mat := make([][]rune, nl)

	// Lê a matriz
	for i := range nl {
		if !scanner.Scan() {
			break
		}
		mat[i] = []rune(scanner.Text())
	}

	// Procura posições de início e fim e conserta para _
	inicio, fim := procurarPosicoes(mat)
	procurarSaida(mat, inicio, fim)
	removerErros(mat)

	// Imprime o labirinto final
	for _, line := range mat {
		fmt.Println(string(line))
	}
}

// procura as posições de início e fim do labirinto
// depois, as substitui por '_'
func procurarPosicoes(mat [][]rune) (Pos, Pos) {
	nl := len(mat)
	nc := len(mat[0])
	inicio := Pos{}
	fim := Pos{}
	for l := range nl {
		for c := range nc {
			switch mat[l][c] {
			case 'I':
				mat[l][c] = '_'
				inicio = Pos{l, c}
			case 'F':
				mat[l][c] = '_'
				fim = Pos{l, c}
			}
		}
	}
	return inicio, fim
}
