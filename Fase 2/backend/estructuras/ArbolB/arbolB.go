package ArbolB

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"proyecto/estructuras/ArbolMerkle"
	"proyecto/estructuras/GenerarArchivos"
	"strconv"
)

type ArbolB struct {
	Raiz  *RamasArbolB
	Orden int
}

func (a *ArbolB) Insertar(carnet int, nombre string, curso string, password string) {
	nuevoTutor := &TutoresB{Carnet: carnet, Nombre: nombre, Curso: curso, Password: password}
	nuevoNodo := &NodoArbolB{TutorB: nuevoTutor}
	if a.Raiz == nil {
		a.Raiz = &RamasArbolB{Primero: nil, Hoja: true, Contador: 0}
		a.Raiz.Insertar(nuevoNodo)
	} else {
		obj := a.insertar_rama(nuevoNodo, a.Raiz)
		if obj != nil {
			a.Raiz = &RamasArbolB{Primero: nil, Hoja: true, Contador: 0}
			a.Raiz.Insertar(obj)
			a.Raiz.Hoja = false
		}
	}
}

func (a *ArbolB) GraficarMerkle(raiz *NodoArbolB, arbol *ArbolMerkle.ArbolMerkle) *ArbolMerkle.ArbolMerkle {
	if a.Raiz == nil {
		return nil
	}
	aux := raiz
	for aux != nil {
		if aux.Izq != nil {
			arbol = a.GraficarMerkle(aux.Izq.Primero, arbol)
		}
		for _, libro := range aux.TutorB.Libro {
			estado := ""
			if libro.Estado == 0 {
				estado = "Pendiente"
			} else if libro.Estado == 1 {
				estado = "Aprobado"
			} else {
				estado = "Rechazado"
			}
			arbol.AgregarBloque(estado, libro.Nombre, libro.Carnet)
		}
		if aux.Sig == nil {
			if aux.Der != nil {
				arbol = a.GraficarMerkle(aux.Der.Primero, arbol)
			}
		}
		aux = aux.Sig
	}
	return arbol
}

func (a *ArbolB) GuardarPubli(carnet int, contenido string, raiz *NodoArbolB) {
	if raiz == nil {
		return
	}

	aux := raiz
	for aux != nil {
		if aux.Izq != nil {
			a.GuardarPubli(carnet, contenido, aux.Izq.Primero)
		}
		if aux.TutorB.Carnet == carnet {
			aux.TutorB.Publicacion = append(aux.TutorB.Publicacion, &Publicaciones{Contenido: contenido})
			return
		}
		if aux.Sig == nil {
			if aux.Der != nil {
				a.GuardarPubli(carnet, contenido, aux.Der.Primero)
			}
		}
		aux = aux.Sig
	}
}

func (a *ArbolB) GuardarLibro(carnet int, nombre string, contenido string, raiz *NodoArbolB) {
	if raiz == nil {
		return
	}

	aux := raiz
	for aux != nil {
		if aux.Izq != nil {
			a.GuardarLibro(carnet, nombre, contenido, aux.Izq.Primero)
		}
		if aux.TutorB.Carnet == carnet {
			aux.TutorB.Libro = append(aux.TutorB.Libro, &Libros{Nombre: nombre, Contenido: contenido, Estado: 0, Carnet: carnet})
			return
		}
		if aux.Sig == nil {
			if aux.Der != nil {
				a.GuardarLibro(carnet, nombre, contenido, aux.Der.Primero)
			}
		}
		aux = aux.Sig
	}
}

func (a *ArbolB) BuscarPublicacionesCarnet(curso string, raiz *NodoArbolB) []*Publicaciones {
	if raiz == nil {
		return nil
	}
	var result []*Publicaciones
	aux := raiz
	for aux != nil {
		if aux.Izq != nil {
			result = append(result, a.BuscarPublicacionesCarnet(curso, aux.Izq.Primero)...)
		}
		if aux.TutorB.Curso == curso {
			result = append(result, aux.TutorB.Publicacion...)
		}
		if aux.Sig == nil {
			if aux.Der != nil {
				result = append(result, a.BuscarPublicacionesCarnet(curso, aux.Der.Primero)...)
			}
		}
		aux = aux.Sig
	}
	return result
}

func (a *ArbolB) ActualizarEstado(carnet int, nombre string, estado int, raiz *NodoArbolB) {
	if raiz == nil {
		return
	}
	aux := raiz
	for aux != nil {
		if aux.Izq != nil {
			a.ActualizarEstado(carnet, nombre, estado, aux.Izq.Primero)
		}
		if aux.TutorB.Carnet == carnet {
			for _, libro := range aux.TutorB.Libro {
				if libro.Nombre == nombre {
					libro.Estado = estado
					return
				}
			}
		}
		if aux.Sig == nil {
			if aux.Der != nil {
				a.ActualizarEstado(carnet, nombre, estado, aux.Der.Primero)
			}
		}
		aux = aux.Sig
	}
}

func (a *ArbolB) DevolverLibros(raiz *NodoArbolB) []*Libros {
	if raiz == nil {
		return nil
	}
	var result []*Libros
	aux := raiz
	for aux != nil {
		if aux.Izq != nil {
			result = append(result, a.DevolverLibros(aux.Izq.Primero)...)
		}
		result = append(result, aux.TutorB.Libro...)
		if aux.Sig == nil {
			if aux.Der != nil {
				result = append(result, a.DevolverLibros(aux.Der.Primero)...)
			}
		}
		aux = aux.Sig
	}
	return result
}

func (a *ArbolB) BuscarLibrosCarnet(curso string, raiz *NodoArbolB) []*Libros {
	if raiz == nil {
		return nil
	}
	var result []*Libros
	aux := raiz
	for aux != nil {
		if aux.Izq != nil {
			result = append(result, a.BuscarLibrosCarnet(curso, aux.Izq.Primero)...)
		}
		if aux.TutorB.Curso == curso {
			for _, libro := range aux.TutorB.Libro {
				if libro.Estado == 1 {
					result = append(result, libro)
				}
			}
		}
		if aux.Sig == nil {
			if aux.Der != nil {
				result = append(result, a.BuscarLibrosCarnet(curso, aux.Der.Primero)...)
			}
		}
		aux = aux.Sig
	}
	return result
}

func (a *ArbolB) Buscar(carnetStr string, pass string) bool {
	if a.Raiz == nil {
		return false
	}
	carnet, _ := strconv.Atoi(carnetStr)
	return a.buscarCarnet(a.Raiz.Primero, carnet, pass)
}

func (a *ArbolB) buscarCarnet(raiz *NodoArbolB, carnet int, pass string) bool {
	if raiz == nil {
		return false
	}
	aux := raiz
	for aux != nil {
		if aux.Izq != nil {
			if a.buscarCarnet(aux.Izq.Primero, carnet, pass) {
				return true
			}
		}
		if aux.TutorB.Carnet == carnet && aux.TutorB.Password == pass {
			return true
		}
		if aux.Sig == nil {
			if aux.Der != nil {
				if a.buscarCarnet(aux.Der.Primero, carnet, pass) {
					return true
				}
			}
		}
		aux = aux.Sig
	}
	return false
}

func (a *ArbolB) insertar_rama(nodo *NodoArbolB, rama *RamasArbolB) *NodoArbolB {
	if rama.Hoja {
		rama.Insertar(nodo)
		if rama.Contador == a.Orden {
			return a.dividir(rama)
		} else {
			return nil
		}
	} else {
		temp := rama.Primero
		for temp != nil {
			if nodo.TutorB.Curso == temp.TutorB.Curso {
				return nil
			} else if nodo.TutorB.Curso < temp.TutorB.Curso {
				obj := a.insertar_rama(nodo, temp.Izq)
				if obj != nil {
					rama.Insertar(obj)
					if rama.Contador == a.Orden {
						return a.dividir(rama)
					}
				}
				return nil
			} else if temp.Sig == nil {
				obj := a.insertar_rama(nodo, temp.Der)
				if obj != nil {
					rama.Insertar(obj)
					if rama.Contador == a.Orden {
						return a.dividir(rama)
					}
				}
				return nil
			}
			temp = temp.Sig
		}
	}
	return nil
}

func (a *ArbolB) dividir(rama *RamasArbolB) *NodoArbolB {
	tutor := &TutoresB{Carnet: 0, Nombre: "", Curso: "", Password: ""}
	val := &NodoArbolB{TutorB: tutor}
	aux := rama.Primero
	rder := &RamasArbolB{Primero: nil, Hoja: true, Contador: 0}
	rizq := &RamasArbolB{Primero: nil, Hoja: true, Contador: 0}
	contador := 0
	for aux != nil {
		contador++
		if contador < 2 {
			temp := &NodoArbolB{TutorB: aux.TutorB}
			temp.Izq = aux.Izq
			if contador == 1 {
				temp.Der = aux.Sig.Izq
			}
			if temp.Der != nil && temp.Izq != nil {
				rizq.Hoja = false
			}
			rizq.Insertar(temp)
		} else if contador == 2 {
			val.TutorB = aux.TutorB
		} else {
			temp := &NodoArbolB{TutorB: aux.TutorB}
			temp.Izq = aux.Izq
			temp.Der = aux.Der
			if temp.Der != nil && temp.Izq != nil {
				rder.Hoja = false
			}
			rder.Insertar(temp)
		}
		aux = aux.Sig
	}

	nuevo := &NodoArbolB{TutorB: val.TutorB}
	nuevo.Der = rder
	nuevo.Izq = rizq

	return nuevo
}

/***************************************/
func (a *ArbolB) Graficar() {
	cadena := ""
	nombre_archivo := "./report/ArbolB.dot"
	nombre_imagen := "./report/ArbolB.jpg"
	if a.Raiz != nil {
		cadena += "digraph arbol { \nnode[shape=record]\n"
		cadena += a.grafo(a.Raiz.Primero)
		cadena += a.conexionRamas(a.Raiz.Primero)
		cadena += "}"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolB) grafo(rama *NodoArbolB) string {
	dot := ""
	if rama != nil {
		dot += a.grafoRamas(rama)
		aux := rama
		for aux != nil {
			if aux.Izq != nil {
				dot += a.grafo(aux.Izq.Primero)
			}
			if aux.Sig == nil {
				if aux.Der != nil {
					dot += a.grafo(aux.Der.Primero)
				}
			}
			aux = aux.Sig
		}
	}
	return dot
}

func (a *ArbolB) grafoRamas(rama *NodoArbolB) string {
	dot := ""
	if rama != nil {
		aux := rama
		dot = dot + "R" + rama.TutorB.Curso + "[label=\""
		r := 1
		for aux != nil {
			if aux.Izq != nil {
				dot = dot + "<C" + strconv.Itoa(r) + ">|"
				r++
			}
			if aux.Sig != nil {
				dot = dot + aux.TutorB.Curso + "|"
			} else {
				dot = dot + aux.TutorB.Curso
				if aux.Der != nil {
					dot = dot + "|<C" + strconv.Itoa(r) + ">"
				}
			}
			aux = aux.Sig
		}
		dot = dot + "\"];\n"
	}
	return dot
}

func (a *ArbolB) conexionRamas(rama *NodoArbolB) string {
	dot := ""
	if rama != nil {
		aux := rama
		actual := "R" + rama.TutorB.Curso
		r := 1
		for aux != nil {
			if aux.Izq != nil {
				dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + aux.Izq.Primero.TutorB.Curso + ";\n"
				r++
				dot += a.conexionRamas(aux.Izq.Primero)
			}
			if aux.Sig == nil {
				if aux.Der != nil {
					dot += actual + ":C" + strconv.Itoa(r) + " -> " + "R" + aux.Der.Primero.TutorB.Curso + ";\n"
					r++
					dot += a.conexionRamas(aux.Der.Primero)
				}
			}
			aux = aux.Sig
		}
	}
	return dot
}

func (a *ArbolB) LeerCSV(ruta string) bool {
	file, err := os.Open(ruta)
	if err != nil {
		return false
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
			fmt.Println("No se pudo leer la linea del csv.")
			continue
		}
		if encabezado {
			encabezado = false
			continue
		}
		valor, _ := strconv.Atoi(linea[0])
		a.Insertar(valor, linea[1], linea[2], linea[3])
	}
	return true
}
