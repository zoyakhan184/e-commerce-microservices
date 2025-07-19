"use client"

import { useEffect, useState } from "react"
import { orderApi } from "@/lib/api/order"
import { productsApi } from "@/lib/api/products"
import type { Order, Product } from "@/types"

import { Header } from "@/components/layout/header"
import { Footer } from "@/components/layout/footer"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Package, Eye, Calendar, DollarSign, Truck } from "lucide-react"

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"

export default function OrdersPage() {
  const [orders, setOrders] = useState<Order[]>([])
  const [loading, setLoading] = useState(true)
  const [cancellingOrderId, setCancellingOrderId] = useState<string | null>(null)
  const [selectedOrder, setSelectedOrder] = useState<Order | null>(null)
  const [productDetailsMap, setProductDetailsMap] = useState<Record<string, Product>>({})

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

  const handleCancelOrder = async (orderId: string) => {
    if (!confirm("Are you sure you want to cancel this order?")) return

    setCancellingOrderId(orderId)
    try {
      await orderApi.cancelOrder(orderId)
      setOrders((prev) =>
        prev.map((order) =>
          order.id === orderId ? { ...order, status: "cancelled" } : order
        )
      )
    } catch (err) {
      console.error("❌ Cancel order failed:", err)
      alert("Failed to cancel order.")
    } finally {
      setCancellingOrderId(null)
    }
  }

  const handleViewDetails = async (order: Order) => {
    const details: Record<string, Product> = {}

    await Promise.all(
      order.items.map(async (item) => {
        if (!productDetailsMap[item.product_id]) {
          try {
            const product = await productsApi.getProduct(item.product_id)
            details[item.product_id] = product
          } catch (err) {
            console.error("❌ Failed to fetch product:", item.product_id, err)
          }
        }
      })
    )

    setProductDetailsMap((prev) => ({ ...prev, ...details }))
    setSelectedOrder(order)
  }

  const getImageSrc = (imageUrl: string | undefined) => {
    if (!imageUrl) return "/placeholder.svg"
    if (imageUrl.startsWith("data:image")) {
      return imageUrl.replace(
        /^data:image\/jpeg;base64,?data:image\/jpeg;base64,?/,
        "data:image/jpeg;base64,"
      )
    }
    return `/images/${imageUrl}`
  }

  const formatDate = (date: Date) =>
    date.toLocaleDateString("en-US", {
      year: "numeric",
      month: "long",
      day: "numeric",
    })

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
                <p className="text-muted-foreground mb-6">
                  When you place your first order, it will appear here.
                </p>
                <Button asChild>
                  <a href="/products">Start Shopping</a>
                </Button>
              </CardContent>
            </Card>
          ) : (
            <div className="space-y-6">
              {orders.map((order) => {
                const orderDate = new Date(order.created_at)
                const expectedDelivery = new Date(orderDate.getTime() + 7 * 24 * 60 * 60 * 1000)

                return (
                  <Card key={order.id}>
                    <CardHeader>
                      <div className="flex items-center justify-between">
                        <CardTitle className="flex items-center gap-2">
                          <Package className="h-5 w-5" />
                          Order #{order.id}
                        </CardTitle>
                        <Badge
                          variant={order.status === "delivered" ? "default" : "secondary"}
                          className={
                            order.status === "delivered"
                              ? "bg-green-100 text-green-700"
                              : order.status === "cancelled"
                              ? "bg-red-100 text-red-700"
                              : ""
                          }
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
                            <p className="font-medium">{formatDate(orderDate)}</p>
                          </div>
                        </div>

                        <div className="flex items-center gap-2">
                          <Truck className="h-4 w-4 text-muted-foreground" />
                          <div>
                            <p className="text-sm text-muted-foreground">Expected Delivery</p>
                            <p className="font-medium">{formatDate(expectedDelivery)}</p>
                          </div>
                        </div>

                        <div className="flex items-center gap-2">
                          <DollarSign className="h-4 w-4 text-muted-foreground" />
                          <div>
                            <p className="text-sm text-muted-foreground">Total Amount</p>
                            <p className="font-medium">${order.total_amount?.toFixed(2)}</p>
                          </div>
                        </div>
                      </div>
                      <div className="flex justify-end mt-4 gap-2">
                        <Button
                          variant="outline"
                          size="sm"
                          onClick={() => handleViewDetails(order)}
                        >
                          <Eye className="h-4 w-4 mr-2" />
                          View Details
                        </Button>

                        {["pending", "processing"].includes(order.status) && (
                          <Button
                            variant="destructive"
                            size="sm"
                            onClick={() => handleCancelOrder(order.id)}
                            disabled={cancellingOrderId === order.id}
                          >
                            {cancellingOrderId === order.id ? "Cancelling..." : "Cancel Order"}
                          </Button>
                        )}
                      </div>
                    </CardContent>
                  </Card>
                )
              })}
            </div>
          )}
        </div>
      </main>
      <Footer />

      {/* View Details Modal */}
      <Dialog open={!!selectedOrder} onOpenChange={() => setSelectedOrder(null)}>
        <DialogContent className="max-w-3xl max-h-[90vh] overflow-auto">
          <DialogHeader>
            <DialogTitle>Order #{selectedOrder?.id} - Items</DialogTitle>
          </DialogHeader>

          <div className="space-y-4">
            {selectedOrder?.items.map((item, idx) => {
              const product = productDetailsMap[item.product_id]
              const rawImage =
                Array.isArray(product?.image_urls) && product.image_urls.length > 0
                  ? product.image_urls[0]
                  : typeof product?.image_urls === "string"
                  ? product.image_urls
                  : ""
              const imageSrc = getImageSrc(rawImage)
              const price = item.price || product?.price || 0

              return (
                <div
                  key={idx}
                  className="flex items-center gap-4 p-4 border rounded-lg"
                >
                  <div className="w-24 h-24 relative flex-shrink-0 rounded overflow-hidden">
                    <img
                      src={imageSrc}
                      alt={product?.name || item.product_id}
                      className="object-cover w-full h-full"
                    />
                  </div>
                  <div className="flex-1">
                    <p className="font-semibold">{product?.name || item.product_id}</p>
                    <p className="text-sm text-muted-foreground">
                      Size: {item.size || "-"} | Color: {item.color || "-"} | Qty: {item.quantity}
                    </p>
                    <p className="font-medium mt-1">${(price * item.quantity).toFixed(2)}</p>
                  </div>
                </div>
              )
            })}
          </div>
        </DialogContent>
      </Dialog>
    </div>
  )
}
