import axios from "../client"
import type { AdminDashboard, User, Product, Order, Payment, Category, ActivityItem } from "@/types"

export const adminApi = {
  // Dashboard data: total users, orders, revenue, low stock products
  getDashboard: async (): Promise<AdminDashboard> => {
    const res = await axios.get("api/admin/dashboard")
    return res.data
  },

  // Recent activity log (user registered, orders placed, etc.)
  getRecentActivity: async (): Promise<ActivityItem[]> => {
    const res = await axios.get("api/admin/activity")
    return res.data
  },

  // User management
  getUsers: async (): Promise<User[]> => {
    const res = await axios.get("api/admin/user-profiles")
    return res.data
  },

  updateUserStatus: async (userId: string, status: string): Promise<void> => {
    await axios.put(`api/admin/users/${userId}/status`, { status })
  },

  deleteUser: async (userId: string): Promise<void> => {
    await axios.delete(`api/admin/users/${userId}`)
  },

  // Product & category management
  getProducts: async (): Promise<Product[]> => {
    const res = await axios.get("api/products")
    return res.data.products
  },

  getCategories: async (): Promise<Category[]> => {
    const res = await axios.get("api/categories")
    return res.data
  },

  createProduct: async (product: Omit<Product, "id">): Promise<Product> => {
    const res = await axios.post("api/products", product)
    return res.data
  },

  updateProduct: async (productId: string, updates: Partial<Product>): Promise<Product> => {
    const res = await axios.put(`api/products/${productId}`, updates)
    return res.data
  },

  deleteProduct: async (productId: string): Promise<void> => {
    await axios.delete(`api/products/${productId}`)
  },

  // Order management
  getOrders: async (): Promise<Order[]> => {
    const res = await axios.get("api/admin/orders")
    return res.data.map((o: any) => ({
      id: o.order_id,
      user_id: o.user_id,
      status: o.status || o.order_status,
      payment_status: o.payment_status || "unpaid",
      total_amount: o.total_amount,
      created_at: o.created_at,
      items: o.items || [],
    }))
  },

  updateOrderStatus: async (orderId: string, status: string): Promise<void> => {
    await axios.put(`api/orders/${orderId}/status`, { status })
  },

  // Payment tracking
  getPayments: async (): Promise<Payment[]> => {
    const res = await axios.get("api/payments")
    return res.data.payments
  },

  updatePaymentStatus: async (paymentId: string, status: string): Promise<void> => {
    await axios.put(`api/payments/${paymentId}/status`, { status })
  },

  processRefund: async (paymentId: string): Promise<void> => {
    await axios.post(`api/payment/refund`, { payment_id: paymentId })
  },

  // Admin profile
  getAdminProfile: async (): Promise<User> => {
    const res = await axios.get("api/users/profile")
    return res.data
  },

  updateAdminProfile: async (updates: Partial<User>): Promise<User> => {
    const res = await axios.put("api/users/profile", updates)
    return res.data
  },

  changePassword: async (currentPassword: string, newPassword: string): Promise<void> => {
  await axios.post("api/auth/change-password", { currentPassword, newPassword }) // ✅ Correct path
}

}


