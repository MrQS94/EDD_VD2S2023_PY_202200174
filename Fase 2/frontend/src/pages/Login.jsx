import React, {useState} from 'react'

function Login() {
    const [isChecked, setIsChecked] = useState(false);
    const [user, setUser] = useState('');
    const [pass, setPass] = useState('');

    const handleSubmit = async(e) => {
        e.preventDefault();
        const response = await fetch('http://localhost:4000/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                User : user,
                Pass : pass,
                Tutor : isChecked,
            }),
        });

        const result = await response.json();
        console.log(result);
        if (result.rol == 0) {
            alert("Credenciales incorrectas");
        } else if (result.rol == 1) {
            window.open("/principal/admin", "_self");
            localStorage.setItem('tipo', "1");
            localStorage.setItem('user', user);
        } else if (result.rol == 2) {
            window.open("/principal/tutor", "_self");
            localStorage.setItem('tipo', "2");
            localStorage.setItem('user', user);
        } else if (result.rol == 3) {
            window.open("/principal/estudiante", "_self");
            localStorage.setItem('tipo', "3");
            localStorage.setItem('user', user);
        }
    }

    return (
        <div className="container mt-5 col-3">
            <div className="text-center">
                <form className="card card-body" onSubmit={handleSubmit}>
                <h1 className="h3 mb-3 fw-normal">Inicio de Sesión</h1>
                <h1 className="h3 mb-3 fw-normal">Tutorias ECYS</h1>
                <label htmlFor="inputEmail" className="visually-hidden">
                    Usuario
                </label>
                <input
                    type="text"
                    id="userI"
                    className="form-control"
                    placeholder="Usuario:"
                    value={user}
                    onChange={(e) => setUser(e.target.value)}
                    required
                />
                <br />
                <label htmlFor="inputPassword" className="visually-hidden">
                    Password
                </label>
                <input
                    type="password"
                    id="passI"
                    className="form-control"
                    placeholder="Contraseña:"
                    aria-describedby="passwordHelpInline" 
                    value={pass}
                    onChange={(e) => setPass(e.target.value)}
                />
                <br />
                <div className="form-check form-switch col-4">
                    <input
                    className="form-check-input"
                    type="checkbox"
                    role="switch"
                    id="flexSwitchCheckDefault"
                    checked={isChecked}
                    onChange={() => setIsChecked(!isChecked)}
                    />
                    <label
                    className="form-check-label"
                    htmlFor="flexSwitchCheckDefault">
                    Tutor
                    </label>
                </div>
                <br />
                <button
                    className="w-100 btn btn-lg btn btn-outline-success"
                    type="submit">
                    Iniciar Sesion
                </button>
                <p className="mt-5 mb-3 text-muted">EDD 202200174</p>
                <br />
                </form>
            </div>
        </div>
  )
}

export default Login