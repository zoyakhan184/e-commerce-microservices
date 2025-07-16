import axios from "../client"
import type { AdminDashboard, User, Product, Order, Payment, Category , ActivityItem } from "@/types"

export const adminApi = {
  getDashboard: async (): Promise<AdminDashboard> => {
    const res = await axios.get("/admin/dashboard")
    return res.data
  },

  getRecentActivity: async (): Promise<ActivityItem[]> => {
    const res = await axios.get("/admin/activity")
    return res.data
  },

  getUsers: async (): Promise<User[]> => {
    const res = await axios.get("/admin/user-profiles")
    return res.data
  },

  deleteUser: async (userId: string): Promise<void> => {
    await axios.delete(`/admin/users/${userId}`)
  },

  updateUserStatus: async (userId: string, status: string): Promise<void> => {
    await axios.put(`/admin/users/${userId}/status`, { status })
  },

  getProducts: async (): Promise<Product[]> => {
    const res = await axios.get("/products")
    return res.data.products
  },

  getCategories: async (): Promise<Category[]> => {
    const res = await axios.get("/categories")
    return res.data
  },

  createProduct: async (product: Omit<Product, "id">): Promise<Product> => {
    const res = await axios.post("/products", { ...product })
    return res.data
  },

  updateProduct: async (productId: string, updates: Partial<Product>): Promise<Product> => {
    const res = await axios.put(`/products/${productId}`, updates)
    return res.data
  },

  deleteProduct: async (productId: string): Promise<void> => {
    await axios.delete(`/products/${productId}`)
  },

 getOrders: async (): Promise<Order[]> => {
  const res = await axios.get("/admin/orders")

  return res.data.map((o: any) => ({
    id: o.order_id,
    user_id: o.user_id,
    status: o.status || o.order_status,  // 👈 key fix
    payment_status: o.payment_status || "unpaid",
    total_amount: o.total_amount,
    created_at: o.created_at,
    items: o.items || [],
  }))
},

  updateOrderStatus: async (orderId: string, status: string): Promise<void> => {
  await axios.put(`/orders/${orderId}/status`, { status })
}
,

  getPayments: async (): Promise<Payment[]> => {
    const res = await axios.get("/payments")
    return res.data.payments
  },

  updatePaymentStatus: async (paymentId: string, status: string): Promise<void> => {
    await axios.put(`/payments/${paymentId}/status`, { status })
  },

  processRefund: async (paymentId: string): Promise<void> => {
    await axios.post(`/payment/refund`, { payment_id: paymentId })
  },

  getAdminProfile: async (): Promise<User> => {
    const res = await axios.get("/users/profile")
    return res.data
  },

  updateAdminProfile: async (updates: Partial<User>): Promise<User> => {
    const res = await axios.put("/users/profile", updates)
    return res.data
  },

  changePassword: async (currentPassword: string, newPassword: string): Promise<void> => {
    await axios.post("/users/change-password", { currentPassword, newPassword })
  },
}
