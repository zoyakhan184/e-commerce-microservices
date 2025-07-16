"use client"

import { createContext, useContext, useEffect, useState } from "react"
import { authApi } from "@/lib/api/auth"
import { useRouter } from "next/navigation"
import Cookies from "js-cookie"
import type { User } from "@/types"

interface AuthContextType {
  user: User | null
  login: (email: string, password: string) => Promise<boolean>
  register: (name: string, email: string, password: string) => Promise<boolean>
  forgotPassword: (email: string) => Promise<boolean>
  resetPassword: (token: string, newPassword: string) => Promise<boolean>
  logout: () => void
  isLoading: boolean
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const router = useRouter()

    useEffect(() => {
    const token = Cookies.get("token")
    if (token) {
      authApi
        .validateToken(token)
        .then((data) => {
          setUser({ id: data.userId, role: data.role as "user" | "admin", name: "", email: "" })
        })
        .catch(() => {
          logout()
        })
        .finally(() => setIsLoading(false))
    } else {
      setIsLoading(false)
    }
  }, [])



  const login = async (email: string, password: string): Promise<boolean> => {
  try {
    
    const data = await authApi.login(email, password)
     Cookies.set("token", data.token, {
      expires: 7,
      path: "/",
      sameSite: "Lax",
      secure: process.env.NODE_ENV === "production",
    })
    setUser({ id: data.userId, role: data.role as "user" | "admin", name: "", email })
    return true
  } catch {
    return false
  }
}

  const register = async (name: string, email: string, password: string): Promise<boolean> => {
    try {
      const data = await authApi.register(name, email, password)
      console.log("[AuthContext] 📝 Registered user:", data)

      Cookies.set("token", data.token, {
        expires: 7,
        path: "/",
        sameSite: "Lax",
        secure: process.env.NODE_ENV === "production",
      })
      setUser({ id: data.userId, role: data.role as "user" | "admin", name, email })
      return true
    } catch {
      return false
    }
  }


  const forgotPassword = async (email: string): Promise<boolean> => {
    try {
      await authApi.forgotPassword(email)
      return true
    } catch {
      return false
    }
  }

  const resetPassword = async (token: string, newPassword: string): Promise<boolean> => {
    try {
      await authApi.resetPassword({ token, newPassword })
      return true
    } catch {
      return false
    }
  }

  const logout = () => {
    Cookies.remove("token")
    setUser(null)
    router.push("/auth/login")
  }

  return (
    <AuthContext.Provider
      value={{
        user,
        login,
        register,
        forgotPassword,
        resetPassword,
        logout,
        isLoading,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth() {
  const context = useContext(AuthContext)
  if (!context) throw new Error("useAuth must be used within AuthProvider")
  return context
}
