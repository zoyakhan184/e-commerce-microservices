"use client"

import { useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { AdminLayout } from "@/components/admin/admin-layout"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Badge } from "@/components/ui/badge"
import {
  Table, TableBody, TableCell, TableHead, TableHeader, TableRow
} from "@/components/ui/table"
import {
  Tabs, TabsContent, TabsList, TabsTrigger
} from "@/components/ui/tabs"
import {
  DropdownMenu, DropdownMenuContent, DropdownMenuItem,
  DropdownMenuSeparator, DropdownMenuTrigger
} from "@/components/ui/dropdown-menu"
import {
  Search, MoreHorizontal, Eye, Edit, Package, Clock, CheckCircle, RotateCcw
} from "lucide-react"
import { orderApi } from "@/lib/api/order"
import { useToast } from "@/hooks/use-toast"
import type { Order } from "@/types"

export default function AdminOrdersPage() {
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedTab, setSelectedTab] = useState("all")
  const { toast } = useToast()

  const {
    data: orders,
    isLoading,
    refetch,
  } = useQuery({
    queryKey: ["admin-orders"],
    queryFn: orderApi.getOrders, // ✅ Updated
  })

  const filteredOrders = orders?.filter((order) => {
    const matchesSearch =
      order.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
      order.items.some((item) => item.product_name?.toLowerCase().includes(searchQuery.toLowerCase()))

    const matchesTab =
      selectedTab === "all" ||
      (selectedTab === "pending" && order.status === "pending") ||
      (selectedTab === "in_progress" && order.status === "in_progress") ||
      (selectedTab === "completed" && order.status === "completed") ||
      (selectedTab === "delivered" && order.status === "delivered") ||
      (selectedTab === "returned" && order.status === "returned")

    return matchesSearch && matchesTab
  }) || []

  const handleUpdateOrderStatus = async (orderId: string, newStatus: string) => {
    try {
      // You may need to rewire this if `orderApi.updateOrderStatus` is not exposed
      // You can replug `adminApi.updateOrderStatus` here if only admin can change statuses
      toast({
        title: "Order update not implemented",
        description: "Hook up updateOrderStatus to orderApi or adminApi.",
        variant: "destructive",
      })
    } catch (error: any) {
      toast({
        title: "Error",
        description: error?.response?.data?.message || "Failed to update order status. Please try again.",
        variant: "destructive",
      })
    }
  }

  const getStatusBadge = (status: string) => {
    type BadgeVariant = "default" | "destructive" | "secondary" | "outline";
    const statusConfig: Record<string, { variant: BadgeVariant, color: string, icon: React.ElementType }> = {
      pending: { variant: "secondary", color: "bg-yellow-100 text-yellow-700", icon: Clock },
      in_progress: { variant: "default", color: "bg-blue-100 text-blue-700", icon: Package },
      completed: { variant: "default", color: "bg-green-100 text-green-700", icon: CheckCircle },
      delivered: { variant: "default", color: "bg-green-100 text-green-700", icon: CheckCircle },
      returned: { variant: "destructive", color: "bg-red-100 text-red-700", icon: RotateCcw },
    }

    const config = statusConfig[status as keyof typeof statusConfig] || statusConfig.pending
    const Icon = config.icon

    return (
      <Badge variant={config.variant} className={config.color}>
        <Icon className="h-3 w-3 mr-1" />
        {status.replace("_", " ")}
      </Badge>
    )
  }

  const getUserName = (userId: string) => userId // You can enhance this with a user lookup if needed

  const getOrderStats = () => {
    if (!orders) return { total: 0, pending: 0, in_progress: 0, completed: 0, delivered: 0, returned: 0 }
    return {
      total: orders.length,
      pending: orders.filter((o) => o.status === "pending").length,
      in_progress: orders.filter((o) => o.status === "in_progress").length,
      completed: orders.filter((o) => o.status === "completed").length,
      delivered: orders.filter((o) => o.status === "delivered").length,
      returned: orders.filter((o) => o.status === "returned").length,
    }
  }

  const stats = getOrderStats()

  return (
    <AdminLayout>
      <div className="space-y-6">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold">Orders Management</h1>
            <p className="text-muted-foreground">Monitor and manage all customer orders ({stats.total} orders)</p>
          </div>
        </div>

        {/* Stats */}
        <div className="grid grid-cols-1 md:grid-cols-6 gap-6">
          {[
            { label: "Total", value: stats.total, icon: Package, color: "text-blue-500" },
            { label: "Pending", value: stats.pending, icon: Clock, color: "text-yellow-500" },
            { label: "In Progress", value: stats.in_progress, icon: Package, color: "text-blue-500" },
            { label: "Completed", value: stats.completed, icon: CheckCircle, color: "text-green-500" },
            { label: "Delivered", value: stats.delivered, icon: CheckCircle, color: "text-green-500" },
            { label: "Returned", value: stats.returned, icon: RotateCcw, color: "text-red-500" },
          ].map((stat, i) => (
            <Card key={i}>
              <CardContent className="p-6">
                <div className="flex items-center justify-between">
                  <div>
                    <p className="text-sm font-medium text-muted-foreground">{stat.label}</p>
                    <p className={`text-2xl font-bold ${stat.color}`}>{stat.value}</p>
                  </div>
                  <stat.icon className={`h-8 w-8 ${stat.color}`} />
                </div>
              </CardContent>
            </Card>
          ))}
        </div>

        {/* Orders Table */}
        <Card>
          <CardHeader>
            <CardTitle>All Orders</CardTitle>
          </CardHeader>
          <CardContent>
            <Tabs value={selectedTab} onValueChange={setSelectedTab} className="space-y-4">
              <div className="flex flex-col sm:flex-row gap-4 items-start sm:items-center justify-between">
                <TabsList>
                  {["all", "pending", "in_progress", "completed", "delivered", "returned"].map((status) => (
                    <TabsTrigger key={status} value={status}>
                      {status.replace("_", " ")}
                    </TabsTrigger>
                  ))}
                </TabsList>
                <div className="relative">
                  <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
                  <Input
                    placeholder="Search orders..."
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                    className="pl-8 w-64"
                  />
                </div>
              </div>

              <TabsContent value={selectedTab} className="space-y-4">
                <div className="rounded-md border">
                  <Table>
                    <TableHeader>
                      <TableRow>
                        <TableHead>Order ID</TableHead>
                        <TableHead>Customer</TableHead>
                        <TableHead>Items</TableHead>
                        <TableHead>Total Amount</TableHead>
                        <TableHead>Status</TableHead>
                        <TableHead>Payment</TableHead>
                        <TableHead>Date</TableHead>
                        <TableHead className="text-right">Actions</TableHead>
                      </TableRow>
                    </TableHeader>
                    <TableBody>
                      {filteredOrders.length === 0 ? (
                        <TableRow>
                          <TableCell colSpan={8} className="text-center py-8 text-muted-foreground">
                            No orders found matching your criteria.
                          </TableCell>
                        </TableRow>
                      ) : (
                        filteredOrders.map((order) => (
                          <TableRow key={order.id}>
                            <TableCell className="font-medium">#{order.id}</TableCell>
                            <TableCell>{getUserName(order.user_id)}</TableCell>
                            <TableCell>
                              <div className="space-y-1">
                                {order.items.slice(0, 2).map((item, i) => (
                                  <div key={i} className="text-sm">
                                    {item.product_name} x{item.quantity}
                                  </div>
                                ))}
                                {order.items.length > 2 && (
                                  <div className="text-xs text-muted-foreground">
                                    +{order.items.length - 2} more items
                                  </div>
                                )}
                              </div>
                            </TableCell>
                            <TableCell className="font-semibold">${order.total_amount.toFixed(2)}</TableCell>
                            <TableCell>{getStatusBadge(order.status)}</TableCell>
                            <TableCell>
                              <Badge
                                variant={order.payment_status === "paid" ? "default" : "secondary"}
                                className={
                                  order.payment_status === "paid"
                                    ? "bg-green-100 text-green-700"
                                    : "bg-yellow-100 text-yellow-700"
                                }
                              >
                                {order.payment_status}
                              </Badge>
                            </TableCell>
                            <TableCell>{new Date(order.created_at).toLocaleDateString()}</TableCell>
                            <TableCell className="text-right">
                              <DropdownMenu>
                                <DropdownMenuTrigger asChild>
                                  <Button variant="ghost" size="icon">
                                    <MoreHorizontal className="h-4 w-4" />
                                  </Button>
                                </DropdownMenuTrigger>
                                <DropdownMenuContent align="end">
                                  <DropdownMenuItem>
                                    <Eye className="h-4 w-4 mr-2" />
                                    View Details
                                  </DropdownMenuItem>
                                  <DropdownMenuItem>
                                    <Edit className="h-4 w-4 mr-2" />
                                    Edit Order
                                  </DropdownMenuItem>
                                  <DropdownMenuSeparator />
                                  {/* You can rewire below handlers based on actual status change support */}
                                </DropdownMenuContent>
                              </DropdownMenu>
                            </TableCell>
                          </TableRow>
                        ))
                      )}
                    </TableBody>
                  </Table>
                </div>
              </TabsContent>
            </Tabs>
          </CardContent>
        </Card>
      </div>
    </AdminLayout>
  )
}
