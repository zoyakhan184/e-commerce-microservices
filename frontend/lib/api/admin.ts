import axios from "../client"
import type { AdminDashboard, User, Product, Order, Payment, Category , ActivityItem } from "@/types"

export const adminApi = {
  // DASHBOARD
  getDashboard: async (): Promise<AdminDashboard> => {
    const res = await axios.get("/admin/dashboard")
    return res.data
  },

  getRecentActivity: async (): Promise<ActivityItem[]> => {
    const res = await axios.get("/admin/activity")
    return res.data
  },
  
  // USERS
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

  // PRODUCTS
  getProducts: async (): Promise<Product[]> => {
    const res = await axios.get("/products")
    return res.data.products
  },

  getCategories: async (): Promise<Category[]> => {
    const res = await axios.get<Category[]>("/api/categories")
    return res.data
  },

  createProduct: async (product: Omit<Product, "id">): Promise<Product> => {
    const res = await axios.post("/products", {
      name: product.name,
      description: product.description,
      price: product.price,
      brand: product.brand,
      category_id: product.categoryId,
      image_urls: product.imageUrls,
      sizes: product.sizes,
      colors: product.colors,
    })
    return res.data
  },

  updateProduct: async (productId: string, updates: Partial<Product>): Promise<Product> => {
    const res = await axios.put(`/products/${productId}`, updates)
    return res.data
  },

  deleteProduct: async (productId: string): Promise<void> => {
    await axios.delete(`/products/${productId}`)
  },

  // ORDERS
  getOrders: async (): Promise<Order[]> => {
    const res = await axios.get("/admin/orders")
    return res.data.orders
  },

  updateOrderStatus: async (orderId: string, status: string): Promise<void> => {
    await axios.put(`/orders/${orderId}/status`, { status })
  },

  // PAYMENTS
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

  // ADMIN PROFILE
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