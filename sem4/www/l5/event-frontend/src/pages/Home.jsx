import { useEffect, useState } from "react";
import API from "../api/api";
import { Link } from "react-router-dom";

export default function Home() {
  const [events, setEvents] = useState([]);

  useEffect(() => {
    API.get("/events/").then(res => setEvents(res.data));
  }, []);

  return (
    <div>
      <h1 className="text-xl font-bold mb-4">Wydarzenia</h1>
      <ul className="space-y-2">
        {events.map(event => (
          <li key={event.id}>
            <Link to={`/events/${event.id}`} className="underline text-blue-600">{event.title}</Link>
          </li>
        ))}
      </ul>
    </div>
  );
}
