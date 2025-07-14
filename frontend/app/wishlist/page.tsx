"use client"

import { useEffect, useState } from "react"
import { Header } from "@/components/layout/header"
import { Footer } from "@/components/layout/footer"
import { Card, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Heart, X } from "lucide-react"
import Link from "next/link"
import { userApi } from "@/lib/api/user"
import type { Product } from "@/types"
import Image from "next/image"

export default function WishlistPage() {
  const [wishlistProducts, setWishlistProducts] = useState<Product[]>([])
  const [loading, setLoading] = useState(true)
  const [removing, setRemoving] = useState<string | null>(null)

  useEffect(() => {
    const fetchWishlist = async () => {
      try {
        const products = await userApi.getWishlist()

        if (Array.isArray(products)) {
          setWishlistProducts(products)
        } else {
          console.warn("⚠️ Wishlist data is not an array:", products)
          setWishlistProducts([])
        }
      } catch (error) {
        console.error("❌ Failed to fetch wishlist:", error)
        setWishlistProducts([])
      } finally {
        setLoading(false)
      }
    }

    fetchWishlist()
  }, [])

  const handleRemove = async (productId: string) => {
    try {
      setRemoving(productId)
      await userApi.removeFromWishlist(productId)
      setWishlistProducts((prev) => prev.filter((p) => p.id !== productId))
    } catch (error) {
      console.error("❌ Failed to remove from wishlist:", error)
    } finally {
      setRemoving(null)
    }
  }

  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      <main className="flex-1">
        <div className="container py-8">
          <div className="mb-8">
            <h1 className="text-3xl font-bold mb-2">My Wishlist</h1>
            <p className="text-muted-foreground">Items you've saved for later</p>
          </div>

          {loading ? (
            <p className="text-center text-muted-foreground">Loading wishlist...</p>
          ) : wishlistProducts.length === 0 ? (
            <Card>
              <CardContent className="text-center py-16">
                <Heart className="h-16 w-16 text-muted-foreground mx-auto mb-4" />
                <h2 className="text-2xl font-semibold mb-2">Your wishlist is empty</h2>
                <p className="text-muted-foreground mb-6">
                  Save items you love for later by clicking the heart icon.
                </p>
                <Button asChild>
                  <Link href="/products">Browse Products</Link>
                </Button>
              </CardContent>
            </Card>
          ) : (
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
              {wishlistProducts.map((product) => (
                <Card key={product.id} className="relative group">
                  <div className="relative aspect-square">
                    <Image
                      src={product.imageUrls?.[0] || "/placeholder.svg"}
                      alt={product.name}
                      fill
                      className="object-cover rounded-t"
                    />
                  </div>
                  <CardContent className="p-4 space-y-2">
                    <h3 className="text-lg font-semibold line-clamp-2">{product.name}</h3>
                    <p className="text-muted-foreground text-sm line-clamp-2">
                      {product.description}
                    </p>
                    <div className="font-bold text-lg">${product.price}</div>
                  </CardContent>
                  <Button
                    variant="ghost"
                    size="icon"
                    className="absolute top-2 right-2 text-destructive hover:bg-destructive/10"
                    onClick={() => handleRemove(product.id)}
                    disabled={removing === product.id}
                  >
                    <X className="w-4 h-4" />
                  </Button>
                </Card>
              ))}
            </div>
          )}
        </div>
      </main>
      <Footer />
    </div>
  )
}
