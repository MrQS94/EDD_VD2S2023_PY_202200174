import React, { useEffect, useState} from 'react'

function LibrosAdmin() {
    const [libros, setLibros] = useState([]);
    const [contenido, setContenido] = useState([]);
    const [carnet1, setCarnet1] = useState("");

    useEffect(() => {
        async function fetchData() {
            const response = await fetch("http://localhost:4000/devolver-libros")
            const data = await response.json()
            setLibros(data.libros)    
        }
        fetchData();
    }, []);

    const handleContenido = (selectedLibro) => {
        const libroSeleccionado = libros.find((libro) => libro.Nombre === selectedLibro);
        if (libroSeleccionado) {
            setCarnet1(libroSeleccionado.Carnet);
            return libroSeleccionado.Contenido;
        }
        return '';
    };

    const handleOnChange = (e) => {
        const selectedLibro = e.target.value;
        const selectedContenido = handleContenido(selectedLibro);
        setContenido(selectedContenido);
    };

    const handleAceptar = async (e) => {
        var opcion = document.getElementById("opcionSeleccionada").value;

        const response = await fetch("http://localhost:4000/actualizar-estado", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                carnet : String(carnet1),
                nombre : opcion,
                estado : 1,
            }),
        });
        alert("Libro aceptado");
    }

    const handleRechazar = async () => {
        var opcion = document.getElementById("opcionSeleccionada").value;

        await fetch("http://localhost:4000/actualizar-estado", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                carnet : String(carnet1),
                nombre : opcion,
                estado : 2,
            }),
        });
        alert("Libro rechazado");
        fetchData();
    }

    const volver = (e) => {
        e.preventDefault()
        window.open("/principal/admin", "_self")
    }

    return (
        <div className="container col-6 mt-5">
            <div className="text-center">
                <form className="card card-body">
                    <h2 className="mb-3 fw-normal">Bienvenido Administrador</h2>
                    <br />
                    <br />
                    <div className="container col-4">
                        <select className="form-control" onChange={handleOnChange} id="opcionSeleccionada">
                            <option disabled selected value="" >Seleccione una opción:</option>
                            {libros && libros.map((libro, i) => {
                                return (
                                    <>
                                        <option key={`libro${i}`}>{libro.Nombre}</option>
                                    </>
                                );
                            })}
                        </select>
                    </div>
                    <br />
                    <br />
                    <center>
                        <iframe 
                        src={contenido} 
                        width="600"
                        height="400"
                        title="Contenido del libro"
                        />
                    </center>
                    <div className="mt-3">
                        <button type="button" className="btn btn-success me-2" 
                        onClick={handleAceptar}
                        >
                            Aceptar
                        </button>
                        <button type="button" className="btn btn-success me-2"
                        onClick={handleRechazar}
                        >
                            Rechazar
                        </button>
                    </div>
                    <div className="mt-3">
                        <button className="w-25 btn btn-outline-success"
                        onClick={volver}
                        >
                            Volver
                        </button>
                    </div>
                </form>
            </div>
        </div>
    )
}

export default LibrosAdmin