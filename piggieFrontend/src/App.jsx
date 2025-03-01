import viteLogo from "/vite.svg";
import "./App.css";
import React, { useState } from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router";
import Welcome from "./Welcome";
import Login from "./Login";
import Register from "./Register";

function App() {
    const [message, setMessage] = useState("");

    fetch("http://localhost:8080/welcome")
        .then((response) => response.json())
        .then((data) => setMessage(data));

    return (
        <Router>
            <Routes>
                <Route path="/" element={<Welcome />} />
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
            </Routes>
        </Router>
    );
}

export default App;
