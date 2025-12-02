import { useState } from "react";
import API from "../api/api";
import { useNavigate } from "react-router-dom";

export default function CreateEvent() {
  const [title, setTitle] = useState("");
  const [date, setDate] = useState("");
  const navigate = useNavigate();

  const handleCreate = async (e) => {
    e.preventDefault();
    try {
      await API.post("/events/", { title, date });
      navigate("/");
    } catch (err) {
      alert("Nie udało się utworzyć wydarzenia");
    }
  };

  return (
    <form onSubmit={handleCreate} className="flex flex-col gap-2 max-w-md mx-auto mt-8">
      <input type="text" placeholder="Tytuł wydarzenia" value={title} onChange={(e) => setTitle(e.target.value)} required />
      <input type="date" value={date} onChange={(e) => setDate(e.target.value)} required />
      <button type="submit">Utwórz wydarzenie</button>
    </form>
  );
}