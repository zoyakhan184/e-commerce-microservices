"use client"

import { createContext, useContext, useState, useEffect } from "react"
import { cartApi } from "@/lib/api/cart"
import { useToast } from "@/hooks/use-toast"
import type { CartItem } from "@/types"

interface CartContextType {
  items: CartItem[]
  loading: boolean
  totalItems: number
  totalPrice: number // optional (can be removed or mapped to total)
  subtotal: number
  shipping: number
  image_url?: string
  total: number
  addToCart: (item: Omit<CartItem, "quantity"> & { quantity?: number }) => Promise<void>
  removeFromCart: (item: { product_id: string; size: string; color: string }) => Promise<void>
  updateQuantity: (item: { product_id: string; size: string; color: string; quantity: number }) => Promise<void>
}


const CartContext = createContext<CartContextType | undefined>(undefined)

  export function CartProvider({ children }: { children: React.ReactNode }) {
    const [items, setItems] = useState<CartItem[]>([])
    const [loading, setLoading] = useState(true)
    const [subtotal, setSubtotal] = useState(0)
    const [shipping, setShipping] = useState(0)
    const [total, setTotal] = useState(0)

    const { toast } = useToast()

    const fetchCart = async () => {
    setLoading(true)
    try {
      const res = await cartApi.getCart()
      const cleaned = res.items?.filter(
        (item: CartItem) => item.product_id && item.quantity != null
      ) || []

      setItems(cleaned)
      setSubtotal(res.subtotal ?? 0)
      setShipping(res.shipping ?? 0)
      setTotal(res.total ?? 0)

      console.log("[CartContext] 🛒 Cleaned Cart:", cleaned)
    } catch (err) {
      console.error("[CartContext] ❌ Failed to fetch cart:", err)
      toast({
        title: "Error",
        description: "Failed to load cart",
        variant: "destructive",
      })
    } finally {
      setLoading(false)
    }
  }


  useEffect(() => {
    fetchCart()
  }, [])

  const addToCart = async (newItem: Omit<CartItem, "quantity"> & { quantity?: number }) => {
    const quantity = newItem.quantity ?? 1
    try {
      await cartApi.addToCart(newItem.product_id, newItem.size, newItem.color, quantity)
      await fetchCart()
      toast({ title: "Added to Cart", description: `${newItem.product_name} added successfully.` })
    } catch (err) {
      console.error("[CartContext] ❌ Failed to add to cart:", err)
      toast({ title: "Error", description: "Failed to add item to cart", variant: "destructive" })
    }
  }

  const removeFromCart = async ({ product_id, size, color }: { product_id: string; size: string; color: string }) => {
    try {
      await cartApi.removeFromCart(product_id, size, color)
      setItems((prev) => prev.filter(item => !(item.product_id === product_id && item.size === size && item.color === color)))
      toast({ title: "Item Removed", description: "Item removed from cart." })
    } catch (err) {
      console.error("[CartContext] ❌ Failed to remove item:", err)
      toast({ title: "Error", description: "Failed to remove item", variant: "destructive" })
    }
  }

  const updateQuantity = async ({
    product_id,
    size,
    color,
    quantity,
  }: {
    product_id: string
    size: string
    color: string
    quantity: number
  }) => {
    try {
      if (quantity <= 0) {
        // Remove from backend
        await cartApi.removeFromCart(product_id, size, color)
        // Remove locally to instantly reflect in UI
        setItems((prev) => prev.filter(item => !(item.product_id === product_id && item.size === size && item.color === color)))
        toast({ title: "Item Removed", description: "Item removed from cart." })
        return
      }

      await cartApi.updateCartItem(product_id, size, color, quantity)
      await fetchCart()
    } catch (err) {
      console.error("[CartContext] ❌ Failed to update quantity:", err)
      toast({ title: "Error", description: "Failed to update quantity", variant: "destructive" })
    }
  }

  const totalItems = items.reduce((sum, item) => sum + (item.quantity ?? 0), 0)
  const totalPrice = items.reduce((sum, item) => sum + (item.quantity ?? 0) * (item.price ?? 0), 0)

  return (
    <CartContext.Provider
      value={{
        items,
        loading,
        totalItems: items.reduce((sum, item) => sum + (item.quantity ?? 0), 0),
        totalPrice: total, // can be deprecated if you switch to `total` directly
        subtotal,
        shipping,
        total,
        addToCart,
        removeFromCart,
        updateQuantity,
      }}
    >
  {children}
</CartContext.Provider>

  )
}

export function useCart() {
  const context = useContext(CartContext)
  if (!context) {
    throw new Error("useCart must be used within a CartProvider")
  }
  return context
}
