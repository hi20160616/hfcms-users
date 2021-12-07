// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: api/users/v1/hfcms-departments.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListDepartmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The parent resource name
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of items to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListDepartmentsRequest) Reset() {
	*x = ListDepartmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDepartmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDepartmentsRequest) ProtoMessage() {}

func (x *ListDepartmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDepartmentsRequest.ProtoReflect.Descriptor instead.
func (*ListDepartmentsRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_hfcms_departments_proto_rawDescGZIP(), []int{0}
}

func (x *ListDepartmentsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListDepartmentsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDepartmentsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListDepartmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Departments   []*Department `protobuf:"bytes,1,rep,name=departments,proto3" json:"departments,omitempty"`
	NextPageToken string        `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListDepartmentsResponse) Reset() {
	*x = ListDepartmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDepartmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDepartmentsResponse) ProtoMessage() {}

func (x *ListDepartmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDepartmentsResponse.ProtoReflect.Descriptor instead.
func (*ListDepartmentsResponse) Descriptor() ([]byte, []int) {
	return file_api_users_v1_hfcms_departments_proto_rawDescGZIP(), []int{1}
}

func (x *ListDepartmentsResponse) GetDepartments() []*Department {
	if x != nil {
		return x.Departments
	}
	return nil
}

func (x *ListDepartmentsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDepartmentRequest) Reset() {
	*x = GetDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentRequest) ProtoMessage() {}

func (x *GetDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentRequest.ProtoReflect.Descriptor instead.
func (*GetDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_hfcms_departments_proto_rawDescGZIP(), []int{2}
}

func (x *GetDepartmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Department *Department `protobuf:"bytes,2,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *CreateDepartmentRequest) Reset() {
	*x = CreateDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDepartmentRequest) ProtoMessage() {}

func (x *CreateDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDepartmentRequest.ProtoReflect.Descriptor instead.
func (*CreateDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_hfcms_departments_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDepartmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDepartmentRequest) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type UpdateDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Department *Department `protobuf:"bytes,1,opt,name=department,proto3" json:"department,omitempty"`
}

func (x *UpdateDepartmentRequest) Reset() {
	*x = UpdateDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDepartmentRequest) ProtoMessage() {}

func (x *UpdateDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDepartmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_hfcms_departments_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDepartmentRequest) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

type DeleteDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteDepartmentRequest) Reset() {
	*x = DeleteDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDepartmentRequest) ProtoMessage() {}

func (x *DeleteDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDepartmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_hfcms_departments_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDepartmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DepartmentId     string `protobuf:"bytes,2,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	DepartmentName   string `protobuf:"bytes,3,opt,name=department_name,json=departmentName,proto3" json:"department_name,omitempty"`
	DepartmentFather string `protobuf:"bytes,4,opt,name=department_father,json=departmentFather,proto3" json:"department_father,omitempty"`
}

func (x *Department) Reset() {
	*x = Department{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_hfcms_departments_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_api_users_v1_hfcms_departments_proto_rawDescGZIP(), []int{6}
}

func (x *Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Department) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *Department) GetDepartmentName() string {
	if x != nil {
		return x.DepartmentName
	}
	return ""
}

func (x *Department) GetDepartmentFather() string {
	if x != nil {
		return x.DepartmentFather
	}
	return ""
}

var File_api_users_v1_hfcms_departments_proto protoreflect.FileDescriptor

var file_api_users_v1_hfcms_departments_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x66, 0x63, 0x6d, 0x73, 0x2d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6c, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x7f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2d,
	0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9b, 0x01,
	0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x74, 0x68, 0x65, 0x72, 0x32, 0x9c, 0x05, 0x0a, 0x12,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x73, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x24, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x7c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x23, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x97, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x38, 0x32, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x7d,
	0x3a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x7c, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x27, 0x2e, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x2a, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x7d, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x32, 0x30, 0x31, 0x36, 0x30,
	0x36, 0x31, 0x36, 0x2f, 0x68, 0x66, 0x63, 0x6d, 0x73, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_users_v1_hfcms_departments_proto_rawDescOnce sync.Once
	file_api_users_v1_hfcms_departments_proto_rawDescData = file_api_users_v1_hfcms_departments_proto_rawDesc
)

func file_api_users_v1_hfcms_departments_proto_rawDescGZIP() []byte {
	file_api_users_v1_hfcms_departments_proto_rawDescOnce.Do(func() {
		file_api_users_v1_hfcms_departments_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_users_v1_hfcms_departments_proto_rawDescData)
	})
	return file_api_users_v1_hfcms_departments_proto_rawDescData
}

var file_api_users_v1_hfcms_departments_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_users_v1_hfcms_departments_proto_goTypes = []interface{}{
	(*ListDepartmentsRequest)(nil),  // 0: hfcms.users.v1.ListDepartmentsRequest
	(*ListDepartmentsResponse)(nil), // 1: hfcms.users.v1.ListDepartmentsResponse
	(*GetDepartmentRequest)(nil),    // 2: hfcms.users.v1.GetDepartmentRequest
	(*CreateDepartmentRequest)(nil), // 3: hfcms.users.v1.CreateDepartmentRequest
	(*UpdateDepartmentRequest)(nil), // 4: hfcms.users.v1.UpdateDepartmentRequest
	(*DeleteDepartmentRequest)(nil), // 5: hfcms.users.v1.DeleteDepartmentRequest
	(*Department)(nil),              // 6: hfcms.users.v1.Department
	(*emptypb.Empty)(nil),           // 7: google.protobuf.Empty
}
var file_api_users_v1_hfcms_departments_proto_depIdxs = []int32{
	6, // 0: hfcms.users.v1.ListDepartmentsResponse.departments:type_name -> hfcms.users.v1.Department
	6, // 1: hfcms.users.v1.CreateDepartmentRequest.department:type_name -> hfcms.users.v1.Department
	6, // 2: hfcms.users.v1.UpdateDepartmentRequest.department:type_name -> hfcms.users.v1.Department
	0, // 3: hfcms.users.v1.DepartmentsService.ListDepartments:input_type -> hfcms.users.v1.ListDepartmentsRequest
	2, // 4: hfcms.users.v1.DepartmentsService.GetDepartment:input_type -> hfcms.users.v1.GetDepartmentRequest
	3, // 5: hfcms.users.v1.DepartmentsService.CreateDepartment:input_type -> hfcms.users.v1.CreateDepartmentRequest
	4, // 6: hfcms.users.v1.DepartmentsService.UpdateDepartment:input_type -> hfcms.users.v1.UpdateDepartmentRequest
	5, // 7: hfcms.users.v1.DepartmentsService.DeleteDepartment:input_type -> hfcms.users.v1.DeleteDepartmentRequest
	1, // 8: hfcms.users.v1.DepartmentsService.ListDepartments:output_type -> hfcms.users.v1.ListDepartmentsResponse
	6, // 9: hfcms.users.v1.DepartmentsService.GetDepartment:output_type -> hfcms.users.v1.Department
	6, // 10: hfcms.users.v1.DepartmentsService.CreateDepartment:output_type -> hfcms.users.v1.Department
	6, // 11: hfcms.users.v1.DepartmentsService.UpdateDepartment:output_type -> hfcms.users.v1.Department
	7, // 12: hfcms.users.v1.DepartmentsService.DeleteDepartment:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_users_v1_hfcms_departments_proto_init() }
func file_api_users_v1_hfcms_departments_proto_init() {
	if File_api_users_v1_hfcms_departments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_users_v1_hfcms_departments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDepartmentsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_users_v1_hfcms_departments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDepartmentsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_users_v1_hfcms_departments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDepartmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_users_v1_hfcms_departments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDepartmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_users_v1_hfcms_departments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDepartmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_users_v1_hfcms_departments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDepartmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_users_v1_hfcms_departments_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Department); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_users_v1_hfcms_departments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_users_v1_hfcms_departments_proto_goTypes,
		DependencyIndexes: file_api_users_v1_hfcms_departments_proto_depIdxs,
		MessageInfos:      file_api_users_v1_hfcms_departments_proto_msgTypes,
	}.Build()
	File_api_users_v1_hfcms_departments_proto = out.File
	file_api_users_v1_hfcms_departments_proto_rawDesc = nil
	file_api_users_v1_hfcms_departments_proto_goTypes = nil
	file_api_users_v1_hfcms_departments_proto_depIdxs = nil
}
