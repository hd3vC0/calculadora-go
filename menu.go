package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func mostrarMenu() {

	fmt.Println("=== Calculadora ===")
	fmt.Println("Seleccione la operación que desea realizar:")
	fmt.Println("\t1. Suma")
	fmt.Println("\t2. Resta")
	fmt.Println("\t3. Multiplicación")
	fmt.Println("\t4. División")
	fmt.Println("\t5. Convertidor de base")
	fmt.Println("\t9. Salir")
	fmt.Print("?: ")
}

func leerBase(mesnaje string) int {
	for {
		fmt.Print(mesnaje, ": ")
		scanner.Scan()
		base := scanner.Text()
		baseInt, err := strconv.ParseInt(base, 10, 0)

		if err != nil {
			fmt.Println("[*] valor invalido")
			continue
		}

		return int(baseInt)
	}
}

func leerTexto(mensaje string) string {
	for {
		fmt.Print(mensaje, ": ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			fmt.Println("[*] No ingreso ningun valor")
		}

		return input
	}

}

func leerNumero() Numero {
	for {
		fmt.Print("Ingrese número: ")
		scanner.Scan()
		input := scanner.Text()
		numero, err := strconv.ParseFloat(input, 0)

		if err != nil {
			fmt.Println("[*] El número que ingresó es invalido")
			continue
		}

		return *NewNumero(float64(numero))
	}
}

func flujoMenu() {

	for {
		mostrarMenu()
		scanner.Scan()
		opcion := scanner.Text()
		intval, err := strconv.ParseInt(opcion, 10, 0)

		if err != nil {
			fmt.Println("[*] El valor ingresado es invalido")
		}

		switch intval {
		case 1:
			numero1 := leerNumero()
			numero2 := leerNumero()
			NewOperacion(numero1, numero2).Sumar()
		case 2:
			numero1 := leerNumero()
			numero2 := leerNumero()
			NewOperacion(numero1, numero2).Restar()
		case 3:
			numero1 := leerNumero()
			numero2 := leerNumero()
			NewOperacion(numero1, numero2).Multiplicar()
		case 4:
			numero1 := leerNumero()
			numero2 := leerNumero()
			NewOperacion(numero1, numero2).Dividir()
		case 5:
			base1 := leerBase("Base origen")
			base2 := leerBase("Base destino")
			numero := leerTexto("Ingrese número")
			ob, err := NewOperacionBase(base1, numero, base2)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				ob.Convertir()
			}
		case 9:
			fmt.Println("Bye bye")
			os.Exit(0)
		default:
			fmt.Println("[*] Opcion invalida :(")
		}

	}
}
