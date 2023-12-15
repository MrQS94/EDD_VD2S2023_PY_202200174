package Listas

import (
	"fmt"
	"proyecto/estructuras/GenerarArchivos"
	"strconv"
)

type ListaDobleCircular struct {
	Inicio   *NodoListaCircular
	Longitud int
}

func (lista *ListaDobleCircular) Insertar(carnet int, nombre string, curso string, nota int) {
	nuevoTutor := &Tutores{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoListaCircular{Tutor: nuevoTutor, Siguiente: nil, Anterior: nil}

	if lista.Longitud == 0 {
		lista.Inicio = nuevoNodo
		nuevoNodo.Siguiente = nuevoNodo
		nuevoNodo.Anterior = nuevoNodo
		lista.Longitud++
	} else {
		aux := lista.Inicio
		contador := 1
		for contador < lista.Longitud {
			if lista.Inicio.Tutor.Carnet > carnet {
				nuevoNodo.Siguiente = lista.Inicio
				nuevoNodo.Anterior = lista.Inicio.Anterior
				lista.Inicio.Anterior = nuevoNodo
				lista.Inicio = nuevoNodo
				lista.Longitud++
				return
			}
			if aux.Tutor.Carnet < carnet {
				aux = aux.Siguiente
			} else {
				nuevoNodo.Anterior = aux.Anterior
				aux.Anterior.Siguiente = nuevoNodo
				nuevoNodo.Siguiente = aux
				aux.Anterior = nuevoNodo
				lista.Longitud++
				return
			}
			contador++
		}
		if aux.Tutor.Carnet > carnet {
			nuevoNodo.Siguiente = aux
			nuevoNodo.Anterior = aux.Anterior
			aux.Anterior.Siguiente = nuevoNodo
			aux.Anterior = nuevoNodo
			lista.Longitud++
			return
		}
		nuevoNodo.Anterior = aux
		nuevoNodo.Siguiente = lista.Inicio
		aux.Siguiente = nuevoNodo
		lista.Inicio.Anterior = nuevoNodo
		lista.Longitud++
	}
}

func (lista *ListaDobleCircular) Mostrar() {
	aux := lista.Inicio
	contador := 0
	fmt.Println("-----------------------------")
	for contador < lista.Longitud {
		fmt.Println(aux.Tutor.Curso, " -> ", aux.Tutor.Nombre)
		aux = aux.Siguiente
		contador++
	}
}

func (lista *ListaDobleCircular) Buscar(curso string) bool {
	if lista.Longitud == 0 {
		return false
	} else {
		aux := lista.Inicio
		contador := 0
		for contador < lista.Longitud {
			if aux.Tutor.Curso == curso {
				return true
			}
			aux = aux.Siguiente
			contador++
		}
	}
	return false
}

func (lista *ListaDobleCircular) BuscarTutor(curso string) *NodoListaCircular {
	aux := lista.Inicio
	contador := 0
	for contador < lista.Longitud {
		if aux.Tutor.Curso == curso {
			return aux
		}
		aux = aux.Siguiente
		contador++
	}
	return nil
}

func (lista *ListaDobleCircular) Reporte(nombre string) {
	if lista.Longitud == 0 {
		fmt.Println("No hay tutores registrados.")
		return
	}
	aux := lista.Inicio
	contador := 0
	texto := "digraph G {\n node[shape=box];\nrankdir=UD;\n{rank=min;\n"
	nombre_archivo := "./out/" + nombre + ".dot"
	nombre_imagen := "./out/" + nombre + ".png"
	for contador <= lista.Longitud {
		texto += "nodo" + strconv.Itoa(contador) + "[label=\"Nombre: " + aux.Tutor.Nombre + ",\\nCarnet: " + strconv.Itoa(aux.Tutor.Carnet) + "\"];\n"
		if contador+1 > lista.Longitud {
			texto += "nodo" + strconv.Itoa(0) + " -> nodo" + strconv.Itoa(lista.Longitud) + ";\n"
		} else {
			texto += "nodo" + strconv.Itoa(contador) + " -> nodo" + strconv.Itoa(contador+1) + ";\n"
		}
		aux = aux.Siguiente
		contador++
	}
	contador = 0
	for contador <= lista.Longitud {
		if contador+1 > lista.Longitud {
			texto += "nodo" + strconv.Itoa(lista.Longitud) + " -> nodo" + strconv.Itoa(0) + ";\n"
		} else {
			texto += "nodo" + strconv.Itoa(contador+1) + " -> nodo" + strconv.Itoa(contador) + ";\n"
		}
		aux = aux.Anterior
		contador++
	}
	texto += "}\n}"
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(texto, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}
