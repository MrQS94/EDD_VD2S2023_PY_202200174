import React, {useState, useEffect} from 'react'

function TablaAlumnos() {
    const [alumnosRegistrados, setAlumnosRegistrados] = useState([]);
    useEffect(() => {
        async function fetchData() {
            const response = await fetch('http://localhost:4000/tabla-estudiantes');
            const result = await response.json();
            setAlumnosRegistrados(result.tabla);
        }
        fetchData();
    }, []);

    return (
        <div className="container col-8 mt-5">
        <div className="text-center">
            <form className="card card-body">
            <h1 className="h3 mb-3 fw-normal">Administrador</h1>
            <br />
            <h4 className="h3 mb-3 fw-normal">Cargar Archivos</h4>
            <br />
            <br />
            <table className="table table-dark table-striped">
                <thead>
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">Posicion</th>
                    <th scope="col">Carnet </th>
                    <th scope="col">Password </th>
                </tr>
                </thead>
                <tbody key={"sd"}>
                    {alumnosRegistrados.map((alumno, i) => {
                        return (
                            <>
                                <tr key={"alumn" + i}>
                                    <th scope="row">{i + 1}</th>
                                    <td>{alumno.Llave}</td>
                                    <td>{alumno.Estudiante.Carnet}</td>
                                    <td>{alumno.Estudiante.Password}</td>
                                </tr>
                            </>
                        );
                    })}
                </tbody>
            </table>
            <br />
            <p className="mt-5 mb-3 text-muted">EDD 202200174</p>
            <br />
            </form>
        </div>
        </div>
    )
    }

export default TablaAlumnos