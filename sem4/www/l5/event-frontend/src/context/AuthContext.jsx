import { createContext, useContext, useState, useEffect } from "react";
import API from "../api/api";

const AuthContext = createContext();

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null);

  const login = async (email, password) => {
    const res = await API.post("/login", new URLSearchParams({
      username: email,
      password: password
    }));
    localStorage.setItem("token", res.data.access_token);
    await fetchProfile();
  };

  const fetchProfile = async () => {
    try {
      const res = await API.get("/users/me");
      setUser(res.data);
    } catch {
      logout();
    }
  };

  const logout = () => {
    localStorage.removeItem("token");
    setUser(null);
    window.location.href = "/";
  };

  useEffect(() => {
    if (localStorage.getItem("token")) fetchProfile();
  }, []);

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => useContext(AuthContext);