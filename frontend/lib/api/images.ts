import axios from "../client"

export const imageApi = {
  uploadImage: async (
    file: File,
    entityId: string,
    entityType: string
  ): Promise<{ image_id: string }> => {
    const formData = new FormData()
    formData.append("file", file)
    formData.append("entity_id", entityId)
    formData.append("entity_type", entityType)

    const res = await axios.post("/images/upload", formData, {
      headers: { "Content-Type": "multipart/form-data" },
    })

    return res.data
  },

  getImage: async (imageId: string): Promise<{ url: string; fileType: string }> => {
    return {
      url: `/images/${imageId}`,
      fileType: "image/jpeg",
    }
  },

  deleteImage: async (imageId: string): Promise<{ status: string }> => {
    const res = await axios.delete(`/images/${imageId}`)
    return res.data
  },
}

