package main

import (
	"crypto/sha256"
	"encoding/hex"
	"proyecto/estructuras/ArbolB"
	"proyecto/estructuras/ArbolMerkle"
	"proyecto/estructuras/GenerarArchivos"
	"proyecto/estructuras/Grafos"
	"proyecto/estructuras/TablaHash"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var arbolB *ArbolB.ArbolB = &ArbolB.ArbolB{Raiz: nil, Orden: 3}
var tablaHash *TablaHash.TablaHash = &TablaHash.TablaHash{Tabla: make(map[int]TablaHash.NodoHash), Capacidad: 7, Utilizacion: 0}
var grafo *Grafos.Grafo = &Grafos.Grafo{Cabecera: &Grafos.NodoGrafos{Codigo: "ECYS"}}
var arbolMerkle *ArbolMerkle.ArbolMerkle = nil

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// Login
	app.Post("/login", LoginFiber)

	// Estudiantes
	app.Post("/cargar-estudiantes", CargarEstudiantes)
	app.Get("/tabla-estudiantes", TablaEstudiantes)
	app.Post("/tabla-cursos", TablaCursos)
	app.Post("/devolver-libros-carnet", DevolverLibrosCarnet)
	app.Post("/devolver-publicaciones", DevolverPublicaciones)

	// Tutores, Libros y Publicaciones
	app.Post("/cargar-tutores", CargarTutores)
	app.Post("/cargar-libro", CargarLibro)
	app.Post("/cargar-publicacion", CargarPubli)
	app.Get("/devolver-libros", DevolverLibros)
	app.Post("/actualizar-estado", ActualizarEstado)

	// Cursos
	app.Post("/cargar-cursos", CargarCursos)

	// Reporte
	app.Get("/generar-reporte", Reporte)

	app.Listen(":4000")
}

func ActualizarEstado(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudActualizar
	c.BodyParser(&solicitud)

	carnet, _ := strconv.Atoi(solicitud.Carnet)
	if solicitud.Estado == 1 {
		arbolB.ActualizarEstado(carnet, solicitud.Nombre, solicitud.Estado, arbolB.Raiz.Primero)
	} else {
		arbolB.ActualizarEstado(carnet, solicitud.Nombre, solicitud.Estado, arbolB.Raiz.Primero)
	}

	return c.JSON(&fiber.Map{
		"message": "Se actualizo el estado correctamente.",
		"status":  200,
	})
}

func DevolverLibros(c *fiber.Ctx) error {
	resultEstado := arbolB.DevolverLibros(arbolB.Raiz.Primero)
	var result []*ArbolB.Libros
	for _, libro := range resultEstado {
		if libro.Estado == 0 {
			result = append(result, libro)
		}
	}

	return c.JSON(&fiber.Map{
		"message": "Se actualizo el estado correctamente.",
		"libros":  result,
		"status":  200,
	})
}

func Reporte(c *fiber.Ctx) error {
	arbolB.Graficar()
	grafo.Reporte()
	arbolMerkle = &ArbolMerkle.ArbolMerkle{RaizMerkle: nil, BloqueDatos: nil, CantidadBloques: 0}
	arbolMerkle = arbolB.GraficarMerkle(arbolB.Raiz.Primero, arbolMerkle)
	arbolMerkle.GenerarArbol()
	arbolMerkle.Graficar()

	return c.JSON(&fiber.Map{
		"message": "Se generaron los archivos correctamente.",
		"status":  200,
	})
}

func CargarLibro(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudLibro
	c.BodyParser(&solicitud)

	carnet, _ := strconv.Atoi(solicitud.Carnet)
	arbolB.GuardarLibro(carnet, solicitud.Nombre, solicitud.Contenido, arbolB.Raiz.Primero)

	return c.JSON(&fiber.Map{
		"message": "Se cargaron los libros correctamente.",
		"status":  200,
	})
}

func DevolverPublicaciones(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudCarnet
	c.BodyParser(&solicitud)

	carnet, _ := strconv.Atoi(solicitud.Carnet)
	curso1, curso2, curso3 := tablaHash.BuscarCursos(carnet)

	PublicacionCurso1 := arbolB.BuscarPublicacionesCarnet(curso1, arbolB.Raiz.Primero)
	PublicacionCurso2 := arbolB.BuscarPublicacionesCarnet(curso2, arbolB.Raiz.Primero)
	PublicacionCurso3 := arbolB.BuscarPublicacionesCarnet(curso3, arbolB.Raiz.Primero)

	return c.JSON(&fiber.Map{
		"message":           "Se devolvieron las publicaciones correctamente.",
		"publicacioncurso1": PublicacionCurso1,
		"publicacioncurso2": PublicacionCurso2,
		"publicacioncurso3": PublicacionCurso3,
		"curso1":            curso1,
		"curso2":            curso2,
		"curso3":            curso3,
		"status":            200,
	})
}

func DevolverLibrosCarnet(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudCarnet
	c.BodyParser(&solicitud)

	carnet, _ := strconv.Atoi(solicitud.Carnet)
	curso1, curso2, curso3 := tablaHash.BuscarCursos(carnet)

	LibroCurso1 := arbolB.BuscarLibrosCarnet(curso1, arbolB.Raiz.Primero)
	LibroCurso2 := arbolB.BuscarLibrosCarnet(curso2, arbolB.Raiz.Primero)
	LibroCurso3 := arbolB.BuscarLibrosCarnet(curso3, arbolB.Raiz.Primero)

	return c.JSON(&fiber.Map{
		"message":     "Se devolvieron los libros correctamente.",
		"librocurso1": LibroCurso1,
		"librocurso2": LibroCurso2,
		"librocurso3": LibroCurso3,
		"curso1":      curso1,
		"curso2":      curso2,
		"curso3":      curso3,
		"status":      200,
	})
}

func TablaCursos(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudCarnet
	c.BodyParser(&solicitud)

	carnet, _ := strconv.Atoi(solicitud.Carnet)
	curso1, curso2, curso3 := tablaHash.BuscarCursos(carnet)
	status := 200

	if curso1 == "" && curso2 == "" && curso3 == "" {
		status = 300
	}

	return c.JSON(&fiber.Map{
		"message": "Se enviaron los cursos correctamente.",
		"curso1":  curso1,
		"curso2":  curso2,
		"curso3":  curso3,
		"status":  status,
	})
}

func CargarPubli(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudPublicacion
	c.BodyParser(&solicitud)

	carnet, _ := strconv.Atoi(solicitud.Carnet)
	arbolB.GuardarPubli(carnet, solicitud.Contenido, arbolB.Raiz.Primero)

	return c.JSON(&fiber.Map{
		"message": "Se cargaron las publicaciones correctamente.",
		"status":  200,
	})
}

func TablaEstudiantes(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"message": "Se enviaron los estudiantes correctamente.",
		"tabla":   tablaHash.ConvertirArreglo(),
		"status":  200,
	})
}

func CargarCursos(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudCursos
	c.BodyParser(&solicitud)

	for _, curso := range solicitud.Curso {
		if len(curso.Post) <= 0 {
			grafo.Insertar("ECYS", curso.Codigo)
		} else {
			for _, post := range curso.Post {
				grafo.Insertar(curso.Codigo, post)
			}
		}
	}

	return c.JSON(&fiber.Map{
		"message": "Se cargaron los cursos correctamente.",
		"status":  200,
	})
}

func CargarTutores(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudTutores
	c.BodyParser(&solicitud)

	for _, tutor := range solicitud.Tutor {
		carnet, _ := strconv.Atoi(tutor.Carnet)
		arbolB.Insertar(carnet, tutor.Nombre, tutor.Curso, SHA256(tutor.Password))
	}

	return c.JSON(&fiber.Map{
		"message": "Se cargaron los estudiantes correctamente.",
		"status":  200,
	})
}

func CargarEstudiantes(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudEstudiante
	c.BodyParser(&solicitud)

	for _, estudiante := range solicitud.Estudiante {
		carnet, _ := strconv.Atoi(estudiante.Carnet)
		tablaHash.Insertar(carnet, estudiante.Nombre, SHA256(estudiante.Password), estudiante.Curso1, estudiante.Curso2, estudiante.Curso3)
	}

	return c.JSON(&fiber.Map{
		"message": "Se cargaron los estudiantes correctamente.",
		"status":  200,
	})
}

func LoginFiber(c *fiber.Ctx) error {
	var solicitud GenerarArchivos.SolicitudLogin
	c.BodyParser(&solicitud)

	if solicitud.User == "admin_202200174" && solicitud.Pass == "admin" {
		return c.JSON(&fiber.Map{
			"message": "Se inicio sesión correctamente.",
			"rol":     1,
			"status":  200,
		})
	} else {
		if solicitud.Tutor {
			// Buscar en el Arbol B
			if arbolB.Buscar(solicitud.User, SHA256(solicitud.Pass)) {
				return c.JSON(&fiber.Map{
					"message": "Se inicio sesión correctamente.",
					"rol":     2,
					"status":  200,
				})
			}
		} else {
			// Buscar en la Tabla Hash
			if tablaHash.Buscar(solicitud.User, SHA256(solicitud.Pass)) {
				return c.JSON(&fiber.Map{
					"message": "Se inicio sesión correctamente.",
					"rol":     3,
					"status":  200,
				})
			}
		}
	}

	return c.JSON(&fiber.Map{
		"message": "No se pudo iniciar sesión.",
		"rol":     0,
		"status":  300,
	})
}

func SHA256(text string) string {
	hash := sha256.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}
