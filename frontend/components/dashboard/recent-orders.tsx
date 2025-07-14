"use client"

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Package, Eye } from "lucide-react"
import Link from "next/link"
import { dummyOrders } from "@/lib/dummy-data"
export function RecentOrders() {
  const recentOrders = dummyOrders.slice(0, 3)

  if (recentOrders.length === 0) {
    return (
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Package className="h-5 w-5" />
            Recent Orders
          </CardTitle>
        </CardHeader>
        <CardContent>
          <div className="text-center py-8 text-muted-foreground">
            <Package className="h-12 w-12 mx-auto mb-4 opacity-50" />
            <p>No orders yet</p>
            <Button asChild className="mt-4">
              <Link href="/products">Start Shopping</Link>
            </Button>
          </div>
        </CardContent>
      </Card>
    )
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          <Package className="h-5 w-5" />
          Recent Orders
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {recentOrders.map((order) => (
          <div key={order.id} className="flex items-center justify-between p-4 border rounded-lg">
            <div className="space-y-1">
              <p className="font-medium">Order #{order.id}</p>
              <p className="text-sm text-muted-foreground">{new Date(order.created_at).toLocaleDateString()}</p>
              <p className="font-semibold">${order.total_amount.toFixed(2)}</p>
            </div>
            <div className="flex items-center gap-2">
              <Badge
                variant={order.status === "delivered" ? "default" : "secondary"}
                className={order.status === "delivered" ? "bg-green-100 text-green-700" : ""}
              >
                {order.status}
              </Badge>
              <Button size="sm" variant="outline">
                <Eye className="h-4 w-4" />
              </Button>
            </div>
          </div>
        ))}

        <Button asChild variant="outline" className="w-full bg-transparent">
          <Link href="/orders">View All Orders</Link>
        </Button>
      </CardContent>
    </Card>
  )
}
