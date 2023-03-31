package main

type Numero float64

func NewNumero(numero float64) *Numero {
	return (*Numero)(&numero)
}

func (n Numero) suma(n2 Numero) float64 {
	return float64(n) + float64(n2)
}

func (n Numero) resta(n2 Numero) float64 {
	return float64(n) - float64(n2)
}

func (n Numero) multiplica(n2 Numero) float64 {
	return float64(n) * float64(n2)
}

func (n Numero) divide(n2 Numero) float64 {
	return float64(n) / float64(n2)
}
