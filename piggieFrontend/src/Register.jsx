import React, { useState } from "react";
import { useNavigate } from "react-router";
import "./Register.css";

function Register() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [email, setEmail] = useState("");
    const [dateOfBirth, setDateOfBirth] = useState("");

    const navigation = useNavigate();

    const registerFunc = async () => {
        const response = await fetch("http://localhost:8080/register", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username: username,
                password: password,
                email: email,
                dateOfBirth: dateOfBirth,
            }),
        });

        if (!response.ok) {
            setUsername("");
            setPassword("");
            setEmail("");
            setDateOfBirth("");
            throw "Error";
        } else {
            setUsername("");
            setPassword("");
            setEmail("");
            setDateOfBirth("");

            navigation("/login");
        }
    };

    return (
        <>
            <div className="register-container">
                <label htmlFor="username" className="credential-box-label">
                    Username:{" "}
                </label>
                <input
                    type="text"
                    className="credential-box"
                    placeholder="Username..."
                    id="username"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                ></input>
                <label htmlFor="password" className="credential-box-label">
                    Password:{" "}
                </label>
                <input
                    type="password"
                    className="credential-box"
                    placeholder="Password..."
                    id="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                ></input>
                <label htmlFor="email" className="credential-box-label">
                    Email:{" "}
                </label>
                <input
                    type="email"
                    className="credential-box"
                    placeholder="Email..."
                    id="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                ></input>
                <label htmlFor="dateOfBirth" className="credential-box-label">
                    Date of birth:{" "}
                </label>
                <input
                    type="date"
                    className="credential-box"
                    id="dateOfBirth"
                    value={dateOfBirth}
                    onChange={(e) => setDateOfBirth(e.target.value)}
                ></input>
                <button
                    type="submit"
                    onClick={registerFunc}
                    className="credential-submit"
                >
                    Submit
                </button>
            </div>
        </>
    );
}

export default Register;
