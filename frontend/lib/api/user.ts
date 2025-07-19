import axios from "../client"
import type { Address } from "@/types"

export const userApi = {
  getProfile: async () => {
    const res = await axios.get("api/users/profile")
    return res.data
  },

  updateProfile: async (data: any) => {
    const res = await axios.put("api/users/profile", data)
    return res.data
  },

  getAddresses: async (): Promise<any> => {
    const res = await axios.get("api/users/addresses")
    console.log("Raw address response:", res.data)
    return Array.isArray(res.data) ? res.data : res.data.addresses || []
  },

  addAddress: async (data: Omit<Address, "id"> & { user_id: string }): Promise<Address> => {
    const res = await axios.post("api/users/address", data)
    return res.data
  },

  updateAddress: async (data: Address): Promise<Address> => {
    const res = await axios.put("api/users/address", data)
    return res.data
  },

  setDefaultAddress: async (addressId: string): Promise<void> => {
    await axios.put(`api/users/address/default/${addressId}`)
  },

  getWishlist: async (): Promise<{ items: any[] }> => {
    const res = await axios.get("api/users/wishlist")
    console.log("Wishlist response:", res.data)
    return res.data
  },

  addToWishlist: async (productId: string): Promise<any> => {
    const res = await axios.post("api/users/wishlist", { product_id: productId })
    return res.data
  },

  removeFromWishlist: async (productId: string): Promise<any> => {
    const res = await axios.delete(`api/users/wishlist/${productId}`)
    return res.data
  },
}
