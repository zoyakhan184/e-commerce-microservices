// "use client"

// import { useState, useEffect } from "react"
// import { useAuth } from "@/contexts/auth-context"
// import { useRouter } from "next/navigation"
// import { Header } from "@/components/layout/header"
// import { Footer } from "@/components/layout/footer"
// import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
// import { Badge } from "@/components/ui/badge"
// import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
// import { RefreshCw, Clock, CheckCircle, XCircle, AlertCircle } from "lucide-react"

// interface Refund {
//   id: string
//   orderId: string
//   amount: number
//   reason: string
//   status: "pending" | "approved" | "processed" | "rejected"
//   requestDate: string
//   processedDate?: string
//   estimatedProcessingTime: string
// }

// const dummyRefunds: Refund[] = [
//   {
//     id: "REF-001",
//     orderId: "ORD-12345",
//     amount: 89.99,
//     reason: "Item not as described",
//     status: "processed",
//     requestDate: "2024-01-15",
//     processedDate: "2024-01-18",
//     estimatedProcessingTime: "3-5 business days",
//   },
//   {
//     id: "REF-002",
//     orderId: "ORD-12346",
//     amount: 129.99,
//     reason: "Wrong size received",
//     status: "approved",
//     requestDate: "2024-01-20",
//     estimatedProcessingTime: "5-7 business days",
//   },
//   {
//     id: "REF-003",
//     orderId: "ORD-12347",
//     amount: 59.99,
//     reason: "Damaged during shipping",
//     status: "pending",
//     requestDate: "2024-01-22",
//     estimatedProcessingTime: "7-10 business days",
//   },
//   {
//     id: "REF-004",
//     orderId: "ORD-12348",
//     amount: 199.99,
//     reason: "Changed mind",
//     status: "rejected",
//     requestDate: "2024-01-10",
//     processedDate: "2024-01-12",
//     estimatedProcessingTime: "N/A",
//   },
// ]

// export default function RefundsPage() {
//   const { user, isLoading } = useAuth()
//   const router = useRouter()
//   const [selectedTab, setSelectedTab] = useState("all")

//   useEffect(() => {
//     if (!isLoading && !user) {
//       router.push("/auth/login")
//     }
//   }, [user, isLoading, router])

//   const getStatusIcon = (status: string) => {
//     switch (status) {
//       case "pending":
//         return <Clock className="h-4 w-4" />
//       case "approved":
//         return <AlertCircle className="h-4 w-4" />
//       case "processed":
//         return <CheckCircle className="h-4 w-4" />
//       case "rejected":
//         return <XCircle className="h-4 w-4" />
//       default:
//         return <Clock className="h-4 w-4" />
//     }
//   }

//   const getStatusColor = (status: string) => {
//     switch (status) {
//       case "pending":
//         return "bg-yellow-100 text-yellow-800 dark:bg-yellow-900/20 dark:text-yellow-400"
//       case "approved":
//         return "bg-blue-100 text-blue-800 dark:bg-blue-900/20 dark:text-blue-400"
//       case "processed":
//         return "bg-green-100 text-green-800 dark:bg-green-900/20 dark:text-green-400"
//       case "rejected":
//         return "bg-red-100 text-red-800 dark:bg-red-900/20 dark:text-red-400"
//       default:
//         return "bg-gray-100 text-gray-800 dark:bg-gray-900/20 dark:text-gray-400"
//     }
//   }

//   const filteredRefunds =
//     selectedTab === "all" ? dummyRefunds : dummyRefunds.filter((refund) => refund.status === selectedTab)

//   if (isLoading) {
//     return (
//       <div className="min-h-screen flex items-center justify-center">
//         <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-primary"></div>
//       </div>
//     )
//   }

//   if (!user) {
//     return null
//   }

//   return (
//     <div className="min-h-screen flex flex-col">
//       <Header />
//       <main className="flex-1 py-8">
//         <div className="container max-w-4xl">
//           <div className="flex items-center gap-2 mb-8">
//             <RefreshCw className="h-6 w-6" />
//             <h1 className="text-3xl font-bold">Refund Tracking</h1>
//           </div>

//           <Tabs value={selectedTab} onValueChange={setSelectedTab}>
//             <TabsList className="grid w-full grid-cols-5">
//               <TabsTrigger value="all">All Refunds</TabsTrigger>
//               <TabsTrigger value="pending">Pending</TabsTrigger>
//               <TabsTrigger value="approved">Approved</TabsTrigger>
//               <TabsTrigger value="processed">Processed</TabsTrigger>
//               <TabsTrigger value="rejected">Rejected</TabsTrigger>
//             </TabsList>

//             <TabsContent value={selectedTab} className="mt-6">
//               <div className="space-y-4">
//                 {filteredRefunds.length === 0 ? (
//                   <Card>
//                     <CardContent className="text-center py-16">
//                       <RefreshCw className="h-16 w-16 text-muted-foreground mx-auto mb-4" />
//                       <h3 className="text-xl font-semibold mb-2">No refunds found</h3>
//                       <p className="text-muted-foreground">
//                         {selectedTab === "all"
//                           ? "You haven't requested any refunds yet."
//                           : `No ${selectedTab} refunds found.`}
//                       </p>
//                     </CardContent>
//                   </Card>
//                 ) : (
//                   filteredRefunds.map((refund) => (
//                     <Card key={refund.id}>
//                       <CardHeader>
//                         <div className="flex items-center justify-between">
//                           <CardTitle className="text-lg">Refund {refund.id}</CardTitle>
//                           <Badge className={getStatusColor(refund.status)}>
//                             <div className="flex items-center gap-1">
//                               {getStatusIcon(refund.status)}
//                               {refund.status.charAt(0).toUpperCase() + refund.status.slice(1)}
//                             </div>
//                           </Badge>
//                         </div>
//                       </CardHeader>
//                       <CardContent className="space-y-4">
//                         <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
//                           <div>
//                             <p className="text-sm font-medium text-muted-foreground">Order ID</p>
//                             <p className="font-semibold">{refund.orderId}</p>
//                           </div>
//                           <div>
//                             <p className="text-sm font-medium text-muted-foreground">Refund Amount</p>
//                             <p className="font-semibold text-lg text-primary">${refund.amount.toFixed(2)}</p>
//                           </div>
//                           <div>
//                             <p className="text-sm font-medium text-muted-foreground">Request Date</p>
//                             <p className="font-semibold">
//                               {new Date(refund.requestDate).toLocaleDateString("en-US", {
//                                 year: "numeric",
//                                 month: "long",
//                                 day: "numeric",
//                               })}
//                             </p>
//                           </div>
//                         </div>

//                         <div>
//                           <p className="text-sm font-medium text-muted-foreground mb-1">Reason</p>
//                           <p>{refund.reason}</p>
//                         </div>

//                         {refund.status === "processed" && refund.processedDate && (
//                           <div className="p-4 bg-green-50 dark:bg-green-950/20 rounded-lg border border-green-200 dark:border-green-800">
//                             <div className="flex items-center gap-2 mb-2">
//                               <CheckCircle className="h-5 w-5 text-green-600" />
//                               <p className="font-medium text-green-800 dark:text-green-200">Refund Processed</p>
//                             </div>
//                             <p className="text-sm text-green-600 dark:text-green-400">
//                               Processed on{" "}
//                               {new Date(refund.processedDate).toLocaleDateString("en-US", {
//                                 year: "numeric",
//                                 month: "long",
//                                 day: "numeric",
//                               })}
//                               . The refund should appear in your account within 3-5 business days.
//                             </p>
//                           </div>
//                         )}

//                         {refund.status === "approved" && (
//                           <div className="p-4 bg-blue-50 dark:bg-blue-950/20 rounded-lg border border-blue-200 dark:border-blue-800">
//                             <div className="flex items-center gap-2 mb-2">
//                               <AlertCircle className="h-5 w-5 text-blue-600" />
//                               <p className="font-medium text-blue-800 dark:text-blue-200">Refund Approved</p>
//                             </div>
//                             <p className="text-sm text-blue-600 dark:text-blue-400">
//                               Your refund has been approved and is being processed. Estimated processing time:{" "}
//                               {refund.estimatedProcessingTime}
//                             </p>
//                           </div>
//                         )}

//                         {refund.status === "pending" && (
//                           <div className="p-4 bg-yellow-50 dark:bg-yellow-950/20 rounded-lg border border-yellow-200 dark:border-yellow-800">
//                             <div className="flex items-center gap-2 mb-2">
//                               <Clock className="h-5 w-5 text-yellow-600" />
//                               <p className="font-medium text-yellow-800 dark:text-yellow-200">Under Review</p>
//                             </div>
//                             <p className="text-sm text-yellow-600 dark:text-yellow-400">
//                               Your refund request is being reviewed by our team. Estimated processing time:{" "}
//                               {refund.estimatedProcessingTime}
//                             </p>
//                           </div>
//                         )}

//                         {refund.status === "rejected" && (
//                           <div className="p-4 bg-red-50 dark:bg-red-950/20 rounded-lg border border-red-200 dark:border-red-800">
//                             <div className="flex items-center gap-2 mb-2">
//                               <XCircle className="h-5 w-5 text-red-600" />
//                               <p className="font-medium text-red-800 dark:text-red-200">Refund Rejected</p>
//                             </div>
//                             <p className="text-sm text-red-600 dark:text-red-400">
//                               Your refund request has been rejected. Please contact support for more information.
//                             </p>
//                           </div>
//                         )}
//                       </CardContent>
//                     </Card>
//                   ))
//                 )}
//               </div>
//             </TabsContent>
//           </Tabs>
//         </div>
//       </main>
//       <Footer />
//     </div>
//   )
// }
