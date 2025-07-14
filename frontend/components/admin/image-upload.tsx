"use client"

import { useState, useRef } from "react"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Upload, X, ImageIcon } from "lucide-react"
import { useToast } from "@/hooks/use-toast"
import Image from "next/image"
import { imageApi } from "@/lib/api/images"

interface ImageUploadProps {
  onImageUpload: (imageId: string, imageUrl: string) => void
  onImageRemove: (index: number) => void
  images: string[]
  entityType: string
  entityId: string
}

export function ImageUpload({ onImageUpload, onImageRemove, images, entityType, entityId }: ImageUploadProps) {
  const [uploading, setUploading] = useState(false)
  const fileInputRef = useRef<HTMLInputElement>(null)
  const { toast } = useToast()

  const handleFileSelect = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files
    if (!files || files.length === 0) return

    setUploading(true)

    try {
      for (const file of files) {
        // Convert file to base64
        const reader = new FileReader()
        reader.onload = async (e) => {
          const base64Data = e.target?.result as string
          const imageData = base64Data.split(",")[1] // Remove data:image/jpeg;base64, prefix

          try {
            // Simulate image upload to your image service
            const response = await fetch("/api/images/upload", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify({
                entityId,
                entityType,
                fileType: file.type,
                imageData,
              }),
            })

            const result = await response.json()

            if (response.ok) {
              // Create a temporary URL for preview (in real app, you'd get this from your image service)
              const imageUrl = URL.createObjectURL(file)
              onImageUpload(result.imageId, imageUrl)

              toast({
                title: "Image uploaded",
                description: "Image has been successfully uploaded.",
              })
            } else {
              throw new Error(result.message || "Upload failed")
            }
          } catch (error) {
            toast({
              title: "Upload failed",
              description: "Failed to upload image. Please try again.",
              variant: "destructive",
            })
          }
        }
        reader.readAsDataURL(file)
      }
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to process images.",
        variant: "destructive",
      })
    } finally {
      setUploading(false)
      if (fileInputRef.current) {
        fileInputRef.current.value = ""
      }
    }
  }

  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (!file) return
    setUploading(true)
    try {
      const { image_id } = await imageApi.uploadImage(file, entityId, entityType)
      onImageUpload(image_id, URL.createObjectURL(file))
    } catch (err) {
      toast({
        title: "Upload failed",
        description: "Failed to upload image. Please try again.",
        variant: "destructive",
      })
    }
    setUploading(false)
  }

  return (
    <div className="space-y-4">
      <Label>Product Images</Label>

      {/* Upload Button */}
      <div className="flex items-center gap-4">
        <Button
          type="button"
          variant="outline"
          onClick={() => fileInputRef.current?.click()}
          disabled={uploading}
          className="bg-transparent"
        >
          <Upload className="h-4 w-4 mr-2" />
          {uploading ? "Uploading..." : "Upload Images"}
        </Button>
        <Input
          ref={fileInputRef}
          type="file"
          accept="image/*"
          multiple
          onChange={handleFileSelect}
          className="hidden"
        />
        <span className="text-sm text-muted-foreground">Support: JPG, PNG, GIF (Max 5MB each)</span>
      </div>

      {/* Image Preview Grid */}
      {images.length > 0 && (
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          {images.map((imageUrl, index) => (
            <div key={index} className="relative group">
              <div className="aspect-square rounded-lg border-2 border-dashed border-muted-foreground/25 overflow-hidden">
                <Image
                  src={imageUrl || "/placeholder.svg?height=200&width=200"}
                  alt={`Product image ${index + 1}`}
                  fill
                  className="object-cover"
                />
              </div>
              <Button
                type="button"
                variant="destructive"
                size="icon"
                className="absolute -top-2 -right-2 h-6 w-6 opacity-0 group-hover:opacity-100 transition-opacity"
                onClick={() => onImageRemove(index)}
              >
                <X className="h-3 w-3" />
              </Button>
            </div>
          ))}
        </div>
      )}

      {/* Empty State */}
      {images.length === 0 && (
        <div className="border-2 border-dashed border-muted-foreground/25 rounded-lg p-8 text-center">
          <ImageIcon className="h-12 w-12 text-muted-foreground mx-auto mb-4" />
          <p className="text-muted-foreground mb-2">No images uploaded yet</p>
          <p className="text-sm text-muted-foreground">Click "Upload Images" to add product photos</p>
        </div>
      )}
    </div>
  )
}
