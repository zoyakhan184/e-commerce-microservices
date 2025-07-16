"use client"

import { useAuth } from "@/contexts/auth-context"
import { useRouter, usePathname } from "next/navigation"
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
  const pathname = usePathname()

  useEffect(() => {
    if (!isLoading) {
      if (!user && pathname !== "/auth/login") {
        router.replace(`/auth/login?redirect=${pathname}`)
      } else if (requireAdmin && user?.role !== "admin") {
        router.replace("/dashboard")
      }
    }
  }, [user, isLoading, requireAdmin, pathname, router])

  if (isLoading) return <div className="text-center mt-20">Loading...</div>
  return <>{children}</>
}
