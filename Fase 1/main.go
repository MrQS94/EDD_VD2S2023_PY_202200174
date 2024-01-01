package main

import (
	"fmt"
	"proyecto/estructuras/ArbolAVL"
	"proyecto/estructuras/ColaPrioridad"
	"proyecto/estructuras/Listas"
	"proyecto/estructuras/MatrizDispersa"
	"strconv"
)

var listaDobleCircular *Listas.ListaDobleCircular = &Listas.ListaDobleCircular{Inicio: nil, Longitud: 0}
var listaDoble *Listas.ListaDoble = &Listas.ListaDoble{Inicio: nil, Longitud: 0}
var colaPrioridad *ColaPrioridad.ColaPrioridad = &ColaPrioridad.ColaPrioridad{Primero: nil, Longitud: 0}
var matrizDispersa *MatrizDispersa.MatrizDispersa = &MatrizDispersa.MatrizDispersa{Raiz: &MatrizDispersa.NodoMatriz{PosX: -1, PosY: -1, Dato: &MatrizDispersa.Dato{Carnet_Tutor: 0, Carnet_Estudiante: 0, Curso: "RAIZ"}}, Cantidad_Alumnos: 0, Cantidad_Tutores: 0}
var arbolAVL *ArbolAVL.ArbolAVL = &ArbolAVL.ArbolAVL{Raiz: nil}

var user_log string

func main() {
	MenuPrincipal()
}

func MenuPrincipal() {
	var opcion int
	for {
		fmt.Println("--------------------")
		fmt.Println("Login")
		fmt.Println("1. Iniciar sesion")
		fmt.Println("2. Salir")
		fmt.Println("Ingrese la opción que desea: ")
		fmt.Scanln(&opcion)
		fmt.Println("--------------------")
		if opcion == 1 {
			Login()
		} else if opcion == 2 {
			fmt.Println("Saliendo...")
			break
		} else {
			fmt.Println("Opción no valida")
		}
	}
}

func Login() {
	var usuario, pass string
	fmt.Println("Ingrese su usuario: ")
	fmt.Scanln(&usuario)
	fmt.Println("Ingrese su contraseña: ")
	fmt.Scanln(&pass)
	if usuario == "ADMIN_202200174" && pass == "admin" {
		VistaAdministrador()
	} else if listaDoble.Buscar(usuario, pass) {
		user_log = usuario
		VistaEstudiante()
	} else {
		fmt.Println("USUARIO O CONTRASEÑA INCORRECTOS")
	}
}

func VistaAdministrador() {
	var opcion int
	var ruta string
	for {
		fmt.Println("--------------------")
		fmt.Println("1. Carga de Estudiantes Tutores")
		fmt.Println("2. Carga de Estudiantes")
		fmt.Println("3. Cargar Cursos al Sistema")
		fmt.Println("4. Control de Estudiantes Tutores")
		fmt.Println("5. Reportes de Estructuras")
		fmt.Println("6. Salir")
		fmt.Println("--------------------")
		fmt.Println("Ingrese la opción que desea: ")
		fmt.Scanln(&opcion)
		fmt.Println("--------------------")
		if opcion == 1 {
			fmt.Println("Ingrese la ruta del Estudiante Tutores: ")
			fmt.Scanln(&ruta)
			listaDoble.LeerCSV(ruta)
		} else if opcion == 2 {
			fmt.Println("Ingrese la ruta de Estudiantes: ")
			fmt.Scanln(&ruta)
			colaPrioridad.LeerCSV(ruta, arbolAVL)
		} else if opcion == 3 {
			fmt.Println("Ingrese la ruta de Cursos: ")
			fmt.Scanln(&ruta)
			arbolAVL.LeerJSON(ruta)
		} else if opcion == 4 {
			ControlEstudiantes()
		} else if opcion == 5 {
			Reportes()
		} else if opcion == 6 {
			fmt.Println("Saliendo...")
			break
		} else {
			fmt.Println("Opción no valida")
		}
	}
}

func Reportes() {
	if listaDoble.Longitud == 0 {
		fmt.Println("No hay estudiantes registrados.")
	} else {
		listaDoble.Reporte("ListaDoble")
	}

	if listaDobleCircular.Longitud == 0 {
		fmt.Println("No hay tutores registrados.")
	} else {
		listaDobleCircular.Reporte("ListaDobleCircular")
	}

	if matrizDispersa.Cantidad_Alumnos == 0 && matrizDispersa.Cantidad_Tutores == 0 {
		fmt.Println("No hay estudiantes en la matriz dispersa.")
	} else {
		matrizDispersa.Reporte("MatrizDispersa")
	}

	if arbolAVL.Raiz == nil {
		fmt.Println("No hay cursos registrados.")
	} else {
		arbolAVL.Reporte("ArbolAVL")
	}
}

func VistaEstudiante() {
	opcion := 0
	salir := false
	for !salir {
		fmt.Println("-----------------------------")
		fmt.Println("1. Ver Tutores Disponibles")
		fmt.Println("2. Asignarse Tutores")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		if opcion == 1 {
			listaDobleCircular.Mostrar()
		} else if opcion == 2 {
			AsignarCurso()
		} else if opcion == 3 {
			salir = true
		}
	}
}

func AsignarCurso() {
	opcion := ""
	for {
		fmt.Println("-----------------------------")
		fmt.Println("Ingrese el curso que desea: ")
		fmt.Scanln(&opcion)
		fmt.Println("-----------------------------")
		if listaDobleCircular.Buscar(opcion) {
			tutorEncontrado := listaDobleCircular.BuscarTutor(opcion)
			estudiante, err := strconv.Atoi(user_log)
			if err != nil {
				fmt.Println(err)
				break
			}
			matrizDispersa.Insertar(estudiante, tutorEncontrado.Tutor.Carnet, opcion)
			fmt.Println("Se asignó el tutor con éxito.")
		} else {
			fmt.Println("No hay tutores para ese curso.")
			break
		}
	}
}

func ControlEstudiantes() {
	opcion := 0
	salir := false

	if colaPrioridad.Longitud == 0 {
		fmt.Println("No hay estudiantes en cola.")
		return
	}

	for !salir {
		colaPrioridad.Primero_Cola()
		fmt.Println("--------------------")
		fmt.Println("1. Aceptar")
		fmt.Println("2. Rechazar")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)
		fmt.Println("--------------------")
		if opcion == 1 {
			curso := colaPrioridad.Primero.Tutor.Curso
			nota := colaPrioridad.Primero.Tutor.Nota
			carnet := colaPrioridad.Primero.Tutor.Carnet
			nombre := colaPrioridad.Primero.Tutor.Nombre

			aux_tutor := listaDobleCircular.BuscarTutor(curso)
			if aux_tutor == nil {
				listaDobleCircular.Insertar(carnet, nombre, curso, nota)
				colaPrioridad.Dequeue()
				fmt.Println("Se registró tutor con éxito.")
			} else {
				if aux_tutor.Tutor.Nota < nota {
					aux_tutor.Tutor.Carnet = carnet
					aux_tutor.Tutor.Nombre = nombre
					aux_tutor.Tutor.Nota = nota
					fmt.Println("Se sustituyó tutor de curso actual")
					colaPrioridad.Dequeue()
				} else {
					fmt.Println("Tutor rechazado.")
					colaPrioridad.Dequeue()
				}
			}
		} else if opcion == 2 {
			fmt.Println("Tutor rechazado.")
			colaPrioridad.Dequeue()
		} else if opcion == 3 {
			fmt.Println("Saliendo...")
			salir = true
		} else {
			fmt.Println("Opción no valida")
		}
	}
	listaDobleCircular.Mostrar()
}
