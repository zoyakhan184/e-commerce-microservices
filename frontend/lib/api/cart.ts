import axios from "../client"
import type { CartItem } from "@/types"

export type CartSummaryResponse = {
  items: CartItem[]
  subtotal: number
  shipping: number
  total: number
}

export const cartApi = {
  // ✅ Get full cart summary
  async getCart(): Promise<CartSummaryResponse> {
    const res = await axios.get("/cart")
    return res.data
  },

  // ✅ Add item to cart
  async addToCart(productId: string, size: string, color: string, quantity: number): Promise<void> {
    await axios.post("/cart", {
      product_id: productId,
      size,
      color,
      quantity,
    })
  },

  // ✅ Remove single item from cart (FIXED: should use DELETE)
  async removeFromCart(productId: string, size: string, color: string): Promise<void> {
    await axios.delete("/cart", {
      data: {
        product_id: productId,
        size,
        color,
      },
    })
  },

  // ✅ Update cart item
  async updateCartItem(productId: string, size: string, color: string, quantity: number): Promise<void> {
    await axios.put("/cart", {
      product_id: productId,
      size,
      color,
      quantity,
    })
  },

  // ✅ Clear the entire cart
  async clearCart(): Promise<void> {
  try {
    await axios.delete("/cart/clear")
  } catch (err) {
    console.error("❌ Failed to clear cart:", err)
  }
}

}
