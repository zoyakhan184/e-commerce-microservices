import Link from "next/link"
import { Button } from "@/components/ui/button"

export function HeroSection() {
  return (
    <section className="relative overflow-hidden bg-gradient-to-br from-primary/10 via-purple-500/10 to-pink-500/10 py-20 md:py-32">
      {/* Background Pattern */}
      <div className="absolute inset-0 bg-grid-white/[0.02] bg-[size:50px_50px]" />
      <div className="absolute inset-0 bg-gradient-to-t from-background/80 to-transparent" />

      <div className="container relative">
        <div className="max-w-4xl mx-auto text-center">
          <div className="inline-flex items-center rounded-full border px-4 py-2 text-sm mb-6 bg-background/50 backdrop-blur-sm">
            <span className="mr-2">✨</span>
            New Collection Available Now
          </div>

          <h1 className="text-5xl md:text-7xl font-bold tracking-tight mb-6">
            <span className="bg-gradient-to-r from-primary via-purple-600 to-pink-600 bg-clip-text text-transparent">
              Fashion
            </span>
            <br />
            <span className="text-foreground">That Speaks</span>
            <br />
            <span className="bg-gradient-to-r from-pink-600 via-purple-600 to-primary bg-clip-text text-transparent">
              Your Language
            </span>
          </h1>

          <p className="text-xl text-muted-foreground mb-8 max-w-2xl mx-auto leading-relaxed">
            Discover premium fashion that combines contemporary style with timeless elegance. Express your unique
            personality through our curated collections.
          </p>

          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Button
              size="lg"
              className="bg-gradient-to-r from-primary to-purple-600 hover:from-primary/90 hover:to-purple-600/90 text-white shadow-lg hover:shadow-xl transition-all duration-300"
              asChild
            >
              <Link href="/products">
                Explore Collection
                <span className="ml-2">→</span>
              </Link>
            </Button>
            <Button
              size="lg"
              variant="outline"
              className="border-2 hover:bg-muted/50 transition-all duration-300 bg-transparent"
              asChild
            >
              <Link href="/categories/women">Browse Categories</Link>
            </Button>
          </div>

          <div className="mt-12 grid grid-cols-3 gap-8 max-w-md mx-auto text-center">
            <div>
              <div className="text-2xl font-bold text-primary">10K+</div>
              <div className="text-sm text-muted-foreground">Happy Customers</div>
            </div>
            <div>
              <div className="text-2xl font-bold text-purple-600">500+</div>
              <div className="text-sm text-muted-foreground">Premium Products</div>
            </div>
            <div>
              <div className="text-2xl font-bold text-pink-600">24/7</div>
              <div className="text-sm text-muted-foreground">Customer Support</div>
            </div>
          </div>
        </div>
      </div>
    </section>
  )
}
