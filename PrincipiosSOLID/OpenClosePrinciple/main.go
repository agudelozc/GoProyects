package main
//Las entidades deben estar abiertas para la extensión, pero cerradas para la modificación.
//En este caso, se crean dos tipos de descuentos, uno que no aplica descuento y otro que aplica un descuento porcentual.
//Si se quisiera agregar un nuevo tipo de descuento, se podría hacer sin modificar el código existente.
//Para ello, se crea una interfaz Discount con un método Apply que recibe un precio y devuelve el precio final.
//Se crean dos tipos de descuento, NoDiscount y PercentageDiscount, que implementan la interfaz Discount.
//En el main, se crea un precio y un slice de descuentos, y se recorre el slice aplicando cada descuento al precio.
//Finalmente, se imprime el precio final.
import "fmt"

type Discount interface {
	Apply(price float64) float64
}

type NoDiscount struct{}

func (d NoDiscount) Apply(price float64) float64 {
	return price
}

type PercentageDiscount struct {
	Percentage float64
}

func (d PercentageDiscount) Apply(price float64) float64 {
	return price * (1 - d.Percentage/100)
}

func main() {
	price := 100.0
	discounts := []Discount{NoDiscount{}, PercentageDiscount{Percentage: 10}}

	for _, discount := range discounts {
		fmt.Printf("Final price: %.2f\n", discount.Apply(price))
	}
}
