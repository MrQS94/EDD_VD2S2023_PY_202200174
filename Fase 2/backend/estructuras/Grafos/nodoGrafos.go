package Grafos

type NodoGrafos struct {
	Sig    *NodoGrafos
	Abajo  *NodoGrafos
	Codigo string
}
