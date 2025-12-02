import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import API from "../api/api";

export default function EventTickets() {
  const { id } = useParams();
  const [tickets, setTickets] = useState([]);

  useEffect(() => {
    API.get(`/events/${id}/tickets`).then(res => setTickets(res.data)).catch(() => {
      alert("Brak uprawnień do przeglądania biletów tego wydarzenia");
    });
  }, [id]);

  return (
    <div>
      <h1 className="text-xl font-bold">Bilety dla wydarzenia #{id}</h1>
      <ul>
        {tickets.map(t => (
          <li key={t.id}>Bilet #{t.id} – Użytkownik: {t.user_id}</li>
        ))}
      </ul>
    </div>
  );
}