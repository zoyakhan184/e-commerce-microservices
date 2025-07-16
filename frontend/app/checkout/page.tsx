"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import Image from "next/image"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"
import { Checkbox } from "@/components/ui/checkbox"
import { useCart } from "@/contexts/cart-context"
import { useAuth } from "@/contexts/auth-context"
import { Header } from "@/components/layout/header"
import { MapPin, Plus, Check } from "lucide-react"
import type { Address } from "@/types"
import { userApi } from "@/lib/api/user"
import { toast } from "@/components/ui/use-toast"

export default function CheckoutPage() {
  const { items, totalPrice } = useCart()
  const { user } = useAuth()
  const router = useRouter()

  const [selectedAddress, setSelectedAddress] = useState("")
  const [showAddressForm, setShowAddressForm] = useState(false)
  const [addresses, setAddresses] = useState<Address[]>([])
  const [newAddress, setNewAddress] = useState({
    name: user?.name || "",
    phone: user?.phone || "",
    address_line: "",
    city: "",
    state: "",
    zip: "",
    country: "United States",
    is_default: false,
  })

  if (!user) {
    router.push("/auth/login")
    return null
  }

  if (items.length === 0) {
    router.push("/cart")
    return null
  }

  useEffect(() => {
    const fetchAddresses = async () => {
      try {
        const res = await userApi.getAddresses()
        setAddresses(res || [])
        if (res?.length > 0) {
          setSelectedAddress(res.find((a: Address) => a.is_default)?.id || res[0].id)
        }
      } catch (err) {
        console.error("Failed to fetch addresses", err)
        toast({ title: "Error loading addresses", variant: "destructive" })
      }
    }
    fetchAddresses()
  }, [])

  const handleAddAddress = async () => {
    try {
      const res = await userApi.addAddress({ ...newAddress, user_id: user.id })
      const updated = await userApi.getAddresses()
      setAddresses(updated)
      setSelectedAddress(res.id)
      setShowAddressForm(false)
      toast({ title: "Address added successfully." })
      setNewAddress({
        name: user?.name || "",
        phone: user?.phone || "",
        address_line: "",
        city: "",
        state: "",
        zip: "",
        country: "United States",
        is_default: false,
      })
    } catch (err) {
      console.error("Failed to add address", err)
      toast({ title: "Failed to add address", variant: "destructive" })
    }
  }

  const handleSetDefault = (addressId: string) => {
    setAddresses((prev) =>
      prev.map((addr) => ({ ...addr, is_default: addr.id === addressId }))
    )
  }

  const handleContinueToPayment = () => {
    router.push("/payment")
  }

  const shippingCost = 9.99
  const tax = totalPrice * 0.08
  const finalTotal = totalPrice + shippingCost + tax

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <Header />
      <div className="container mx-auto px-4 py-8">
        <div className="max-w-4xl mx-auto">
          <h1 className="text-3xl font-bold mb-8">Checkout</h1>

          <div className="grid lg:grid-cols-3 gap-8">
            <div className="lg:col-span-2 space-y-6">
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center gap-2">
                    <MapPin className="h-5 w-5" /> Delivery Address
                  </CardTitle>
                </CardHeader>
                <CardContent className="space-y-4">
                  <RadioGroup value={selectedAddress} onValueChange={setSelectedAddress}>
                    {addresses.map((address) => (
                      <div key={address.id} className="flex items-start space-x-3 p-4 border rounded-lg">
                        <RadioGroupItem value={address.id} id={address.id} className="mt-1" />
                        <div className="flex-1">
                          <Label htmlFor={address.id} className="cursor-pointer">
                            <div className="flex items-center gap-2 mb-1">
                              <span className="font-semibold">{address.name}</span>
                              {address.is_default && (
                                <span className="bg-green-100 text-green-800 text-xs px-2 py-1 rounded-full">
                                  Default
                                </span>
                              )}
                            </div>
                            <p className="text-sm text-gray-600 dark:text-gray-400">{address.address_line}</p>
                            <p className="text-sm text-gray-600 dark:text-gray-400">
                              {address.city}, {address.state} {address.zip}
                            </p>
                            <p className="text-sm text-gray-600 dark:text-gray-400">{address.phone}</p>
                          </Label>
                          <div className="mt-2 flex gap-2">
                            {!address.is_default && (
                              <Button variant="outline" size="sm" onClick={() => handleSetDefault(address.id)}>
                                Set as Default
                              </Button>
                            )}
                          </div>
                        </div>
                      </div>
                    ))}
                  </RadioGroup>

                  {!showAddressForm ? (
                    <Button variant="outline" onClick={() => setShowAddressForm(true)} className="w-full">
                      <Plus className="h-4 w-4 mr-2" /> Add New Address
                    </Button>
                  ) : (
                    <div className="border rounded-lg p-4 space-y-4">
                      <h3 className="font-semibold">Add New Address</h3>
                      <div className="grid grid-cols-2 gap-4">
                        <div>
                          <Label htmlFor="name">Full Name</Label>
                          <Input
                            id="name"
                            value={newAddress.name}
                            onChange={(e) => setNewAddress((prev) => ({ ...prev, name: e.target.value }))}
                          />
                        </div>
                        <div>
                          <Label htmlFor="phone">Phone Number</Label>
                          <Input
                            id="phone"
                            value={newAddress.phone}
                            onChange={(e) => setNewAddress((prev) => ({ ...prev, phone: e.target.value }))}
                          />
                        </div>
                      </div>
                      <div>
                        <Label htmlFor="address">Address Line</Label>
                        <Input
                          id="address"
                          value={newAddress.address_line}
                          onChange={(e) => setNewAddress((prev) => ({ ...prev, address_line: e.target.value }))}
                        />
                      </div>
                      <div className="grid grid-cols-3 gap-4">
                        <div>
                          <Label htmlFor="city">City</Label>
                          <Input
                            id="city"
                            value={newAddress.city}
                            onChange={(e) => setNewAddress((prev) => ({ ...prev, city: e.target.value }))}
                          />
                        </div>
                        <div>
                          <Label htmlFor="state">State</Label>
                          <Input
                            id="state"
                            value={newAddress.state}
                            onChange={(e) => setNewAddress((prev) => ({ ...prev, state: e.target.value }))}
                          />
                        </div>
                        <div>
                          <Label htmlFor="zip">ZIP Code</Label>
                          <Input
                            id="zip"
                            value={newAddress.zip}
                            onChange={(e) => setNewAddress((prev) => ({ ...prev, zip: e.target.value }))}
                          />
                        </div>
                      </div>
                      <div className="flex items-center space-x-2">
                        <Checkbox
                          id="default"
                          checked={newAddress.is_default}
                          onCheckedChange={(checked) =>
                            setNewAddress((prev) => ({ ...prev, is_default: checked as boolean }))
                          }
                        />
                        <Label htmlFor="default">Set as default address</Label>
                      </div>
                      <div className="flex gap-2">
                        <Button onClick={handleAddAddress}>
                          <Check className="h-4 w-4 mr-2" /> Add Address
                        </Button>
                        <Button variant="outline" onClick={() => setShowAddressForm(false)}>
                          Cancel
                        </Button>
                      </div>
                    </div>
                  )}
                </CardContent>
              </Card>
            </div>

            <div>
              <Card>
                <CardHeader>
                  <CardTitle>Order Summary</CardTitle>
                </CardHeader>
                <CardContent className="space-y-4">
                  {items.map((item) => (
                    <div key={`${item.product_id}-${item.size}-${item.color}`} className="flex gap-3">
                      <div className="w-16 h-16 bg-gray-100 dark:bg-gray-800 rounded-lg flex items-center justify-center overflow-hidden relative">
                        <Image
                          src={
                            item.image_url?.startsWith("data:image")
                              ? item.image_url.replace(
                                  /^data:image\/jpeg;base64,?data:image\/jpeg;base64,?/,
                                  "data:image/jpeg;base64,"
                                )
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
                        <p className="text-xs text-gray-600 dark:text-gray-400">
                          {item.size} • {item.color}
                        </p>
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
                    onClick={handleContinueToPayment}
                    className="w-full bg-gradient-to-r from-purple-600 to-pink-600 hover:from-purple-700 hover:to-pink-700"
                  >
                    Continue to Payment
                  </Button>
                </CardContent>
              </Card>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}