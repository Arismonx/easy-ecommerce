export async function login(email: string, password: string) {
    const res = await fetch("http://localhost:5000/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });
  
    if (!res.ok) throw new Error("Invalid credentials");
    return res.json();
  }