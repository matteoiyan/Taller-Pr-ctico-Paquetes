package converciones

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
