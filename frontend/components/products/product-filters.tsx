"use client"

import { useQuery } from "@tanstack/react-query"
import { productsApi } from "@/lib/api/products"
import { Label } from "@/components/ui/label"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import type { Category } from "@/types/category" // Declare the Category variable

interface ProductFiltersProps {
  filters: {
    categoryId: string
    brand: string
    priceRange: string
    sortBy: string
    search: string
    gender?: string
  }
  onFiltersChange: (filters: any) => void
  categories?: Category[]
}

export function ProductFilters({ filters, onFiltersChange, categories }: ProductFiltersProps) {
  const { data: categoriesData } = useQuery({
    queryKey: ["categories"],
    queryFn: productsApi.getCategories,
    enabled: !categories,
  })

  const allCategories = categories || categoriesData

  const handleFilterChange = (key: string, value: string) => {
    onFiltersChange({ ...filters, [key]: value })
  }

  const clearFilters = () => {
    onFiltersChange({
      categoryId: "",
      brand: "",
      priceRange: "",
      sortBy: "",
      search: "",
    })
  }

  return (
    <div className="space-y-6">
      <Card>
        <CardHeader>
          <CardTitle className="text-lg">Filters</CardTitle>
        </CardHeader>
        <CardContent className="space-y-4">
          {/* Search */}
          <div className="space-y-2">
            <Label>Search</Label>
            <Input
              placeholder="Search products..."
              value={filters.search}
              onChange={(e) => handleFilterChange("search", e.target.value)}
            />
          </div>

          {/* Category - only show if categories are provided */}
          {allCategories && allCategories.length > 0 && (
            <div className="space-y-2">
              <Label>Category</Label>
              <Select value={filters.categoryId} onValueChange={(value) => handleFilterChange("categoryId", value)}>
                <SelectTrigger>
                  <SelectValue placeholder="All categories" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All categories</SelectItem>{" "}
                  {/* Modify default value prop to be a non-empty string */}
                  {allCategories.map((category) => (
                    <SelectItem key={category.id} value={category.id}>
                      {category.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          )}

          {/* Brand */}
          <div className="space-y-2">
            <Label>Brand</Label>
            <Select value={filters.brand} onValueChange={(value) => handleFilterChange("brand", value)}>
              <SelectTrigger>
                <SelectValue placeholder="All brands" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">All brands</SelectItem>
                <SelectItem value="Nike">Nike</SelectItem>
                <SelectItem value="Adidas">Adidas</SelectItem>
                <SelectItem value="Zara">Zara</SelectItem>
                <SelectItem value="H&M">H&M</SelectItem>
              </SelectContent>
            </Select>
          </div>

          {/* Price Range */}
          <div className="space-y-2">
            <Label>Price Range</Label>
            <Select value={filters.priceRange} onValueChange={(value) => handleFilterChange("priceRange", value)}>
              <SelectTrigger>
                <SelectValue placeholder="All prices" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">All prices</SelectItem>
                <SelectItem value="0-50">$0 - $50</SelectItem>
                <SelectItem value="50-100">$50 - $100</SelectItem>
                <SelectItem value="100-200">$100 - $200</SelectItem>
                <SelectItem value="200+">$200+</SelectItem>
              </SelectContent>
            </Select>
          </div>

          {/* Sort By */}
          <div className="space-y-2">
            <Label>Sort By</Label>
            <Select value={filters.sortBy} onValueChange={(value) => handleFilterChange("sortBy", value)}>
              <SelectTrigger>
                <SelectValue placeholder="Default" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="default">Default</SelectItem>
                <SelectItem value="price-low">Price: Low to High</SelectItem>
                <SelectItem value="price-high">Price: High to Low</SelectItem>
                <SelectItem value="name">Name: A to Z</SelectItem>
                <SelectItem value="popular">Most Popular</SelectItem>
              </SelectContent>
            </Select>
          </div>

          <Button variant="outline" onClick={clearFilters} className="w-full bg-transparent">
            Clear Filters
          </Button>
        </CardContent>
      </Card>
    </div>
  )
}
