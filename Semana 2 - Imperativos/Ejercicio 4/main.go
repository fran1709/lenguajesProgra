package main

import (
	"fmt"
	"strings"
)

type Calzado struct {
	marca    string
	precio   int
	talla    int
	cantidad int
}
type ListaZapatos []Calzado

var lCalzados ListaZapatos

/*
*
Busca el elemento de la lista, si lo encuentra retorna el indice, sino -1.
*/
func (pLista *ListaZapatos) buscarProducto(pMarca string, pTalla int) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*pLista); i++ {
		if strings.ToLower((*pLista)[i].marca) == strings.ToLower(pMarca) && (*pLista)[i].talla == pTalla && (*pLista)[i].cantidad == 0 {
			result = -2
		} else if strings.ToLower((*pLista)[i].marca) == strings.ToLower(pMarca) && (*pLista)[i].talla == pTalla {
			result = i
		}
	}
	return result
}

/*
*Elimina de stock la cantidad de elementos solicitados si existe la cantidad solicitado sino notifica.
 */
func (pLista *ListaZapatos) venderCalzado(pMarca string, pCantidad int, pTallla int) {
	var prod = pLista.buscarProducto(pMarca, pTallla)
	if prod >= 0 && (*pLista)[prod].cantidad >= pCantidad {
		fmt.Println("Producto vendido")
		(*pLista)[prod].cantidad = (*pLista)[prod].cantidad - pCantidad

	} else if prod == -2 {
		fmt.Println("Producto existencia 0.")
	} else {
		fmt.Println("Producto inexistente ó cantidad excede la existente.")
	}
}

/*
*
Agrega un nuevo elemento "calzado" a la lista de calzandos->lCalzados
*/
func (pLista *ListaZapatos) agregarCalzado(pMarca string, pPrecio int, pTalla int, pCantidad int) {
	lCalzados = append(lCalzados, Calzado{marca: pMarca, precio: pPrecio, talla: pTalla, cantidad: pCantidad})
}

func cargarDatos() {
	lCalzados.agregarCalzado("Nike", 40000, 42, 2)
	lCalzados.agregarCalzado("Nike", 35000, 37, 4)
	lCalzados.agregarCalzado("Adidas", 43000, 41, 3)
	lCalzados.agregarCalzado("Adidas", 45000, 38, 5)
	lCalzados.agregarCalzado("Mizuno", 65000, 44, 1)
	lCalzados.agregarCalzado("Mizuno", 60000, 34, 3)
	lCalzados.agregarCalzado("Procs", 20000, 36, 2)
	lCalzados.agregarCalzado("Procs", 32000, 39, 3)
	lCalzados.agregarCalzado("Euros", 20000, 42, 4)
	lCalzados.agregarCalzado("Euros", 30000, 35, 6)
	lCalzados.agregarCalzado("Converse", 49000, 40, 7)
	lCalzados.agregarCalzado("Converse", 50000, 43, 3)
}

func main() {
	cargarDatos()

	fmt.Println(lCalzados)

	lCalzados.venderCalzado("mizuno", 1, 44) //vende con exito
	lCalzados.venderCalzado("mizuno", 1, 44) //excede, no vende y notifica existencia 0.

	fmt.Println(lCalzados)
}

/**
-------------------------
|		RESULTADOS		|
-------------------------
[{Nike 40000 42 2} {Nike 35000 37 4} {Adidas 43000 41 3} {Adidas 45000 38 5} {Mi
zuno 65000 44 1} {Mizuno 60000 34 3} {Procs 20000 36 2} {Procs 32000 39 3} {Euro
s 20000 42 4} {Euros 30000 35 6} {Converse 49000 40 7} {Converse 50000 43 3}]
Producto vendido
Producto existencia 0.
[{Nike 40000 42 2} {Nike 35000 37 4} {Adidas 43000 41 3} {Adidas 45000 38 5} {Mi
zuno 65000 44 0} {Mizuno 60000 34 3} {Procs 20000 36 2} {Procs 32000 39 3} {Euro
s 20000 42 4} {Euros 30000 35 6} {Converse 49000 40 7} {Converse 50000 43 3}]
*/
