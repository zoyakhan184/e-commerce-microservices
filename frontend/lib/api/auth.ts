import axios from "../client"
import type { User } from "@/types"
import Cookies from "js-cookie"

interface AuthResponse {
  token: string
  userId: string
  role: string
}

export const authApi = {
  async login(email: string, password: string): Promise<AuthResponse> {
    const res = await axios.post("api/auth/login", { email, password })
    return res.data
  },

  async register(name: string, email: string, password: string): Promise<AuthResponse> {
    const res = await axios.post("api/auth/register", { name, email, password })
    return res.data
  },

  async validateToken(token: string): Promise<{ userId: string; role: string }> {
    const res = await axios.get("api/auth/validate", {
      headers: {
        Authorization: `Bearer ${Cookies.get("token")}`
      }
    })
    return res.data
  },

  async forgotPassword(email: string): Promise<{ message: string }> {
    const res = await axios.post("api/auth/forgot", { email })
    return res.data
  },

  async resetPassword(data: { token: string; newPassword: string }): Promise<{ message: string }> {
    const res = await axios.post("api/auth/reset", data)
    return res.data
  },
}
