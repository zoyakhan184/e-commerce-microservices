"use client"

import { useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { Header } from "@/components/layout/header"
import { Footer } from "@/components/layout/footer"
import { ProductCard } from "@/components/products/product-card"
import { ProductCardSkeleton } from "@/components/products/product-card-skeleton"
import { ProductFilters } from "@/components/products/product-filters"
import { productsApi } from "@/lib/api/products"
import { cartApi } from "@/lib/api/cart"
import { Button } from "@/components/ui/button"
import { Filter } from "lucide-react"
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet"
import { toast } from "sonner"

export default function ProductsPage() {
  const [filters, setFilters] = useState({
    category_id: "",
    brand: "",
    priceRange: "",
    sortBy: "",
    search: "",
  })

  const {
    data: products = [],
    isLoading,
    isError,
  } = useQuery({
    queryKey: ["products", filters],
    queryFn: () => productsApi.getProducts(filters),
    staleTime: 1000 * 60 * 5, // optional: cache for 5 mins
    refetchOnWindowFocus: false,
  })

  const handleQuickAdd = async (productId: string) => {
    try {
      await cartApi.addToCart(productId, "default", "default", 1)
      toast.success("✅ Product added to cart!")
    } catch (err) {
      console.error("❌ Error adding to cart:", err)
      toast.error("Failed to add product to cart")
    }
  }

  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      <main className="flex-1">
        <div className="container py-8">
          <div className="flex items-center justify-between mb-8">
            <div>
              <h1 className="text-3xl font-bold">All Products</h1>
              <p className="text-muted-foreground mt-2">
                {isLoading ? "Loading products..." : `${products?.length} products found`}
              </p>
            </div>

            {/* Mobile Filters */}
            <Sheet>
              <SheetTrigger asChild>
                <Button variant="outline" className="lg:hidden bg-transparent">
                  <Filter className="h-4 w-4 mr-2" />
                  Filters
                </Button>
              </SheetTrigger>
              <SheetContent side="left">
                <ProductFilters filters={filters} onFiltersChange={setFilters} />
              </SheetContent>
            </Sheet>
          </div>

          <div className="flex gap-8">
            {/* Desktop Filters */}
            <aside className="hidden lg:block w-64 flex-shrink-0">
              <ProductFilters filters={filters} onFiltersChange={setFilters} />
            </aside>

            {/* Products Grid */}
            <section className="flex-1">
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {isLoading
                  ? Array.from({ length: 12 }).map((_, i) => <ProductCardSkeleton key={i} />)
                  : Array.isArray(products)
                    ? products.map((product) => (
                        <ProductCard
                          key={product.id}
                          product={product}
                          onQuickAdd={() => handleQuickAdd(product.id)}
                        />
                      ))
                    : null}
              </div>

              {!isLoading && products?.length === 0 && (
                <div className="text-center py-12">
                  <p className="text-muted-foreground">No products match your filters.</p>
                </div>
              )}

              {isError && (
                <div className="text-center py-12">
                  <p className="text-red-500">Something went wrong. Please try again later.</p>
                </div>
              )}
            </section>
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}
