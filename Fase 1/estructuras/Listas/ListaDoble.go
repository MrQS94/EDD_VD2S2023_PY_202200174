package Listas

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"proyecto/estructuras/GenerarArchivos"
	"strconv"
)

type ListaDoble struct {
	Inicio   *NodoListaDoble
	Fin      *NodoListaDoble
	Longitud int
}

func (lista *ListaDoble) Insertar(carnet int, nombre string) {
	nuevoAlumno := &Alumno{Carnet: carnet, Nombre: nombre}
	nuevoNodo := &NodoListaDoble{Alumno: nuevoAlumno}
	if lista.Longitud == 0 {
		lista.Inicio = nuevoNodo
		lista.Fin = nuevoNodo
		lista.Longitud++
	} else {
		nuevoNodo.Anterior = lista.Fin
		lista.Fin.Siguiente = nuevoNodo
		lista.Fin = nuevoNodo
		lista.Longitud++
	}
}

func (lista *ListaDoble) Buscar(carnet string, pass string) bool {
	if lista.Longitud == 0 {
		return false
	} else {
		aux := lista.Inicio
		for aux != nil {
			if strconv.Itoa(aux.Alumno.Carnet) == carnet && carnet == pass {
				return true
			}
			aux = aux.Siguiente
		}
	}
	return false
}

func (lista *ListaDoble) LeerCSV(ruta string) {
	fmt.Println("-------------------------------")
	file, err := os.Open(ruta)
	if err != nil {
		fmt.Println("No pude abrir el archivo")
		return
	}
	defer file.Close()

	lectura := csv.NewReader(file)
	lectura.Comma = ','
	encabezado := true
	for {
		linea, err := lectura.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("No pude leer la linea del csv.")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		valor, _ := strconv.Atoi(linea[0])
		lista.Insertar(valor, linea[1])
	}
	fmt.Println("Carga de Estudiantes Tutores exitosa.")
	fmt.Println("-----------------------------------")
}

func (lista *ListaDoble) Reporte(nombre string) {
	if lista.Longitud == 0 {
		fmt.Println("No hay tutores registrados.")
		return
	}
	aux := lista.Inicio
	contador := 0
	texto := "digraph G {\n node[shape=box];\nrankdir=UD;\n{rank=min;\n"
	nombre_archivo := "./out/" + nombre + ".dot"
	nombre_imagen := "./out/" + nombre + ".png"
	for aux != nil {
		texto += "nodo" + strconv.Itoa(contador) + "[label=\"Nombre: " + aux.Alumno.Nombre + ",\\nCarnet: " + strconv.Itoa(aux.Alumno.Carnet) + "\"];\n"
		if contador+1 >= lista.Longitud {
			break
		} else {
			texto += "nodo" + strconv.Itoa(contador) + " -> nodo" + strconv.Itoa(contador+1) + ";\n"
		}
		aux = aux.Siguiente
		contador++
	}
	aux2 := lista.Fin
	contador = 0
	for aux2 != nil {
		if contador+1 >= lista.Longitud {
			break
		} else {
			texto += "nodo" + strconv.Itoa(contador+1) + " -> nodo" + strconv.Itoa(contador) + ";\n"
		}
		aux2 = aux2.Anterior
		contador++
	}

	texto += "}\n}"
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(texto, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}
