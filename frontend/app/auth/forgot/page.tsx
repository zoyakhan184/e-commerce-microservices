"use client"

import { useState } from "react"
import { authApi } from "@/lib/api/auth"

export default function ForgotPasswordPage() {
  const [email, setEmail] = useState("")
  const [message, setMessage] = useState("")
  const [error, setError] = useState("")

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setMessage("")
    setError("")

    try {
      const res = await authApi.forgotPassword(email)
      setMessage(res.message || "Reset link sent to your email.")
    } catch (err) {
      setError("Email not found or failed to send reset link.")
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4 max-w-md mx-auto mt-10">
      <h2 className="text-xl font-bold">Forgot Password</h2>
      <input
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Enter your email"
        className="border p-2 w-full"
        required
      />
      <button type="submit" className="bg-black text-white px-4 py-2 rounded">
        Send Reset Link
      </button>
      {message && <p className="text-green-600">{message}</p>}
      {error && <p className="text-red-600">{error}</p>}
    </form>
  )
}
