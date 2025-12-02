import { createBrowserRouter } from "react-router-dom";
import App from "./App";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Profile from "./pages/Profile";
import Tickets from "./pages/Tickets";
import EventDetails from "./pages/EventDetails";
import CreateEvent from "./pages/CreateEvent";
import ManageEvents from "./pages/ManageEvents";
import EventTickets from "./pages/EventTickets";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />, // layout wspólny (z Navbar)
    children: [
        { path: "", element: <Home /> },
        { path: "login", element: <Login /> },
        { path: "register", element: <Register /> },
        { path: "profile", element: <Profile /> },
        { path: "tickets", element: <Tickets /> },
        { path: "events/:id", element: <EventDetails /> },
        { path: "events/create", element: <CreateEvent /> },
        { path: "events/manage", element: <ManageEvents /> },
        { path: "events/:id/tickets", element: <EventTickets /> },
    ],
  },
]);