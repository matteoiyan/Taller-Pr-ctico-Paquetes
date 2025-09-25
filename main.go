/*Codigo hecho por mateito*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*funcion para hacer el cambio de dolares a euros multiplicando el cambio actual por el valor de los dolares*/
func Euros(valor int) int {
	/*convertioms el valor entero en un valor flotante para por multiplicar por el cambio */
	return int(float64(valor) * 0.84)
}

/*funcion para hacer el cambio de dolares a libras esterlinas*/
func Libra(valor int) int {
	return int(float64(valor) * 0.74)
}

/*funcion para hacer el cambio de dolares a won*/
func Won(valor int) int {
	return int(float64(valor) * 1402.21)
}

/*funcion para hacer el cambio de dolares a BTC*/
func BTC(valor float64) float64 {
	return ((valor) * 0.0000088)
}

/*funcion para contar las vocales en una frase*/
func contador(palabra string) string {
	cont := 0
	vocales := "aeiouAEIOU"
	slicepalabra := []rune(palabra)

	for _, letra := range slicepalabra {
		if strings.Contains(vocales, string(letra)) {
			cont++
		}
	}

	if cont > 0 {
		return strconv.Itoa(cont)
	} else {
		return "No contiene vocales"
	}
}
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
