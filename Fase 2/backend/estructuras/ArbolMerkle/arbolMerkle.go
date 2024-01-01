package ArbolMerkle

import (
	"encoding/hex"
	"math"
	"proyecto/estructuras/GenerarArchivos"
	"strconv"
	"time"

	"golang.org/x/crypto/sha3"
)

type ArbolMerkle struct {
	RaizMerkle      *NodoMerkle
	BloqueDatos     *NodoBloqueDatos
	CantidadBloques int
}

func fechaActual() string {
	now := time.Now()
	formato := "02-01-2006::15:04:05"
	fechaFormato := now.Format(formato)
	return fechaFormato
}

func (a *ArbolMerkle) AgregarBloque(estado string, nombreLibro string, carnet int) {
	nuevoRegistro := &InformacionBloque{Fecha: fechaActual(), Accion: estado, Nombre: nombreLibro, Tutor: carnet}
	nuevoBloque := &NodoBloqueDatos{Valor: nuevoRegistro}
	if a.BloqueDatos == nil {
		a.BloqueDatos = nuevoBloque
		a.CantidadBloques++
	} else {
		aux := a.BloqueDatos
		for aux.Sig != nil {
			aux = aux.Sig
		}
		nuevoBloque.Ant = aux
		aux.Sig = nuevoBloque
		a.CantidadBloques++
	}
}

func (a *ArbolMerkle) GenerarArbol() {
	nivel := 1
	for int(math.Pow(2, float64(nivel))) < a.CantidadBloques {
		nivel++
	}
	for i := a.CantidadBloques; i < int(math.Pow(2, float64(nivel))); i++ {
		a.AgregarBloque(strconv.Itoa(i), "null", 0)
	}
	a.generarHash()
}

func (a *ArbolMerkle) generarHash() {
	var arrayNodos []*NodoMerkle
	aux := a.BloqueDatos
	for aux != nil {
		concatenacion := aux.Valor.Fecha + aux.Valor.Accion + aux.Valor.Nombre + strconv.Itoa(aux.Valor.Tutor)
		encriptado := a.encriptarSha3(concatenacion)
		nodoHoja := &NodoMerkle{Valor: encriptado, Bloque: aux}
		arrayNodos = append(arrayNodos, nodoHoja)
		aux = aux.Sig
	}
	a.RaizMerkle = a.crearArbol(arrayNodos)
}

func (a *ArbolMerkle) crearArbol(arrayNodos []*NodoMerkle) *NodoMerkle {
	var auxNodos []*NodoMerkle
	var raiz *NodoMerkle
	if len(arrayNodos) == 2 {
		encriptado := a.encriptarSha3(arrayNodos[0].Valor + arrayNodos[1].Valor)
		raiz = &NodoMerkle{Valor: encriptado, Izq: arrayNodos[0], Der: arrayNodos[1]}
		return raiz
	} else {
		for i := 0; i < len(arrayNodos); i += 2 {
			encriptado := a.encriptarSha3(arrayNodos[i].Valor + arrayNodos[i+1].Valor)
			nodoRaiz := &NodoMerkle{Valor: encriptado, Izq: arrayNodos[i], Der: arrayNodos[i+1]}
			auxNodos = append(auxNodos, nodoRaiz)
		}
		return a.crearArbol(auxNodos)
	}
}

func (a *ArbolMerkle) encriptarSha3(cadena string) string {
	hash := sha3.New256()
	hash.Write([]byte(cadena))
	encriptacion := hex.EncodeToString(hash.Sum(nil))
	return encriptacion
}

/*------------------------------------------------------*/
func (a *ArbolMerkle) Graficar() {
	cadena := ""
	nombre_archivo := "./report/ArbolMerkle.dot"
	nombre_imagen := "./report/ArbolMerkle.png"
	if a.RaizMerkle != nil {
		cadena += "digraph G {\nnode [shape = record];\n"
		cadena += a.returnBloques(a.RaizMerkle, 0)
		cadena += "\n}"
	}
	GenerarArchivos.CrearArchivo(nombre_archivo)
	GenerarArchivos.EscribirArchivo(cadena, nombre_archivo)
	GenerarArchivos.Ejecutar(nombre_imagen, nombre_archivo)
}

func (a *ArbolMerkle) returnBloques(raiz *NodoMerkle, indice int) string {
	cadena := ""
	numero := indice + 1
	if raiz != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" [dir=back];\n"
		if raiz.Izq != nil && raiz.Der != nil {
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += a.returnBloques(raiz.Izq, numero)
			cadena += "\""
			cadena += raiz.Valor[:20]
			cadena += "\" -> "
			cadena += a.returnBloques(raiz.Der, numero)
			cadena += "{rank=same" + "\"" + (raiz.Izq.Valor[:20]) + "\"" + " -> " + "\"" + (raiz.Der.Valor[:20]) + "\"" + " [style=invis]}; \n"
		}
	}
	if raiz.Bloque != nil {
		cadena += "\""
		cadena += raiz.Valor[:20]
		cadena += "\" -> "
		cadena += "\""
		cadena += raiz.Bloque.Valor.Fecha + "\n" + raiz.Bloque.Valor.Accion + "\n" + raiz.Bloque.Valor.Nombre + "\n" + strconv.Itoa(raiz.Bloque.Valor.Tutor)
		cadena += "\" [dir=back];\n "
	}
	return cadena
}
