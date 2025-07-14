"use client"

import { useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { AdminLayout } from "@/components/admin/admin-layout"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Badge } from "@/components/ui/badge"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
  Search,
  MoreHorizontal,
  Eye,
  RefreshCw,
  CreditCard,
  DollarSign,
  TrendingUp,
  AlertCircle,
  CheckCircle,
  XCircle,
} from "lucide-react"
import { adminApi } from "@/lib/api/admin"
import { dummyUsers } from "@/lib/dummy-data"
import { useToast } from "@/hooks/use-toast"

export default function AdminPaymentsPage() {
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedTab, setSelectedTab] = useState("all")
  const { toast } = useToast()

  const {
    data: payments,
    isLoading,
    refetch,
  } = useQuery({
    queryKey: ["admin-payments"],
    queryFn: adminApi.getPayments,
  })

  // Filter payments based on search and status
  const filteredPayments =
    payments?.filter((payment) => {
      const matchesSearch =
        payment.id.toLowerCase().includes(searchQuery.toLowerCase()) ||
        payment.order_id.toLowerCase().includes(searchQuery.toLowerCase()) ||
        payment.txn_ref?.toLowerCase().includes(searchQuery.toLowerCase())

      const matchesTab =
        selectedTab === "all" ||
        (selectedTab === "successful" && payment.status === "success") ||
        (selectedTab === "pending" && payment.status === "pending") ||
        (selectedTab === "failed" && payment.status === "failed") ||
        (selectedTab === "refunded" && payment.status === "refunded")

      return matchesSearch && matchesTab
    }) || []

  const handleRefund = async (paymentId: string, amount: number) => {
    if (window.confirm(`Are you sure you want to refund $${amount.toFixed(2)}?`)) {
      try {
        await adminApi.processRefund(paymentId, amount)
        toast({
          title: "Refund initiated",
          description: "The refund has been processed successfully.",
        })
        refetch()
      } catch (error) {
        toast({
          title: "Error",
          description: "Failed to process refund. Please try again.",
          variant: "destructive",
        })
      }
    }
  }

  const handleUpdatePaymentStatus = async (paymentId: string, newStatus: string) => {
    try {
      await adminApi.updatePaymentStatus(paymentId, newStatus)
      toast({
        title: "Payment status updated",
        description: `Payment status has been updated to ${newStatus}.`,
      })
      refetch()
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to update payment status. Please try again.",
        variant: "destructive",
      })
    }
  }

  const getStatusBadge = (status: string) => {
    const statusConfig = {
      success: { variant: "default" as const, color: "bg-green-100 text-green-700", icon: CheckCircle },
      pending: { variant: "secondary" as const, color: "bg-yellow-100 text-yellow-700", icon: AlertCircle },
      failed: { variant: "destructive" as const, color: "bg-red-100 text-red-700", icon: XCircle },
      refunded: { variant: "outline" as const, color: "bg-blue-100 text-blue-700", icon: RefreshCw },
    }

    const config = statusConfig[status as keyof typeof statusConfig] || statusConfig.pending
    const Icon = config.icon

    return (
      <Badge variant={config.variant} className={config.color}>
        <Icon className="h-3 w-3 mr-1" />
        {status}
      </Badge>
    )
  }

  const getUserName = (userId: string) => {
    const user = dummyUsers.find((u) => u.id === userId)
    return user ? user.name : `User ${userId.slice(-4)}`
  }

  const getPaymentStats = () => {
    if (!payments) return { total: 0, successful: 0, pending: 0, failed: 0, refunded: 0, totalAmount: 0 }

    return {
      total: payments.length,
      successful: payments.filter((p) => p.status === "success").length,
      pending: payments.filter((p) => p.status === "pending").length,
      failed: payments.filter((p) => p.status === "failed").length,
      refunded: payments.filter((p) => p.status === "refunded").length,
      totalAmount: payments.filter((p) => p.status === "success").reduce((sum, p) => sum + p.amount, 0),
    }
  }

  const stats = getPaymentStats()

  if (isLoading) {
    return (
      <AdminLayout>
        <div className="space-y-6">
          <h1 className="text-3xl font-bold">Payments Management</h1>
          <Card>
            <CardContent className="p-6">
              <div className="animate-pulse space-y-4">
                {Array.from({ length: 5 }).map((_, i) => (
                  <div key={i} className="h-16 bg-muted rounded"></div>
                ))}
              </div>
            </CardContent>
          </Card>
        </div>
      </AdminLayout>
    )
  }

  return (
    <AdminLayout>
      <div className="space-y-6">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold">Payments Management</h1>
            <p className="text-muted-foreground">
              Monitor payments, process refunds, and manage transactions ({stats.total} payments)
            </p>
          </div>
        </div>

        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-6 gap-6">
          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Total Revenue</p>
                  <p className="text-2xl font-bold text-green-600">${stats.totalAmount.toFixed(2)}</p>
                </div>
                <DollarSign className="h-8 w-8 text-green-500" />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Successful</p>
                  <p className="text-2xl font-bold text-green-600">{stats.successful}</p>
                </div>
                <CheckCircle className="h-8 w-8 text-green-500" />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Pending</p>
                  <p className="text-2xl font-bold text-yellow-600">{stats.pending}</p>
                </div>
                <AlertCircle className="h-8 w-8 text-yellow-500" />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Failed</p>
                  <p className="text-2xl font-bold text-red-600">{stats.failed}</p>
                </div>
                <XCircle className="h-8 w-8 text-red-500" />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Refunded</p>
                  <p className="text-2xl font-bold text-blue-600">{stats.refunded}</p>
                </div>
                <RefreshCw className="h-8 w-8 text-blue-500" />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Success Rate</p>
                  <p className="text-2xl font-bold text-blue-600">
                    {stats.total > 0 ? ((stats.successful / stats.total) * 100).toFixed(1) : 0}%
                  </p>
                </div>
                <TrendingUp className="h-8 w-8 text-blue-500" />
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Payments Table with Tabs */}
        <Card>
          <CardHeader>
            <CardTitle>Payment Transactions</CardTitle>
          </CardHeader>
          <CardContent>
            <Tabs value={selectedTab} onValueChange={setSelectedTab} className="space-y-4">
              <div className="flex flex-col sm:flex-row gap-4 items-start sm:items-center justify-between">
                <TabsList>
                  <TabsTrigger value="all">All Payments</TabsTrigger>
                  <TabsTrigger value="successful">Successful</TabsTrigger>
                  <TabsTrigger value="pending">Pending</TabsTrigger>
                  <TabsTrigger value="failed">Failed</TabsTrigger>
                  <TabsTrigger value="refunded">Refunded</TabsTrigger>
                </TabsList>

                <div className="relative">
                  <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
                  <Input
                    placeholder="Search payments..."
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
                        <TableHead>Payment ID</TableHead>
                        <TableHead>Order ID</TableHead>
                        <TableHead>Customer</TableHead>
                        <TableHead>Amount</TableHead>
                        <TableHead>Gateway</TableHead>
                        <TableHead>Status</TableHead>
                        <TableHead>Transaction Ref</TableHead>
                        <TableHead>Date</TableHead>
                        <TableHead className="text-right">Actions</TableHead>
                      </TableRow>
                    </TableHeader>
                    <TableBody>
                      {filteredPayments.length === 0 ? (
                        <TableRow>
                          <TableCell colSpan={9} className="text-center py-8 text-muted-foreground">
                            No payments found matching your criteria.
                          </TableCell>
                        </TableRow>
                      ) : (
                        filteredPayments.map((payment) => (
                          <TableRow key={payment.id}>
                            <TableCell className="font-medium">#{payment.id}</TableCell>
                            <TableCell>#{payment.order_id}</TableCell>
                            <TableCell>{getUserName(payment.user_id)}</TableCell>
                            <TableCell className="font-semibold">${payment.amount.toFixed(2)}</TableCell>
                            <TableCell>
                              <div className="flex items-center gap-2">
                                <CreditCard className="h-4 w-4" />
                                {payment.gateway}
                              </div>
                            </TableCell>
                            <TableCell>{getStatusBadge(payment.status)}</TableCell>
                            <TableCell className="font-mono text-sm">{payment.txn_ref || "N/A"}</TableCell>
                            <TableCell>{new Date(payment.created_at).toLocaleDateString()}</TableCell>
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
                                  {payment.status === "pending" && (
                                    <>
                                      <DropdownMenuSeparator />
                                      <DropdownMenuItem
                                        onClick={() => handleUpdatePaymentStatus(payment.id, "success")}
                                      >
                                        Mark as Success
                                      </DropdownMenuItem>
                                      <DropdownMenuItem onClick={() => handleUpdatePaymentStatus(payment.id, "failed")}>
                                        Mark as Failed
                                      </DropdownMenuItem>
                                    </>
                                  )}
                                  {payment.status === "success" && (
                                    <>
                                      <DropdownMenuSeparator />
                                      <DropdownMenuItem
                                        className="text-orange-600"
                                        onClick={() => handleRefund(payment.id, payment.amount)}
                                      >
                                        <RefreshCw className="h-4 w-4 mr-2" />
                                        Process Refund
                                      </DropdownMenuItem>
                                    </>
                                  )}
                                  {payment.status === "failed" && (
                                    <>
                                      <DropdownMenuSeparator />
                                      <DropdownMenuItem
                                        onClick={() => handleUpdatePaymentStatus(payment.id, "pending")}
                                      >
                                        <RefreshCw className="h-4 w-4 mr-2" />
                                        Retry Payment
                                      </DropdownMenuItem>
                                    </>
                                  )}
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
