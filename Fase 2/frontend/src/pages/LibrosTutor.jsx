import React, { useState } from 'react'

function LibrosTutor() {
    const [contenidoPDF, setContenidoPDF] = useState("");

    const uploadFileTutor = (e) => {
        const file = e.target.files[0];
        const reader = new FileReader();
        const name = file.name;

        reader.onload = async (e) => {
            const content = e.target.result;
            setContenidoPDF(content);

            const userLocal = localStorage.getItem("user");
            const response = await fetch("http://localhost:4000/cargar-libro", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    contenido : content,
                    carnet : userLocal,
                    nombre : name,
                }),
            });

            if (response.status === 200) {
                alert("Libro cargado exitosamente");
            } else {
                alert("Error al cargar libro");
            }
        };

        reader.readAsDataURL(file);
    }

    return (
        <div className="container col-6 mt-5">
            <div className="text-center">
                <form className="card card-body">
                <h1 className="h3 mb-3 fw-normal">Tutor</h1>
                <br />
                <h4 className="h3 mb-3 fw-normal">Cargar PDF</h4>
                <br />
                <div className="input-group mb-3">
                    <label className="input-group-text">Cargar Tutores</label>
                    <input
                    className="form-control"
                    id="inputGroupFile01"
                    type="file"
                    accept=".pdf"
                    onChange={uploadFileTutor}
                    />
                </div>
                <div className="mb-3 fw-normal">
                    <iframe src={contenidoPDF} width="800" height="800" />
                </div>
                </form>
            </div>
        </div>
    );
}

export default LibrosTutor