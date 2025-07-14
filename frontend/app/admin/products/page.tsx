"use client"

import { useState } from "react"
import { useQuery } from "@tanstack/react-query"
import { AdminLayout } from "@/components/admin/admin-layout"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Badge } from "@/components/ui/badge"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"
import { Textarea } from "@/components/ui/textarea"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import {
  Search,
  MoreHorizontal,
  Eye,
  Edit,
  Trash2,
  Plus,
  Package,
  DollarSign,
  TrendingUp,
  AlertTriangle,
} from "lucide-react"
import { adminApi } from "@/lib/api/admin"
import { productsApi } from "@/lib/api/products"
import { useToast } from "@/hooks/use-toast"
import Image from "next/image"
import { imageApi } from "@/lib/api/images"

export default function AdminProductsPage() {
  const [searchQuery, setSearchQuery] = useState("")
  const [selectedCategory, setSelectedCategory] = useState("all")
  const [isAddDialogOpen, setIsAddDialogOpen] = useState(false)
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false)
  const [editingProduct, setEditingProduct] = useState<any>(null)
  const { toast } = useToast()

  // Form state for add/edit product
  const [productForm, setProductForm] = useState({
    name: "",
    description: "",
    price: "",
    brand: "",
    categoryId: "",
    imageUrls: [""],
    sizes: [""],
    colors: [""],
  })

  const {
    data: products,
    isLoading: productsLoading,
    refetch: refetchProducts,
  } = useQuery({
    queryKey: ["admin-products"],
    queryFn: () => productsApi.getProducts(),
  })
  console.log("products",products)
  const { data: categories } = useQuery({
  queryKey: ["categories"],
  queryFn: productsApi.getCategories,
  staleTime: 0,           // Always consider stale (force refetch)
  refetchOnMount: true,   // Refetch on every mount
  refetchOnWindowFocus: true, // Optional: useful during debugging
})
console.log("Categories:", categories)

  // Filter products based on search and category
  const filteredProducts:any =
    products?.filter((product:any) => {
      const matchesSearch =
        product?.name?.toLowerCase().includes(searchQuery?.toLowerCase()) ||
        product?.brand?.toLowerCase().includes(searchQuery?.toLowerCase()) ||
        product?.description?.toLowerCase().includes(searchQuery?.toLowerCase())

      const matchesCategory = selectedCategory === "all" || product.categoryId === selectedCategory

      return matchesSearch && matchesCategory
    }) || []

  const handleAddProduct = async () => {
    try {
      await adminApi.createProduct({
        name: productForm.name,
        description: productForm.description,
        price: Number.parseFloat(productForm.price),
        brand: productForm.brand,
        categoryId: productForm.categoryId,
        imageUrls: productForm.imageUrls.filter((url) => url.trim() !== ""),
        sizes: productForm.sizes.filter((size) => size.trim() !== ""),
        colors: productForm.colors.filter((color) => color.trim() !== ""),
      })

      toast({
        title: "Product added",
        description: "The product has been successfully added.",
      })

      setIsAddDialogOpen(false)
      resetForm()
      refetchProducts()
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to add product. Please try again.",
        variant: "destructive",
      })
    }
  }

  const handleEditProduct = async () => {
    try {
      await adminApi.updateProduct(editingProduct.id, {
        name: productForm.name,
        description: productForm.description,
        price: Number.parseFloat(productForm.price),
        brand: productForm.brand,
        categoryId: productForm.categoryId,
        imageUrls: productForm.imageUrls.filter((url) => url.trim() !== ""),
        sizes: productForm.sizes.filter((size) => size.trim() !== ""),
        colors: productForm.colors.filter((color) => color.trim() !== ""),
      })

      toast({
        title: "Product updated",
        description: "The product has been successfully updated.",
      })

      setIsEditDialogOpen(false)
      setEditingProduct(null)
      resetForm()
      refetchProducts()
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to update product. Please try again.",
        variant: "destructive",
      })
    }
  }

  const handleDeleteProduct = async (productId: string, productName: string) => {
    if (window.confirm(`Are you sure you want to delete "${productName}"?`)) {
      try {
        await adminApi.deleteProduct(productId)
        toast({
          title: "Product deleted",
          description: `${productName} has been successfully deleted.`,
        })
        refetchProducts()
      } catch (error) {
        toast({
          title: "Error",
          description: "Failed to delete product. Please try again.",
          variant: "destructive",
        })
      }
    }
  }

  const openEditDialog = (product: any) => {
    setEditingProduct(product)
    setProductForm({
      name: product.name,
      description: product.description,
      price: product.price.toString(),
      brand: product.brand,
      categoryId: product.categoryId,
      imageUrls: product.imageUrls.length > 0 ? product.imageUrls : [""],
      sizes: product.sizes?.length > 0 ? product.sizes : [""],
      colors: product.colors?.length > 0 ? product.colors : [""],
    })
    setIsEditDialogOpen(true)
  }

  const resetForm = () => {
    setProductForm({
      name: "",
      description: "",
      price: "",
      brand: "",
      categoryId: "",
      imageUrls: [""],
      sizes: [""],
      colors: [""],
    })
  }

  const addArrayField = (field: "imageUrls" | "sizes" | "colors") => {
    setProductForm((prev) => ({
      ...prev,
      [field]: [...prev[field], ""],
    }))
  }

  const updateArrayField = (field: "imageUrls" | "sizes" | "colors", index: number, value: string) => {
    setProductForm((prev) => ({
      ...prev,
      [field]: prev[field].map((item, i) => (i === index ? value : item)),
    }))
  }

  const removeArrayField = (field: "imageUrls" | "sizes" | "colors", index: number) => {
    setProductForm((prev) => ({
      ...prev,
      [field]: prev[field].filter((_, i) => i !== index),
    }))
  }

  const getProductStats = () => {
    if (!products) return { total: 0, lowStock: 0, totalValue: 0, categories: 0 }

    return {
      total: products.length,
      lowStock: Math.floor(products.length * 0.1), // Simulate 10% low stock
      totalValue: products.reduce((sum, p) => sum + p.price, 0),
      categories: categories?.length || 0,
    }
  }

// const handleFileUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
//   const file = e.target.files?.[0]
//   if (!file) return

//   // ✅ Optional: check file type (backend only accepts .jpg and image/jpeg)
//   const isJpg = file.name.toLowerCase().endsWith(".jpg") && file.type === "image/jpeg"
//   if (!isJpg) {
//     toast({
//       title: "Invalid image",
//       description: "Only JPEG (.jpg) files are allowed.",
//       variant: "destructive",
//     })
//     return
//   }

//   try {
//     const res = await imageApi.uploadImage(file, "temp", "product") // replace "temp" with real entity ID if needed

//     setProductForm((prev) => ({
//       ...prev,
//       imageUrls: [...prev.imageUrls, res.image_id],
//     }))

//     toast({
//       title: "Image uploaded",
//       description: "Image successfully added to the product.",
//     })
//   } catch (err) {
//     console.error("Image upload failed:", err)
//     toast({
//       title: "Upload failed",
//       description: "Unable to upload image. Please try again.",
//       variant: "destructive",
//     })
//   }
// }

const handleFileUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
  const file = e.target.files?.[0]
  if (!file) return

  // Optional validation
  const isJpg = file.name.toLowerCase().endsWith(".jpg") && file.type === "image/jpeg"
  if (!isJpg) {
    toast({
      title: "Invalid image",
      description: "Only JPEG (.jpg) files are allowed.",
      variant: "destructive",
    })
    return
  }

  const reader = new FileReader()
  reader.onloadend = () => {
    const base64String = (reader.result as string).split(",")[1] // Remove data:mime;base64,
    setProductForm((prev) => ({
      ...prev,
      imageUrls: [...prev.imageUrls, base64String],
    }))
  }

  reader.readAsDataURL(file)
}


  const stats = getProductStats()

  if (productsLoading) {
    return (
      <AdminLayout>
        <div className="space-y-6">
          <h1 className="text-3xl font-bold">Products Management</h1>
          <Card>
            <CardContent className="p-6">
              <div className="animate-pulse space-y-4">
                {Array.from({ length: 5 }).map((_, i) => (
                  <div key={i} className="h-16 bg-muted rounded"></div>
                ))}
              </div>
            </CardContent>
          </Card>
        </div>
      </AdminLayout>
    )
  }

  return (
    <AdminLayout>
      <div className="space-y-6">
        {/* Header */}
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold">Products Management</h1>
            <p className="text-muted-foreground">Manage your product catalog and inventory ({stats.total} products)</p>
          </div>
          <Dialog open={isAddDialogOpen} onOpenChange={setIsAddDialogOpen}>
            <DialogTrigger asChild>
              <Button className="bg-primary hover:bg-primary/90">
                <Plus className="h-4 w-4 mr-2" />
                Add Product
              </Button>
            </DialogTrigger>
            <DialogContent className="max-w-2xl max-h-[80vh] overflow-y-auto">
            <DialogHeader>
              <DialogTitle>Add New Product</DialogTitle>
              <DialogDescription>Create a new product in your catalog.</DialogDescription>
            </DialogHeader>

            <div className="grid gap-4 py-4">
              {/* Product Name & Brand */}
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="name">Product Name *</Label>
                  <Input
                    id="name"
                    required
                    value={productForm.name}
                    onChange={(e) => setProductForm((prev) => ({ ...prev, name: e.target.value }))}
                    placeholder="Enter product name"
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="brand">Brand</Label>
                  <Input
                    id="brand"
                    value={productForm.brand}
                    onChange={(e) => setProductForm((prev) => ({ ...prev, brand: e.target.value }))}
                    placeholder="Enter brand name"
                  />
                </div>
              </div>

              {/* Price & Category */}
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="price">Price ($)</Label>
                  <Input
                    id="price"
                    type="number"
                    step="0.01"
                    value={productForm.price}
                    onChange={(e) => setProductForm((prev) => ({ ...prev, price: e.target.value }))}
                    placeholder="0.00"
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="category">Category *</Label>
                  <Select
                    value={productForm.categoryId}
                    onValueChange={(value) => setProductForm((prev) => ({ ...prev, categoryId: value }))}
                    required
                  >
                    <SelectTrigger>
                      <SelectValue placeholder="Select category" />
                    </SelectTrigger>
                    <SelectContent>
                      {categories
                        ?.filter((cat) => cat.parent_id)
                        .map((category) => {
                          const parent = categories.find((p) => p.id === category.parent_id)
                          const label = `${parent?.name || ""} / ${category.name}`
                          return (
                            <SelectItem key={category.id} value={category.id}>
                              {label} ({category.gender})
                            </SelectItem>
                          )
                        })}
                    </SelectContent>
                  </Select>
                </div>

              </div>

              {/* Description */}
              <div className="space-y-2">
                <Label htmlFor="description">Description</Label>
                <Textarea
                  id="description"
                  value={productForm.description}
                  onChange={(e) => setProductForm((prev) => ({ ...prev, description: e.target.value }))}
                  placeholder="Enter product description"
                  rows={3}
                />
              </div>

              {/* Upload Image */}
              <div className="space-y-2">
                <Label>Product Image</Label>
                <Input type="file" accept="image/jpeg" onChange={handleFileUpload} />
                {productForm.imageUrls.length > 0 && (
                  <div className="flex flex-wrap gap-2 mt-2">
                    {productForm.imageUrls
                      .filter((url) => url.trim() !== "")
                      .map((id, index) => (
                        <div
                          key={index}
                          className="relative h-16 w-16 border rounded overflow-hidden bg-muted"
                        >
                          <Image
                            src={`/images/${id}`}
                            alt={`Uploaded image ${index + 1}`}
                            fill
                            className="object-cover rounded-md"
                          />
                        </div>
                      ))}
                  </div>
                )}
              </div>

              {/* Sizes */}
              <div className="space-y-2">
                <Label>Available Sizes</Label>
                {productForm.sizes.map((size, index) => (
                  <div key={index} className="flex gap-2">
                    <Input
                      value={size}
                      onChange={(e) => updateArrayField("sizes", index, e.target.value)}
                      placeholder="e.g., S, M, L"
                    />
                    {productForm.sizes.length > 1 && (
                      <Button
                        type="button"
                        variant="outline"
                        size="icon"
                        onClick={() => removeArrayField("sizes", index)}
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    )}
                  </div>
                ))}
                <Button type="button" variant="outline" onClick={() => addArrayField("sizes")}>
                  <Plus className="h-4 w-4 mr-2" />
                  Add Size
                </Button>
              </div>

              {/* Colors */}
              <div className="space-y-2">
                <Label>Available Colors</Label>
                {productForm.colors.map((color, index) => (
                  <div key={index} className="flex gap-2">
                    <Input
                      value={color}
                      onChange={(e) => updateArrayField("colors", index, e.target.value)}
                      placeholder="Enter color"
                    />
                    {productForm.colors.length > 1 && (
                      <Button
                        type="button"
                        variant="outline"
                        size="icon"
                        onClick={() => removeArrayField("colors", index)}
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    )}
                  </div>
                ))}
                <Button type="button" variant="outline" onClick={() => addArrayField("colors")}>
                  <Plus className="h-4 w-4 mr-2" />
                  Add Color
                </Button>
              </div>
            </div>

            <DialogFooter>
              <Button variant="outline" onClick={() => setIsAddDialogOpen(false)}>
                Cancel
              </Button>
              <Button onClick={handleAddProduct}>Add Product</Button>
            </DialogFooter>
          </DialogContent>

          </Dialog>
        </div>

        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Total Products</p>
                  <p className="text-2xl font-bold">{stats.total}</p>
                </div>
                <Package className="h-8 w-8 text-blue-500" />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Categories</p>
                  <p className="text-2xl font-bold">{stats.categories}</p>
                </div>
                <TrendingUp className="h-8 w-8 text-green-500" />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Total Value</p>
                  <p className="text-2xl font-bold">${stats.totalValue.toFixed(2)}</p>
                </div>
                <DollarSign className="h-8 w-8 text-purple-500" />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardContent className="p-6">
              <div className="flex items-center justify-between">
                <div>
                  <p className="text-sm font-medium text-muted-foreground">Low Stock</p>
                  <p className="text-2xl font-bold text-orange-600">{stats.lowStock}</p>
                </div>
                <AlertTriangle className="h-8 w-8 text-orange-500" />
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Products Table */}
        <Card>
          <CardHeader>
            <CardTitle>Product Catalog</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="flex flex-col sm:flex-row gap-4 mb-6">
              <div className="relative flex-1">
                <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
                <Input
                  placeholder="Search products..."
                  value={searchQuery}
                  onChange={(e) => setSearchQuery(e.target.value)}
                  className="pl-8"
                />
              </div>
              <Select value={selectedCategory} onValueChange={setSelectedCategory}>
                <SelectTrigger className="w-48 bg-transparent">
                  <SelectValue placeholder="Filter by category" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All Categories</SelectItem>
                  {categories?.map((category) => (
                    <SelectItem key={category.id} value={category.id}>
                      {category.name} ({category.gender})
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>

            <div className="rounded-md border">
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead>Product</TableHead>
                    <TableHead>Category</TableHead>
                    <TableHead>Brand</TableHead>
                    <TableHead>Price</TableHead>
                    <TableHead>Stock</TableHead>
                    <TableHead>Status</TableHead>
                    <TableHead className="text-right">Actions</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {filteredProducts.length === 0 ? (
                    <TableRow>
                      <TableCell colSpan={7} className="text-center py-8 text-muted-foreground">
                        No products found matching your criteria.
                      </TableCell>
                    </TableRow>
                  ) : (
                    filteredProducts.map((product:any) => (
                      <TableRow key={product.id}>
                        <TableCell>
                          <div className="flex items-center space-x-3">
                            <div className="relative h-12 w-12 flex-shrink-0">
                              {product.image_urls?.[0] ? (
                                <Image
                                  src={product.image_urls[0]}
                                  alt={product.name}
                                  fill
                                  className="object-cover rounded-md"
                                />
                              ) : (
                                <div className="h-12 w-12 bg-gray-200 flex items-center justify-center rounded-md text-xs text-muted-foreground">
                                  No Image
                                </div>
                              )}
                            </div>
                            <div>
                              <p className="font-medium">{product.name}</p>
                              <p className="text-sm text-muted-foreground">ID: {product.id}</p>
                            </div>
                          </div>
                        </TableCell>
                        <TableCell>
                          {categories?.find((cat) => cat.id === product.categoryId)?.name || "Unknown"}
                        </TableCell>
                        <TableCell>{product.brand}</TableCell>
                        <TableCell className="font-semibold">${product.price.toFixed(2)}</TableCell>
                        <TableCell>
                          <Badge variant="outline" className="bg-green-50 text-green-700 border-green-200">
                            In Stock
                          </Badge>
                        </TableCell>
                        <TableCell>
                          <Badge variant="outline" className="bg-blue-50 text-blue-700 border-blue-200">
                            Active
                          </Badge>
                        </TableCell>
                        <TableCell className="text-right">
                          <DropdownMenu>
                            <DropdownMenuTrigger asChild>
                              <Button variant="ghost" size="icon">
                                <MoreHorizontal className="h-4 w-4" />
                              </Button>
                            </DropdownMenuTrigger>
                            <DropdownMenuContent align="end">
                              <DropdownMenuItem>
                                <Eye className="h-4 w-4 mr-2" />
                                View Details
                              </DropdownMenuItem>
                              <DropdownMenuItem onClick={() => openEditDialog(product)}>
                                <Edit className="h-4 w-4 mr-2" />
                                Edit Product
                              </DropdownMenuItem>
                              <DropdownMenuSeparator />
                              <DropdownMenuItem
                                className="text-red-600"
                                onClick={() => handleDeleteProduct(product.id, product.name)}
                              >
                                <Trash2 className="h-4 w-4 mr-2" />
                                Delete Product
                              </DropdownMenuItem>
                            </DropdownMenuContent>
                          </DropdownMenu>
                        </TableCell>
                      </TableRow>
                    ))
                  )}
                </TableBody>
              </Table>
            </div>
          </CardContent>
        </Card>

        {/* Edit Product Dialog */}
        <Dialog open={isEditDialogOpen} onOpenChange={setIsEditDialogOpen}>
          <DialogContent className="max-w-2xl max-h-[80vh] overflow-y-auto">
            <DialogHeader>
              <DialogTitle>Edit Product</DialogTitle>
              <DialogDescription>Update the product information.</DialogDescription>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="edit-name">Product Name</Label>
                  <Input
                    id="edit-name"
                    value={productForm.name}
                    onChange={(e) => setProductForm((prev) => ({ ...prev, name: e.target.value }))}
                    placeholder="Enter product name"
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="edit-brand">Brand</Label>
                  <Input
                    id="edit-brand"
                    value={productForm.brand}
                    onChange={(e) => setProductForm((prev) => ({ ...prev, brand: e.target.value }))}
                    placeholder="Enter brand name"
                  />
                </div>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <div className="space-y-2">
                  <Label htmlFor="edit-price">Price ($)</Label>
                  <Input
                    id="edit-price"
                    type="number"
                    step="0.01"
                    value={productForm.price}
                    onChange={(e) => setProductForm((prev) => ({ ...prev, price: e.target.value }))}
                    placeholder="0.00"
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="edit-category">Category</Label>
                  <Select
                    value={productForm.categoryId}
                    onValueChange={(value) => setProductForm((prev) => ({ ...prev, categoryId: value }))}
                  >
                    <SelectTrigger>
                      <SelectValue placeholder="Select category" />
                    </SelectTrigger>
                    <SelectContent>
                      {categories?.map((category) => (
                        <SelectItem key={category.id} value={category.id}>
                          {category.name} ({category.gender})
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>
              </div>

              <div className="space-y-2">
                <Label htmlFor="edit-description">Description</Label>
                <Textarea
                  id="edit-description"
                  value={productForm.description}
                  onChange={(e) => setProductForm((prev) => ({ ...prev, description: e.target.value }))}
                  placeholder="Enter product description"
                  rows={3}
                />
              </div>

              {/* Image URLs */}
              <div className="space-y-2">
                <Label>Image URLs</Label>
                {productForm.imageUrls.map((url, index) => (
                  <div key={index} className="flex gap-2">
                    <Input
                      value={url}
                      onChange={(e) => updateArrayField("imageUrls", index, e.target.value)}
                      placeholder="Enter image URL"
                    />
                    {productForm.imageUrls.length > 1 && (
                      <Button
                        type="button"
                        variant="outline"
                        size="icon"
                        onClick={() => removeArrayField("imageUrls", index)}
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    )}
                  </div>
                ))}
                <Button type="button" variant="outline" onClick={() => addArrayField("imageUrls")}>
                  <Plus className="h-4 w-4 mr-2" />
                  Add Image URL
                </Button>
              </div>

              {/* Sizes */}
              <div className="space-y-2">
                <Label>Available Sizes</Label>
                {productForm.sizes.map((size, index) => (
                  <div key={index} className="flex gap-2">
                    <Input
                      value={size}
                      onChange={(e) => updateArrayField("sizes", index, e.target.value)}
                      placeholder="Enter size (e.g., S, M, L)"
                    />
                    {productForm.sizes.length > 1 && (
                      <Button
                        type="button"
                        variant="outline"
                        size="icon"
                        onClick={() => removeArrayField("sizes", index)}
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    )}
                  </div>
                ))}
                <Button type="button" variant="outline" onClick={() => addArrayField("sizes")}>
                  <Plus className="h-4 w-4 mr-2" />
                  Add Size
                </Button>
              </div>

              {/* Colors */}
              <div className="space-y-2">
                <Label>Available Colors</Label>
                {productForm.colors.map((color, index) => (
                  <div key={index} className="flex gap-2">
                    <Input
                      value={color}
                      onChange={(e) => updateArrayField("colors", index, e.target.value)}
                      placeholder="Enter color"
                    />
                    {productForm.colors.length > 1 && (
                      <Button
                        type="button"
                        variant="outline"
                        size="icon"
                        onClick={() => removeArrayField("colors", index)}
                      >
                        <Trash2 className="h-4 w-4" />
                      </Button>
                    )}
                  </div>
                ))}
                <Button type="button" variant="outline" onClick={() => addArrayField("colors")}>
                  <Plus className="h-4 w-4 mr-2" />
                  Add Color
                </Button>
              </div>
            </div>
            <DialogFooter>
              <Button variant="outline" onClick={() => setIsEditDialogOpen(false)}>
                Cancel
              </Button>
              <Button onClick={handleEditProduct}>Update Product</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      </div>
    </AdminLayout>
  )
}
