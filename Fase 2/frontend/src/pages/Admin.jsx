import React from 'react'
import Papa from 'papaparse'

function Admin() {

    const salir = (e) => {
        e.preventDefault();
        window.open("/", "_self");
    }

    const changeLibrosAdmin = (e) => {
        e.preventDefault();
        window.open("/principal/admin/libros-admin", "_self");
    }

    const changeTablaAlumnos = (e) => {
        e.preventDefault();
        window.open("/principal/admin/tabla-alumnos", "_self");
    }

    const generarReporte = async (e) =>{
        e.preventDefault();
        await fetch('http://localhost:4000/generar-reporte')
        alert('Reporte Generado Correctamente');
    }

    const uploadFileEstudiantes = async (event) => {
        const file = event.target.files[0];
        Papa.parse(file, {
            complete: (result) => {
                fetch ('http://localhost:4000/cargar-estudiantes', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        estudiante : result.data,
                    }),
                })
            },
            header: true,
            skipEmptyLines: true,
        });
    };

    const uploadFileTutores = async (event) => {
        const file = event.target.files[0];
        Papa.parse(file, {
            complete: (result) => {
                fetch ('http://localhost:4000/cargar-tutores', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        Tutor : result.data,
                    }),
                })
            },
            header: true,
            skipEmptyLines: true,
        });
    };

    const uploadFileCursos = async (event) => {
        var reader = new FileReader();
        reader.onload = async function(event) {
            var obj = JSON.parse(event.target.result);
            console.log(obj.Cursos);
            
            const response = await fetch('http://localhost:4000/cargar-cursos', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    Curso : obj.Cursos,
                }),
            })
        };
        
        reader.readAsText(event.target.files[0]);
        alert('Cursos Cargados Correctamente');
    }

    return (
        <div className="container mt-5 col-4">
            <div className="text-center">
                <form className="card card-body">
                <h1 className="h3 mb-3 fw-normal">Administrador</h1>
                <br />
                <h4 className="h3 mb-3 fw-normal">Cargar Archivos</h4>
                <br />
                <div className="input-group mb-3">
                    <label className="input-group-text"> Cargar Estudiantes | TablaHash </label>
                    <input
                    className="form-control"
                    id="inputGroupFile01"
                    type="file"
                    accept=".csv"
                    onChange={uploadFileEstudiantes}
                    />
                </div>
                <br />
                <div className="input-group mb-3">
                    <label className="input-group-text"> Cargar Tutores | ArbolB </label>
                    <input
                    className="form-control"
                    id="inputGroupFile02"
                    type="file"
                    accept=".csv"
                    onChange={uploadFileTutores}
                    />
                </div>
                <br />
                <div className="input-group mb-3">
                    <label className="input-group-text">Cargar Cursos | Grafos</label>
                    <input
                    className="form-control"
                    id="inputGroupFile02"
                    type="file"
                    accept="application/json"
                    onChange={uploadFileCursos}
                    />
                </div>
                <br />
                <center>
                    <button className="w-50 btn btn-outline-primary"
                    onClick={generarReporte}
                    >
                    Reportes
                    </button>
                </center>
                <br />
                <center>
                    <button className="w-50 btn btn-outline-primary"
                    onClick={changeTablaAlumnos}
                    >
                    Tabla Alumnos
                    </button>
                </center>
                <br />
                <center>
                    <button className="w-50 btn btn-outline-primary" 
                    onClick={changeLibrosAdmin}
                    >
                    Aceptar Libros
                    </button>
                </center>
                <br />
                <center>
                    <button className="w-50 btn btn-outline-success" 
                    onClick={salir}
                    >
                    Salir
                    </button>
                </center>
                <p className="mt-5 mb-3 text-muted">EDD 202200174</p>
                <br />
                </form>
            </div>
        </div>
    )
}

export default Admin