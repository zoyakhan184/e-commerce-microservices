import axios from "../client"
import type { Order } from "@/types"

export type CartItem = {
  product_id: string
  quantity: number
}

export const orderApi = {
  async placeOrder(items: CartItem[]): Promise<{ orderId: string }> {
    const res = await axios.post("api/orders", { items })
    return res.data
  },

  async clearCart(): Promise<void> {
    await axios.delete("api/cart/clear")
  },

  async getOrders(): Promise<Order[]> {
    const res = await axios.get("api/orders")
    return res.data
  },

  async getOrderDetails(orderId: string): Promise<Order> {
    const res = await axios.get(`api/orders/${orderId}`)
    return res.data
  },
}
