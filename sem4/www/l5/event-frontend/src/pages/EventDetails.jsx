import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import API from "../api/api";

export default function EventDetails() {
  const { id } = useParams();
  const [event, setEvent] = useState(null);

  useEffect(() => {
    API.get(`/events/${id}`).then(res => setEvent(res.data));
  }, [id]);

  const handleBuyTicket = async () => {
    try {
      await API.post(`/events/${id}/tickets`);
      alert("Kupiono bilet!");
    } catch (err) {
      alert("Błąd zakupu biletu");
    }
  };

  if (!event) return <p>Wczytywanie...</p>;

  return (
    <div>
      <h2 className="text-xl font-bold">{event.title}</h2>
      <p>Data: {event.date}</p>
      <button onClick={handleBuyTicket} className="mt-2">Kup bilet</button>
    </div>
  );
}