package ArbolAVL

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"proyecto/estructuras/GenerarArchivos"
	"strconv"
)

type ArbolAVL struct {
	Raiz *NodoArbolAVL
}

func (arbol *ArbolAVL) getAltura(raiz *NodoArbolAVL) int {
	if raiz == nil {
		return 0
	}
	return raiz.Altura
}

func (arbol *ArbolAVL) getEquilibrio(raiz *NodoArbolAVL) int {
	if raiz == nil {
		return 0
	}
	return arbol.getAltura(raiz.Derecha) - arbol.getAltura(raiz.Izquierda)
}

func (arbol *ArbolAVL) rotacionIzquierda(raiz *NodoArbolAVL) *NodoArbolAVL {
	raiz_derecho := raiz.Derecha
	hijo_izquierdo := raiz_derecho.Izquierda
	raiz_derecho.Izquierda = raiz
	raiz.Derecha = hijo_izquierdo
	// Actualizar alturas
	numeroMax := math.Max(float64(arbol.getAltura(raiz.Izquierda)), float64(arbol.getAltura(raiz.Derecha)))
	raiz.Altura = int(numeroMax) + 1
	raiz.FactorEquilibrio = arbol.getEquilibrio(raiz)
	// Cambiar alturas
	numeroMax = math.Max(float64(arbol.getAltura(raiz_derecho.Izquierda)), float64(arbol.getAltura(raiz_derecho.Derecha)))
	raiz_derecho.Altura = int(numeroMax) + 1
	raiz_derecho.FactorEquilibrio = arbol.getEquilibrio(raiz_derecho)
	return raiz_derecho
}

func (arbol *ArbolAVL) rotacionDerecha(raiz *NodoArbolAVL) *NodoArbolAVL {
	raiz_izquierdo := raiz.Izquierda
	hijo_derecho := raiz_izquierdo.Derecha
	raiz_izquierdo.Derecha = raiz
	raiz.Izquierda = hijo_derecho
	// Actualizar alturas
	numeroMax := math.Max(float64(arbol.getAltura(raiz.Izquierda)), float64(arbol.getAltura(raiz.Derecha)))
	raiz.Altura = int(numeroMax) + 1
	raiz.FactorEquilibrio = arbol.getEquilibrio(raiz)
	// Cambiar alturas
	numeroMax = math.Max(float64(arbol.getAltura(raiz_izquierdo.Izquierda)), float64(arbol.getAltura(raiz_izquierdo.Derecha)))
	raiz_izquierdo.Altura = int(numeroMax) + 1
	raiz_izquierdo.FactorEquilibrio = arbol.getEquilibrio(raiz_izquierdo)
	return raiz_izquierdo
}

func (arbol *ArbolAVL) insertarNodo(raiz *NodoArbolAVL, nuevoNodo *NodoArbolAVL) *NodoArbolAVL {
	if raiz == nil {
		raiz = nuevoNodo
	} else {
		if raiz.Curso.Codigo > nuevoNodo.Curso.Codigo {
			raiz.Izquierda = arbol.insertarNodo(raiz.Izquierda, nuevoNodo)
		} else {
			raiz.Derecha = arbol.insertarNodo(raiz.Derecha, nuevoNodo)
		}
	}
	numeroMax := math.Max(float64(arbol.getAltura(raiz.Izquierda)), float64(arbol.getAltura(raiz.Derecha)))
	raiz.Altura = int(numeroMax) + 1
	balanceo := arbol.getEquilibrio(raiz)
	raiz.FactorEquilibrio = balanceo
	if balanceo > 1 && nuevoNodo.Curso.Codigo > raiz.Derecha.Curso.Codigo {
		// Rotacion Simple a la Izquierda
		return arbol.rotacionIzquierda(raiz)
	} else if balanceo < -1 && nuevoNodo.Curso.Codigo < raiz.Izquierda.Curso.Codigo {
		// Rotacion Simple a la Derecha
		return arbol.rotacionDerecha(raiz)
	} else if balanceo > 1 && nuevoNodo.Curso.Codigo < raiz.Derecha.Curso.Codigo {
		// Rotacion Doble a la Izquierda
		raiz.Derecha = arbol.rotacionDerecha(raiz.Derecha)
		return arbol.rotacionIzquierda(raiz)
	} else if balanceo < -1 && nuevoNodo.Curso.Codigo > raiz.Izquierda.Curso.Codigo {
		// Rotacion Doble a la Derecha
		raiz.Izquierda = arbol.rotacionIzquierda(raiz.Izquierda)
		return arbol.rotacionDerecha(raiz)
	}
	return raiz
}

func (arbol *ArbolAVL) InsertarElemento(codigo string, nombre string) {
	nuevoCurso := &Cursos{Codigo: codigo, Nombre: nombre}
	nuevoNodo := &NodoArbolAVL{Curso: nuevoCurso}
	arbol.Raiz = arbol.insertarNodo(arbol.Raiz, nuevoNodo)
}

func (arbol *ArbolAVL) Reporte(nombre string) {
	cadena := ""
	nombre_archivo := "./out/" + nombre + ".dot"
	nombre_imagen := "./out/" + nombre + ".jpg"
	if arbol.Raiz != nil {
		cadena += "digraph G {\n"
		cadena += arbol.retornarDot(arbol.Raiz, 0)
		cadena += "}\n"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (arbol *ArbolAVL) retornarDot(raiz *NodoArbolAVL, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\"" + raiz.Curso.Codigo + "\"\n"
		if raiz.Izquierda != nil && raiz.Derecha != nil {
			cadena += "x" + strconv.Itoa(numero) + " [label=\"\", width=.1, style=invis];\n"
			cadena += "\"" + raiz.Curso.Codigo + "\" -> "
			cadena += arbol.retornarDot(raiz.Izquierda, numero)
			cadena += "\"" + raiz.Curso.Codigo + "\" -> "
			cadena += arbol.retornarDot(raiz.Derecha, numero)
			cadena += "{rank=same" + "\"" + raiz.Izquierda.Curso.Codigo + "\"" + " -> " + "\"" + raiz.Derecha.Curso.Codigo + "\"" + " [style=invis]};\n"
		} else if raiz.Izquierda != nil && raiz.Derecha == nil {
			cadena += "x" + strconv.Itoa(numero) + " [label=\"\", width=.1, style=invis];\n"
			cadena += "\"" + raiz.Curso.Codigo + "\" -> "
			cadena += arbol.retornarDot(raiz.Izquierda, numero)
			cadena += "\"" + raiz.Curso.Codigo + "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis];\n"
			cadena += "{rank=same" + "\"" + raiz.Izquierda.Curso.Codigo + "\"" + " -> " + "x" + strconv.Itoa(numero) + "[style=invis]};\n"
		} else if raiz.Izquierda == nil && raiz.Derecha != nil {
			cadena += "x" + strconv.Itoa(numero) + " [label=\"\", width=.1, style=invis];\n"
			cadena += "\"" + raiz.Curso.Codigo + "\" -> "
			cadena += arbol.retornarDot(raiz.Derecha, numero)
			cadena += "\"" + raiz.Curso.Codigo + "\" -> "
			cadena += "x" + strconv.Itoa(numero) + "[style=invis];\n"
			cadena += "{rank=same" + "\"" + raiz.Derecha.Curso.Codigo + "\"" + " -> " + "x" + strconv.Itoa(numero) + "[style=invis]};\n"
		}
	}
	return cadena
}

func (arbol *ArbolAVL) LeerJSON(ruta string) {
	file, err := os.ReadFile(ruta)
	if err != nil {
		fmt.Println("Error al abrir el archivo")
		return
	}

	var cursosData CursosData
	err2 := json.Unmarshal(file, &cursosData)
	if err2 != nil {
		fmt.Println("Error al leer el archivo")
		return
	}

	for _, curso := range cursosData.Cursos {
		arbol.InsertarElemento(curso.Codigo, curso.Nombre)
	}
	fmt.Println("Cursos cargados correctamente.")
}

func (arbol *ArbolAVL) Buscar(curso string) bool {
	buscarElemento := arbol.buscarNodo(curso, arbol.Raiz)
	return buscarElemento != nil
}

func (arbol *ArbolAVL) buscarNodo(curso string, raiz *NodoArbolAVL) *NodoArbolAVL {
	var encontrado *NodoArbolAVL
	if raiz != nil {
		if raiz.Curso.Codigo == curso {
			encontrado = raiz
		} else {
			if raiz.Curso.Codigo > curso {
				encontrado = arbol.buscarNodo(curso, raiz.Izquierda)
			} else {
				encontrado = arbol.buscarNodo(curso, raiz.Derecha)
			}
		}
	}
	return encontrado
}
