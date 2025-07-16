import axios from "../client"
import type { CartItem } from "@/types"

export type CartSummaryResponse = {
  items: CartItem[]
  subtotal: number
  shipping: number
  total: number
}

export const cartApi = {
  async getCart(): Promise<CartSummaryResponse> {
    const res = await axios.get("/cart")
    return res.data
  },

  async addToCart(productId: string, size: string, color: string, quantity: number): Promise<void> {
    await axios.post("/cart", {
      product_id: productId,
      size,
      color,
      quantity,
    })
  },

  async removeFromCart(productId: string, size: string, color: string): Promise<void> {
    await axios.delete("/cart", {
      data: { product_id: productId, size, color },
    })
  },

  async updateCartItem(productId: string, size: string, color: string, quantity: number): Promise<void> {
    await axios.put("/cart", {
      product_id: productId,
      size,
      color,
      quantity,
    })
  },

  async clearCart(): Promise<void> {
    await axios.delete("/cart/clear")
  },
}

