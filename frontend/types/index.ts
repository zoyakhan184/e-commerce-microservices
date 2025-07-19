export interface User {
  user_id: string
  full_name: string
  email: string
  role: "user" | "admin"
  phone?: string
  gender?: string
  dob?: string
  avatar_url?: string
  created_at?: string
}



export interface Product {
  id: string
  name: string
  description: string
  price: number
  originalPrice?: number
  brand: string
  category_id?: string
  image_urls?: string[]
  sizes?: string[]
  colors?: string[]
  quantity: number // ✅ Add this
}


export interface Category {
  id: string
  name: string
  gender: string
  parent_id?: string
}

export interface CartItem {
  product_id: string
  product_name: string
  product_image: string
  size: string
  color: string
  quantity: number
  price: number
}

export interface Order {
  id: string
  user_id: string
  status: string
  payment_status: string
  total_amount: number
  created_at: string
  items: OrderItem[]
}

export interface OrderItem {
  product_id: string
  product_name: string
  quantity: number
  price: number
  size: string
  color: string
}

export interface Payment {
  id: string
  order_id: string
  user_id: string
  amount: number
  gateway: string
  status: string
  txn_ref?: string
  created_at: string
}

export interface Review {
  id: string
  user_id: string
  user_name: string
  product_id: string
  rating: number
  comment: string
  created_at: string
}

export interface AdminDashboard {
  total_users: number
  total_orders: number
  total_revenue: number
  low_stock_items: {
    product_id: string
    quantity: number
  }[]
}

export interface Address {
  id: string
  user_id: string
  name: string
  phone: string
  address_line: string
  city: string
  state: string
  zip: string
  country: string
  is_default: boolean
}

export interface ActivityItem {
  type: "user" | "order" | "inventory | product"
  message: string
  timestamp: string // ISO string
}

