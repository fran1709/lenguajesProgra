package main

import (
	"fmt"
	"strings"
)

func resultado(pChars int, pWords int, pLines int) {
	// Imprimiendo información.
	fmt.Println("Cantidad de caracteres -> ", pChars)
	fmt.Println("Cantidad de palabras -> ", pWords)
	fmt.Println("Cantidad de lineas -> ", pLines)
}

func contadorTexto(pTexto string) {
	spliText := strings.Fields(pTexto)     // Separado en palabras.
	caracters := strings.Split(pTexto, "") // Separando en caracteres.
	spaces := strings.Count(pTexto, " ")
	asks := strings.Count(pTexto, "?")
	exclamtion := strings.Count(pTexto, "!")
	dat := strings.Count(pTexto, ".")
	arrov := strings.Count(pTexto, "@")
	com := strings.Count(pTexto, ",")

	totalWords := len(spliText)
	totalChars := len(caracters) - spaces - asks - exclamtion - dat - arrov - com - strings.Count(pTexto, "\n")
	lines := strings.Count(pTexto, "\n") + 1

	resultado(totalChars, totalWords, lines)
}

func main() {
	var texto = "@"
	texto += "Hola mundo!. Te pregunto,\n" +
		"como estás?"
	contadorTexto(strings.ToLower(texto))
}

/**
-------------------------
|		RESULTADOS		|
-------------------------
Cantidad de caracteres ->  28
Cantidad de palabras ->  6
Cantidad de lineas ->  2
*/
