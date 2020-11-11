// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: tb.proto

package tb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sensor int32

const (
	Sensor_CAMERA  Sensor = 0
	Sensor_RADAR   Sensor = 1
	Sensor_LIDAR   Sensor = 2
	Sensor_THERMAL Sensor = 3
)

// Enum value maps for Sensor.
var (
	Sensor_name = map[int32]string{
		0: "CAMERA",
		1: "RADAR",
		2: "LIDAR",
		3: "THERMAL",
	}
	Sensor_value = map[string]int32{
		"CAMERA":  0,
		"RADAR":   1,
		"LIDAR":   2,
		"THERMAL": 3,
	}
)

func (x Sensor) Enum() *Sensor {
	p := new(Sensor)
	*p = x
	return p
}

func (x Sensor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sensor) Descriptor() protoreflect.EnumDescriptor {
	return file_tb_proto_enumTypes[0].Descriptor()
}

func (Sensor) Type() protoreflect.EnumType {
	return &file_tb_proto_enumTypes[0]
}

func (x Sensor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Sensor) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Sensor(num)
	return nil
}

// Deprecated: Use Sensor.Descriptor instead.
func (Sensor) EnumDescriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{0}
}

type Node_NodeStatus int32

const (
	Node_ONLINE  Node_NodeStatus = 1
	Node_OFFLINE Node_NodeStatus = 2
)

// Enum value maps for Node_NodeStatus.
var (
	Node_NodeStatus_name = map[int32]string{
		1: "ONLINE",
		2: "OFFLINE",
	}
	Node_NodeStatus_value = map[string]int32{
		"ONLINE":  1,
		"OFFLINE": 2,
	}
)

func (x Node_NodeStatus) Enum() *Node_NodeStatus {
	p := new(Node_NodeStatus)
	*p = x
	return p
}

func (x Node_NodeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Node_NodeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_tb_proto_enumTypes[1].Descriptor()
}

func (Node_NodeStatus) Type() protoreflect.EnumType {
	return &file_tb_proto_enumTypes[1]
}

func (x Node_NodeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Node_NodeStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Node_NodeStatus(num)
	return nil
}

// Deprecated: Use Node_NodeStatus.Descriptor instead.
func (Node_NodeStatus) EnumDescriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{1, 0}
}

type NodeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SensorType []Sensor `protobuf:"varint,1,rep,name=sensorType,enum=tb.Sensor" json:"sensorType,omitempty"`
	NodeName   *string  `protobuf:"bytes,2,opt,name=nodeName" json:"nodeName,omitempty"`
}

func (x *NodeSpec) Reset() {
	*x = NodeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeSpec) ProtoMessage() {}

func (x *NodeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeSpec.ProtoReflect.Descriptor instead.
func (*NodeSpec) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{0}
}

func (x *NodeSpec) GetSensorType() []Sensor {
	if x != nil {
		return x.SensorType
	}
	return nil
}

func (x *NodeSpec) GetNodeName() string {
	if x != nil && x.NodeName != nil {
		return *x.NodeName
	}
	return ""
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             *string          `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Status           *Node_NodeStatus `protobuf:"varint,2,req,name=status,enum=tb.Node_NodeStatus" json:"status,omitempty"`
	OnlineSensorType []Sensor         `protobuf:"varint,3,rep,name=onlineSensorType,enum=tb.Sensor" json:"onlineSensorType,omitempty"`
	OnlineSensor     []string         `protobuf:"bytes,4,rep,name=onlineSensor" json:"onlineSensor,omitempty"`
	RunningTask      *int32           `protobuf:"varint,5,opt,name=runningTask" json:"runningTask,omitempty"` // @TODO
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Node) GetStatus() Node_NodeStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Node_ONLINE
}

func (x *Node) GetOnlineSensorType() []Sensor {
	if x != nil {
		return x.OnlineSensorType
	}
	return nil
}

func (x *Node) GetOnlineSensor() []string {
	if x != nil {
		return x.OnlineSensor
	}
	return nil
}

func (x *Node) GetRunningTask() int32 {
	if x != nil && x.RunningTask != nil {
		return *x.RunningTask
	}
	return 0
}

type NodeNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node []*Node `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"` // @TODO
}

func (x *NodeNetwork) Reset() {
	*x = NodeNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeNetwork) ProtoMessage() {}

func (x *NodeNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeNetwork.ProtoReflect.Descriptor instead.
func (*NodeNetwork) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{2}
}

func (x *NodeNetwork) GetNode() []*Node {
	if x != nil {
		return x.Node
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string          `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Image    *string          `protobuf:"bytes,2,req,name=image" json:"image,omitempty"`
	Cmd      *string          `protobuf:"bytes,3,opt,name=cmd" json:"cmd,omitempty"`
	Device   map[string]int32 `protobuf:"bytes,4,rep,name=device" json:"device,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Sensor   []Sensor         `protobuf:"varint,5,rep,name=sensor,enum=tb.Sensor" json:"sensor,omitempty"`
	NodeSpec *NodeSpec        `protobuf:"bytes,6,opt,name=nodeSpec" json:"nodeSpec,omitempty"`
	Status   *string          `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{3}
}

func (x *Task) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Task) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *Task) GetCmd() string {
	if x != nil && x.Cmd != nil {
		return *x.Cmd
	}
	return ""
}

func (x *Task) GetDevice() map[string]int32 {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *Task) GetSensor() []Sensor {
	if x != nil {
		return x.Sensor
	}
	return nil
}

func (x *Task) GetNodeSpec() *NodeSpec {
	if x != nil {
		return x.NodeSpec
	}
	return nil
}

func (x *Task) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

type NewTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskName *string `protobuf:"bytes,1,req,name=taskName" json:"taskName,omitempty"`
	UserName *string `protobuf:"bytes,2,req,name=userName" json:"userName,omitempty"`
	// required string userCert = 3;
	UserEmail *string `protobuf:"bytes,4,req,name=userEmail" json:"userEmail,omitempty"`
	TaskInfo  *Task   `protobuf:"bytes,5,req,name=taskInfo" json:"taskInfo,omitempty"`
}

func (x *NewTaskRequest) Reset() {
	*x = NewTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTaskRequest) ProtoMessage() {}

func (x *NewTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTaskRequest.ProtoReflect.Descriptor instead.
func (*NewTaskRequest) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{4}
}

func (x *NewTaskRequest) GetTaskName() string {
	if x != nil && x.TaskName != nil {
		return *x.TaskName
	}
	return ""
}

func (x *NewTaskRequest) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *NewTaskRequest) GetUserEmail() string {
	if x != nil && x.UserEmail != nil {
		return *x.UserEmail
	}
	return ""
}

func (x *NewTaskRequest) GetTaskInfo() *Task {
	if x != nil {
		return x.TaskInfo
	}
	return nil
}

type NewTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If fail
	Error []string `protobuf:"bytes,1,rep,name=error" json:"error,omitempty"`
	// Otherwise success
	TaskId   *string `protobuf:"bytes,2,opt,name=taskId" json:"taskId,omitempty"`
	TaskInfo *Task   `protobuf:"bytes,3,opt,name=taskInfo" json:"taskInfo,omitempty"`
}

func (x *NewTaskResponse) Reset() {
	*x = NewTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewTaskResponse) ProtoMessage() {}

func (x *NewTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewTaskResponse.ProtoReflect.Descriptor instead.
func (*NewTaskResponse) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{5}
}

func (x *NewTaskResponse) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *NewTaskResponse) GetTaskId() string {
	if x != nil && x.TaskId != nil {
		return *x.TaskId
	}
	return ""
}

func (x *NewTaskResponse) GetTaskInfo() *Task {
	if x != nil {
		return x.TaskInfo
	}
	return nil
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskName *string `protobuf:"bytes,1,opt,name=taskName" json:"taskName,omitempty"`
	TaskId   *string `protobuf:"bytes,2,opt,name=taskId" json:"taskId,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{6}
}

func (x *GetTaskRequest) GetTaskName() string {
	if x != nil && x.TaskName != nil {
		return *x.TaskName
	}
	return ""
}

func (x *GetTaskRequest) GetTaskId() string {
	if x != nil && x.TaskId != nil {
		return *x.TaskId
	}
	return ""
}

type GetTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If fail
	Error []string `protobuf:"bytes,1,rep,name=error" json:"error,omitempty"`
	// Otherwise success
	TaskInfo *Task `protobuf:"bytes,2,opt,name=taskInfo" json:"taskInfo,omitempty"`
}

func (x *GetTaskResponse) Reset() {
	*x = GetTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskResponse) ProtoMessage() {}

func (x *GetTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskResponse.ProtoReflect.Descriptor instead.
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{7}
}

func (x *GetTaskResponse) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *GetTaskResponse) GetTaskInfo() *Task {
	if x != nil {
		return x.TaskInfo
	}
	return nil
}

type ListNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListNodesRequest) Reset() {
	*x = ListNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodesRequest) ProtoMessage() {}

func (x *ListNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodesRequest.ProtoReflect.Descriptor instead.
func (*ListNodesRequest) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{8}
}

type ListNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   []string     `protobuf:"bytes,1,rep,name=error" json:"error,omitempty"`
	Network *NodeNetwork `protobuf:"bytes,2,opt,name=network" json:"network,omitempty"`
}

func (x *ListNodesResponse) Reset() {
	*x = ListNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodesResponse) ProtoMessage() {}

func (x *ListNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodesResponse.ProtoReflect.Descriptor instead.
func (*ListNodesResponse) Descriptor() ([]byte, []int) {
	return file_tb_proto_rawDescGZIP(), []int{9}
}

func (x *ListNodesResponse) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *ListNodesResponse) GetNetwork() *NodeNetwork {
	if x != nil {
		return x.Network
	}
	return nil
}

var File_tb_proto protoreflect.FileDescriptor

var file_tb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x74, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x74, 0x62, 0x22, 0x52,
	0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2a, 0x0a, 0x0a, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x74, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x74, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x10,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x74, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x52, 0x10, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x25, 0x0a, 0x0a, 0x4e, 0x6f,
	0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49,
	0x4e, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10,
	0x02, 0x22, 0x2b, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x74, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x91,
	0x02, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x0a, 0x2e, 0x74, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x62, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x8c, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x08, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x74, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x65, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x08,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x44, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61,
	0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x4d,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x74, 0x62, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x12, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x54, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x74, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2a, 0x37, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x41, 0x4d, 0x45, 0x52, 0x41, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x41, 0x44, 0x41, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x44, 0x41,
	0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x48, 0x45, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x03,
	0x32, 0xb7, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x42, 0x65, 0x64, 0x4d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x34, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x2e,
	0x74, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x74, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x12, 0x2e, 0x74, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x74, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x74, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x69, 0x74, 0x46, 0x75, 0x6e, 0x67,
	0x2f, 0x74, 0x62, 0x2d, 0x69, 0x6f, 0x74, 0x2f, 0x74, 0x62,
}

var (
	file_tb_proto_rawDescOnce sync.Once
	file_tb_proto_rawDescData = file_tb_proto_rawDesc
)

func file_tb_proto_rawDescGZIP() []byte {
	file_tb_proto_rawDescOnce.Do(func() {
		file_tb_proto_rawDescData = protoimpl.X.CompressGZIP(file_tb_proto_rawDescData)
	})
	return file_tb_proto_rawDescData
}

var file_tb_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tb_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tb_proto_goTypes = []interface{}{
	(Sensor)(0),               // 0: tb.Sensor
	(Node_NodeStatus)(0),      // 1: tb.Node.NodeStatus
	(*NodeSpec)(nil),          // 2: tb.NodeSpec
	(*Node)(nil),              // 3: tb.Node
	(*NodeNetwork)(nil),       // 4: tb.NodeNetwork
	(*Task)(nil),              // 5: tb.Task
	(*NewTaskRequest)(nil),    // 6: tb.NewTaskRequest
	(*NewTaskResponse)(nil),   // 7: tb.NewTaskResponse
	(*GetTaskRequest)(nil),    // 8: tb.GetTaskRequest
	(*GetTaskResponse)(nil),   // 9: tb.GetTaskResponse
	(*ListNodesRequest)(nil),  // 10: tb.ListNodesRequest
	(*ListNodesResponse)(nil), // 11: tb.ListNodesResponse
	nil,                       // 12: tb.Task.DeviceEntry
}
var file_tb_proto_depIdxs = []int32{
	0,  // 0: tb.NodeSpec.sensorType:type_name -> tb.Sensor
	1,  // 1: tb.Node.status:type_name -> tb.Node.NodeStatus
	0,  // 2: tb.Node.onlineSensorType:type_name -> tb.Sensor
	3,  // 3: tb.NodeNetwork.node:type_name -> tb.Node
	12, // 4: tb.Task.device:type_name -> tb.Task.DeviceEntry
	0,  // 5: tb.Task.sensor:type_name -> tb.Sensor
	2,  // 6: tb.Task.nodeSpec:type_name -> tb.NodeSpec
	5,  // 7: tb.NewTaskRequest.taskInfo:type_name -> tb.Task
	5,  // 8: tb.NewTaskResponse.taskInfo:type_name -> tb.Task
	5,  // 9: tb.GetTaskResponse.taskInfo:type_name -> tb.Task
	4,  // 10: tb.ListNodesResponse.network:type_name -> tb.NodeNetwork
	6,  // 11: tb.TestBedMaster.NewTask:input_type -> tb.NewTaskRequest
	8,  // 12: tb.TestBedMaster.GetTask:input_type -> tb.GetTaskRequest
	10, // 13: tb.TestBedMaster.ListNodes:input_type -> tb.ListNodesRequest
	7,  // 14: tb.TestBedMaster.NewTask:output_type -> tb.NewTaskResponse
	9,  // 15: tb.TestBedMaster.GetTask:output_type -> tb.GetTaskResponse
	11, // 16: tb.TestBedMaster.ListNodes:output_type -> tb.ListNodesResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_tb_proto_init() }
func file_tb_proto_init() {
	if File_tb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeSpec); i {
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
		file_tb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_tb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeNetwork); i {
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
		file_tb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_tb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTaskRequest); i {
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
		file_tb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewTaskResponse); i {
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
		file_tb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRequest); i {
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
		file_tb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskResponse); i {
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
		file_tb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodesRequest); i {
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
		file_tb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodesResponse); i {
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
			RawDescriptor: file_tb_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tb_proto_goTypes,
		DependencyIndexes: file_tb_proto_depIdxs,
		EnumInfos:         file_tb_proto_enumTypes,
		MessageInfos:      file_tb_proto_msgTypes,
	}.Build()
	File_tb_proto = out.File
	file_tb_proto_rawDesc = nil
	file_tb_proto_goTypes = nil
	file_tb_proto_depIdxs = nil
}
