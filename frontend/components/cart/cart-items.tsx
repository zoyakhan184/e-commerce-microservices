"use client"

import Image from "next/image"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { useCart } from "@/contexts/cart-context"
import { Minus, Plus, Trash2 } from "lucide-react"
import { useToast } from "@/hooks/use-toast"

export function CartItems() {
  const { items, updateQuantity, removeFromCart } = useCart()
  const { toast } = useToast()
  console.log("Cart Items:", items) 
  const handleQuantityChange = async (
    productId: string,
    size: string,
    color: string,
    newQuantity: number
  ) => {
    if (newQuantity < 1) return

    try {
      await updateQuantity({ product_id: productId, size, color, quantity: newQuantity })
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to update quantity",
        variant: "destructive",
      })
    }
  }

  const handleRemoveItem = async (productId: string, size: string, color: string) => {
    try {
      await removeFromCart({ product_id: productId, size, color })
      toast({
        title: "Item removed",
        description: "Item has been removed from your cart",
      })
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to remove item",
        variant: "destructive",
      })
    }
  }

  return (
    <div className="space-y-4">
      {items?.map((item: any) => (
        <Card key={`${item.product_id}-${item.size}-${item.color}`}>
          <CardContent className="p-6">
            <div className="flex items-center space-x-4">
              <div className="relative h-20 w-20 flex-shrink-0">
                <img src={item.image_url} alt="" />
          <Image
          src={
              item.image_url?.startsWith("data:image")
                ? item.image_url.replace(/^data:image\/jpeg;base64,?data:image\/jpeg;base64,?/, "data:image/jpeg;base64,")
                : `/images/${item.image_url}` // fallback if not base64
            }
          alt={item.product_name}
          fill
          className="object-cover group-hover:scale-105 transition-transform duration-300"
          unoptimized // if using Next.js Image for API routes
          />
              </div>

              <div className="flex-1 min-w-0">
                <h3 className="font-semibold text-lg">{item.product_name}</h3>
                <div className="flex items-center space-x-4 text-sm text-muted-foreground">
                  <span>Size: {item.size}</span>
                  <span>Color: {item.color}</span>
                </div>
                <p className="text-lg font-bold text-primary mt-1">${item?.price?.toFixed(2)}</p>
              </div>

              <div className="flex items-center space-x-2">
                <Button
                  variant="outline"
                  size="icon"
                  onClick={() =>
                    handleQuantityChange(item.product_id, item.size, item.color, item.quantity - 1)
                  }
                  disabled={item.quantity <= 1}
                >
                  <Minus className="h-4 w-4" />
                </Button>
                <Input
                  type="number"
                  value={item.quantity}
                  onChange={(e) =>
                    handleQuantityChange(
                      item.product_id,
                      item.size,
                      item.color,
                      Number.parseInt(e.target.value) || 1
                    )
                  }
                  className="w-16 text-center"
                  min="1"
                />
                <Button
                  variant="outline"
                  size="icon"
                  onClick={() =>
                    handleQuantityChange(item.product_id, item.size, item.color, item.quantity + 1)
                  }
                >
                  <Plus className="h-4 w-4" />
                </Button>
              </div>

              <div className="text-right">
                <p className="font-semibold text-lg">${(item.price * item.quantity).toFixed(2)}</p>
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={() => handleRemoveItem(item.product_id, item.size, item.color)}
                  className="text-red-600 hover:text-red-700 hover:bg-red-50"
                >
                  <Trash2 className="h-4 w-4 mr-1" />
                  Remove
                </Button>
              </div>
            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  )
}
