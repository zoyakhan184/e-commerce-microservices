import axios from "../client"
import type { Product, Category } from "@/types"

export const productsApi = {
  getProducts: async (params?: {
    category_id?: string
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

    const res = await axios.get<Product[]>(`api/products?${searchParams.toString()}`)
    return res.data
  },

  getProduct: async (id: string): Promise<Product> => {
    const res = await axios.get<Product>(`api/products/${id}`)
    return res.data
  },

  getCategories: async (): Promise<Category[]> => {
    const res = await axios.get<Category[]>(`api/categories`)
    return res.data
  }
}
