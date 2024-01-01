import {Routes, Route} from 'react-router-dom'
import "bootstrap/dist/css/bootstrap.min.css";
import './App.css'
import Login from './pages/Login'
import Admin from './pages/Admin'
import TablaAlumnos from './pages/TablaAlumnos'
import Estudiante from './pages/Estudiante'
import Tutor from './pages/Tutor'
import LibroTutor from './pages/LibrosTutor'
import PubliTutor from './pages/PublicacionesTutor'
import VerLibrosEstudiante from './pages/VerLibrosEstudiante'
import VerPubliEstudiante from './pages/VerPubliEstudiante';
import LibrosAdmin from './pages/LibrosAdmin';

function App() {

  return (
    <>
      <Routes>
        <Route  path="/" element={<Login />} />
        <Route  path="/principal/admin" element={<Admin />} />
        <Route  path="/principal/admin/tabla-alumnos" element={<TablaAlumnos />} />
        <Route  path="/principal/admin/libros-admin" element={<LibrosAdmin />} />
        
        <Route  path="/principal/tutor" element={<Tutor />} />
        <Route  path="/principal/tutor/libro" element={<LibroTutor />} />
        <Route  path="/principal/tutor/publicacion" element={<PubliTutor />} />

        <Route  path="/principal/estudiante" element={<Estudiante />} />
        <Route  path="/principal/estudiante/libro" element={<VerLibrosEstudiante />} />
        <Route  path="/principal/estudiante/publicacion" element={<VerPubliEstudiante />} />
      </Routes>
    </>
  )
}

export default App
