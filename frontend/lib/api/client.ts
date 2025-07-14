import {
  dummyProducts,
  dummyCategories,
  dummyUsers,
  dummyCartItems,
  dummyOrders,
  dummyPayments,
  dummyAdminDashboard,
  testCredentials,
} from "@/lib/dummy-data"

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080"

class ApiClient {
  private baseURL: string
  private useDummyData: boolean

  constructor(baseURL: string) {
    this.baseURL = baseURL
    // Use dummy data if API is not available
    this.useDummyData = false // Set to false when real API is ready
  }

  private async request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    // If using dummy data, simulate API responses
    if (this.useDummyData) {
      return this.handleDummyRequest<T>(endpoint, options)
    }

    const url = `${this.baseURL}${endpoint}`
    const token = localStorage.getItem("token")

    const config: RequestInit = {
      headers: {
        "Content-Type": "application/json",
        ...(token && { Authorization: `Bearer ${token}` }),
        ...options.headers,
      },
      ...options,
    }

    const response = await fetch(url, config)

    if (!response.ok) {
      const error = await response.text()
      throw new Error(error || `HTTP error! status: ${response.status}`)
    }

    return response.json()
  }

  private async handleDummyRequest<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
    // Simulate network delay
    await new Promise((resolve) => setTimeout(resolve, 500))

    const method = options.method || "GET"
    const body = options.body ? JSON.parse(options.body as string) : null

    // Handle different endpoints
    if (endpoint === "/api/auth/login" && method === "POST") {
      const { email, password } = body
      const user = dummyUsers.find((u) => u.email === email)

      if (
        user &&
        ((email === testCredentials.user.email && password === testCredentials.user.password) ||
          (email === testCredentials.admin.email && password === testCredentials.admin.password))
      ) {
        return {
          token: "dummy-jwt-token",
          user_id: user.id,
          role: user.role,
        } as T
      } else {
        throw new Error("Invalid credentials")
      }
    }

    if (endpoint === "/api/auth/register" && method === "POST") {
      const { name, email } = body
      return {
        token: "dummy-jwt-token",
        user_id: "new-user-id",
        role: "user",
      } as T
    }

    if (endpoint.startsWith("/api/products")) {
      if (endpoint === "/api/products" || endpoint.includes("/api/products?")) {
        const urlParts = endpoint.split("?")
        const params = new URLSearchParams(urlParts[1] || "")
        const gender = params.get("gender")
        const search = params.get("search")
        const categoryId = params.get("categoryId")
        const brand = params.get("brand")

        let filteredProducts = [...dummyProducts]

        if (gender) {
          const genderCategories = dummyCategories.filter((cat) => cat.gender.toLowerCase() === gender.toLowerCase())
          const categoryIds = genderCategories.map((cat) => cat.id)
          filteredProducts = filteredProducts.filter((product) => categoryIds.includes(product.categoryId))
        }

        if (search) {
          filteredProducts = filteredProducts.filter(
            (product) =>
              product.name.toLowerCase().includes(search.toLowerCase()) ||
              product.description.toLowerCase().includes(search.toLowerCase()),
          )
        }

        if (categoryId && categoryId !== "all") {
          filteredProducts = filteredProducts.filter((product) => product.categoryId === categoryId)
        }

        if (brand && brand !== "all") {
          filteredProducts = filteredProducts.filter((product) => product.brand === brand)
        }

        return filteredProducts as T
      }

      if (endpoint.match(/\/api\/products\/[\w-]+$/)) {
        const productId = endpoint.split("/").pop()
        const product = dummyProducts.find((p) => p.id === productId)
        if (product) {
          return product as T
        } else {
          throw new Error("Product not found")
        }
      }
    }

    if (endpoint === "/api/categories") {
      return dummyCategories as T
    }

    if (endpoint === "/api/cart") {
      return dummyCartItems as T
    }

    if (endpoint === "/api/cart/add" && method === "POST") {
      return { message: "Item added to cart" } as T
    }

    if (endpoint === "/api/cart/remove" && method === "POST") {
      return { message: "Item removed from cart" } as T
    }

    if (endpoint === "/api/cart/update" && method === "POST") {
      return { message: "Cart updated" } as T
    }

    if (endpoint === "/api/orders") {
      return dummyOrders as T
    }

    if (endpoint === "/api/admin/dashboard") {
      return dummyAdminDashboard as T
    }

    if (endpoint === "/api/admin/users") {
      return dummyUsers as T
    }

    if (endpoint === "/api/admin/orders") {
      return dummyOrders as T
    }

    if (endpoint.match(/\/api\/admin\/users\/[\w-]+$/) && method === "DELETE") {
      return { message: "User deleted successfully" } as T
    }

    if (endpoint.match(/\/api\/admin\/users\/[\w-]+\/status$/) && method === "PUT") {
      return { message: "User status updated successfully" } as T
    }

    if (endpoint === "/api/user/profile") {
      const token = localStorage.getItem("token")
      if (token) {
        return dummyUsers[0] as T // Return first user as current user
      } else {
        throw new Error("Not authenticated")
      }
    }

    if (endpoint === "/api/admin/payments") {
      return dummyPayments as T
    }

    if (endpoint.match(/\/api\/admin\/payments\/[\w-]+\/refund$/) && method === "POST") {
      return { message: "Refund processed successfully" } as T
    }

    if (endpoint.match(/\/api\/admin\/orders\/[\w-]+\/status$/) && method === "PUT") {
      return { message: "Order status updated successfully" } as T
    }

    if (endpoint === "/api/admin/products/add" && method === "POST") {
      return { message: "Product added successfully" } as T
    }

    if (endpoint === "/api/admin/products/edit" && method === "PUT") {
      return { message: "Product updated successfully" } as T
    }

    // Default response for unhandled endpoints
    return {} as T
  }

async postForm<T>(endpoint: string, formData: FormData): Promise<T> {
  if (this.useDummyData) {
    // Convert FormData to JSON-like object for dummy simulation
    const dummyBody: Record<string, any> = {}
    formData.forEach((value, key) => {
      dummyBody[key] = value
    })
    return this.handleDummyRequest<T>(endpoint, {
      method: "POST",
      body: JSON.stringify(dummyBody),
    })
  }

  const url = `${this.baseURL}${endpoint}`
  const token = localStorage.getItem("token")

  const response = await fetch(url, {
    method: "POST",
    headers: {
      ...(token && { Authorization: `Bearer ${token}` }),
      // DON'T set Content-Type for FormData — browser handles it
    },
    body: formData,
  })

  if (!response.ok) {
    const error = await response.text()
    throw new Error(error || `HTTP error! status: ${response.status}`)
  }

  return response.json()
}
}
export const apiClient = new ApiClient(API_BASE_URL)
