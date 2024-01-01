package ArbolB

type RamasArbolB struct {
	Primero  *NodoArbolB
	Hoja     bool
	Contador int
}

func (r *RamasArbolB) Insertar(nuevoNodo *NodoArbolB) {
	if r.Primero == nil {
		r.Primero = nuevoNodo
		r.Contador++
	} else {
		if nuevoNodo.TutorB.Curso < r.Primero.TutorB.Curso {
			nuevoNodo.Sig = r.Primero
			r.Primero.Izq = nuevoNodo.Der
			r.Primero.Ant = nuevoNodo
			r.Primero = nuevoNodo
			r.Contador++
		} else if r.Primero.Sig != nil {
			if r.Primero.Sig.TutorB.Curso > nuevoNodo.TutorB.Curso {
				nuevoNodo.Sig = r.Primero.Sig
				nuevoNodo.Ant = r.Primero
				r.Primero.Sig.Izq = nuevoNodo.Der
				r.Primero.Der = nuevoNodo.Izq
				r.Primero.Sig.Ant = nuevoNodo
				r.Primero.Sig = nuevoNodo
				r.Contador++
			} else {
				aux := r.Primero.Sig
				nuevoNodo.Ant = aux
				aux.Der = nuevoNodo.Izq
				aux.Sig = nuevoNodo
				r.Contador++
			}
		} else if r.Primero.Sig == nil {
			nuevoNodo.Ant = r.Primero
			r.Primero.Der = nuevoNodo.Izq
			r.Primero.Sig = nuevoNodo
			r.Contador++
		}
	}
}
