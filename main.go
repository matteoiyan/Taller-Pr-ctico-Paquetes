/*Codigo hecho por mateito*/
package main

import (
	"fmt"
)

func main() {
	var valor int
	var frase string
	var opcion int

	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("                          Bienvenido usuario                                    ")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("Escoja una de las opciones")
	fmt.Println("1. Conversión de Dólares a Euros")
	fmt.Println("2. Conversión de Dólares a Libras Esterlinas")
	fmt.Println("3. Conversión de Dólares a Won")
	fmt.Println("4. Conversión de Dólares a BTC")
	fmt.Println("5. Contador de vocales")

	fmt.Print("Ingrese la opción: ")
	fmt.Scanln(&opcion)

	switch opcion {
	case 1:
		fmt.Print("Ingrese un valor en dólares: ")
		fmt.Scanln(&valor)
		fmt.Println("Equivale a:", Euros(valor))
	case 2:
		fmt.Print("Ingrese un valor en dólares: ")
		fmt.Scanln(&valor)
		fmt.Println("Equivale a:", Libra(valor))
	case 3:
		fmt.Print("Ingrese un valor en dólares: ")
		fmt.Scanln(&valor)
		fmt.Println("Equivale a:", Won(valor))
	case 4:
		fmt.Print("Ingrese un valor en dólares: ")
		fmt.Scanln(&valor)
		fmt.Println("Equivale a:", BTC(float64(valor)))
	case 5:
		fmt.Print("Ingrese una frase: ")
		fmt.Scanln(&frase)
		fmt.Println("Número de vocales:", contador(frase))
	default:
		break
	}
}
