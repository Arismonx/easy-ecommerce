'use client'
import { useState } from "react";
import { login } from "../utils/api";
import { saveToken } from "../utils/auth";
import { logout } from "@/utils/auth";
export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    try {
      const data = await login(email, password);
      saveToken(data.token);
      alert("Login successful!");
    } catch (err) {
      alert("Login failed");
    }
  };

  const headleLogout = async () => {
    try {
      await logout()
      alert("Logout successful!");
    } catch {
      alert("Logout failed");
    }
  }

  return (
    <div>
      <h1>Login</h1>
      <input type="email" onChange={(e) => setEmail(e.target.value)} />
      <input type="password" onChange={(e) => setPassword(e.target.value)} />
      <button onClick={handleLogin}>Login</button>
      <button onClick={headleLogout}>logout</button>
    </div>
  );
}
