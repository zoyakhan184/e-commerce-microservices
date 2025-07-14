"use client"

import { useState, useEffect } from "react"
import { useQuery } from "@tanstack/react-query"
import { Card, CardContent } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Badge } from "@/components/ui/badge"
import { ProductCard } from "@/components/products/product-card"
import { dummyOffers } from "@/lib/dummy-data"
import { useAuth } from "@/contexts/auth-context"
import { Search, Filter, ShoppingBag, Heart, Star, TrendingUp } from "lucide-react"
import { productsApi } from "@/lib/api/products"
import { toast } from "sonner"
import { cartApi } from "@/lib/api/cart"
import { userApi } from "@/lib/api/user"

export function UserDashboard() {
  const { user } = useAuth()

  const [searchQuery, setSearchQuery] = useState("")
  const [selectedCategory, setSelectedCategory] = useState("all")
  const [currentOfferIndex, setCurrentOfferIndex] = useState(0)
  const [wishlistCount, setWishlistCount] = useState(0)

  const [filters, setFilters] = useState({
    categoryId: "",
    brand: "",
    priceRange: "",
    sortBy: "",
    search: "",
  })

  const {
    data: products = [],
    isLoading,
    error,
  } = useQuery({
    queryKey: ["products", filters],
    queryFn: () => productsApi.getProducts(filters),
  })

  const handleQuickAdd = async (productId: string) => {
    try {
      await cartApi.addToCart(productId, "default", "default", 1)
      toast.success("Added to cart")
    } catch (err) {
      console.log("Error adding to cart:", err)
      toast.error("Failed to add to cart")
    }
  }

  const filteredProducts = products.filter((product : any) => {
    const matchesSearch =
      product?.name?.toLowerCase().includes(searchQuery.toLowerCase()) ||
      product?.description?.toLowerCase().includes(searchQuery.toLowerCase())
    const matchesCategory =
      selectedCategory === "all" || product.categoryId?.toLowerCase().includes(selectedCategory)
    return matchesSearch && matchesCategory
  })

  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentOfferIndex((prev) => (prev + 1) % dummyOffers.length)
    }, 5000)
    return () => clearInterval(interval)
  }, [])

  useEffect(() => {
    const fetchWishlist = async () => {
      try {
        const wishlist = await userApi.getWishlist()
        setWishlistCount(wishlist?.length || 0)
      } catch (error) {
        console.error("Failed to fetch wishlist count", error)
      }
    }
    fetchWishlist()
  }, [])

  const currentOffer = dummyOffers[currentOfferIndex]

  const categories = [
    { id: "all", name: "All Products" },
    { id: "men", name: "Men's Fashion" },
    { id: "women", name: "Women's Fashion" },
    { id: "kids", name: "Kids' Wear" },
  ]

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mb-8">
        <div className="flex flex-col lg:flex-row gap-6">
          <Card className="flex-1 bg-gradient-to-r from-purple-500 to-pink-500 text-white border-0">
            <CardContent className="p-8">
              <h1 className="text-3xl font-bold mb-2">Welcome back, {user?.name}! 👋</h1>
              <p className="text-purple-100 mb-4">Discover amazing deals and new arrivals just for you</p>
              <div className="flex gap-4">
                <div className="text-center">
                  <div className="text-2xl font-bold">25%</div>
                  <div className="text-sm text-purple-100">Avg. Savings</div>
                </div>
                <div className="text-center">
                  <div className="text-2xl font-bold">150+</div>
                  <div className="text-sm text-purple-100">New Items</div>
                </div>
                <div className="text-center">
                  <div className="text-2xl font-bold">4.8★</div>
                  <div className="text-sm text-purple-100">Rating</div>
                </div>
              </div>
            </CardContent>
          </Card>

          <Card
            className={`lg:w-80 bg-gradient-to-r ${currentOffer.color} text-white border-0 transition-all duration-500`}
          >
            <CardContent className="p-6">
              <div className="flex items-center justify-between mb-4">
                <Badge variant="secondary" className="bg-white/20 text-white">
                  Limited Time
                </Badge>
                <div className="text-3xl font-bold">{currentOffer.discount}% OFF</div>
              </div>
              <h3 className="text-xl font-bold mb-2">{currentOffer.title}</h3>
              <p className="text-sm opacity-90 mb-4">{currentOffer.description}</p>
              <Button variant="secondary" size="sm" className="w-full">
                Shop Now
              </Button>
            </CardContent>
          </Card>
        </div>
      </div>

      <Card className="mb-8">
        <CardContent className="p-6">
          <div className="flex flex-col lg:flex-row gap-4">
            <div className="flex-1 relative">
              <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400 h-4 w-4" />
              <Input
                placeholder="Search for products, brands, or categories..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="pl-10 pr-4"
              />
            </div>
            <Button variant="outline" size="icon">
              <Filter className="h-4 w-4" />
            </Button>
          </div>

          <div className="flex flex-wrap gap-2 mt-4">
            {categories.map((category) => (
              <Button
                key={category.id}
                variant={selectedCategory === category.id ? "default" : "outline"}
                size="sm"
                onClick={() => setSelectedCategory(category.id)}
                className="flex items-center gap-2"
              >
                {category.name}
              </Button>
            ))}
          </div>
        </CardContent>
      </Card>

      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        <Card className="bg-gradient-to-r from-blue-500 to-cyan-500 text-white border-0">
          <CardContent className="p-4 text-center">
            <ShoppingBag className="h-8 w-8 mx-auto mb-2" />
            <div className="text-2xl font-bold">12</div>
            <div className="text-sm opacity-90">Total Orders</div>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-r from-pink-500 to-rose-500 text-white border-0">
          <CardContent className="p-4 text-center">
            <Heart className="h-8 w-8 mx-auto mb-2" />
            <div className="text-2xl font-bold">{wishlistCount}</div>
            <div className="text-sm opacity-90">Wishlist Items</div>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-r from-green-500 to-emerald-500 text-white border-0">
          <CardContent className="p-4 text-center">
            <Star className="h-8 w-8 mx-auto mb-2" />
            <div className="text-2xl font-bold">4.9</div>
            <div className="text-sm opacity-90">Your Rating</div>
          </CardContent>
        </Card>

        <Card className="bg-gradient-to-r from-orange-500 to-amber-500 text-white border-0">
          <CardContent className="p-4 text-center">
            <TrendingUp className="h-8 w-8 mx-auto mb-2" />
            <div className="text-2xl font-bold">$1,234</div>
            <div className="text-sm opacity-90">Total Saved</div>
          </CardContent>
        </Card>
      </div>

      <div className="mb-8">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-2xl font-bold">
            {selectedCategory === "all" ? "All Products" : categories.find((c) => c.id === selectedCategory)?.name}
          </h2>
          <Badge variant="outline">{filteredProducts.length} items found</Badge>
        </div>

        {isLoading ? (
          <Card>
            <CardContent className="p-12 text-center text-muted-foreground">Loading products...</CardContent>
          </Card>
        ) : error ? (
          <Card>
            <CardContent className="p-12 text-center text-red-500">Error loading products</CardContent>
          </Card>
        ) : filteredProducts.length === 0 ? (
          <Card>
            <CardContent className="p-12 text-center text-muted-foreground">
              No products found matching your search.
            </CardContent>
          </Card>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {filteredProducts.map((product) => (
              <ProductCard
                key={product.id}
                product={product}
                onQuickAdd={() => handleQuickAdd(product.id)}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  )
}
