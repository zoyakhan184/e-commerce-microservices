// "use client"

// import { useQuery } from "@tanstack/react-query"
// import { ProductCard } from "@/components/products/product-card"
// import { ProductCardSkeleton } from "@/components/products/product-card-skeleton"

// export function FeaturedProducts() {
//   const { data: products, isLoading } = useQuery({
//     queryKey: ["featured-products"],
//     queryFn: async () => {
//       // Simulate API call that returns featured products
//       await new Promise((resolve) => setTimeout(resolve, 500))
//       return dummyProducts.slice(0, 8)
//     },
//   })

//   return (
//     <section className="py-16">
//       <div className="container">
//         <div className="text-center mb-12">
//           <h2 className="text-3xl font-bold mb-4">Featured Products</h2>
//           <p className="text-muted-foreground max-w-2xl mx-auto">
//             Discover our handpicked selection of trending items that are loved by our customers.
//           </p>
//         </div>

//         <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
//           {isLoading
//             ? Array.from({ length: 8 }).map((_, i) => <ProductCardSkeleton key={i} />)
//             : products && Array.isArray(products)
//               ? products.map((product) => <ProductCard key={product.id} product={product} />)
//               : null}
//         </div>
//       </div>
//     </section>
//   )
// }
