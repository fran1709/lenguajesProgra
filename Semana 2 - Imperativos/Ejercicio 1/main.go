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

	totalWords := len(spliText)
	totalChars := len(caracters) - spaces - asks - exclamtion - dat - arrov
	lines := strings.Count(pTexto, "\n")

	resultado(totalChars, totalWords, lines)
}

/**
func repetition(pTexto string) map[string]int {

	// using strings.Field Function
	spliText := strings.Fields(pTexto)
	dictionary := make(map[string]int)

	for _, word := range spliText {
		_, matched := dictionary[word]
		if matched {
			dictionary[word] += 1
		} else {
			dictionary[word] = 1
		}
	}
	return dictionary
}*/

func main() {
	var texto = "@"
	texto += "Hola mundo!. Te pregunto,\n" +
		"como estás?"
	contadorTexto(strings.ToLower(texto))
}
