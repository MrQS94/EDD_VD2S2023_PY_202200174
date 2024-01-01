package GenerarArchivos

type SolicitudLogin struct {
	User  string `json:"user"`
	Pass  string `json:"pass"`
	Tutor bool   `json:"tutor"`
}

type SolicitudRuta struct {
	Ruta string `json:"ruta"`
}

type SolicitudLibro struct {
	Carnet    string `json:"carnet"`
	Nombre    string `json:"nombre"`
	Contenido string `json:"contenido"`
}

type SolicitudPublicacion struct {
	Carnet    string `json:"carnet"`
	Contenido string `json:"contenido"`
}

type SolicitudCursos struct {
	Curso []*Cursos `json:"curso"`
}

type Cursos struct {
	Codigo string   `json:"codigo"`
	Post   []string `json:"post"`
}

type SolicitudTutores struct {
	Tutor []*Tutores `json:"tutor"`
}

type Tutores struct {
	Carnet   string `json:"carnet"`
	Nombre   string `json:"nombre"`
	Curso    string `json:"curso"`
	Password string `json:"password"`
}

type SolicitudEstudiante struct {
	Estudiante []*Estudiantes `json:"estudiante"`
}

type Estudiantes struct {
	Carnet   string `json:"carnet"`
	Nombre   string `json:"nombre"`
	Password string `json:"password"`
	Curso1   string `json:"curso1"`
	Curso2   string `json:"curso2"`
	Curso3   string `json:"curso3"`
}

type SolicitudCarnet struct {
	Carnet string `json:"carnet"`
}

type SolicitudActualizar struct {
	Carnet string `json:"carnet"`
	Nombre string `json:"nombre"`
	Estado int    `json:"estado"`
}
