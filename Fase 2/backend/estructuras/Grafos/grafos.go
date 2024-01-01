package Grafos

import (
	"encoding/json"
	"fmt"
	"os"
	"proyecto/estructuras/GenerarArchivos"
)

type Grafo struct {
	Cabecera *NodoGrafos
}

type Cursos struct {
	Codigo        string   `json:"Codigo"`
	PostRequisito []string `json:"Post"`
}

type DatosCurso struct {
	Curso []Cursos `json:"Cursos"`
}

func (g *Grafo) insertarFila(codigo string) {
	nuevoNodo := &NodoGrafos{Codigo: codigo}
	if g.Cabecera == nil {
		g.Cabecera = nuevoNodo
	} else {
		aux := g.Cabecera
		for aux.Abajo != nil {
			if aux.Codigo == codigo {
				return
			}
			aux = aux.Abajo
		}
		aux.Abajo = nuevoNodo
	}
}

func (g *Grafo) insertarColumna(codigo string, post string) {
	nuevoNodo := &NodoGrafos{Codigo: post}
	if g.Cabecera != nil && g.Cabecera.Codigo == codigo {
		g.insertarFila(post)
		aux := g.Cabecera
		for aux.Sig != nil {
			aux = aux.Sig
		}
		aux.Sig = nuevoNodo
	} else {
		g.insertarFila(codigo)
		aux := g.Cabecera
		for aux != nil {
			if aux.Codigo == codigo {
				break
			}
			aux = aux.Abajo
		}
		if aux != nil {
			for aux.Sig != nil {
				aux = aux.Sig
			}
			aux.Sig = nuevoNodo
		}
	}
}

func (g *Grafo) Insertar(codigo string, post string) {
	if g.Cabecera == nil {
		g.insertarFila(codigo)
		g.insertarColumna(codigo, post)
	} else {
		g.insertarColumna(codigo, post)
	}
}

func (g *Grafo) LeerJSON(ruta string) bool {
	data, err := os.ReadFile(ruta)
	if err != nil {
		fmt.Println("Error al leer el archivo")
		return false
	}
	var datos DatosCurso
	err = json.Unmarshal(data, &datos)
	if err != nil {
		fmt.Println("Error al convertir el archivo")
		return false
	}
	for _, curso := range datos.Curso {
		if len(curso.PostRequisito) > 0 {
			for i := 0; i < len(curso.PostRequisito); i++ {
				g.Insertar(curso.Codigo, curso.PostRequisito[i])
			}
		} else {
			g.Insertar("ECYS", curso.Codigo)
		}
	}
	return true
}

func (g *Grafo) Reporte() {
	cadena := ""
	nombre_archivo := "./report/Grafos.dot"
	nombre_imagen := "./report/Grafos.jpg"
	if g.Cabecera != nil {
		cadena += "digraph grafoDirigido{ \n rankdir=LR; \n node [shape=box]; layout=neato; \n nodo" + g.Cabecera.Codigo + "[label=\"" + g.Cabecera.Codigo + "\"]; \n"
		cadena += "node [shape = ellipse]; \n"
		cadena += g.retornarValoresMatriz()
		cadena += "\n}"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (g *Grafo) retornarValoresMatriz() string {
	cadena := ""
	/*CREACION DE NODOS*/
	aux := g.Cabecera.Abajo //Filas
	aux1 := aux             //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux1.Codigo + "[label=\"" + aux1.Codigo + "\" ]; \n"
			aux1 = aux1.Sig
		}
		if aux != nil {
			aux = aux.Abajo
			aux1 = aux
		}
	}
	/*CONEXION DE NODOS*/
	aux = g.Cabecera //Filas
	aux1 = aux.Sig   //Columnas
	/*CREACION DE NODOS CON LABELS*/
	for aux != nil {
		for aux1 != nil {
			cadena += "nodo" + aux.Codigo + " -> "
			cadena += "nodo" + aux1.Codigo + "[len=1.00]; \n"
			aux1 = aux1.Sig
		}
		if aux.Abajo != nil {
			aux = aux.Abajo
			aux1 = aux.Sig
		} else {
			aux = aux.Abajo
		}
	}
	return cadena
}
