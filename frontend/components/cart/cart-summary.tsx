"use client"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { useCart } from "@/contexts/cart-context"
import { useRouter } from "next/navigation"

export function CartSummary() {
  const { items, totalItems, totalPrice } = useCart()
  const router = useRouter()

  const shipping = totalPrice > 100 ? 0 : 10
  const tax = totalPrice * 0.08
  const finalTotal = totalPrice + shipping + tax

  const handleCheckout = () => {
    router.push("/checkout")
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Order Summary</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        <div className="flex justify-between">
          <span>Subtotal ({totalItems} items)</span>
          <span>${totalPrice.toFixed(2)}</span>
        </div>

        <div className="flex justify-between">
          <span>Shipping</span>
          <span>{shipping === 0 ? "Free" : `$${shipping.toFixed(2)}`}</span>
        </div>

        <div className="flex justify-between">
          <span>Tax</span>
          <span>${tax.toFixed(2)}</span>
        </div>

        <Separator />

        <div className="flex justify-between font-semibold text-lg">
          <span>Total</span>
          <span>${finalTotal.toFixed(2)}</span>
        </div>

        {totalPrice > 100 && <p className="text-sm text-green-600">🎉 You qualify for free shipping!</p>}

        <Button className="w-full" size="lg" onClick={handleCheckout} disabled={items.length === 0}>
          Proceed to Checkout
        </Button>

        <p className="text-xs text-muted-foreground text-center">Secure checkout powered by SSL encryption</p>
      </CardContent>
    </Card>
  )
}
