"use client"

import { useAuth } from "@/contexts/auth-context"
import { useRouter } from "next/navigation"
import { useEffect } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import Link from "next/link"
import { ShoppingBag, Users, Shield } from "lucide-react"

export default function HomePage() {
  const { user } = useAuth()
  const router = useRouter()

  useEffect(() => {
    if (user) {
      if (user.role === "admin") {
        router.push("/admin")
      } else {
        router.push("/dashboard")
      }
    }
  }, [user, router])

  if (user) {
    return null // Will redirect
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-50 via-white to-pink-50 dark:from-gray-900 dark:via-gray-800 dark:to-purple-900">
      <div className="container mx-auto px-4 py-16">
        <div className="text-center mb-16">
          <h1 className="text-6xl font-bold bg-gradient-to-r from-purple-600 to-pink-600 bg-clip-text text-transparent mb-6">
            StyleNest
          </h1>
          <p className="text-xl text-gray-600 dark:text-gray-300 mb-8 max-w-2xl mx-auto">
            Discover premium fashion that speaks your style. From casual wear to luxury pieces, find everything you need
            to express yourself.
          </p>
          <div className="flex gap-4 justify-center">
            <Link href="/auth/login">
              <Button
                size="lg"
                className="bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700"
              >
                Get Started
              </Button>
            </Link>
            <Link href="/auth/register">
              <Button variant="outline" size="lg">
                Create Account
              </Button>
            </Link>
          </div>
        </div>

        <div className="grid md:grid-cols-3 gap-8 mb-16">
          <Card className="glassmorphism border-0">
            <CardContent className="p-8 text-center">
              <ShoppingBag className="h-12 w-12 mx-auto mb-4 text-purple-600" />
              <h3 className="text-xl font-semibold mb-2">Premium Collection</h3>
              <p className="text-gray-600 dark:text-gray-300">
                Curated selection of high-quality fashion items from top brands
              </p>
            </CardContent>
          </Card>

          <Card className="glassmorphism border-0">
            <CardContent className="p-8 text-center">
              <Users className="h-12 w-12 mx-auto mb-4 text-pink-600" />
              <h3 className="text-xl font-semibold mb-2">Community Driven</h3>
              <p className="text-gray-600 dark:text-gray-300">
                Join thousands of fashion enthusiasts sharing their style
              </p>
            </CardContent>
          </Card>

          <Card className="glassmorphism border-0">
            <CardContent className="p-8 text-center">
              <Shield className="h-12 w-12 mx-auto mb-4 text-blue-600" />
              <h3 className="text-xl font-semibold mb-2">Secure Shopping</h3>
              <p className="text-gray-600 dark:text-gray-300">
                Safe and secure transactions with multiple payment options
              </p>
            </CardContent>
          </Card>
        </div>

        <div className="text-center">
          <h2 className="text-3xl font-bold mb-4">Test Credentials</h2>
          <div className="grid md:grid-cols-2 gap-6 max-w-2xl mx-auto">
            <Card>
              <CardContent className="p-6">
                <h3 className="font-semibold mb-2">User Account</h3>
                <p className="text-sm text-gray-600 dark:text-gray-300">
                  Email: user@example.com
                  <br />
                  Password: password123
                </p>
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-6">
                <h3 className="font-semibold mb-2">Admin Account</h3>
                <p className="text-sm text-gray-600 dark:text-gray-300">
                  Email: admin@example.com
                  <br />
                  Password: admin123
                </p>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  )
}
