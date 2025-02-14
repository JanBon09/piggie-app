import viteLogo from "/vite.svg";
import "./App.css";
import React, { useState } from "react";

function App() {
    const [message, setMessage] = useState("");

    // const getMessage = async () => {
    //     const response = await fetch("http://localhost:8080/welcome");
    //     const data = response.json();
    //     setMessage(data);
    // };

    fetch("http://localhost:8080/welcome")
        .then((response) => response.json())
        .then((data) => setMessage(data));

    return <div className="mainContainter">{message}</div>;
}

export default App;
