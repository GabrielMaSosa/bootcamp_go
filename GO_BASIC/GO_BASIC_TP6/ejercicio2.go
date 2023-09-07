package main

import "fmt"

const (
	Tam1 = "Pequeño"
	Tam2 = "Mediano"
	Tam3 = "Grande"
)

type SKU struct {
	Amount float64
	Tipo   string
}
type Producto interface {
	Precio() float64
}

func (data *SKU) Precio() float64 {
	switch (*data).Tipo {
	case Tam1:
		return (*data).Amount
	case Tam2:
		return (*data).Amount * 1.03
	case Tam3:
		return (*data).Amount * 1.06
	default:
		return 0.0
	}

}
func NewSKU(tipo string, cost float64) (r *SKU) {
	r = &SKU{
		Amount: cost,
		Tipo:   tipo,
	}
	return

}

func Factory(tipo string, cost float64) Producto {

	return NewSKU(tipo, cost)

}

func Factory2(val *SKU) Producto {

	return val

}

func main() {
	var mount float64 = 1000.0
	var tip string = Tam3
	nsku := Factory(tip, mount)
	fmt.Println(nsku.Precio())

	var val2 SKU

	val2.Amount = mount
	val2.Tipo = tip
	var nsk5 Producto
	nsk5 = Factory2(&val2)
	fmt.Println(nsk5.Precio())

}
