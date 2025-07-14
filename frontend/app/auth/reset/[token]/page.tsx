"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { authApi } from "@/lib/api/auth"

export default function ResetPasswordPage({ params }: { params: { token: string } }) {
  const [newPassword, setNewPassword] = useState("")
  const [confirmPassword, setConfirmPassword] = useState("")
  const [error, setError] = useState("")
  const [message, setMessage] = useState("")
  const router = useRouter()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setError("")
    setMessage("")

    if (newPassword !== confirmPassword) {
      setError("Passwords do not match")
      return
    }

    try {
      const res = await authApi.resetPassword({ token: params.token, newPassword })
      setMessage(res.message || "Password reset successful")
      setTimeout(() => router.push("/auth/login"), 2000)
    } catch (err: any) {
        const errorMsg = err?.response?.data?.message || "Invalid or expired token"
        setError(errorMsg)
}


  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4 max-w-md mx-auto mt-10">
      <h2 className="text-xl font-bold">Reset Password</h2>
      <input
        type="password"
        value={newPassword}
        onChange={(e) => setNewPassword(e.target.value)}
        placeholder="New Password"
        className="border p-2 w-full"
        required
      />
      <input
        type="password"
        value={confirmPassword}
        onChange={(e) => setConfirmPassword(e.target.value)}
        placeholder="Confirm Password"
        className="border p-2 w-full"
        required
      />
      <button type="submit" className="bg-black text-white px-4 py-2 rounded">
        Reset Password
      </button>
      {message && <p className="text-green-600">{message}</p>}
      {error && <p className="text-red-600">{error}</p>}
    </form>
  )
}
