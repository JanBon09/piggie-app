import React, { useState } from "react";
import { useNavigate } from "react-router";
import "./Login.css";

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const navigation = useNavigate();

    const loginFunc = async () => {
        const response = await fetch("http://localhost:8080/login", {
            method: "POST",
            credentials: "include",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username: username,
                password: password,
            }),
        });

        if (!response.ok) {
            setPassword("");
            throw "Wrong username or password";
        } else {
            setUsername("");
            setPassword("");

            navigation("/");
        }
    };

    return (
        <>
            <div className="login-container">
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
                <button
                    type="submit"
                    onClick={loginFunc}
                    className="credential-submit"
                >
                    Submit
                </button>
            </div>
        </>
    );
}

export default Login;
