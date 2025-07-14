"use client"

import { useEffect } from "react"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { useAuth } from "@/contexts/auth-context"
import { Header } from "@/components/layout/header"
import { CheckCircle, Package, Truck, Calendar, ArrowRight } from "lucide-react"
import Link from "next/link"

export default function OrderConfirmationPage() {
  const { user } = useAuth()

  useEffect(() => {
    // Show success notification
    const timer = setTimeout(() => {
      // You could show a toast notification here
    }, 1000)

    return () => clearTimeout(timer)
  }, [])

  const orderId = `ORD-${Date.now()}`
  const expectedDelivery = new Date()
  expectedDelivery.setDate(expectedDelivery.getDate() + 5)

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <Header />
      <div className="container mx-auto px-4 py-8">
        <div className="max-w-2xl mx-auto">
          {/* Success Animation */}
          <div className="text-center mb-8">
            <div className="inline-flex items-center justify-center w-20 h-20 bg-green-100 dark:bg-green-900/20 rounded-full mb-4">
              <CheckCircle className="h-12 w-12 text-green-600" />
            </div>
            <h1 className="text-3xl font-bold text-green-800 dark:text-green-400 mb-2">Order Confirmed!</h1>
            <p className="text-gray-600 dark:text-gray-400">Thank you for your purchase, {user?.name}</p>
          </div>

          {/* Order Details */}
          <Card className="mb-6">
            <CardHeader>
              <CardTitle className="flex items-center justify-between">
                <span>Order Details</span>
                <Badge variant="outline" className="text-green-600 border-green-600">
                  Confirmed
                </Badge>
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid grid-cols-2 gap-4">
                <div>
                  <p className="text-sm text-gray-600 dark:text-gray-400">Order ID</p>
                  <p className="font-semibold">{orderId}</p>
                </div>
                <div>
                  <p className="text-sm text-gray-600 dark:text-gray-400">Order Date</p>
                  <p className="font-semibold">{new Date().toLocaleDateString()}</p>
                </div>
              </div>

              <div className="border-t pt-4">
                <p className="text-sm text-gray-600 dark:text-gray-400 mb-2">Delivery Address</p>
                <div className="bg-gray-50 dark:bg-gray-800 p-3 rounded-lg">
                  <p className="font-semibold">{user?.name}</p>
                  <p className="text-sm text-gray-600 dark:text-gray-400">
                    123 Main Street, Apt 4B
                    <br />
                    New York, NY 10001
                    <br />
                    {user?.phone}
                  </p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Delivery Timeline */}
          <Card className="mb-6">
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Truck className="h-5 w-5" />
                Delivery Timeline
              </CardTitle>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="flex items-center gap-4">
                  <div className="w-8 h-8 bg-green-100 dark:bg-green-900/20 rounded-full flex items-center justify-center">
                    <CheckCircle className="h-4 w-4 text-green-600" />
                  </div>
                  <div className="flex-1">
                    <p className="font-semibold">Order Confirmed</p>
                    <p className="text-sm text-gray-600 dark:text-gray-400">{new Date().toLocaleString()}</p>
                  </div>
                </div>

                <div className="flex items-center gap-4">
                  <div className="w-8 h-8 bg-blue-100 dark:bg-blue-900/20 rounded-full flex items-center justify-center">
                    <Package className="h-4 w-4 text-blue-600" />
                  </div>
                  <div className="flex-1">
                    <p className="font-semibold">Processing</p>
                    <p className="text-sm text-gray-600 dark:text-gray-400">Expected within 24 hours</p>
                  </div>
                </div>

                <div className="flex items-center gap-4">
                  <div className="w-8 h-8 bg-gray-100 dark:bg-gray-800 rounded-full flex items-center justify-center">
                    <Truck className="h-4 w-4 text-gray-400" />
                  </div>
                  <div className="flex-1">
                    <p className="font-semibold text-gray-600 dark:text-gray-400">Shipped</p>
                    <p className="text-sm text-gray-600 dark:text-gray-400">Expected in 2-3 days</p>
                  </div>
                </div>

                <div className="flex items-center gap-4">
                  <div className="w-8 h-8 bg-gray-100 dark:bg-gray-800 rounded-full flex items-center justify-center">
                    <Calendar className="h-4 w-4 text-gray-400" />
                  </div>
                  <div className="flex-1">
                    <p className="font-semibold text-gray-600 dark:text-gray-400">Delivered</p>
                    <p className="text-sm text-gray-600 dark:text-gray-400">
                      Expected by {expectedDelivery.toLocaleDateString()}
                    </p>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Next Steps */}
          <Card>
            <CardHeader>
              <CardTitle>What's Next?</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div className="grid gap-3">
                <Link href="/orders">
                  <Button variant="outline" className="w-full justify-between bg-transparent">
                    Track Your Order
                    <ArrowRight className="h-4 w-4" />
                  </Button>
                </Link>

                <Link href="/dashboard">
                  <Button variant="outline" className="w-full justify-between bg-transparent">
                    Continue Shopping
                    <ArrowRight className="h-4 w-4" />
                  </Button>
                </Link>
              </div>

              <div className="bg-blue-50 dark:bg-blue-900/20 p-4 rounded-lg">
                <p className="text-sm text-blue-800 dark:text-blue-400">
                  📧 A confirmation email has been sent to {user?.email}
                </p>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
