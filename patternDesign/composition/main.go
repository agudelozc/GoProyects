package main

import "fmt"

type liquid struct {
	color string
}

func (l *liquid) isColor(){
	fmt.Println(l.color)
}

type beer struct{
	name string
	alcohol float32
	liquid liquid
}

func (b *beer) isBeer(){
	fmt.Println("Mi cerveza es", b.name, "y tiene", b.alcohol, "de alcohol", "y es de color", b.liquid.color)
}

func main() {
	l := liquid{color: "red"}
	b := beer{name: "Cerveza", alcohol: 5.5, liquid: l}
	b.isBeer()
}