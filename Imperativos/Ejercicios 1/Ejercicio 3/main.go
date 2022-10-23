package main

import (
	"fmt"
)

/*
*
Rota los elementos una posición hacia la derecha.
*/
func dereRotar(pArray *[10]int, size int) {
	var i int
	var last int = pArray[size-1]
	var aux int
	for i = size - 1; i > 0; i-- {
		if i-1 < 0 {
			break
		}
		aux = pArray[i-1]
		pArray[i] = aux
		//fmt.Println(pArray)
	}
	pArray[0] = last
}

/*
*
Rota los elementos una posición hacia la izquierda.
*/
func izquiRotar(pArray *[10]int, size int) {
	var i int
	var first int = pArray[0]
	var aux int
	for i = 0; i < size; i++ {
		if i+1 == size {
			break
		}
		aux = pArray[i+1]
		pArray[i] = aux
		//fmt.Println(pArray)
	}
	pArray[size-1] = first
}

/*
*
Rota los elementos n cantidad de veces especificadas respectivamente.
*/
func rotacion(pArray *[10]int, size int) {
	var i int
	var cantIzq int = 3
	var cantDer int = 7
	for i = 0; i < cantIzq; i++ {
		izquiRotar(pArray, size)
	}
	for i = 0; i < cantDer; i++ {
		dereRotar(pArray, size)
	}

}

func main() {
	var arrayOne = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	puntero := arrayOne

	rotacion(&puntero, len(arrayOne))
	fmt.Println(arrayOne)
	fmt.Println(puntero)
}

/**
-------------------------
|		RESULTADOS		|
-------------------------
Rotacones hacia la izquierda = 3
Rotaciones hacia la derecha = 7
[0 1 2 3 4 5 6 7 8 9]
[6 7 8 9 0 1 2 3 4 5]
*/
