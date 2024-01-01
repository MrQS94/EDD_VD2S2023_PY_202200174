package ArbolMerkle

type NodoMerkle struct {
	Izq    *NodoMerkle
	Der    *NodoMerkle
	Bloque *NodoBloqueDatos
	Valor  string
}

type NodoBloqueDatos struct {
	Sig   *NodoBloqueDatos
	Ant   *NodoBloqueDatos
	Valor *InformacionBloque
}

type InformacionBloque struct {
	Fecha  string
	Accion string
	Nombre string
	Tutor  int
}
