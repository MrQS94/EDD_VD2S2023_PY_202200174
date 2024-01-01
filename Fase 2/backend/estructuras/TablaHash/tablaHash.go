package TablaHash

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type TablaHash struct {
	Tabla       map[int]NodoHash
	Capacidad   int
	Utilizacion int
}

func (t *TablaHash) reInsertar(capacidadAnt int) {
	auxTabla := t.Tabla
	t.Tabla = make(map[int]NodoHash)
	for i := 0; i < capacidadAnt; i++ {
		if usuario, ok := auxTabla[i]; ok {
			t.Insertar(usuario.Estudiante.Carnet, usuario.Estudiante.Nombre, usuario.Estudiante.Password, usuario.Estudiante.Curso1, usuario.Estudiante.Curso2, usuario.Estudiante.Curso3)
		}
	}
}

func (t *TablaHash) nuevaCapacidad() int {
	contador := 0
	a, b := 0, 1
	for contador < 100 {
		contador++
		if a > t.Capacidad {
			return a
		}
		a, b = b, a+b
	}
	return a
}

func (t *TablaHash) nuevoIndice(indice int) int {
	nuevaPos := 0
	if indice < t.Capacidad {
		nuevaPos = indice
	} else {
		nuevaPos = indice - t.Capacidad
		nuevaPos = t.nuevoIndice(nuevaPos)
	}
	return nuevaPos
}

func (t *TablaHash) reCalculoIndice(carnet int, contador int) int {
	nuevoIndice := t.calculoIndice(carnet) + (contador * contador)
	return t.nuevoIndice(nuevoIndice)
}

func (t *TablaHash) capacidadTabla() {
	auxCap := float64(t.Capacidad) * 0.7
	if t.Utilizacion > int(auxCap) {
		auxAnterior := t.Capacidad
		t.Capacidad = t.nuevaCapacidad()
		t.Utilizacion = 0
		t.reInsertar(auxAnterior)
	}
}

func (t *TablaHash) calculoIndice(carnet int) int {
	var numeros []int
	for {
		if carnet > 0 {
			digito := carnet % 10
			numeros = append([]int{digito}, numeros...)
			carnet /= 10
		} else {
			break
		}
	}

	var numeros_ascii []rune
	for _, numero := range numeros {
		valor := rune(numero + 48)
		numeros_ascii = append(numeros_ascii, valor)
	}
	final := 0
	for _, numero := range numeros_ascii {
		final += int(numero)
	}
	indice := final % t.Capacidad
	return indice
}

func (t *TablaHash) Insertar(carnet int, nombre string, password string, curso1 string, curso2 string, curso3 string) {
	indice := t.calculoIndice(carnet)

	nuevoEstudiante := &Estudiantes{Carnet: carnet, Nombre: nombre, Password: password, Curso1: curso1, Curso2: curso2, Curso3: curso3}
	nuevoNodo := &NodoHash{Llave: indice, Estudiante: nuevoEstudiante}
	if indice < t.Capacidad {
		if _, ok := t.Tabla[indice]; !ok {
			t.Tabla[indice] = *nuevoNodo
			t.Utilizacion++
			t.capacidadTabla()
		} else {
			contador := 1
			indice = t.reCalculoIndice(carnet, contador)
			for {
				if _, ok := t.Tabla[indice]; ok {
					contador++
					indice = t.reCalculoIndice(carnet, contador)
				} else {
					nuevoNodo.Llave = indice
					t.Tabla[indice] = *nuevoNodo
					t.Utilizacion++
					t.capacidadTabla()
					break
				}
			}
		}
	}
}

func (t *TablaHash) Buscar(carnetStr string, pass string) bool {
	carnet, _ := strconv.Atoi(carnetStr)
	indice := t.calculoIndice(carnet)
	if indice < t.Capacidad {
		if usuario, existe := t.Tabla[indice]; existe {
			if usuario.Estudiante.Carnet == carnet && usuario.Estudiante.Password == pass {
				return true
			} else {
				contador := 1
				indice = t.reCalculoIndice(carnet, contador)
				for {
					if us, existe := t.Tabla[indice]; existe {
						if us.Estudiante.Carnet == carnet && us.Estudiante.Password == pass {
							return true
						} else {
							contador++
							indice = t.reCalculoIndice(carnet, contador)
						}
					} else {
						return false
					}
				}
			}
		}
	}
	return false
}

func (t *TablaHash) BuscarCursos(carnet int) (string, string, string) {
	indice := t.calculoIndice(carnet)
	if indice < t.Capacidad {
		if usuario, existe := t.Tabla[indice]; existe {
			if usuario.Estudiante.Carnet == carnet {
				return usuario.Estudiante.Curso1, usuario.Estudiante.Curso2, usuario.Estudiante.Curso3
			} else {
				contador := 1
				indice = t.reCalculoIndice(carnet, contador)
				for {
					if us, existe := t.Tabla[indice]; existe {
						if us.Estudiante.Carnet == carnet {
							return usuario.Estudiante.Curso1, usuario.Estudiante.Curso2, usuario.Estudiante.Curso3
						} else {
							contador++
							indice = t.reCalculoIndice(carnet, contador)
						}
					} else {
						return "", "", ""
					}
				}
			}
		}
	}
	return "", "", ""
}

func (t *TablaHash) ConvertirArreglo() []NodoHash {
	var arrays []NodoHash
	if t.Utilizacion > 0 {
		for i := 0; i < t.Capacidad; i++ {
			if usuario, ok := t.Tabla[i]; ok {
				arrays = append(arrays, usuario)
			}
		}
	}
	return arrays
}

func (t *TablaHash) Mostrar() {
	for i := 0; i < t.Capacidad; i++ {
		if usuario, ok := t.Tabla[i]; ok {
			fmt.Println("Carnet: ", usuario.Estudiante.Carnet, " Nombre: ", usuario.Estudiante.Nombre, " Password: ", usuario.Estudiante.Password, " Curso1: ", usuario.Estudiante.Curso1, " Curso2: ", usuario.Estudiante.Curso2, " Curso3: ", usuario.Estudiante.Curso3)
		}
	}
}

func (t *TablaHash) LeerCSV(ruta string) bool {
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
			return false
		}
		if encabezado {
			encabezado = false
			continue
		}
		carnet, _ := strconv.Atoi(linea[0])
		t.Insertar(carnet, linea[1], linea[2], linea[3], linea[4], linea[5])
	}
	return true
}
