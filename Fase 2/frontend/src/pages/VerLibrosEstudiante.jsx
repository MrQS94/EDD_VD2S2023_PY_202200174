import React, {useState } from 'react'

function VerLibrosEstudiante() {
    const [libro1, setLibro1] = useState([]);
    const [libro2, setLibro2] = useState([]);
    const [libro3, setLibro3] = useState([]);
    const [curso1, setCurso1] = useState("");
    const [curso2, setCurso2] = useState("");
    const [curso3, setCurso3] = useState("");
    const userLocal = localStorage.getItem("user");
    const [libro, setLibro] = useState("");
    
    
    const fetchData = async() => {
        try {
            const response = await fetch("http://localhost:4000/devolver-libros-carnet", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                        carnet : userLocal,
                }),
            });

            const result = await response.json();   

            setCurso1(result.curso1);
            setCurso2(result.curso2);
            setCurso3(result.curso3);
            
            let libros1 = [];
            let libros2 = [];
            let libros3 = [];

            if (result.librocurso1 != null) {
                libros1 = result.librocurso1.map((libro) => ({
                    nombre : libro.Nombre,
                    contenido : libro.Contenido,
                }));
            } 
            if (result.librocurso2 != null) {
                libros2 = result.librocurso2.map((libro) => ({
                    nombre : libro.Nombre,
                    contenido : libro.Contenido,
                }));
            } 
            if (result.librocurso3 != null) {
                libros3 = result.librocurso3.map((libro) => ({
                    nombre : libro.Nombre,
                    contenido : libro.Contenido,
                }));
            }
            setLibro1(libros1);
            setLibro2(libros2);
            setLibro3(libros3); 
        } catch (error) {
            console.log('Error fetching data:', error);
        }     
    };

    const handleVerLibro = (contenido) => {
        setLibro(contenido);
    }

    return (
        <div className="container col-6 mt-5">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="h3 mb-3 fw-normal">Estudiante, {userLocal}</h1>
                    <br />
                    <h4 className="h3 mb-3 fw-normal">Libros en PDF</h4>
                    <br />
                    <center>
                        <button
                        onClick={fetchData}
                        className="w-50 btn btn-outline-primary"
                        type="button"
                        >
                            Cargar Libros
                        </button>
                    </center>
                    <br />
                    <div className="col-md-6">
                        <div className="text-center mb-3">
                            <h2>{curso1}</h2>
                        </div>
                        {libro1.map((libro, i) => (                            
                            <div key={i} className="col-md-8">
                                <div className="text-center mb-5">
                                    <label htmlFor="tituloLibro">{libro.nombre}</label>
                                    <button className="btn btn-primary"
                                    onClick={() => handleVerLibro(libro.contenido)}
                                    type="button"
                                    >Ver</button>
                                </div>
                            </div>                               
                        ))}
                    </div>
                    <br />
                    <div className="col-md-6">
                        <div className="text-center mb-3">
                            <h2>{curso2}</h2>
                        </div>
                        {libro2.map((libro, i) => (                            
                            <div key={i} className="col-md-8">
                                <div className="text-center mb-5">
                                    <label htmlFor="tituloLibro">{libro.nombre}</label>
                                    <button className="btn btn-primary"
                                    onClick={() => handleVerLibro(libro.contenido)}
                                    type="button"
                                    >Ver</button>
                                </div>
                            </div>                               
                        ))}
                    </div>
                    <br />
                    <div className="col-md-6">
                        <div className="text-center mb-3">
                            <h2>{curso3}</h2>
                        </div>
                        {libro3.map((libro, i) => (                            
                            <div key={i} className="col-md-8">
                                <div className="text-center mb-5">
                                    <label htmlFor="tituloLibro">{libro.nombre}</label>
                                    <button className="btn btn-primary"
                                    onClick={() => handleVerLibro(libro.contenido)}
                                    type="button"
                                    >Ver</button>
                                </div>
                            </div>                               
                        ))}
                    </div>
                    <center>
                        <iframe src={libro} 
                        width="600"
                        height="400"
                        />
                    </center>
                </form>
            </div>
        </div>
    )
}

export default VerLibrosEstudiante