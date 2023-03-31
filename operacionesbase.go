package main

import (
	"errors"
	"fmt"
	"strconv"
)

type OperacionBase struct {
	baseOrigen   int
	numeroOrigen string
	baseDestino  int
}

func NewOperacionBase(baseOrigen int, numero string, baseDestino int) (*OperacionBase, error) {
	if (baseDestino < 2 || baseDestino > 36) && (baseOrigen < 2 || baseOrigen > 36) {
		return nil, errors.New("[*] El valor de las bases debe estar entre 2 y 36")
	}

	inst := &OperacionBase{
		baseOrigen:   baseOrigen,
		numeroOrigen: numero,
		baseDestino:  baseDestino,
	}

	return inst, nil
}

func (o OperacionBase) Convertir() {
	valint, err := strconv.ParseInt(o.numeroOrigen, o.baseOrigen, 64)

	if err != nil {
		fmt.Println("[*] número ingresado invalido.")
	} else {
		val2base := strconv.FormatInt(valint, o.baseDestino)
		fmt.Println("Resultado: ", o.numeroOrigen, " base ", o.baseOrigen, " a base ", o.baseDestino, " = ", val2base)
	}
}
