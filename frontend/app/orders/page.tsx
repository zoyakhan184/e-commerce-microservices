"use client"

import { useEffect, useState } from "react"
import { orderApi } from "@/lib/api/order"
import type { Order } from "@/types"

import { Header } from "@/components/layout/header"
import { Footer } from "@/components/layout/footer"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Package, Eye, Calendar, DollarSign } from "lucide-react"
import Link from "next/link"

export default function OrdersPage() {
  const [orders, setOrders] = useState<Order[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const data = await orderApi.getOrders()
        setOrders(data)
      } catch (err) {
        console.error("❌ Failed to fetch orders:", err)
      } finally {
        setLoading(false)
      }
    }

    fetchOrders()
  }, [])

  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      <main className="flex-1">
        <div className="container py-8">
          <div className="mb-8">
            <h1 className="text-3xl font-bold mb-2">Order History</h1>
            <p className="text-muted-foreground">Track and manage your orders</p>
          </div>

          {loading ? (
            <p className="text-center py-10">Loading orders...</p>
          ) : orders?.length === 0 ? (
            <Card>
              <CardContent className="text-center py-16">
                <Package className="h-16 w-16 text-muted-foreground mx-auto mb-4" />
                <h2 className="text-2xl font-semibold mb-2">No orders yet</h2>
                <p className="text-muted-foreground mb-6">When you place your first order, it will appear here.</p>
                <Button asChild>
                  <Link href="/products">Start Shopping</Link>
                </Button>
              </CardContent>
            </Card>
          ) : (
            <div className="space-y-6">
              {orders?.map((order) => (
                <Card key={order.id}>
                  <CardHeader>
                    <div className="flex items-center justify-between">
                      <CardTitle className="flex items-center gap-2">
                        <Package className="h-5 w-5" />
                        Order #{order.id}
                      </CardTitle>
                      <Badge
                        variant={order.status === "delivered" ? "default" : "secondary"}
                        className={order.status === "delivered" ? "bg-green-100 text-green-700" : ""}
                      >
                        {order.status}
                      </Badge>
                    </div>
                  </CardHeader>
                  <CardContent>
                    <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
                      <div className="flex items-center gap-2">
                        <Calendar className="h-4 w-4 text-muted-foreground" />
                        <div>
                          <p className="text-sm text-muted-foreground">Order Date</p>
                          <p className="font-medium">
                            {new Date(order.created_at).toLocaleDateString()}
                          </p>
                        </div>
                      </div>
                      <div className="flex items-center gap-2">
                        <DollarSign className="h-4 w-4 text-muted-foreground" />
                        <div>
                          <p className="text-sm text-muted-foreground">Total Amount</p>
                          <p className="font-medium">${order.total_amount?.toFixed(2)}</p>
                        </div>
                      </div>
                      <div className="flex items-center gap-2">
                        <Package className="h-4 w-4 text-muted-foreground" />
                        <div>
                          <p className="text-sm text-muted-foreground">Payment Status</p>
                          <p className="font-medium capitalize">{order.payment_status}</p>
                        </div>
                      </div>
                    </div>

                    <div className="space-y-3">
                      <h4 className="font-semibold">Items Ordered</h4>
                      {(order.items ?? []).map((item, index) => (
                        <div
                          key={index}
                          className="flex items-center justify-between p-3 bg-muted/30 rounded-lg"
                        >
                          <div>
                            <p className="font-medium">{item.product_name || item.product_id}</p>
                            <p className="text-sm text-muted-foreground">
                              Size: {item.size || "-"} | Color: {item.color || "-"} | Qty: {item.quantity}
                            </p>
                          </div>
                          <p className="font-semibold">
                            ${(item.price * item.quantity).toFixed(2)}
                          </p>
                        </div>
                      ))}
                    </div>

                    <div className="flex justify-end mt-4">
                      <Link href={`/orders/${order.id}`}>
                        <Button variant="outline" size="sm">
                          <Eye className="h-4 w-4 mr-2" />
                          View Details
                        </Button>
                      </Link>
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          )}
        </div>
      </main>
      <Footer />
    </div>
  )
}
