"use client"

import { useAuth } from "@/contexts/auth-context"
import { useRouter } from "next/navigation"
import { useEffect } from "react"

export default function AuthGuard({
  children,
  requireAdmin = false,
}: {
  children: React.ReactNode
  requireAdmin?: boolean
}) {
  const { user, isLoading } = useAuth()
  const router = useRouter()

  useEffect(() => {
    if (!isLoading) {
      if (!user) router.replace("/auth/login")
      else if (requireAdmin && user.role !== "admin") router.replace("/dashboard")
    }
  }, [user, isLoading, requireAdmin, router])

  if (isLoading) return <div className="text-center mt-20">Loading...</div>
  return <>{children}</>
}