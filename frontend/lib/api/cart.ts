import axios from "../client"
import type { CartItem } from "@/types"

export type CartSummaryResponse = {
  items: CartItem[]
  subtotal: number
  shipping: number
  total: number
}

export const cartApi = {
  // ✅ Updated
  async getCart(): Promise<CartSummaryResponse> {
    const res = await axios.get("/cart")
    return res.data // now expects { items, subtotal, shipping, total }
  },

  // ✅ Already fine
  async addToCart(productId: string, size: string, color: string, quantity: number): Promise<void> {
    await axios.post("/cart", {
      product_id: productId,
      size,
      color,
      quantity,
    })
  },

  async removeFromCart(productId: string, size: string, color: string): Promise<void> {
    await axios.post("/cart", {
      product_id: productId,
      size,
      color,
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
}
