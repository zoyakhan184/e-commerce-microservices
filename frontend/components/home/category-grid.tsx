"use client"

import Link from "next/link"
import { useQuery } from "@tanstack/react-query"
import { productsApi } from "@/lib/api/products"
import { Card, CardContent } from "@/components/ui/card"

export function CategoryGrid() {
  const { data: categories } = useQuery({
    queryKey: ["categories"],
    queryFn: productsApi.getCategories,
  })

  const mainCategories = [
    { id: "men", name: "Men", gradient: "from-blue-500 to-cyan-500", icon: "👔" },
    { id: "women", name: "Women", gradient: "from-pink-500 to-rose-500", icon: "👗" },
    { id: "kids", name: "Kids", gradient: "from-yellow-500 to-orange-500", icon: "🧸" },
  ]

  return (
    <section className="py-16 bg-gradient-to-b from-background to-muted/30">
      <div className="container">
        <div className="text-center mb-12">
          <h2 className="text-3xl font-bold mb-4 bg-gradient-to-r from-primary to-purple-600 bg-clip-text text-transparent">
            Shop by Category
          </h2>
          <p className="text-muted-foreground max-w-2xl mx-auto">
            Find exactly what you're looking for in our carefully curated collections for every style and occasion.
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {mainCategories.map((category) => (
            <Link key={category.id} href={`/categories/${category.id}`}>
              <Card className="group hover:shadow-2xl transition-all duration-300 cursor-pointer border-0 bg-gradient-to-br from-background to-muted/50 hover:scale-105">
                <CardContent className="p-8 text-center relative overflow-hidden">
                  <div
                    className={`absolute inset-0 bg-gradient-to-br ${category.gradient} opacity-5 group-hover:opacity-10 transition-opacity`}
                  />
                  <div className="relative z-10">
                    <div className="text-4xl mb-4">{category.icon}</div>
                    <h3 className="text-2xl font-bold mb-2 group-hover:text-primary transition-colors">
                      {category.name}
                    </h3>
                    <p className="text-muted-foreground mb-4">Explore {category.name.toLowerCase()}'s fashion</p>
                    <div
                      className={`inline-flex items-center text-sm font-medium bg-gradient-to-r ${category.gradient} bg-clip-text text-transparent group-hover:scale-110 transition-transform`}
                    >
                      Shop Now →
                    </div>
                  </div>
                </CardContent>
              </Card>
            </Link>
          ))}
        </div>
      </div>
    </section>
  )
}
