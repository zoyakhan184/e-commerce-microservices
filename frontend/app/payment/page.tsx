"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import Image from "next/image"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"
import { useCart } from "@/contexts/cart-context"
import { useAuth } from "@/contexts/auth-context"
import { Header } from "@/components/layout/header"
import { CreditCard, Truck, Calendar } from "lucide-react"
import { orderApi } from "@/lib/api/order"

export default function PaymentPage() {
  const { items, totalPrice } = useCart()
  const { user } = useAuth()
  const router = useRouter()

  const [paymentMethod, setPaymentMethod] = useState("card")
  const [isProcessing, setIsProcessing] = useState(false)
  const [showConfirmation, setShowConfirmation] = useState(false)
  const [orderId, setOrderId] = useState<string | null>(null)

  if (!user) {
    router.push("/auth/login")
    return null
  }

  if (items.length === 0) {
    router.push("/cart")
    return null
  }

  const shippingCost = 9.99
  const tax = totalPrice * 0.08
  const finalTotal = totalPrice + shippingCost + tax

  const handlePayment = async () => {
    setIsProcessing(true)
    try {
      if (paymentMethod === "card") {
        // TODO: Stripe integration
        return
      }

      const response = await orderApi.placeOrder(
        items.map((item) => ({
          product_id: item.product_id,
          quantity: item.quantity,
        }))
      )
      console.log("response: ",response)
      //await orderApi.clearCart()
      //clearCart()
      setOrderId(response.orderId)
      setShowConfirmation(true)
    } catch (err) {
      console.error("❌ Payment error:", err)
      alert("Something went wrong while placing your order.")
    } finally {
      setTimeout(() => {
        setIsProcessing(false)
      }, 1500)
    }
  }

  const expectedDelivery = new Date()
  expectedDelivery.setDate(expectedDelivery.getDate() + 5)

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <Header />
      <div className="container mx-auto px-4 py-8">
        <div className="max-w-4xl mx-auto">
          <h1 className="text-3xl font-bold mb-8">Payment</h1>

          <div className="grid lg:grid-cols-3 gap-8">
            {/* Payment Methods */}
            <div className="lg:col-span-2 space-y-6">
              <Card>
                <CardHeader>
                  <CardTitle>Payment Method</CardTitle>
                </CardHeader>
                <CardContent>
                  <RadioGroup value={paymentMethod} onValueChange={setPaymentMethod}>
                    <div className="flex items-center space-x-3 p-4 border rounded-lg">
                      <RadioGroupItem value="card" id="card" />
                      <Label htmlFor="card" className="flex items-center gap-2 cursor-pointer">
                        <CreditCard className="h-5 w-5" />
                        Stripe (Card)
                      </Label>
                    </div>
                    {paymentMethod === "card" && (
                      <div className="mt-4 space-y-4 px-4">
                        <Input type="text" placeholder="Cardholder Name" required />
                        <Input type="text" placeholder="Card Number" required />
                        <div className="flex gap-4">
                          <Input type="text" placeholder="MM/YY" required className="flex-1" />
                          <Input type="text" placeholder="CVC" required className="flex-1" />
                        </div>
                      </div>
                    )}
                    <div className="flex items-center space-x-3 p-4 border rounded-lg">
                      <RadioGroupItem value="cod" id="cod" />
                      <Label htmlFor="cod" className="flex items-center gap-2 cursor-pointer">
                        <Truck className="h-5 w-5" />
                        Cash on Delivery
                      </Label>
                    </div>
                  </RadioGroup>
                </CardContent>
              </Card>

              {/* Delivery Info */}
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center gap-2">
                    <Truck className="h-5 w-5" />
                    Delivery Information
                  </CardTitle>
                </CardHeader>
                <CardContent>
                  <div className="flex items-center gap-3 p-4 bg-green-50 dark:bg-green-900/20 rounded-lg">
                    <Calendar className="h-5 w-5 text-green-600" />
                    <div>
                      <p className="font-medium text-green-800 dark:text-green-400">Expected Delivery</p>
                      <p className="text-sm text-green-600 dark:text-green-300">
                        {expectedDelivery.toLocaleDateString("en-US", {
                          weekday: "long",
                          year: "numeric",
                          month: "long",
                          day: "numeric",
                        })}
                      </p>
                    </div>
                  </div>
                </CardContent>
              </Card>
            </div>

            {/* Order Summary */}
            <div>
              <Card>
                <CardHeader>
                  <CardTitle>Order Summary</CardTitle>
                </CardHeader>
                <CardContent className="space-y-4">
                  {items.map((item:any) => (
                    <div key={`${item.product_id}-${item.size}-${item.color}`} className="flex gap-3">
                      <div className="w-16 h-16 bg-gray-100 dark:bg-gray-800 rounded-lg overflow-hidden relative">
                        <Image
                          src={
                            item?.image_url?.startsWith("data:image")
                              ? item.image_url.replace(/^data:image\/jpeg;base64,?data:image\/jpeg;base64,?/, "data:image/jpeg;base64,")
                              : `/images/${item.image_url || "placeholder.svg"}`
                          }
                          alt={item.product_name}
                          fill
                          className="object-cover rounded"
                          unoptimized
                        />
                      </div>
                      <div className="flex-1">
                        <h4 className="font-medium text-sm">{item.product_name}</h4>
                        <p className="text-xs text-gray-600 dark:text-gray-400">{item.size} • {item.color}</p>
                        <p className="text-xs text-gray-600 dark:text-gray-400">Qty: {item.quantity}</p>
                      </div>
                      <div className="text-right">
                        <p className="font-medium">${(item.price * item.quantity).toFixed(2)}</p>
                      </div>
                    </div>
                  ))}
                  <div className="border-t pt-4 space-y-2">
                    <div className="flex justify-between text-sm">
                      <span>Subtotal</span>
                      <span>${totalPrice.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between text-sm">
                      <span>Shipping</span>
                      <span>${shippingCost.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between text-sm">
                      <span>Tax</span>
                      <span>${tax.toFixed(2)}</span>
                    </div>
                    <div className="flex justify-between font-bold text-lg border-t pt-2">
                      <span>Total</span>
                      <span>${finalTotal.toFixed(2)}</span>
                    </div>
                  </div>

                  <Button
                    onClick={handlePayment}
                    disabled={isProcessing}
                    className="w-full bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700"
                  >
                    {isProcessing ? (
                      <div className="flex items-center gap-2">
                        <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white" />
                        {paymentMethod === "card" ? "Processing Payment..." : "Placing Order..."}
                      </div>
                    ) : paymentMethod === "card" ? (
                      `Pay $${finalTotal.toFixed(2)}`
                    ) : (
                      "Place Orderkkkk"
                    )}
                  </Button>
                </CardContent>
              </Card>
            </div>
          </div>
        </div>
      </div>

      {/* Order Confirmation */}
      {showConfirmation && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white dark:bg-gray-800 rounded-xl p-6 max-w-sm text-center space-y-4 shadow-lg">
            <h2 className="text-xl font-semibold text-green-600">🎉 Congratulations!</h2>
            <p className="text-gray-600 dark:text-gray-300">Your order has been placed successfully.</p>
            <p className="text-sm text-gray-500 dark:text-gray-400">
              Order ID: <span className="font-medium">{orderId}</span>
            </p>
            <Button
              className="w-full bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700"
              onClick={() => {
                setShowConfirmation(false)
                router.push("/orders")
              }}
            >
              View My Orders
            </Button>
          </div>
        </div>
      )}
    </div>
  )
}
