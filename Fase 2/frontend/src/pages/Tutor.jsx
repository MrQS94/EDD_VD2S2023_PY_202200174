import React from 'react'

function Tutor() {

  const userLocal = localStorage.getItem('user');
  const changeCargarLibro = (e) => {
    e.preventDefault();
    window.open("/principal/tutor/libro", "_self");
  }

  const salir = (e) => {
    e.preventDefault();
    window.open("/", "_self");
  }

  const changeCargarPublicacion = (e) => {
    e.preventDefault();
    window.open("/principal/tutor/publicacion", "_self");
  }

  return (
    <div className="container mt-5 col-3">
      <div className="text-center card card-body">
          <h1 className="h3 mb-3 fw-normal">Bienvenido Tutor, {userLocal}</h1>
          <br />
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-primary"
              onClick={changeCargarLibro}
              >
              Cargar Libro
            </button>
          </center>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-primary"
              onClick={changeCargarPublicacion}
              >
              Cargar Publicación
            </button>
          </center>
          <br />
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-success"
              onClick={salir}
              >
              Salir
            </button>
          </center>
          <p className="mt-5 mb-3 text-muted">EDD 202200174</p>
          <br />
      </div>
    </div>
  )
}

export default Tutor