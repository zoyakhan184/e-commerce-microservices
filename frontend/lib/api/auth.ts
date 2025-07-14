import axios from "@/lib/client"
import type { User } from "@/types"

interface AuthResponse {
  token: string
  userId: string
  role: string
}

export const authApi = {
  async login(email: string, password: string): Promise<AuthResponse> {
    const res = await axios.post("/auth/login", { email, password })
    return res.data
  },

  async register(name: string, email: string, password: string): Promise<AuthResponse> {
    const res = await axios.post("/auth/register", { name, email, password })
    return res.data
  },

  async validateToken(token: string): Promise<{ userId: string; role: string }> {
    const res = await axios.get(`/auth/validate?token=${token}`)
    return res.data
  },

  async forgotPassword(email: string): Promise<{ message: string }> {
    const res = await axios.post("/auth/forgot", { email })
    return res.data
  },

  async resetPassword(data: { token: string; newPassword: string }): Promise<{ message: string }> {
    const res = await axios.post("/auth/reset", data)
    return res.data
  },
}
