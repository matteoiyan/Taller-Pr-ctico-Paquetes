package utilidades

import (
	"strconv"
	"strings"
)

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
