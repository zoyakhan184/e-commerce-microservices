"use client"

import { useState, useEffect } from "react"
import { Card, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { useCart } from "@/contexts/cart-context"
import { useAuth } from "@/contexts/auth-context"
import { Heart, ShoppingBag, Star } from "lucide-react"
import Image from "next/image"
import Link from "next/link"
import type { Product } from "@/types"
import { userApi } from "@/lib/api/user"

interface ProductCardProps {
  product: Product
  onQuickAdd?: (productId: string) => void
}

export function ProductCard({ product, onQuickAdd }: ProductCardProps) {
  const { addToCart } = useCart()
  const { user } = useAuth()
  const [isWishlisted, setIsWishlisted] = useState(false)
  const [isLoading, setIsLoading] = useState(false)

  useEffect(() => {
    // Optional: Fetch wishlist status initially if needed
    // You could enhance this to pull all wishlist items in a global context
  }, [])

  const handleAddToCart = async () => {
    if (!user) {
      window.location.href = "/auth/login"
      return
    }

    setIsLoading(true)
    try {
      onQuickAdd && (await onQuickAdd(product.id))
    } finally {
      setIsLoading(false)
    }
  }

  const handleWishlistToggle = async () => {
  if (!user) {
    window.location.href = "/auth/login"
    return
  }

  try {
    console.log("💡 product.id =", product.id)

    if (isWishlisted) {
      await userApi.removeFromWishlist(product.id)  // 👈 correct usage
    } else {
      if (!product.id) {
        console.warn("⚠️ product.id is missing!")
        return
      }
      await userApi.addToWishlist(product.id)       // 👈 correct usage
    }

    setIsWishlisted(!isWishlisted)
  } catch (err) {
    console.error("❌ Wishlist toggle failed", err)
  }
}


  const discountPercentage = product.originalPrice
    ? Math.round(((product.originalPrice - product.price) / product.originalPrice) * 100)
    : 0

  return (
    <Card className="group hover:shadow-xl transition-all duration-300 border-0 bg-white dark:bg-gray-800 overflow-hidden">
      <div className="relative overflow-hidden">
        <div className="aspect-square relative">
          {product && Array.isArray(product.image_urls) && product.image_urls[0] && 
          <Image
          src={product.image_urls[0]}
          alt={product.name}
          fill
          className="object-cover group-hover:scale-105 transition-transform duration-300"
          unoptimized // if using Next.js Image for API routes
          />
        }
        </div>

        {/* Discount Badge */}
        {discountPercentage > 0 && (
          <Badge className="absolute top-3 left-3 bg-red-500 hover:bg-red-600 text-white">
            -{discountPercentage}%
          </Badge>
        )}

        {/* Wishlist Button */}
        <Button
          size="icon"
          variant="secondary"
          className="absolute top-3 right-3 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
          onClick={handleWishlistToggle}
        >
          <Heart className={`h-4 w-4 ${isWishlisted ? "fill-red-500 text-red-500" : ""}`} />
        </Button>

        {/* Quick Add to Cart */}
        <div className="absolute bottom-3 left-3 right-3 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
          <Button
            onClick={handleAddToCart}
            disabled={isLoading}
            className="w-full bg-black/80 hover:bg-black text-white"
            size="sm"
          >
            {isLoading ? (
              <div className="animate-spin rounded-full h-4 w-4 border-b-2 border-white" />
            ) : (
              <>
                <ShoppingBag className="h-4 w-4 mr-2" />
                Quick Add
              </>
            )}
          </Button>
        </div>
      </div>

      <CardContent className="p-4">
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <Badge variant="outline" className="text-xs">
              {product.brand}
            </Badge>
            <div className="flex items-center gap-1">
              <Star className="h-3 w-3 fill-yellow-400 text-yellow-400" />
              <span className="text-xs text-gray-600 dark:text-gray-400">4.8</span>
            </div>
          </div>

          <Link href={`/products/${product.id}`}>
            <h3 className="font-semibold text-sm hover:text-purple-600 transition-colors line-clamp-2">
              {product.name}
            </h3>
          </Link>

          <p className="text-xs text-gray-600 dark:text-gray-400 line-clamp-2">{product.description}</p>

          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <span className="font-bold text-lg">${product.price}</span>
              {product.originalPrice && (
                <span className="text-sm text-gray-500 line-through">${product.originalPrice}</span>
              )}
            </div>
          </div>

          {/* Size and Color Options */}
          {product.sizes && product.sizes.length > 0 && (
            <div className="flex gap-2 text-xs">
              {product.sizes.slice(0, 3).map((size) => (
                <Badge key={size} variant="outline" className="text-xs px-2 py-0">
                  {size}
                </Badge>
              ))}
              {product.sizes.length > 3 && (
                <Badge variant="outline" className="text-xs px-2 py-0">
                  +{product.sizes.length - 3}
                </Badge>
              )}
            </div>
          )}
        </div>
      </CardContent>
    </Card>
  )
}
