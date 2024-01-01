import React, {useState} from 'react'

function VerPubliEstudiante() {

    const [publi1, setPubli1] = useState([]);
    const [publi2, setPubli2] = useState([]);
    const [publi3, setPubli3] = useState([]);
    const [curso1, setCurso1] = useState("");
    const [curso2, setCurso2] = useState("");
    const [curso3, setCurso3] = useState("");
    const userLocal = localStorage.getItem("user");
    
    const fetchData = async() => {
        try {
            const response = await fetch("http://localhost:4000/devolver-publicaciones", {
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
            
            let publicaciones1 = [];
            let publicaciones2 = [];
            let publicaciones3 = [];

            if (result.publicacioncurso1 != null) {
                publicaciones1 = result.publicacioncurso1.map((publi) => ({
                    contenido : publi.Contenido,
                }));
            } 
            if (result.publicacioncurso2 != null) {
                publicaciones2 = result.publicacioncurso2.map((publi) => ({
                    contenido : publi.Contenido,
                }));
            } 
            if (result.publicacioncurso3 != null) {
                publicaciones3 = result.publicacioncurso3.map((publi) => ({
                    contenido : publi.Contenido,
                }));
            }
            setPubli1(publicaciones1);
            setPubli2(publicaciones2);
            setPubli3(publicaciones3); 
        } catch (error) {
            console.log('Error fetching data:', error);
        }     
    }

    return (
        <div className="container col-6 mt-5">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="h3 mb-3 fw-normal">Estudiante, {userLocal}</h1>
                    <br />
                    <h4 className="h3 mb-3 fw-normal">Publicaciones</h4>
                    <br />
                    <center>
                        <button
                        onClick={fetchData}
                        className="w-50 btn btn-outline-primary"
                        type="button"
                        >
                            Cargar Publicaciones
                        </button>
                    </center>
                    <br />
                    <div className="col-md-6">
                        <div className="text-center mb-3">
                            <h2> {curso1} </h2>
                        </div>
                        {publi1.map((publi) => (
                            <div className="col-md-8">
                                <div className="text-center mb-5">
                                <textarea className="form-control" 
                                    rows="15"
                                    value={publi.contenido}
                                    readOnly
                                    />
                                </div>
                            </div>
                        ))} 
                    </div>
                    <br />
                    <div className="col-md-6">
                        <div className="text-center mb-3">
                            <h2> {curso2} </h2>
                        </div>
                        {publi2.map((publi) => (
                            <div className="col-md-8">
                                <div className="text-center mb-5">
                                <textarea className="form-control" 
                                    rows="15"
                                    value={publi.contenido}
                                    readOnly
                                    />
                                </div>
                            </div>
                        ))} 
                    </div>
                    <br />
                    <div className="col-md-6">
                        <div className="text-center mb-3">
                            <h2> {curso3} </h2>
                        </div>
                        {publi3.map((publi) => (
                            <div className="col-md-8">
                                <div className="text-center mb-5">
                                <textarea className="form-control" 
                                    rows="15"
                                    value={publi.contenido}
                                    readOnly
                                    />
                                </div>
                            </div>
                        ))} 
                    </div>
                </form>
            </div>
        </div>
    )
}

export default VerPubliEstudiante