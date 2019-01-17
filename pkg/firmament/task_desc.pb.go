/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task_desc.proto

package firmament

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskDescriptor_TaskState int32

const (
	TaskDescriptor_CREATED   TaskDescriptor_TaskState = 0
	TaskDescriptor_BLOCKING  TaskDescriptor_TaskState = 1
	TaskDescriptor_RUNNABLE  TaskDescriptor_TaskState = 2
	TaskDescriptor_ASSIGNED  TaskDescriptor_TaskState = 3
	TaskDescriptor_RUNNING   TaskDescriptor_TaskState = 4
	TaskDescriptor_COMPLETED TaskDescriptor_TaskState = 5
	TaskDescriptor_FAILED    TaskDescriptor_TaskState = 6
	TaskDescriptor_ABORTED   TaskDescriptor_TaskState = 7
	TaskDescriptor_DELEGATED TaskDescriptor_TaskState = 8
	TaskDescriptor_UNKNOWN   TaskDescriptor_TaskState = 9
)

var TaskDescriptor_TaskState_name = map[int32]string{
	0: "CREATED",
	1: "BLOCKING",
	2: "RUNNABLE",
	3: "ASSIGNED",
	4: "RUNNING",
	5: "COMPLETED",
	6: "FAILED",
	7: "ABORTED",
	8: "DELEGATED",
	9: "UNKNOWN",
}
var TaskDescriptor_TaskState_value = map[string]int32{
	"CREATED":   0,
	"BLOCKING":  1,
	"RUNNABLE":  2,
	"ASSIGNED":  3,
	"RUNNING":   4,
	"COMPLETED": 5,
	"FAILED":    6,
	"ABORTED":   7,
	"DELEGATED": 8,
	"UNKNOWN":   9,
}

func (x TaskDescriptor_TaskState) String() string {
	return proto.EnumName(TaskDescriptor_TaskState_name, int32(x))
}
func (TaskDescriptor_TaskState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_task_desc_5d418c25e1058e8b, []int{0, 0}
}

type TaskDescriptor_TaskType int32

const (
	TaskDescriptor_SHEEP  TaskDescriptor_TaskType = 0
	TaskDescriptor_RABBIT TaskDescriptor_TaskType = 1
	TaskDescriptor_DEVIL  TaskDescriptor_TaskType = 2
	TaskDescriptor_TURTLE TaskDescriptor_TaskType = 3
)

var TaskDescriptor_TaskType_name = map[int32]string{
	0: "SHEEP",
	1: "RABBIT",
	2: "DEVIL",
	3: "TURTLE",
}
var TaskDescriptor_TaskType_value = map[string]int32{
	"SHEEP":  0,
	"RABBIT": 1,
	"DEVIL":  2,
	"TURTLE": 3,
}

func (x TaskDescriptor_TaskType) String() string {
	return proto.EnumName(TaskDescriptor_TaskType_name, int32(x))
}
func (TaskDescriptor_TaskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_task_desc_5d418c25e1058e8b, []int{0, 1}
}

// TaskDescriptor describes a task in firmament scheduler.
type TaskDescriptor struct {
	Uid   uint64                   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name  string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State TaskDescriptor_TaskState `protobuf:"varint,3,opt,name=state,proto3,enum=firmament.TaskDescriptor_TaskState" json:"state,omitempty"`
	JobId string                   `protobuf:"bytes,4,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Index uint64                   `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	// Inputs/outputs
	Dependencies []*ReferenceDescriptor `protobuf:"bytes,6,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Outputs      []*ReferenceDescriptor `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// Command and arguments
	Binary string   `protobuf:"bytes,8,opt,name=binary,proto3" json:"binary,omitempty"`
	Args   []string `protobuf:"bytes,9,rep,name=args,proto3" json:"args,omitempty"`
	// spawned is the list of children tasks
	Spawned []*TaskDescriptor `protobuf:"bytes,10,rep,name=spawned,proto3" json:"spawned,omitempty"`
	// Runtime meta-data
	ScheduledToResource   string `protobuf:"bytes,11,opt,name=scheduled_to_resource,json=scheduledToResource,proto3" json:"scheduled_to_resource,omitempty"`
	LastHeartbeatLocation string `protobuf:"bytes,12,opt,name=last_heartbeat_location,json=lastHeartbeatLocation,proto3" json:"last_heartbeat_location,omitempty"`
	LastHeartbeatTime     uint64 `protobuf:"varint,13,opt,name=last_heartbeat_time,json=lastHeartbeatTime,proto3" json:"last_heartbeat_time,omitempty"`
	// Delegation info
	DelegatedTo   string `protobuf:"bytes,14,opt,name=delegated_to,json=delegatedTo,proto3" json:"delegated_to,omitempty"`
	DelegatedFrom string `protobuf:"bytes,15,opt,name=delegated_from,json=delegatedFrom,proto3" json:"delegated_from,omitempty"`
	// Timestamps
	SubmitTime uint64 `protobuf:"varint,16,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	StartTime  uint64 `protobuf:"varint,17,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	FinishTime uint64 `protobuf:"varint,18,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	// The total time spent unscheduled before previous runs.
	TotalUnscheduledTime uint64 `protobuf:"varint,19,opt,name=total_unscheduled_time,json=totalUnscheduledTime,proto3" json:"total_unscheduled_time,omitempty"`
	// The total time spent in previous runs. This field only gets updated when
	//   the task finishes running.
	TotalRunTime uint64 `protobuf:"varint,20,opt,name=total_run_time,json=totalRunTime,proto3" json:"total_run_time,omitempty"`
	// Deadlines
	RelativeDeadline uint64 `protobuf:"varint,21,opt,name=relative_deadline,json=relativeDeadline,proto3" json:"relative_deadline,omitempty"`
	AbsoluteDeadline uint64 `protobuf:"varint,22,opt,name=absolute_deadline,json=absoluteDeadline,proto3" json:"absolute_deadline,omitempty"`
	// Application-specific fields
	// TODO(malte): move these to sub-messages
	Port      uint64 `protobuf:"varint,23,opt,name=port,proto3" json:"port,omitempty"`
	InputSize uint64 `protobuf:"varint,24,opt,name=input_size,json=inputSize,proto3" json:"input_size,omitempty"`
	// TaskLib related stuff
	InjectTaskLib bool `protobuf:"varint,25,opt,name=inject_task_lib,json=injectTaskLib,proto3" json:"inject_task_lib,omitempty"`
	// Task resource request and priority
	ResourceRequest *ResourceVector `protobuf:"bytes,26,opt,name=resource_request,json=resourceRequest,proto3" json:"resource_request,omitempty"`
	Priority        uint32          `protobuf:"varint,27,opt,name=priority,proto3" json:"priority,omitempty"`
	// TODO(malte): move this to a policy-specific sub-message
	TaskType TaskDescriptor_TaskType `protobuf:"varint,28,opt,name=task_type,json=taskType,proto3,enum=firmament.TaskDescriptor_TaskType" json:"task_type,omitempty"`
	// Final report of a task after successful execution
	FinalReport *TaskFinalReport `protobuf:"bytes,29,opt,name=final_report,json=finalReport,proto3" json:"final_report,omitempty"`
	// Simulation related fields
	TraceJobId  uint64 `protobuf:"varint,30,opt,name=trace_job_id,json=traceJobId,proto3" json:"trace_job_id,omitempty"`
	TraceTaskId uint64 `protobuf:"varint,31,opt,name=trace_task_id,json=traceTaskId,proto3" json:"trace_task_id,omitempty"`
	// Task labels
	Labels []*Label `protobuf:"bytes,32,rep,name=labels,proto3" json:"labels,omitempty"`
	// Resource label selectors
	LabelSelectors []*LabelSelector `protobuf:"bytes,33,rep,name=label_selectors,json=labelSelectors,proto3" json:"label_selectors,omitempty"`
	// Affinity
	Affinity *Affinity `protobuf:"bytes,34,opt,name=affinity,proto3" json:"affinity,omitempty"`
	// NameSpace
	Namespace string `protobuf:"bytes,35,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Toleration
	Toleration []*Toleration `protobuf:"bytes,36,rep,name=toleration,proto3" json:"toleration,omitempty"`
	// Owner reference kind
	OwnerRefKind string `protobuf:"bytes,37,opt,name=owner_ref_kind,json=ownerRefKind,proto3" json:"owner_ref_kind,omitempty"`
	// Owner reference uid
	OwnerRefUid          string   `protobuf:"bytes,38,opt,name=owner_ref_uid,json=ownerRefUid,proto3" json:"owner_ref_uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskDescriptor) Reset()         { *m = TaskDescriptor{} }
func (m *TaskDescriptor) String() string { return proto.CompactTextString(m) }
func (*TaskDescriptor) ProtoMessage()    {}
func (*TaskDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_desc_5d418c25e1058e8b, []int{0}
}
func (m *TaskDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDescriptor.Unmarshal(m, b)
}
func (m *TaskDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDescriptor.Marshal(b, m, deterministic)
}
func (dst *TaskDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDescriptor.Merge(dst, src)
}
func (m *TaskDescriptor) XXX_Size() int {
	return xxx_messageInfo_TaskDescriptor.Size(m)
}
func (m *TaskDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDescriptor proto.InternalMessageInfo

func (m *TaskDescriptor) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *TaskDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskDescriptor) GetState() TaskDescriptor_TaskState {
	if m != nil {
		return m.State
	}
	return TaskDescriptor_CREATED
}

func (m *TaskDescriptor) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *TaskDescriptor) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TaskDescriptor) GetDependencies() []*ReferenceDescriptor {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *TaskDescriptor) GetOutputs() []*ReferenceDescriptor {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *TaskDescriptor) GetBinary() string {
	if m != nil {
		return m.Binary
	}
	return ""
}

func (m *TaskDescriptor) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *TaskDescriptor) GetSpawned() []*TaskDescriptor {
	if m != nil {
		return m.Spawned
	}
	return nil
}

func (m *TaskDescriptor) GetScheduledToResource() string {
	if m != nil {
		return m.ScheduledToResource
	}
	return ""
}

func (m *TaskDescriptor) GetLastHeartbeatLocation() string {
	if m != nil {
		return m.LastHeartbeatLocation
	}
	return ""
}

func (m *TaskDescriptor) GetLastHeartbeatTime() uint64 {
	if m != nil {
		return m.LastHeartbeatTime
	}
	return 0
}

func (m *TaskDescriptor) GetDelegatedTo() string {
	if m != nil {
		return m.DelegatedTo
	}
	return ""
}

func (m *TaskDescriptor) GetDelegatedFrom() string {
	if m != nil {
		return m.DelegatedFrom
	}
	return ""
}

func (m *TaskDescriptor) GetSubmitTime() uint64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *TaskDescriptor) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *TaskDescriptor) GetFinishTime() uint64 {
	if m != nil {
		return m.FinishTime
	}
	return 0
}

func (m *TaskDescriptor) GetTotalUnscheduledTime() uint64 {
	if m != nil {
		return m.TotalUnscheduledTime
	}
	return 0
}

func (m *TaskDescriptor) GetTotalRunTime() uint64 {
	if m != nil {
		return m.TotalRunTime
	}
	return 0
}

func (m *TaskDescriptor) GetRelativeDeadline() uint64 {
	if m != nil {
		return m.RelativeDeadline
	}
	return 0
}

func (m *TaskDescriptor) GetAbsoluteDeadline() uint64 {
	if m != nil {
		return m.AbsoluteDeadline
	}
	return 0
}

func (m *TaskDescriptor) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *TaskDescriptor) GetInputSize() uint64 {
	if m != nil {
		return m.InputSize
	}
	return 0
}

func (m *TaskDescriptor) GetInjectTaskLib() bool {
	if m != nil {
		return m.InjectTaskLib
	}
	return false
}

func (m *TaskDescriptor) GetResourceRequest() *ResourceVector {
	if m != nil {
		return m.ResourceRequest
	}
	return nil
}

func (m *TaskDescriptor) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TaskDescriptor) GetTaskType() TaskDescriptor_TaskType {
	if m != nil {
		return m.TaskType
	}
	return TaskDescriptor_SHEEP
}

func (m *TaskDescriptor) GetFinalReport() *TaskFinalReport {
	if m != nil {
		return m.FinalReport
	}
	return nil
}

func (m *TaskDescriptor) GetTraceJobId() uint64 {
	if m != nil {
		return m.TraceJobId
	}
	return 0
}

func (m *TaskDescriptor) GetTraceTaskId() uint64 {
	if m != nil {
		return m.TraceTaskId
	}
	return 0
}

func (m *TaskDescriptor) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TaskDescriptor) GetLabelSelectors() []*LabelSelector {
	if m != nil {
		return m.LabelSelectors
	}
	return nil
}

func (m *TaskDescriptor) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *TaskDescriptor) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TaskDescriptor) GetToleration() []*Toleration {
	if m != nil {
		return m.Toleration
	}
	return nil
}

func (m *TaskDescriptor) GetOwnerRefKind() string {
	if m != nil {
		return m.OwnerRefKind
	}
	return ""
}

func (m *TaskDescriptor) GetOwnerRefUid() string {
	if m != nil {
		return m.OwnerRefUid
	}
	return ""
}

func init() {
	proto.RegisterType((*TaskDescriptor)(nil), "firmament.TaskDescriptor")
	proto.RegisterEnum("firmament.TaskDescriptor_TaskState", TaskDescriptor_TaskState_name, TaskDescriptor_TaskState_value)
	proto.RegisterEnum("firmament.TaskDescriptor_TaskType", TaskDescriptor_TaskType_name, TaskDescriptor_TaskType_value)
}

func init() { proto.RegisterFile("task_desc.proto", fileDescriptor_task_desc_5d418c25e1058e8b) }

var fileDescriptor_task_desc_5d418c25e1058e8b = []byte{
	// 1014 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6f, 0x1a, 0x47,
	0x10, 0x0f, 0xc1, 0x60, 0x18, 0x3e, 0x7c, 0x5e, 0x1b, 0x7b, 0x43, 0x93, 0x98, 0x90, 0x0f, 0x21,
	0x55, 0x72, 0x25, 0xa7, 0xad, 0xd2, 0x87, 0xaa, 0x02, 0x73, 0x76, 0xa8, 0x29, 0x8e, 0x16, 0x9c,
	0x3e, 0x9e, 0x16, 0x6e, 0x89, 0xd7, 0x3e, 0xee, 0xe8, 0xee, 0x5e, 0x52, 0xe7, 0x4f, 0xe8, 0x6b,
	0xff, 0xe1, 0x6a, 0xe7, 0xee, 0x00, 0x5b, 0x6a, 0xd5, 0xb7, 0x9d, 0xdf, 0xc7, 0xcc, 0x32, 0xcc,
	0xce, 0xc1, 0x8e, 0xe1, 0xfa, 0xd6, 0xf3, 0x85, 0x9e, 0x1d, 0x2f, 0x55, 0x64, 0x22, 0x52, 0x9e,
	0x4b, 0xb5, 0xe0, 0x0b, 0x11, 0x9a, 0x66, 0x25, 0xe0, 0x53, 0x11, 0x24, 0x78, 0x73, 0x1f, 0x03,
	0x4f, 0x8b, 0x40, 0xcc, 0x4c, 0xa4, 0x32, 0x54, 0x89, 0xb9, 0x50, 0x22, 0x9c, 0x89, 0x8d, 0x1c,
	0xcd, 0x86, 0x12, 0x3a, 0x8a, 0xd5, 0x4c, 0x78, 0x9f, 0x37, 0xc5, 0x87, 0x58, 0x6b, 0x2e, 0x43,
	0x1e, 0x78, 0x4a, 0x2c, 0x23, 0x65, 0x52, 0xa2, 0xce, 0xe7, 0x73, 0x19, 0x4a, 0x73, 0x97, 0xc6,
	0xbb, 0x26, 0x0a, 0x84, 0xe2, 0x46, 0x46, 0xa1, 0x4e, 0xa0, 0xf6, 0x5f, 0x75, 0xa8, 0x4f, 0xb8,
	0xbe, 0xed, 0x0b, 0x3d, 0x53, 0x72, 0x69, 0x22, 0x45, 0x1c, 0xc8, 0xc7, 0xd2, 0xa7, 0xb9, 0x56,
	0xae, 0xb3, 0xc5, 0xec, 0x91, 0x10, 0xd8, 0x0a, 0xf9, 0x42, 0xd0, 0xc7, 0xad, 0x5c, 0xa7, 0xcc,
	0xf0, 0x4c, 0x7e, 0x82, 0x82, 0x36, 0xdc, 0x08, 0x9a, 0x6f, 0xe5, 0x3a, 0xf5, 0x93, 0x97, 0xc7,
	0xab, 0xdf, 0x77, 0x7c, 0x3f, 0x1f, 0x86, 0x63, 0x2b, 0x65, 0x89, 0x83, 0x34, 0xa0, 0x78, 0x13,
	0x4d, 0x3d, 0xe9, 0xd3, 0x2d, 0x4c, 0x58, 0xb8, 0x89, 0xa6, 0x03, 0x9f, 0xec, 0x43, 0x41, 0x86,
	0xbe, 0xf8, 0x93, 0x16, 0xb0, 0x72, 0x12, 0x90, 0x1e, 0x54, 0x7d, 0xb1, 0x14, 0xa1, 0x2f, 0xc2,
	0x99, 0x14, 0x9a, 0x16, 0x5b, 0xf9, 0x4e, 0xe5, 0xe4, 0xf9, 0x46, 0x39, 0x96, 0xb5, 0x6a, 0x5d,
	0x93, 0xdd, 0xf3, 0x90, 0x77, 0xb0, 0x1d, 0xc5, 0x66, 0x19, 0x1b, 0x4d, 0xb7, 0xff, 0x97, 0x3d,
	0x93, 0x93, 0x03, 0x28, 0x4e, 0x65, 0xc8, 0xd5, 0x1d, 0x2d, 0xe1, 0x55, 0xd3, 0xc8, 0x76, 0x84,
	0xab, 0x4f, 0x9a, 0x96, 0x5b, 0x79, 0xdb, 0x11, 0x7b, 0x26, 0x6f, 0x61, 0x5b, 0x2f, 0xf9, 0x97,
	0x50, 0xf8, 0x14, 0xb0, 0xca, 0x93, 0x7f, 0xed, 0x09, 0xcb, 0x94, 0xe4, 0x04, 0x1a, 0x7a, 0x76,
	0x2d, 0xfc, 0x38, 0x10, 0xbe, 0x67, 0x22, 0x2f, 0xfb, 0x87, 0x69, 0x05, 0xeb, 0xed, 0xad, 0xc8,
	0x49, 0xc4, 0x52, 0x8a, 0xfc, 0x08, 0x87, 0x01, 0xd7, 0xc6, 0xbb, 0x16, 0x5c, 0x99, 0xa9, 0xe0,
	0xc6, 0x0b, 0xa2, 0x19, 0xfe, 0xab, 0xb4, 0x8a, 0xae, 0x86, 0xa5, 0xdf, 0x67, 0xec, 0x30, 0x25,
	0xc9, 0x31, 0xec, 0x3d, 0xf0, 0x19, 0xb9, 0x10, 0xb4, 0x86, 0xed, 0xde, 0xbd, 0xe7, 0x99, 0xc8,
	0x85, 0x20, 0x2f, 0x6c, 0xeb, 0x03, 0xf1, 0x89, 0x1b, 0xbc, 0x1b, 0xad, 0x63, 0xf2, 0xca, 0x0a,
	0x9b, 0x44, 0xe4, 0x35, 0xd4, 0xd7, 0x92, 0xb9, 0x8a, 0x16, 0x74, 0x07, 0x45, 0xb5, 0x15, 0x7a,
	0xa6, 0xa2, 0x05, 0x39, 0x82, 0x8a, 0x8e, 0xa7, 0x0b, 0x99, 0x56, 0x74, 0xb0, 0x22, 0x24, 0x10,
	0x96, 0x7a, 0x06, 0xa0, 0x0d, 0x57, 0x29, 0xbf, 0x8b, 0x7c, 0x19, 0x11, 0xa4, 0x8f, 0xa0, 0x62,
	0x07, 0x59, 0x5f, 0x27, 0x3c, 0x49, 0xfc, 0x09, 0x84, 0x82, 0xef, 0xe1, 0xc0, 0x44, 0x86, 0x07,
	0x5e, 0x1c, 0x6e, 0xb4, 0xd3, 0x6a, 0xf7, 0x50, 0xbb, 0x8f, 0xec, 0xd5, 0x9a, 0x44, 0xd7, 0x2b,
	0xa8, 0x27, 0x2e, 0x15, 0x87, 0x89, 0x7a, 0x1f, 0xd5, 0x55, 0x44, 0x59, 0x1c, 0xa2, 0xea, 0x5b,
	0xd8, 0x55, 0x22, 0xe0, 0x46, 0x7e, 0xb6, 0x8f, 0x91, 0xfb, 0x81, 0x0c, 0x05, 0x6d, 0xa0, 0xd0,
	0xc9, 0x88, 0x7e, 0x8a, 0x5b, 0x31, 0x9f, 0xea, 0x28, 0x88, 0xcd, 0x86, 0xf8, 0x20, 0x11, 0x67,
	0xc4, 0x4a, 0x4c, 0x60, 0xcb, 0xbe, 0x56, 0x7a, 0x88, 0x3c, 0x9e, 0x6d, 0x27, 0x64, 0xb8, 0x8c,
	0x8d, 0xa7, 0xe5, 0x57, 0x41, 0x69, 0xd2, 0x09, 0x44, 0xc6, 0xf2, 0xab, 0x20, 0x6f, 0x60, 0x47,
	0x86, 0x37, 0x62, 0x66, 0x3c, 0x7c, 0xf4, 0x81, 0x9c, 0xd2, 0x27, 0xad, 0x5c, 0xa7, 0xc4, 0x6a,
	0x09, 0x6c, 0xe7, 0x6c, 0x28, 0xa7, 0xa4, 0x0f, 0xce, 0x6a, 0x59, 0x28, 0xf1, 0x47, 0x2c, 0xb4,
	0xa1, 0xcd, 0x56, 0xee, 0xc1, 0x54, 0x66, 0x23, 0xf5, 0x11, 0xd7, 0x09, 0xdb, 0xc9, 0x2c, 0x2c,
	0x71, 0x90, 0x26, 0x94, 0x96, 0x4a, 0x46, 0x4a, 0x9a, 0x3b, 0xfa, 0x4d, 0x2b, 0xd7, 0xa9, 0xb1,
	0x55, 0x4c, 0x7e, 0x81, 0x32, 0x5e, 0xc1, 0xdc, 0x2d, 0x05, 0x7d, 0x8a, 0x4b, 0xa0, 0xfd, 0xdf,
	0x4b, 0x60, 0x72, 0xb7, 0x14, 0xac, 0x64, 0xd2, 0x13, 0xf9, 0x19, 0xaa, 0x9b, 0x3b, 0x8b, 0x3e,
	0xc3, 0xeb, 0x35, 0x1f, 0xe4, 0x38, 0xb3, 0x12, 0x86, 0x0a, 0x56, 0x99, 0xaf, 0x03, 0xd2, 0x82,
	0xaa, 0x51, 0x7c, 0x26, 0xbc, 0x74, 0x97, 0x3c, 0x4f, 0x86, 0x02, 0xb1, 0x5f, 0x71, 0xa1, 0xb4,
	0xa1, 0x96, 0x28, 0xf0, 0x9e, 0xd2, 0xa7, 0x47, 0x28, 0xa9, 0x20, 0x68, 0x73, 0x0f, 0x7c, 0xd2,
	0x81, 0x22, 0x2e, 0x60, 0x4d, 0x5b, 0xf8, 0x66, 0x9d, 0x8d, 0xf2, 0x43, 0x4b, 0xb0, 0x94, 0x27,
	0x5d, 0xd8, 0xb9, 0xbf, 0xaa, 0x35, 0x7d, 0x81, 0x16, 0xfa, 0xd0, 0x32, 0x4e, 0x05, 0xac, 0x1e,
	0x6c, 0x86, 0x9a, 0x7c, 0x07, 0xa5, 0x6c, 0x23, 0xd3, 0x36, 0xfe, 0xda, 0xbd, 0x0d, 0x6f, 0x37,
	0xa5, 0xd8, 0x4a, 0x44, 0x9e, 0x42, 0xd9, 0x2e, 0x5b, 0xbd, 0xe4, 0x33, 0x41, 0x5f, 0xe2, 0xcb,
	0x5a, 0x03, 0xe4, 0x07, 0x80, 0xf5, 0x42, 0xa7, 0xaf, 0xf0, 0x32, 0x8d, 0xcd, 0xf6, 0xad, 0x48,
	0xb6, 0x21, 0xb4, 0x53, 0x1f, 0x7d, 0x09, 0x85, 0xf2, 0x94, 0x98, 0x7b, 0xb7, 0x32, 0xf4, 0xe9,
	0x6b, 0xcc, 0x5c, 0x45, 0x94, 0x89, 0xf9, 0x85, 0x0c, 0xb1, 0x79, 0x6b, 0x95, 0xfd, 0x1e, 0xbc,
	0x49, 0x5e, 0x7f, 0x26, 0xba, 0x92, 0x7e, 0xfb, 0xef, 0x1c, 0x94, 0x57, 0xdb, 0x9d, 0x54, 0x60,
	0xfb, 0x94, 0xb9, 0xdd, 0x89, 0xdb, 0x77, 0x1e, 0x91, 0x2a, 0x94, 0x7a, 0xc3, 0xcb, 0xd3, 0x8b,
	0xc1, 0xe8, 0xdc, 0xc9, 0xd9, 0x88, 0x5d, 0x8d, 0x46, 0xdd, 0xde, 0xd0, 0x75, 0x1e, 0xdb, 0xa8,
	0x3b, 0x1e, 0x0f, 0xce, 0x47, 0x6e, 0xdf, 0xc9, 0x5b, 0x9b, 0xe5, 0xac, 0x70, 0x8b, 0xd4, 0xa0,
	0x7c, 0x7a, 0xf9, 0xdb, 0x87, 0xa1, 0x6b, 0xb3, 0x14, 0x08, 0x40, 0xf1, 0xac, 0x3b, 0x18, 0xba,
	0x7d, 0xa7, 0x68, 0x75, 0xdd, 0xde, 0x25, 0xb3, 0xc4, 0xb6, 0xd5, 0xf5, 0xdd, 0xa1, 0x7b, 0x8e,
	0xd5, 0x4a, 0x96, 0xbb, 0x1a, 0x5d, 0x8c, 0x2e, 0x7f, 0x1f, 0x39, 0xe5, 0xf6, 0x3b, 0x28, 0x65,
	0xd3, 0x46, 0xca, 0x50, 0x18, 0xbf, 0x77, 0xdd, 0x0f, 0xce, 0x23, 0x9b, 0x8b, 0x75, 0x7b, 0xbd,
	0xc1, 0xc4, 0xc9, 0x59, 0xb8, 0xef, 0x7e, 0x1c, 0x0c, 0x9d, 0xc7, 0x16, 0x9e, 0x5c, 0xb1, 0xc9,
	0xd0, 0x75, 0xf2, 0xd3, 0x22, 0x7e, 0x13, 0xdf, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x1f,
	0xd4, 0x3c, 0xbd, 0x07, 0x00, 0x00,
}
