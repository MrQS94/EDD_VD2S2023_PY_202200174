import React, {useState} from 'react'

function PublicacionesTutor() {
    const [contenidoPubli, setContenidoPubli] = useState("");
    const user = localStorage.getItem("user");

    // Falta implementar la función para guardar la publicación en la base de datos
    const guardarPublicacion = (e) => {
        setContenidoPubli(e.target.value);
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = fetch("http://localhost:4000/cargar-publicacion", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                carnet : user,
                contenido : contenidoPubli,
            }),
        });    
        const result = await (await response).json();
        if (result.status === 200) {
            alert("Publicación guardada exitosamente");
            var text = document.getElementById("textarea1");
            text.value = "";
        } else {
            alert("Error al guardar la publicación");
        }
    }

    const salir = (e) => {
        e.preventDefault();
        window.open("/", "_self");
    }

    return (
        <div className="container col-5 mt-5">
            <div className="text-center">
                <form className="card card-body"
                onSubmit={handleSubmit}
                >
                    <h1 className="h3 mb-3 fw-normal">Tutor, {user}</h1>
                    <br />
                    <h4 className="h3 mb-3 fw-normal">Cargar Publicación</h4>
                    <br />
                    <div className="input-group mb-3">
                        <textarea className="form-control" 
                        rows="15"
                        value={contenidoPubli}
                        onChange={guardarPublicacion}
                        id = "textarea1"
                        />
                    </div>
                    <br />
                    <center>
                        <button 
                        className="w-50 btn btn-lg btn btn-outline-success"
                        type="submit"
                        >
                        Guardar Publicación
                        </button>
                    </center>
                    <center>
                        <p className="mt-5 mb-3 text-muted">EDD 202200174</p>
                    </center>
                    <center>
                        <button className="w-50 btn btn-outline-success" 
                        onClick={salir}
                        >
                        Salir
                        </button>
                    </center>
                </form>
            </div>
        </div>
    )
}

export default PublicacionesTutor