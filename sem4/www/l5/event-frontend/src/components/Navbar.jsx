import { Link } from "react-router-dom";
import { useAuth } from "../context/AuthContext";

export default function Navbar() {
  const { user, logout } = useAuth();
  return (
    <nav className="flex gap-4 bg-gray-800 text-white p-4">
      <Link to="/">Wydarzenia</Link>
      {user ? (
        <>
          <Link to="/profile">Profil</Link>
          <Link to="/tickets">Bilety</Link>
          <Link to="/events/create">+ Nowe wydarzenie</Link>
          <Link to="/events/manage">Zarządzaj</Link>
          <button onClick={logout}>Wyloguj</button>
        </>
      ) : (
        <>
          <Link to="/login">Logowanie</Link>
          <Link to="/register">Rejestracja</Link>
        </>
      )}
    </nav>
  );
}