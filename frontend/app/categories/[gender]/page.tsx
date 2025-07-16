"use client"

import { useState, useEffect } from "react"
import { useQuery } from "@tanstack/react-query"
import { useParams } from "next/navigation"
import { Header } from "@/components/layout/header"
import { Footer } from "@/components/layout/footer"
import { ProductCard } from "@/components/products/product-card"
import { ProductCardSkeleton } from "@/components/products/product-card-skeleton"
import { ProductFilters } from "@/components/products/product-filters"
import { productsApi } from "@/lib/api/products"
import { Button } from "@/components/ui/button"
import { Filter } from "lucide-react"
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet"

export default function CategoryPage() {
  const params = useParams()
  const gender = params.gender as string

  const [filters, setFilters] = useState({
    category_id: "",
    brand: "",
    priceRange: "",
    sortBy: "",
    search: "",
    gender: gender, // Add gender filter
  })

  // Update gender filter when route changes
  useEffect(() => {
    setFilters((prev) => ({ ...prev, gender: gender }))
  }, [gender])

  const {
    data: products,
    isLoading,
    error,
  } = useQuery({
    queryKey: ["products", filters],
    queryFn: () => productsApi.getProducts(filters),
  })

  const { data: categories } = useQuery({
    queryKey: ["categories"],
    queryFn: productsApi.getCategories,
  })

  // Get category name for display
  const categoryName = gender.charAt(0).toUpperCase() + gender.slice(1)
  const categoryDescription = `Discover our ${categoryName.toLowerCase()}'s collection featuring the latest trends and timeless classics.`

  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      <main className="flex-1">
        {/* Hero Section */}
        <section className="bg-gradient-to-r from-primary/20 via-purple-500/20 to-pink-500/20 py-16">
          <div className="container text-center">
            <h1 className="text-4xl md:text-5xl font-bold mb-4 bg-gradient-to-r from-primary via-purple-600 to-pink-600 bg-clip-text text-transparent">
              {categoryName}'s Fashion
            </h1>
            <p className="text-lg text-muted-foreground max-w-2xl mx-auto">{categoryDescription}</p>
          </div>
        </section>

        <div className="container py-8">
          <div className="flex items-center justify-between mb-8">
            <div>
              <h2 className="text-2xl font-bold">Shop {categoryName}'s Collection</h2>
              <p className="text-muted-foreground mt-2">
                {products ? `${products.length} products found` : "Loading products..."}
              </p>
            </div>

            {/* Mobile Filter Button */}
            <Sheet>
              <SheetTrigger asChild>
                <Button variant="outline" className="lg:hidden bg-transparent">
                  <Filter className="h-4 w-4 mr-2" />
                  Filters
                </Button>
              </SheetTrigger>
              <SheetContent side="left">
                <ProductFilters
                  filters={filters}
                  onFiltersChange={setFilters}
                  categories={categories?.filter((cat) => cat.gender.toLowerCase() === gender.toLowerCase())}
                />
              </SheetContent>
            </Sheet>
          </div>

          <div className="flex gap-8">
            {/* Desktop Filters */}
            <div className="hidden lg:block w-64 flex-shrink-0">
              <ProductFilters
                filters={filters}
                onFiltersChange={setFilters}
                categories={categories?.filter((cat) => cat.gender.toLowerCase() === gender.toLowerCase())}
              />
            </div>

            {/* Products Grid */}
            <div className="flex-1">
              {error && (
                <div className="text-center py-12">
                  <p className="text-red-500">Error loading products. Please try again.</p>
                </div>
              )}
              <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
                {isLoading
                  ? Array.from({ length: 12 }).map((_, i) => <ProductCardSkeleton key={i} />)
                  : products && Array.isArray(products)
                    ? products.map((product) => <ProductCard key={product.id} product={product} />)
                    : null}
              </div>

              {products && Array.isArray(products) && products.length === 0 && (
                <div className="text-center py-12">
                  <p className="text-muted-foreground">No products found matching your criteria.</p>
                </div>
              )}
            </div>
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}
