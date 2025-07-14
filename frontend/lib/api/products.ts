import axios from "../client"
import type { Product, Category } from "@/types"

export const productsApi = {
  // ✅ Fetch all products with optional filters
  getProducts: async (params?: {
    categoryId?: string
    brand?: string
    priceRange?: string
    sortBy?: string
    search?: string
    gender?: string
  }): Promise<Product[]> => {
    const searchParams = new URLSearchParams()
    if (params) {
      Object.entries(params).forEach(([key, value]) => {
        if (value) searchParams.append(key, value)
      })
    }

    const res = await axios.get<Product[]>(
      `/products?${searchParams.toString()}`
    )
    return res.data
  },

  // ✅ Fetch single product by ID
  getProduct: async (id: string): Promise<Product> => {
    const res = await axios.get<Product>(`/products/${id}`)
    return res.data
  },

  // ✅ Fetch categories from backend
  getCategories: async (): Promise<Category[]> => {
    const res = await axios.get<Category[]>("/categories")
    return res.data
  }
}
