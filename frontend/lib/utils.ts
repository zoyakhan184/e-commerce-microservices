import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import type { NextApiRequest } from "next"
import Cookies from "js-cookie"
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}


export function getTokenFromRequest(req: NextApiRequest): string | null {
  if (req.headers.cookie) {
    const cookies = req.headers.cookie.split("; ")
    const tokenCookie = cookies.find(cookie => cookie.startsWith("token="))
    return tokenCookie ? tokenCookie.split("=")[1] : null
  }
  return null
}