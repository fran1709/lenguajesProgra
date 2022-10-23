package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre && (*l)[i].cantidad == 0 {
			result = -2
		} else if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
	var prod = l.buscarProducto(nombre)
	if prod >= 0 {
		(*l)[prod].cantidad += cantidad
		(*l)[prod].precio = precio
		//print("modificado")
	} else {
		lProductos = append(lProductos, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

func (l *listaProductos) venderProducto(nombre string, cant int) {
	//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista"
	var prod = l.buscarProducto(nombre)
	if prod >= 0 && (*l)[prod].cantidad >= cant {
		fmt.Println("Producto vendido")
		(*l)[prod].cantidad = (*l)[prod].cantidad - cant
		if (*l)[prod].cantidad == 0 {
			var slice listaProductos
			slice = append((*l)[:prod], (*l)[prod+1:]...)
			lProductos = slice
		}

	} else if prod == -2 {
		fmt.Println("Producto existencia 0.")
		//fmt.Println("entro aca")
		var slice listaProductos
		slice = append((*l)[:prod], (*l)[prod+1:]...)
		lProductos = slice
	} else {
		fmt.Println("Producto inexistente ó cantidad excede la existente.")
	}
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
}

func (l *listaProductos) listarProductosMinimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var slice listaProductos
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			slice = append(slice, producto{(*l)[i].nombre, (*l)[i].cantidad, (*l)[i].precio})
		}
	}
	fmt.Println("Existentes mínimos")
	fmt.Println(slice)
	return slice
}

func (l *listaProductos) aumentarInventarioDeMinimos() {
	var i int
	var aux int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].cantidad <= existenciaMinima {
			fmt.Println((*l)[i])
			aux = existenciaMinima - (*l)[i].cantidad
			l.agregarProducto((*l)[i].nombre, aux+1, (*l)[i].precio)
		}

	}
}

func (l *listaProductos) ordenarSlices(pSize int) {
	//sort.Slice(l, func(i, j int) bool { return (*l)[i].cantidad < (*l)[j].cantidad })
	var dicc = make(map[string]producto)
	var i int
	for i = 0; i < pSize; i++ {
		dicc[(*l)[i].nombre] = (*l)[i]
	}
	keys := make([]string, 0, len(dicc))
	for k := range dicc {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Ordenados")
	for _, k := range keys {
		fmt.Println(k, dicc[k])
	}
}

func main() {
	llenarDatos()
	fmt.Println(lProductos)
	lProductos.venderProducto("arroz", 15)
	lProductos.venderProducto("arroz", 2)
	fmt.Println(lProductos)
	lMinimos := lProductos.listarProductosMinimos()
	lProductos.agregarProducto("azucar", 20, 1500)
	fmt.Println(lProductos)
	lProductos.venderProducto("frijoles", 4)
	fmt.Println(lProductos)
	lProductos.venderProducto("leche", 10)
	lProductos.agregarProducto("azucar", 10, 1700)
	fmt.Println(lProductos)
	lMinimos.aumentarInventarioDeMinimos()
	fmt.Println("Minimos Aumentados")
	fmt.Println(lMinimos)
	fmt.Println("Desordenados")
	fmt.Println(lProductos)
	lProductos.ordenarSlices(len(lProductos))
}

// si ha sobrado tiempo antes del receso, el ejercicio se podría ampliar para que los los productos se almacenen en
// archivos de texto
// que al inicio se carguen del archivo a la lista y que al final se actualice el archivo con el contenido de la lista

/**
-------------------------
|		RESULTADOS		|
-------------------------
[{arroz 15 2500} {frijoles 4 2000} {leche 8 1200} {café 12 4500}]
Producto vendido
Producto inexistente ó cantidad excede la existente.
[{frijoles 4 2000} {leche 8 1200} {café 12 4500}]
Existentes mínimos
[{frijoles 4 2000} {leche 8 1200}]
[{frijoles 4 2000} {leche 8 1200} {café 12 4500} {azucar 20 1500}]
Producto vendido
[{leche 8 1200} {café 12 4500} {azucar 20 1500}]
Producto inexistente ó cantidad excede la existente.
[{leche 8 1200} {café 12 4500} {azucar 30 1700}]
{frijoles 4 2000}
{leche 8 1200}
Minimos Aumentados
[{frijoles 11 2000} {leche 11 1200}]
Desordenados
[{leche 8 1200} {café 12 4500} {azucar 30 1700}]
Ordenados
azucar {azucar 30 1700}
café {café 12 4500}
leche {leche 8 1200}
*/
