import { useAuth } from "../context/AuthContext";

export default function Profile() {
  const { user } = useAuth();

  if (!user) return <p>Nie zalogowano</p>;

  return (
    <div>
      <h1 className="text-xl font-bold">Profil</h1>
      <p>Email: {user.email}</p>
    </div>
  );
}