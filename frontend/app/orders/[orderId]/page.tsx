"use client"

import { useEffect, useState } from "react"
import { useParams } from "next/navigation"
import { orderApi } from "@/lib/api/order"
import type { Order } from "@/types"
import { Header } from "@/components/layout/header"
import { Footer } from "@/components/layout/footer"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Calendar, DollarSign, Package } from "lucide-react"

export default function OrderDetailsPage() {
  const { orderId } = useParams()
  const [order, setOrder] = useState<Order | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchOrder = async () => {
      try {
        const data = await orderApi.getOrderDetails(orderId as string)
        setOrder(data)
      } catch (err) {
        console.error("❌ Failed to fetch order details:", err)
      } finally {
        setLoading(false)
      }
    }

    if (orderId) fetchOrder()
  }, [orderId])

  if (loading) return <p className="text-center py-10">Loading order details...</p>
  if (!order) return <p className="text-center py-10">Order not found.</p>

  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      <main className="flex-1 container py-8">
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center justify-between">
              <span>Order #{order.id}</span>
              <Badge
                variant={order.status === "delivered" ? "default" : "secondary"}
                className={order.status === "delivered" ? "bg-green-100 text-green-700" : ""}
              >
                {order.status}
              </Badge>
            </CardTitle>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mb-6">
              <div className="flex items-center gap-2">
                <Calendar className="h-4 w-4 text-muted-foreground" />
                <div>
                  <p className="text-sm text-muted-foreground">Order Date</p>
                  <p className="font-medium">
                    {order.created_at ? new Date(order.created_at).toLocaleDateString() : "-"}
                  </p>
                </div>
              </div>
              <div className="flex items-center gap-2">
                <DollarSign className="h-4 w-4 text-muted-foreground" />
                <div>
                  <p className="text-sm text-muted-foreground">Total Amount</p>
                  <p className="font-medium">
                    ${(order.total_amount ?? 0).toFixed(2)}
                  </p>
                </div>
              </div>
              <div className="flex items-center gap-2">
                <Package className="h-4 w-4 text-muted-foreground" />
                <div>
                  <p className="text-sm text-muted-foreground">Payment Status</p>
                  <p className="font-medium capitalize">{order.payment_status || "-"}</p>
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
                      Size: {item.size || "-"} | Color: {item.color || "-"} | Qty: {item.quantity ?? 0}
                    </p>
                  </div>
                  <p className="font-semibold">
                    ${((item.price ?? 0) * (item.quantity ?? 0)).toFixed(2)}
                  </p>
                </div>
              ))}
            </div>
          </CardContent>
        </Card>
      </main>
      <Footer />
    </div>
  )
}
