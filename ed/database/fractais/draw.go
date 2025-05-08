package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func tree(p *Pen, tam float64) {
	fator := 0.75
	ang := 35.0

	if tam < 10 {
		if randInt(0, 10) < 5 {
			p.SetColor(255, 0, 0)
			p.DrawCircle(5)
		}
		return
	}
	_, _=fator, ang

	p.SetColor(randInt(0, 255), randInt(0, 255), randInt(0, 255))
	p.SetStrokeWidth(int(tam / 10))
	
	p.Walk(tam)
	p.Rotate(ang)
	tree(p, tam*fator)
	p.Rotate(-2 * ang)
	tree(p, tam*fator)
	p.Rotate(ang)
	p.Walk(-tam)

	//tree(p, tam*fator)
	//p.Rotate(ang)
}


func embua(p *Pen, tam float64){
	if tam < 10{
		return
	}
	p.SetColor(randInt(0, 255), randInt(0, 255), randInt(0, 255))
	p.SetStrokeWidth(max(1, int(tam)/100))
	p.Walk(tam)
	p.Rotate(-90)
	embua(p, tam-3)

}

func circle(p *Pen, raio float64){
	if raio < 1{
		return
	}
	p.SetDown(true)
	p.DrawCircle(raio)
	p.SetDown(false)
	for range 6{
		p.Walk(raio)
		circle(p, raio/3)
		p.Walk(-raio)
		p.Rotate(60)
	}
}

func frozen(p *Pen, raio float64){
	if raio < 1{
		return
	}
	
	for range 5{
		p.Walk(raio)
		frozen(p, raio/3)
		p.Walk(-raio)
		p.Rotate(360/5)
	}
}

func triangulo (p *Pen, tam float64){
	if tam < 5 {
		return 
	}

	for range 3{
		p.Walk(tam)
		p.Rotate(-120)
		triangulo(p, tam/2)


	}
}

func main() {
	pen := NewPen(500, 500)
	//pen.SetColor(255, 0, 0)
	pen.SetPosition(30, 30)
	pen.SetHeading(0)
	//circle(pen, 180)
	//frozen(pen, 160)
	triangulo(pen, 450)
	//tree(pen, 100)
	//embua(pen, 460)
	//pen.SetStrokeWidth(5)
	//tree(pen, 100)
	//pen.Walk(200)
	//pen.Rotate(-60)
	//pen.SetDown(false)
	//pen.Walk(100)
	//pen.SetDown(true)
	//pen.Rotate(60)
	//pen.Walk(100)

	pen.SaveToFile("tree.svg")
	fmt.Println("SVG file created successfully.")
}
