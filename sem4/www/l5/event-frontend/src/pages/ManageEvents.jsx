import { useEffect, useState } from "react";
import API from "../api/api";

export default function ManageEvents() {
  const [events, setEvents] = useState([]);

  const fetchEvents = async () => {
    const res = await API.get("/events/");
    setEvents(res.data);
  };

  const handleDelete = async (id) => {
    try {
      await API.delete(`/events/${id}`);
      fetchEvents();
    } catch (err) {
      alert("Brak uprawnień lub błąd przy usuwaniu");
    }
  };

  useEffect(() => {
    fetchEvents();
  }, []);

  return (
    <div>
      <h1 className="text-xl font-bold">Zarządzaj wydarzeniami</h1>
      <ul className="space-y-2">
        {events.map(event => (
          <li key={event.id} className="flex justify-between">
            <span>{event.title} – {event.date}</span>
            <button onClick={() => handleDelete(event.id)} className="text-red-600">Usuń</button>
          </li>
        ))}
      </ul>
    </div>
  );
}