import type { NextApiRequest, NextApiResponse } from "next"
import { getTokenFromRequest } from "@/lib/utils" // Ensure this utility extracts token from cookies
import { API_URL } from "@/lib/constants" // This should be http://localhost:8080 or your BFF URL

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  const token = getTokenFromRequest(req)

  const response = await fetch(`${API_URL}/admin/activity`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  })

  const data = await response.json()
  res.status(response.status).json(data)
}
