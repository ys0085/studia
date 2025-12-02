import { useState } from "react";
import API from "../api/api";
import { useNavigate } from "react-router-dom";

export default function Register() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    try {
      await API.post("/register", { email, password });
      navigate("/login");
    } catch (err) {
      alert("Rejestracja nie powiodła się");
    }
  };

  return (
    <form onSubmit={handleRegister} className="flex flex-col gap-2 max-w-md mx-auto mt-8">
      <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} required />
      <input type="password" placeholder="Hasło" value={password} onChange={(e) => setPassword(e.target.value)} required />
      <button type="submit">Zarejestruj</button>
    </form>
  );
}