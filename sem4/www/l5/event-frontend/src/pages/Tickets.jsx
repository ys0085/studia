import { useEffect, useState } from "react";
import API from "../api/api";

export default function Tickets() {
  const [tickets, setTickets] = useState([]);
  const [events, setEvents] = useState([]);

  useEffect(() => {
    API.get("/tickets/").then(res => setTickets(res.data));
    API.get("/events/").then(res => setEvents(res.data));
  }, []);

return (
    <div>
        <h1 className="text-xl font-bold">Moje bilety</h1>
        <ul>
            {tickets.map(ticket => {
                const event = events.find(e => e.id === ticket.event_id);
                return event ? (
                    <li key={ticket.id}>
                        Bilet na {event.title}
                    </li>
                ) : null;
            })}
        </ul>
    </div>
);
}