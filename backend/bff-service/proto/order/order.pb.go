// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: proto/order.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlaceOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Items         []*CartItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlaceOrderRequest) Reset() {
	*x = PlaceOrderRequest{}
	mi := &file_proto_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaceOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderRequest) ProtoMessage() {}

func (x *PlaceOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderRequest.ProtoReflect.Descriptor instead.
func (*PlaceOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{0}
}

func (x *PlaceOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PlaceOrderRequest) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CartItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_proto_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{1}
}

func (x *CartItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type PlaceOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlaceOrderResponse) Reset() {
	*x = PlaceOrderResponse{}
	mi := &file_proto_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlaceOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderResponse) ProtoMessage() {}

func (x *PlaceOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderResponse.ProtoReflect.Descriptor instead.
func (*PlaceOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{2}
}

func (x *PlaceOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PlaceOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrdersRequest) Reset() {
	*x = GetOrdersRequest{}
	mi := &file_proto_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersRequest) ProtoMessage() {}

func (x *GetOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrdersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrdersResponse) Reset() {
	*x = GetOrdersResponse{}
	mi := &file_proto_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersResponse) ProtoMessage() {}

func (x *GetOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type GetOrderDetailsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderDetailsRequest) Reset() {
	*x = GetOrderDetailsRequest{}
	mi := &file_proto_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDetailsRequest) ProtoMessage() {}

func (x *GetOrderDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetOrderDetailsRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderDetailsRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GetOrderDetailsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderDetailsResponse) Reset() {
	*x = GetOrderDetailsResponse{}
	mi := &file_proto_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderDetailsResponse) ProtoMessage() {}

func (x *GetOrderDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetOrderDetailsResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrderDetailsResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type UpdateOrderStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusRequest) Reset() {
	*x = UpdateOrderStatusRequest{}
	mi := &file_proto_order_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusRequest) ProtoMessage() {}

func (x *UpdateOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateOrderStatusRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *UpdateOrderStatusRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateOrderStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateOrderStatusResponse) Reset() {
	*x = UpdateOrderStatusResponse{}
	mi := &file_proto_order_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderStatusResponse) ProtoMessage() {}

func (x *UpdateOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*UpdateOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOrderStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderStatus   string                 `protobuf:"bytes,3,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`
	PaymentStatus string                 `protobuf:"bytes,4,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty"`
	TotalAmount   float64                `protobuf:"fixed64,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Items         []*OrderItem           `protobuf:"bytes,7,rep,name=items,proto3" json:"items,omitempty"` // ✅ New
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_proto_order_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{9}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

func (x *Order) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *Order) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Order) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Order) GetItems() []*OrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type GenerateInvoiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateInvoiceRequest) Reset() {
	*x = GenerateInvoiceRequest{}
	mi := &file_proto_order_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateInvoiceRequest) ProtoMessage() {}

func (x *GenerateInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GenerateInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{10}
}

func (x *GenerateInvoiceRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type GenerateInvoiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InvoiceText   string                 `protobuf:"bytes,1,opt,name=invoice_text,json=invoiceText,proto3" json:"invoice_text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateInvoiceResponse) Reset() {
	*x = GenerateInvoiceResponse{}
	mi := &file_proto_order_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateInvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateInvoiceResponse) ProtoMessage() {}

func (x *GenerateInvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateInvoiceResponse.ProtoReflect.Descriptor instead.
func (*GenerateInvoiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{11}
}

func (x *GenerateInvoiceResponse) GetInvoiceText() string {
	if x != nil {
		return x.InvoiceText
	}
	return ""
}

type ListAllOrdersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllOrdersRequest) Reset() {
	*x = ListAllOrdersRequest{}
	mi := &file_proto_order_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllOrdersRequest) ProtoMessage() {}

func (x *ListAllOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListAllOrdersRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{12}
}

type ListAllOrdersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Orders        []*Order               `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllOrdersResponse) Reset() {
	*x = ListAllOrdersResponse{}
	mi := &file_proto_order_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllOrdersResponse) ProtoMessage() {}

func (x *ListAllOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListAllOrdersResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{13}
}

func (x *ListAllOrdersResponse) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type OrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductName   string                 `protobuf:"bytes,2,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"` // populated via DB join or product lookup
	Quantity      int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`                         // ✅ THIS should be the actual ordered quantity
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Size          string                 `protobuf:"bytes,5,opt,name=size,proto3" json:"size,omitempty"`
	Color         string                 `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	mi := &file_proto_order_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{14}
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *OrderItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderItem) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *OrderItem) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	mi := &file_proto_order_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{15}
}

func (x *CancelOrderRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type CancelOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelOrderResponse) Reset() {
	*x = CancelOrderResponse{}
	mi := &file_proto_order_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderResponse) ProtoMessage() {}

func (x *CancelOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_order_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderResponse.ProtoReflect.Descriptor instead.
func (*CancelOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_order_proto_rawDescGZIP(), []int{16}
}

func (x *CancelOrderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_order_proto protoreflect.FileDescriptor

const file_proto_order_proto_rawDesc = "" +
	"\n" +
	"\x11proto/order.proto\x12\x05order\"S\n" +
	"\x11PlaceOrderRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12%\n" +
	"\x05items\x18\x02 \x03(\v2\x0f.order.CartItemR\x05items\"E\n" +
	"\bCartItem\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\"G\n" +
	"\x12PlaceOrderResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\"+\n" +
	"\x10GetOrdersRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"9\n" +
	"\x11GetOrdersResponse\x12$\n" +
	"\x06orders\x18\x01 \x03(\v2\f.order.OrderR\x06orders\"3\n" +
	"\x16GetOrderDetailsRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"=\n" +
	"\x17GetOrderDetailsResponse\x12\"\n" +
	"\x05order\x18\x01 \x01(\v2\f.order.OrderR\x05order\"M\n" +
	"\x18UpdateOrderStatusRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\"3\n" +
	"\x19UpdateOrderStatusResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"\xe4\x01\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12!\n" +
	"\forder_status\x18\x03 \x01(\tR\vorderStatus\x12%\n" +
	"\x0epayment_status\x18\x04 \x01(\tR\rpaymentStatus\x12!\n" +
	"\ftotal_amount\x18\x05 \x01(\x01R\vtotalAmount\x12\x1d\n" +
	"\n" +
	"created_at\x18\x06 \x01(\tR\tcreatedAt\x12&\n" +
	"\x05items\x18\a \x03(\v2\x10.order.OrderItemR\x05items\"3\n" +
	"\x16GenerateInvoiceRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"<\n" +
	"\x17GenerateInvoiceResponse\x12!\n" +
	"\finvoice_text\x18\x01 \x01(\tR\vinvoiceText\"\x16\n" +
	"\x14ListAllOrdersRequest\"=\n" +
	"\x15ListAllOrdersResponse\x12$\n" +
	"\x06orders\x18\x01 \x03(\v2\f.order.OrderR\x06orders\"\xa9\x01\n" +
	"\tOrderItem\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12!\n" +
	"\fproduct_name\x18\x02 \x01(\tR\vproductName\x12\x1a\n" +
	"\bquantity\x18\x03 \x01(\x05R\bquantity\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12\x12\n" +
	"\x04size\x18\x05 \x01(\tR\x04size\x12\x14\n" +
	"\x05color\x18\x06 \x01(\tR\x05color\"/\n" +
	"\x12CancelOrderRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"-\n" +
	"\x13CancelOrderResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\x9f\x04\n" +
	"\fOrderService\x12A\n" +
	"\n" +
	"PlaceOrder\x12\x18.order.PlaceOrderRequest\x1a\x19.order.PlaceOrderResponse\x12>\n" +
	"\tGetOrders\x12\x17.order.GetOrdersRequest\x1a\x18.order.GetOrdersResponse\x12P\n" +
	"\x0fGetOrderDetails\x12\x1d.order.GetOrderDetailsRequest\x1a\x1e.order.GetOrderDetailsResponse\x12V\n" +
	"\x11UpdateOrderStatus\x12\x1f.order.UpdateOrderStatusRequest\x1a .order.UpdateOrderStatusResponse\x12P\n" +
	"\x0fGenerateInvoice\x12\x1d.order.GenerateInvoiceRequest\x1a\x1e.order.GenerateInvoiceResponse\x12J\n" +
	"\rListAllOrders\x12\x1b.order.ListAllOrdersRequest\x1a\x1c.order.ListAllOrdersResponse\x12D\n" +
	"\vCancelOrder\x12\x19.order.CancelOrderRequest\x1a\x1a.order.CancelOrderResponseB\x1dZ\x1border-service/proto;orderpbb\x06proto3"

var (
	file_proto_order_proto_rawDescOnce sync.Once
	file_proto_order_proto_rawDescData []byte
)

func file_proto_order_proto_rawDescGZIP() []byte {
	file_proto_order_proto_rawDescOnce.Do(func() {
		file_proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_order_proto_rawDesc), len(file_proto_order_proto_rawDesc)))
	})
	return file_proto_order_proto_rawDescData
}

var file_proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_proto_order_proto_goTypes = []any{
	(*PlaceOrderRequest)(nil),         // 0: order.PlaceOrderRequest
	(*CartItem)(nil),                  // 1: order.CartItem
	(*PlaceOrderResponse)(nil),        // 2: order.PlaceOrderResponse
	(*GetOrdersRequest)(nil),          // 3: order.GetOrdersRequest
	(*GetOrdersResponse)(nil),         // 4: order.GetOrdersResponse
	(*GetOrderDetailsRequest)(nil),    // 5: order.GetOrderDetailsRequest
	(*GetOrderDetailsResponse)(nil),   // 6: order.GetOrderDetailsResponse
	(*UpdateOrderStatusRequest)(nil),  // 7: order.UpdateOrderStatusRequest
	(*UpdateOrderStatusResponse)(nil), // 8: order.UpdateOrderStatusResponse
	(*Order)(nil),                     // 9: order.Order
	(*GenerateInvoiceRequest)(nil),    // 10: order.GenerateInvoiceRequest
	(*GenerateInvoiceResponse)(nil),   // 11: order.GenerateInvoiceResponse
	(*ListAllOrdersRequest)(nil),      // 12: order.ListAllOrdersRequest
	(*ListAllOrdersResponse)(nil),     // 13: order.ListAllOrdersResponse
	(*OrderItem)(nil),                 // 14: order.OrderItem
	(*CancelOrderRequest)(nil),        // 15: order.CancelOrderRequest
	(*CancelOrderResponse)(nil),       // 16: order.CancelOrderResponse
}
var file_proto_order_proto_depIdxs = []int32{
	1,  // 0: order.PlaceOrderRequest.items:type_name -> order.CartItem
	9,  // 1: order.GetOrdersResponse.orders:type_name -> order.Order
	9,  // 2: order.GetOrderDetailsResponse.order:type_name -> order.Order
	14, // 3: order.Order.items:type_name -> order.OrderItem
	9,  // 4: order.ListAllOrdersResponse.orders:type_name -> order.Order
	0,  // 5: order.OrderService.PlaceOrder:input_type -> order.PlaceOrderRequest
	3,  // 6: order.OrderService.GetOrders:input_type -> order.GetOrdersRequest
	5,  // 7: order.OrderService.GetOrderDetails:input_type -> order.GetOrderDetailsRequest
	7,  // 8: order.OrderService.UpdateOrderStatus:input_type -> order.UpdateOrderStatusRequest
	10, // 9: order.OrderService.GenerateInvoice:input_type -> order.GenerateInvoiceRequest
	12, // 10: order.OrderService.ListAllOrders:input_type -> order.ListAllOrdersRequest
	15, // 11: order.OrderService.CancelOrder:input_type -> order.CancelOrderRequest
	2,  // 12: order.OrderService.PlaceOrder:output_type -> order.PlaceOrderResponse
	4,  // 13: order.OrderService.GetOrders:output_type -> order.GetOrdersResponse
	6,  // 14: order.OrderService.GetOrderDetails:output_type -> order.GetOrderDetailsResponse
	8,  // 15: order.OrderService.UpdateOrderStatus:output_type -> order.UpdateOrderStatusResponse
	11, // 16: order.OrderService.GenerateInvoice:output_type -> order.GenerateInvoiceResponse
	13, // 17: order.OrderService.ListAllOrders:output_type -> order.ListAllOrdersResponse
	16, // 18: order.OrderService.CancelOrder:output_type -> order.CancelOrderResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_order_proto_init() }
func file_proto_order_proto_init() {
	if File_proto_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_order_proto_rawDesc), len(file_proto_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_order_proto_goTypes,
		DependencyIndexes: file_proto_order_proto_depIdxs,
		MessageInfos:      file_proto_order_proto_msgTypes,
	}.Build()
	File_proto_order_proto = out.File
	file_proto_order_proto_goTypes = nil
	file_proto_order_proto_depIdxs = nil
}
