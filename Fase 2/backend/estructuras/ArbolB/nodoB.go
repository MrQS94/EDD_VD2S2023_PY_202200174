package ArbolB

type NodoArbolB struct {
	TutorB *TutoresB
	Sig    *NodoArbolB
	Ant    *NodoArbolB
	Izq    *RamasArbolB
	Der    *RamasArbolB
}

type TutoresB struct {
	Carnet      int
	Nombre      string
	Curso       string
	Password    string
	Libro       []*Libros
	Publicacion []*Publicaciones
}

type Publicaciones struct {
	Contenido string
}

type Libros struct {
	Nombre    string
	Contenido string
	Estado    int
	Carnet    int
}
