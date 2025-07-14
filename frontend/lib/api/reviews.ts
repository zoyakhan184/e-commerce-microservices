import { apiClient } from "./client"

export const reviewsApi = {
  addReview: (userId: string, productId: string, rating: number, comment: string, imageData?: string) =>
    apiClient.post<{ message: string }>("/api/reviews/add", {
      userId,
      productId,
      rating,
      comment,
      imageData,
    }),

  getReviews: (productId: string) => apiClient.get<any[]>(`/api/reviews/${productId}`),

  editReview: (reviewId: string, rating: number, comment: string) =>
    apiClient.put<{ message: string }>(`/api/reviews/${reviewId}`, {
      rating,
      comment,
    }),

  deleteReview: (reviewId: string) => apiClient.delete<{ message: string }>(`/api/reviews/${reviewId}`),
}
