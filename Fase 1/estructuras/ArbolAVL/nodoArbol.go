package ArbolAVL

type Cursos struct {
	Codigo string `json:"Codigo"`
	Nombre string `json:"Nombre"`
}

type NodoArbolAVL struct {
	Curso            *Cursos
	Izquierda        *NodoArbolAVL
	Derecha          *NodoArbolAVL
	Altura           int
	FactorEquilibrio int
}

type CursosData struct {
	Cursos []*Cursos `json:"Cursos"`
}
