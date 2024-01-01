package ColaPrioridad

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"proyecto/estructuras/ArbolAVL"
	"strconv"
)

type ColaPrioridad struct {
	Primero  *NodoCola
	Longitud int
}

func (cola *ColaPrioridad) Queue(carnet int, nombre string, curso string, nota int) {
	var prioridad int
	if nota >= 90 && nota <= 100 {
		prioridad = 1
	} else if nota >= 75 && nota < 90 {
		prioridad = 2
	} else if nota >= 65 && nota < 75 {
		prioridad = 3
	} else if nota >= 60 && nota < 65 {
		prioridad = 4
	} else {
		return
	}

	nuevoTutor := &Tutores{Carnet: carnet, Nombre: nombre, Curso: curso, Nota: nota}
	nuevoNodo := &NodoCola{Tutor: nuevoTutor, Siguiente: nil, Prioridad: prioridad}

	if cola.Primero == nil || nuevoNodo.Prioridad < cola.Primero.Prioridad {
		nuevoNodo.Siguiente = cola.Primero
		cola.Primero = nuevoNodo
		cola.Longitud++
	} else {
		aux := cola.Primero
		for aux.Siguiente != nil && aux.Siguiente.Prioridad <= nuevoNodo.Prioridad {
			aux = aux.Siguiente
		}
		nuevoNodo.Siguiente = aux.Siguiente
		aux.Siguiente = nuevoNodo
		cola.Longitud++
	}
}

func (cola *ColaPrioridad) LeerCSV(ruta string, arbol *ArbolAVL.ArbolAVL) {
	fmt.Println("--------------------------------------------------")
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
			fmt.Println("No pude leer la linea del csv")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		carnet, _ := strconv.Atoi(linea[0])
		nota, _ := strconv.Atoi(linea[3])
		curso := linea[2]

		if arbol.Buscar(curso) {
			cola.Queue(carnet, linea[1], curso, nota)
		} else {
			fmt.Println("El curso " + curso + " no existe.")
		}
	}
	fmt.Println("Carga de Estudiantes Tutores realizada con éxito")
}

func (cola *ColaPrioridad) Dequeue() {
	if cola.Longitud == 0 {
		fmt.Println("La cola está vacía")
	} else {
		cola.Primero = cola.Primero.Siguiente
		cola.Longitud--
	}
}

func (cola *ColaPrioridad) Primero_Cola() {
	if cola.Longitud == 0 {
		fmt.Println("La cola está vacía")
	} else {
		fmt.Println("--------------------------------------------------")
		fmt.Println("Actual: ", cola.Primero.Tutor.Carnet)
		fmt.Println("Nombre: ", cola.Primero.Tutor.Nombre)
		fmt.Println("Curso: ", cola.Primero.Tutor.Curso)
		fmt.Println("Nota: ", cola.Primero.Tutor.Nota)
		fmt.Println("Prioridad: ", cola.Primero.Prioridad)
		if cola.Primero.Siguiente != nil {
			fmt.Println("Siguiente: ", cola.Primero.Siguiente.Tutor.Carnet)
		} else {
			fmt.Println("Siguiente: No hay más tutores por evaluar")
		}
	}
}
