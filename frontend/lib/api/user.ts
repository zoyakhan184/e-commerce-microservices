import axios from "../client"
import type { Address } from "@/types"

export const userApi = {
  getProfile: async () => {
    const res = await axios.get("/users/profile")
    return res.data
  },

  updateProfile: async (data: any) => {
    const res = await axios.put("/users/profile", data)
    return res.data
  },

  getAddresses: async (): Promise<any> => {
  const res = await axios.get("/users/addresses")
  console.log("Raw address response:", res.data) // 👈 add this to confirm structure
  return Array.isArray(res.data) ? res.data : res.data.addresses || []
},


  addAddress: async (data: Omit<Address, "id"> & { user_id: string }): Promise<Address> => {
    const res = await axios.post("/users/address", data)
    return res.data
  },

  updateAddress: async (data: Address): Promise<Address> => {
    const res = await axios.put("/users/address", data)
    return res.data
  },

  setDefaultAddress: async (addressId: string): Promise<void> => {
    await axios.put(`/users/address/default/${addressId}`)
  },

  getWishlist: async (): Promise<any[]> => {
    const res = await axios.get("/users/wishlist")
    const { product_ids } = res.data
    const validIds = product_ids.filter((id: string) => id)
    const productPromises = validIds.map((id: string) =>
      axios.get(`/products/${id}`).then((res) => res.data)
    )
    return Promise.all(productPromises)
  },

  addToWishlist: async (productId: string): Promise<any> => {
    const res = await axios.post("/users/wishlist", { product_id: productId })
    return res.data
  },

  removeFromWishlist: async (productId: string): Promise<any> => {
    const res = await axios.delete(`/users/wishlist/${productId}`)
    return res.data
  },
}
