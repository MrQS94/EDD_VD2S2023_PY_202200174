import React, { useEffect, useState } from 'react'

function Estudiante() {

  const [cursos, setCursos] = useState([]);
  const userLocal = localStorage.getItem('user');

  const salir = (e) => {
    e.preventDefault();
    window.open("/", "_self");
  }

  const changeCargarLibro = (e) => {
    e.preventDefault();
    window.open("/principal/estudiante/libro", "_self");  
  }
  
  const changeCargarPublicacion = (e) => {
    e.preventDefault();
    window.open("/principal/estudiante/publicacion", "_self");  
  }

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:4000/tabla-cursos', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            Carnet: userLocal,
          }),
        });
        const result = await response.json();
        setCursos([result.curso1, result.curso2, result.curso3]);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };
    fetchData(); 
  }, [userLocal]);

  return (
    <div className="container mt-5 col-3">
      <div className="text-center card card-body">
          <h1 className="h3 mb-3 fw-normal">Bienvenido Estudiante, {userLocal}</h1>
          <br />
          <br />
          <table className="table table-dark table-striped">
              <thead>
                <tr>
                  <th scope="col">Curso 1</th>
                  <th scope="col">Curso 2</th>
                  <th scope="col">Curso 3</th>
                </tr>
              </thead>
              <tbody key={"sd"}>
                <tr>
                    <td>{cursos[0]}</td>
                    <td>{cursos[1]}</td>
                    <td>{cursos[2]}</td>
                </tr>
              </tbody>
          </table>
          <br />
          <br />
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-primary"
              onClick={changeCargarLibro}
              >
              Ver Libros
            </button>
          </center>
          <br />
          <center>
            <button
              className="w-50 btn btn-outline-primary"
              onClick={changeCargarPublicacion}
              >
              Ver Publicaciones
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

export default Estudiante