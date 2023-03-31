package main

import "fmt"

type Operacion struct {
	numero1 Numero
	numero2 Numero
}

func NewOperacion(numero1, numero2 Numero) *Operacion {
	inst := new(Operacion)
	inst.numero1 = numero1
	inst.numero2 = numero2
	return inst
}

func (o Operacion) Sumar() {
	resultado := o.numero1.suma(o.numero2)
	fmt.Println("Resultado: ", o.numero1, " + ", o.numero2, " = ", resultado)
}

func (o Operacion) Restar() {
	resultado := o.numero1.resta(o.numero2)
	fmt.Println("Resultado: ", o.numero1, " - ", o.numero2, " = ", resultado)
}

func (o Operacion) Multiplicar() {
	resultado := o.numero1.multiplica(o.numero2)
	fmt.Println("Resultado: ", o.numero1, " x ", o.numero2, " = ", resultado)
}

func (o Operacion) Dividir() {
	if o.numero2 == 0 {
		fmt.Println("[*] El divisor no puede ser 0")
		return
	}
	resultado := o.numero1.divide(o.numero2)
	fmt.Println("Resultado: ", o.numero1, " / ", o.numero2, " = ", resultado)
}
