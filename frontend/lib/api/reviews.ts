import axios from "../client"

export const reviewsApi = {
  addReview: (
    userId: string,
    productId: string,
    rating: number,
    comment: string,
    imageData?: string
  ) =>
    axios.post<{ message: string }>("/reviews", {
      userId,
      productId,
      rating,
      comment,
      imageData,
    }),

  getReviews: (productId: string) =>
    axios.get<any[]>(`/reviews/${productId}`),

  editReview: (reviewId: string, rating: number, comment: string) =>
    axios.put<{ message: string }>(`/reviews/${reviewId}`, {
      rating,
      comment,
    }),

  deleteReview: (reviewId: string) =>
    axios.delete<{ message: string }>(`/reviews/${reviewId}`),
}
