// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: proto/inventory.proto

package inventorypb

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

type StockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Size          string                 `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	Color         string                 `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StockRequest) Reset() {
	*x = StockRequest{}
	mi := &file_proto_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockRequest) ProtoMessage() {}

func (x *StockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockRequest.ProtoReflect.Descriptor instead.
func (*StockRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *StockRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *StockRequest) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *StockRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type StockUpdateRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	SkuId          string                 `protobuf:"bytes,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	QuantityChange int32                  `protobuf:"varint,2,opt,name=quantity_change,json=quantityChange,proto3" json:"quantity_change,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *StockUpdateRequest) Reset() {
	*x = StockUpdateRequest{}
	mi := &file_proto_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockUpdateRequest) ProtoMessage() {}

func (x *StockUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockUpdateRequest.ProtoReflect.Descriptor instead.
func (*StockUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *StockUpdateRequest) GetSkuId() string {
	if x != nil {
		return x.SkuId
	}
	return ""
}

func (x *StockUpdateRequest) GetQuantityChange() int32 {
	if x != nil {
		return x.QuantityChange
	}
	return 0
}

type StockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SkuId         string                 `protobuf:"bytes,1,opt,name=sku_id,json=skuId,proto3" json:"sku_id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Size          string                 `protobuf:"bytes,3,opt,name=size,proto3" json:"size,omitempty"`
	Color         string                 `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Quantity      int32                  `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StockResponse) Reset() {
	*x = StockResponse{}
	mi := &file_proto_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockResponse) ProtoMessage() {}

func (x *StockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockResponse.ProtoReflect.Descriptor instead.
func (*StockResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *StockResponse) GetSkuId() string {
	if x != nil {
		return x.SkuId
	}
	return ""
}

func (x *StockResponse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *StockResponse) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *StockResponse) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *StockResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type LowStockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Threshold     int32                  `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LowStockRequest) Reset() {
	*x = LowStockRequest{}
	mi := &file_proto_inventory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LowStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LowStockRequest) ProtoMessage() {}

func (x *LowStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LowStockRequest.ProtoReflect.Descriptor instead.
func (*LowStockRequest) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *LowStockRequest) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

type LowStockList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*StockResponse       `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LowStockList) Reset() {
	*x = LowStockList{}
	mi := &file_proto_inventory_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LowStockList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LowStockList) ProtoMessage() {}

func (x *LowStockList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LowStockList.ProtoReflect.Descriptor instead.
func (*LowStockList) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *LowStockList) GetItems() []*StockResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

type GenericResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenericResponse) Reset() {
	*x = GenericResponse{}
	mi := &file_proto_inventory_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenericResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericResponse) ProtoMessage() {}

func (x *GenericResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_inventory_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericResponse.ProtoReflect.Descriptor instead.
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return file_proto_inventory_proto_rawDescGZIP(), []int{5}
}

func (x *GenericResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_inventory_proto protoreflect.FileDescriptor

const file_proto_inventory_proto_rawDesc = "" +
	"\n" +
	"\x15proto/inventory.proto\x12\tinventory\"W\n" +
	"\fStockRequest\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x12\n" +
	"\x04size\x18\x02 \x01(\tR\x04size\x12\x14\n" +
	"\x05color\x18\x03 \x01(\tR\x05color\"T\n" +
	"\x12StockUpdateRequest\x12\x15\n" +
	"\x06sku_id\x18\x01 \x01(\tR\x05skuId\x12'\n" +
	"\x0fquantity_change\x18\x02 \x01(\x05R\x0equantityChange\"\x8b\x01\n" +
	"\rStockResponse\x12\x15\n" +
	"\x06sku_id\x18\x01 \x01(\tR\x05skuId\x12\x1d\n" +
	"\n" +
	"product_id\x18\x02 \x01(\tR\tproductId\x12\x12\n" +
	"\x04size\x18\x03 \x01(\tR\x04size\x12\x14\n" +
	"\x05color\x18\x04 \x01(\tR\x05color\x12\x1a\n" +
	"\bquantity\x18\x05 \x01(\x05R\bquantity\"/\n" +
	"\x0fLowStockRequest\x12\x1c\n" +
	"\tthreshold\x18\x01 \x01(\x05R\tthreshold\">\n" +
	"\fLowStockList\x12.\n" +
	"\x05items\x18\x01 \x03(\v2\x18.inventory.StockResponseR\x05items\"+\n" +
	"\x0fGenericResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2\xad\x02\n" +
	"\x10InventoryService\x12=\n" +
	"\bGetStock\x12\x17.inventory.StockRequest\x1a\x18.inventory.StockResponse\x12O\n" +
	"\x12UpdateStockOnOrder\x12\x1d.inventory.StockUpdateRequest\x1a\x1a.inventory.GenericResponse\x12D\n" +
	"\aRestock\x12\x1d.inventory.StockUpdateRequest\x1a\x1a.inventory.GenericResponse\x12C\n" +
	"\fListLowStock\x12\x1a.inventory.LowStockRequest\x1a\x17.inventory.LowStockListB%Z#inventory-service/proto;inventorypbb\x06proto3"

var (
	file_proto_inventory_proto_rawDescOnce sync.Once
	file_proto_inventory_proto_rawDescData []byte
)

func file_proto_inventory_proto_rawDescGZIP() []byte {
	file_proto_inventory_proto_rawDescOnce.Do(func() {
		file_proto_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_inventory_proto_rawDesc), len(file_proto_inventory_proto_rawDesc)))
	})
	return file_proto_inventory_proto_rawDescData
}

var file_proto_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_inventory_proto_goTypes = []any{
	(*StockRequest)(nil),       // 0: inventory.StockRequest
	(*StockUpdateRequest)(nil), // 1: inventory.StockUpdateRequest
	(*StockResponse)(nil),      // 2: inventory.StockResponse
	(*LowStockRequest)(nil),    // 3: inventory.LowStockRequest
	(*LowStockList)(nil),       // 4: inventory.LowStockList
	(*GenericResponse)(nil),    // 5: inventory.GenericResponse
}
var file_proto_inventory_proto_depIdxs = []int32{
	2, // 0: inventory.LowStockList.items:type_name -> inventory.StockResponse
	0, // 1: inventory.InventoryService.GetStock:input_type -> inventory.StockRequest
	1, // 2: inventory.InventoryService.UpdateStockOnOrder:input_type -> inventory.StockUpdateRequest
	1, // 3: inventory.InventoryService.Restock:input_type -> inventory.StockUpdateRequest
	3, // 4: inventory.InventoryService.ListLowStock:input_type -> inventory.LowStockRequest
	2, // 5: inventory.InventoryService.GetStock:output_type -> inventory.StockResponse
	5, // 6: inventory.InventoryService.UpdateStockOnOrder:output_type -> inventory.GenericResponse
	5, // 7: inventory.InventoryService.Restock:output_type -> inventory.GenericResponse
	4, // 8: inventory.InventoryService.ListLowStock:output_type -> inventory.LowStockList
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_inventory_proto_init() }
func file_proto_inventory_proto_init() {
	if File_proto_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_inventory_proto_rawDesc), len(file_proto_inventory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_inventory_proto_goTypes,
		DependencyIndexes: file_proto_inventory_proto_depIdxs,
		MessageInfos:      file_proto_inventory_proto_msgTypes,
	}.Build()
	File_proto_inventory_proto = out.File
	file_proto_inventory_proto_goTypes = nil
	file_proto_inventory_proto_depIdxs = nil
}
