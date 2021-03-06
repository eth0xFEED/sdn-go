// Code generated by protoc-gen-go.
// source: proto/p4/config/p4info.proto
// DO NOT EDIT!

/*
Package p4_config is a generated protocol buffer package.

It is generated from these files:
	proto/p4/config/p4info.proto

It has these top-level messages:
	P4Info
	Preamble
	Extern
	ExternInstance
	MatchField
	Table
	ActionRef
	Action
	ActionProfile
	CounterSpec
	Counter
	DirectCounter
	MeterSpec
	Meter
	DirectMeter
	ControllerPacketMetadata
*/
package p4_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MatchField_MatchType int32

const (
	MatchField_UNSPECIFIED MatchField_MatchType = 0
	MatchField_VALID       MatchField_MatchType = 1
	MatchField_EXACT       MatchField_MatchType = 2
	MatchField_LPM         MatchField_MatchType = 3
	MatchField_TERNARY     MatchField_MatchType = 4
	MatchField_RANGE       MatchField_MatchType = 5
)

var MatchField_MatchType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "VALID",
	2: "EXACT",
	3: "LPM",
	4: "TERNARY",
	5: "RANGE",
}
var MatchField_MatchType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"VALID":       1,
	"EXACT":       2,
	"LPM":         3,
	"TERNARY":     4,
	"RANGE":       5,
}

func (x MatchField_MatchType) String() string {
	return proto.EnumName(MatchField_MatchType_name, int32(x))
}
func (MatchField_MatchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

// Corresponds to 'type' attribute for counter in P4 spec.
type CounterSpec_Unit int32

const (
	CounterSpec_UNSPECIFIED CounterSpec_Unit = 0
	CounterSpec_BYTES       CounterSpec_Unit = 1
	CounterSpec_PACKETS     CounterSpec_Unit = 2
	CounterSpec_BOTH        CounterSpec_Unit = 3
)

var CounterSpec_Unit_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "BYTES",
	2: "PACKETS",
	3: "BOTH",
}
var CounterSpec_Unit_value = map[string]int32{
	"UNSPECIFIED": 0,
	"BYTES":       1,
	"PACKETS":     2,
	"BOTH":        3,
}

func (x CounterSpec_Unit) String() string {
	return proto.EnumName(CounterSpec_Unit_name, int32(x))
}
func (CounterSpec_Unit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9, 0} }

// Corresponds to 'type' attribute for meter in P4 spec.
type MeterSpec_Unit int32

const (
	MeterSpec_UNSPECIFIED MeterSpec_Unit = 0
	MeterSpec_BYTES       MeterSpec_Unit = 1
	MeterSpec_PACKETS     MeterSpec_Unit = 2
)

var MeterSpec_Unit_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "BYTES",
	2: "PACKETS",
}
var MeterSpec_Unit_value = map[string]int32{
	"UNSPECIFIED": 0,
	"BYTES":       1,
	"PACKETS":     2,
}

func (x MeterSpec_Unit) String() string {
	return proto.EnumName(MeterSpec_Unit_name, int32(x))
}
func (MeterSpec_Unit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{12, 0} }

// Not part of the P4 spec yet but will be in the future.
type MeterSpec_Type int32

const (
	MeterSpec_COLOR_UNAWARE MeterSpec_Type = 0
	MeterSpec_COLOR_AWARE   MeterSpec_Type = 1
)

var MeterSpec_Type_name = map[int32]string{
	0: "COLOR_UNAWARE",
	1: "COLOR_AWARE",
}
var MeterSpec_Type_value = map[string]int32{
	"COLOR_UNAWARE": 0,
	"COLOR_AWARE":   1,
}

func (x MeterSpec_Type) String() string {
	return proto.EnumName(MeterSpec_Type_name, int32(x))
}
func (MeterSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{12, 1} }

type P4Info struct {
	Externs                  []*Extern                   `protobuf:"bytes,1,rep,name=externs" json:"externs,omitempty"`
	Tables                   []*Table                    `protobuf:"bytes,2,rep,name=tables" json:"tables,omitempty"`
	Actions                  []*Action                   `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
	ActionProfiles           []*ActionProfile            `protobuf:"bytes,4,rep,name=action_profiles,json=actionProfiles" json:"action_profiles,omitempty"`
	Counters                 []*Counter                  `protobuf:"bytes,5,rep,name=counters" json:"counters,omitempty"`
	DirectCounters           []*DirectCounter            `protobuf:"bytes,6,rep,name=direct_counters,json=directCounters" json:"direct_counters,omitempty"`
	Meters                   []*Meter                    `protobuf:"bytes,7,rep,name=meters" json:"meters,omitempty"`
	DirectMeters             []*DirectMeter              `protobuf:"bytes,8,rep,name=direct_meters,json=directMeters" json:"direct_meters,omitempty"`
	ControllerPacketMetadata []*ControllerPacketMetadata `protobuf:"bytes,9,rep,name=controller_packet_metadata,json=controllerPacketMetadata" json:"controller_packet_metadata,omitempty"`
}

func (m *P4Info) Reset()                    { *m = P4Info{} }
func (m *P4Info) String() string            { return proto.CompactTextString(m) }
func (*P4Info) ProtoMessage()               {}
func (*P4Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *P4Info) GetExterns() []*Extern {
	if m != nil {
		return m.Externs
	}
	return nil
}

func (m *P4Info) GetTables() []*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

func (m *P4Info) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *P4Info) GetActionProfiles() []*ActionProfile {
	if m != nil {
		return m.ActionProfiles
	}
	return nil
}

func (m *P4Info) GetCounters() []*Counter {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *P4Info) GetDirectCounters() []*DirectCounter {
	if m != nil {
		return m.DirectCounters
	}
	return nil
}

func (m *P4Info) GetMeters() []*Meter {
	if m != nil {
		return m.Meters
	}
	return nil
}

func (m *P4Info) GetDirectMeters() []*DirectMeter {
	if m != nil {
		return m.DirectMeters
	}
	return nil
}

func (m *P4Info) GetControllerPacketMetadata() []*ControllerPacketMetadata {
	if m != nil {
		return m.ControllerPacketMetadata
	}
	return nil
}

type Preamble struct {
	// ids share the same number-space; e.g. table ids cannot overlap with counter
	// ids. Even though this is irrelevant to this proto definition, the ids are
	// allocated in such a way that it is possible based on an id to deduce the
	// resource type (e.g. table, action, counter, ...). This means that code
	// using these ids can detect if the wrong resource type is used
	// somewhere. This also means that ids of different types can be mixed
	// (e.g. direct resource list for a table) without ambiguity. Note that id 0
	// is reserved and means "invalid id".
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// fully qualified name of the P4 object, e.g. c1.c2.ipv4_lpm
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// an alias for the P4 object, probably shorter than its name. The only
	// constraint is for it to be unique with respect to other P4 objects of the
	// same type. By default, the compiler uses the shortest suffix of the name
	// that uniquely indentifies the object. For example if the P4 program
	// contains two tables with names s.c1.t and s.c2.t, the default aliases will
	// respectively be c1.t and c2.t. The P4 programmer may also override the
	// default alias for any P4 object (TBD). When resolving a P4 object id, an
	// application should be able to indiscriminately use the name or the alias.
	Alias       string   `protobuf:"bytes,3,opt,name=alias" json:"alias,omitempty"`
	Annotations []string `protobuf:"bytes,4,rep,name=annotations" json:"annotations,omitempty"`
}

func (m *Preamble) Reset()                    { *m = Preamble{} }
func (m *Preamble) String() string            { return proto.CompactTextString(m) }
func (*Preamble) ProtoMessage()               {}
func (*Preamble) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Preamble) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Preamble) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Preamble) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Preamble) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

// used to group all extern instances of the same type in one message
type Extern struct {
	// the extern_type_id is assigned during compilation. It is likely that this
	// id will in fact come from a P4 annotation to the extern declaration and
	// that each vendor will receive a prefix to avoid collisions.
	ExternTypeId   uint32            `protobuf:"varint,1,opt,name=extern_type_id,json=externTypeId" json:"extern_type_id,omitempty"`
	ExternTypeName string            `protobuf:"bytes,2,opt,name=extern_type_name,json=externTypeName" json:"extern_type_name,omitempty"`
	Instances      []*ExternInstance `protobuf:"bytes,3,rep,name=instances" json:"instances,omitempty"`
}

func (m *Extern) Reset()                    { *m = Extern{} }
func (m *Extern) String() string            { return proto.CompactTextString(m) }
func (*Extern) ProtoMessage()               {}
func (*Extern) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Extern) GetExternTypeId() uint32 {
	if m != nil {
		return m.ExternTypeId
	}
	return 0
}

func (m *Extern) GetExternTypeName() string {
	if m != nil {
		return m.ExternTypeName
	}
	return ""
}

func (m *Extern) GetInstances() []*ExternInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type ExternInstance struct {
	Preamble *Preamble `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	// specific to the extern type, declared in a separate vendor-specific proto
	// file
	Info *google_protobuf.Any `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *ExternInstance) Reset()                    { *m = ExternInstance{} }
func (m *ExternInstance) String() string            { return proto.CompactTextString(m) }
func (*ExternInstance) ProtoMessage()               {}
func (*ExternInstance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExternInstance) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *ExternInstance) GetInfo() *google_protobuf.Any {
	if m != nil {
		return m.Info
	}
	return nil
}

// TODO(antonin): define inside the Table message?
type MatchField struct {
	Id          uint32               `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Annotations []string             `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty"`
	Bitwidth    int32                `protobuf:"varint,4,opt,name=bitwidth" json:"bitwidth,omitempty"`
	MatchType   MatchField_MatchType `protobuf:"varint,5,opt,name=match_type,json=matchType,enum=p4.config.MatchField_MatchType" json:"match_type,omitempty"`
}

func (m *MatchField) Reset()                    { *m = MatchField{} }
func (m *MatchField) String() string            { return proto.CompactTextString(m) }
func (*MatchField) ProtoMessage()               {}
func (*MatchField) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MatchField) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MatchField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MatchField) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *MatchField) GetBitwidth() int32 {
	if m != nil {
		return m.Bitwidth
	}
	return 0
}

func (m *MatchField) GetMatchType() MatchField_MatchType {
	if m != nil {
		return m.MatchType
	}
	return MatchField_UNSPECIFIED
}

type Table struct {
	Preamble    *Preamble     `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	MatchFields []*MatchField `protobuf:"bytes,2,rep,name=match_fields,json=matchFields" json:"match_fields,omitempty"`
	// even when the table is indirect (see implementation_id) below, this field
	// includes all possible actions for the table; by using ActionRef instead of
	// a repeated field of action ids, each action reference in a P4 table is able
	// to have its own annotations
	ActionRefs []*ActionRef `protobuf:"bytes,3,rep,name=action_refs,json=actionRefs" json:"action_refs,omitempty"`
	// 0 (default value) means that the table does not have a const default action
	ConstDefaultActionId uint32 `protobuf:"varint,4,opt,name=const_default_action_id,json=constDefaultActionId" json:"const_default_action_id,omitempty"`
	// a table may have a const default action, whose action parameter values can
	// be changed at runtime. However, in most cases the parameters of the default
	// action are also bound at compile-time and cannot be changed by the runtime,
	// which is what this boolean flag indicates.
	ConstDefaultActionHasMutableParams bool `protobuf:"varint,5,opt,name=const_default_action_has_mutable_params,json=constDefaultActionHasMutableParams" json:"const_default_action_has_mutable_params,omitempty"`
	// P4 id of the "implementation" for this table (e.g. action profile id); 0
	// (default value) means that the table is a regular (direct) match table. As
	// of today, only action profiles are supported but other table
	// implementations may be added in the future
	ImplementationId uint32 `protobuf:"varint,6,opt,name=implementation_id,json=implementationId" json:"implementation_id,omitempty"`
	// ids of the direct resources (if any) attached to this table; for now this
	// includes only direct counters and direct meters, but other resources may be
	// added in the future
	DirectResourceIds []uint32 `protobuf:"varint,7,rep,packed,name=direct_resource_ids,json=directResourceIds" json:"direct_resource_ids,omitempty"`
	Size              int64    `protobuf:"varint,8,opt,name=size" json:"size,omitempty"`
	WithEntryTimeout  bool     `protobuf:"varint,9,opt,name=with_entry_timeout,json=withEntryTimeout" json:"with_entry_timeout,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Table) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *Table) GetMatchFields() []*MatchField {
	if m != nil {
		return m.MatchFields
	}
	return nil
}

func (m *Table) GetActionRefs() []*ActionRef {
	if m != nil {
		return m.ActionRefs
	}
	return nil
}

func (m *Table) GetConstDefaultActionId() uint32 {
	if m != nil {
		return m.ConstDefaultActionId
	}
	return 0
}

func (m *Table) GetConstDefaultActionHasMutableParams() bool {
	if m != nil {
		return m.ConstDefaultActionHasMutableParams
	}
	return false
}

func (m *Table) GetImplementationId() uint32 {
	if m != nil {
		return m.ImplementationId
	}
	return 0
}

func (m *Table) GetDirectResourceIds() []uint32 {
	if m != nil {
		return m.DirectResourceIds
	}
	return nil
}

func (m *Table) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Table) GetWithEntryTimeout() bool {
	if m != nil {
		return m.WithEntryTimeout
	}
	return false
}

// used to list all possible actions in a Table
type ActionRef struct {
	Id          uint32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Annotations []string `protobuf:"bytes,2,rep,name=annotations" json:"annotations,omitempty"`
}

func (m *ActionRef) Reset()                    { *m = ActionRef{} }
func (m *ActionRef) String() string            { return proto.CompactTextString(m) }
func (*ActionRef) ProtoMessage()               {}
func (*ActionRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ActionRef) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ActionRef) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type Action struct {
	Preamble *Preamble       `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	Params   []*Action_Param `protobuf:"bytes,2,rep,name=params" json:"params,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Action) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *Action) GetParams() []*Action_Param {
	if m != nil {
		return m.Params
	}
	return nil
}

type Action_Param struct {
	Id          uint32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Annotations []string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty"`
	Bitwidth    int32    `protobuf:"varint,4,opt,name=bitwidth" json:"bitwidth,omitempty"`
}

func (m *Action_Param) Reset()                    { *m = Action_Param{} }
func (m *Action_Param) String() string            { return proto.CompactTextString(m) }
func (*Action_Param) ProtoMessage()               {}
func (*Action_Param) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *Action_Param) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Action_Param) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Action_Param) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Action_Param) GetBitwidth() int32 {
	if m != nil {
		return m.Bitwidth
	}
	return 0
}

type ActionProfile struct {
	Preamble *Preamble `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	// the ids of the tables sharing this action profile
	TableIds []uint32 `protobuf:"varint,2,rep,packed,name=table_ids,json=tableIds" json:"table_ids,omitempty"`
	// true iff the action profile used dynamic selection
	WithSelector bool  `protobuf:"varint,3,opt,name=with_selector,json=withSelector" json:"with_selector,omitempty"`
	Size         int64 `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
}

func (m *ActionProfile) Reset()                    { *m = ActionProfile{} }
func (m *ActionProfile) String() string            { return proto.CompactTextString(m) }
func (*ActionProfile) ProtoMessage()               {}
func (*ActionProfile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ActionProfile) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *ActionProfile) GetTableIds() []uint32 {
	if m != nil {
		return m.TableIds
	}
	return nil
}

func (m *ActionProfile) GetWithSelector() bool {
	if m != nil {
		return m.WithSelector
	}
	return false
}

func (m *ActionProfile) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type CounterSpec struct {
	Unit CounterSpec_Unit `protobuf:"varint,1,opt,name=unit,enum=p4.config.CounterSpec_Unit" json:"unit,omitempty"`
}

func (m *CounterSpec) Reset()                    { *m = CounterSpec{} }
func (m *CounterSpec) String() string            { return proto.CompactTextString(m) }
func (*CounterSpec) ProtoMessage()               {}
func (*CounterSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CounterSpec) GetUnit() CounterSpec_Unit {
	if m != nil {
		return m.Unit
	}
	return CounterSpec_UNSPECIFIED
}

type Counter struct {
	Preamble *Preamble    `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	Spec     *CounterSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// number of entries in the counter array
	Size int64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *Counter) Reset()                    { *m = Counter{} }
func (m *Counter) String() string            { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()               {}
func (*Counter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Counter) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *Counter) GetSpec() *CounterSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Counter) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type DirectCounter struct {
	Preamble *Preamble    `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	Spec     *CounterSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// the id of the table to which the counter is attached
	DirectTableId uint32 `protobuf:"varint,3,opt,name=direct_table_id,json=directTableId" json:"direct_table_id,omitempty"`
}

func (m *DirectCounter) Reset()                    { *m = DirectCounter{} }
func (m *DirectCounter) String() string            { return proto.CompactTextString(m) }
func (*DirectCounter) ProtoMessage()               {}
func (*DirectCounter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DirectCounter) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *DirectCounter) GetSpec() *CounterSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *DirectCounter) GetDirectTableId() uint32 {
	if m != nil {
		return m.DirectTableId
	}
	return 0
}

type MeterSpec struct {
	Unit MeterSpec_Unit `protobuf:"varint,1,opt,name=unit,enum=p4.config.MeterSpec_Unit" json:"unit,omitempty"`
	Type MeterSpec_Type `protobuf:"varint,2,opt,name=type,enum=p4.config.MeterSpec_Type" json:"type,omitempty"`
}

func (m *MeterSpec) Reset()                    { *m = MeterSpec{} }
func (m *MeterSpec) String() string            { return proto.CompactTextString(m) }
func (*MeterSpec) ProtoMessage()               {}
func (*MeterSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *MeterSpec) GetUnit() MeterSpec_Unit {
	if m != nil {
		return m.Unit
	}
	return MeterSpec_UNSPECIFIED
}

func (m *MeterSpec) GetType() MeterSpec_Type {
	if m != nil {
		return m.Type
	}
	return MeterSpec_COLOR_UNAWARE
}

type Meter struct {
	Preamble *Preamble  `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	Spec     *MeterSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// number of entries in the meter array
	Size int64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *Meter) Reset()                    { *m = Meter{} }
func (m *Meter) String() string            { return proto.CompactTextString(m) }
func (*Meter) ProtoMessage()               {}
func (*Meter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Meter) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *Meter) GetSpec() *MeterSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Meter) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type DirectMeter struct {
	Preamble *Preamble  `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	Spec     *MeterSpec `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	// the id of the table to which the meter is attached
	DirectTableId uint32 `protobuf:"varint,3,opt,name=direct_table_id,json=directTableId" json:"direct_table_id,omitempty"`
}

func (m *DirectMeter) Reset()                    { *m = DirectMeter{} }
func (m *DirectMeter) String() string            { return proto.CompactTextString(m) }
func (*DirectMeter) ProtoMessage()               {}
func (*DirectMeter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *DirectMeter) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *DirectMeter) GetSpec() *MeterSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *DirectMeter) GetDirectTableId() uint32 {
	if m != nil {
		return m.DirectTableId
	}
	return 0
}

// Any metadata associated with controller Packet-IO (Packet-In or Packet-Out)
// is modeled as P4 headers carrying special annotations
// @controller_metadata("packet_out") and @controller_metadata("packet_in")
// respectively. There can be at most one header each with these annotations.
// This message captures the info contained within these special headers,
// and used in p4runtime.proto to supply the metadata.
type ControllerPacketMetadata struct {
	// preamble.name and preamble.id will specify header type ("packet_out" or
	// "packet_in" for now).
	Preamble *Preamble `protobuf:"bytes,1,opt,name=preamble" json:"preamble,omitempty"`
	// Ordered based on header layout.
	// This is a constraint on the generator of this P4Info.
	Metadata []*ControllerPacketMetadata_Metadata `protobuf:"bytes,2,rep,name=metadata" json:"metadata,omitempty"`
}

func (m *ControllerPacketMetadata) Reset()                    { *m = ControllerPacketMetadata{} }
func (m *ControllerPacketMetadata) String() string            { return proto.CompactTextString(m) }
func (*ControllerPacketMetadata) ProtoMessage()               {}
func (*ControllerPacketMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ControllerPacketMetadata) GetPreamble() *Preamble {
	if m != nil {
		return m.Preamble
	}
	return nil
}

func (m *ControllerPacketMetadata) GetMetadata() []*ControllerPacketMetadata_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type ControllerPacketMetadata_Metadata struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// This is the name of the header field (not fully-qualified), similar
	// to e.g. Action.Param.name.
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Annotations []string `protobuf:"bytes,3,rep,name=annotations" json:"annotations,omitempty"`
	Bitwidth    int32    `protobuf:"varint,4,opt,name=bitwidth" json:"bitwidth,omitempty"`
}

func (m *ControllerPacketMetadata_Metadata) Reset()         { *m = ControllerPacketMetadata_Metadata{} }
func (m *ControllerPacketMetadata_Metadata) String() string { return proto.CompactTextString(m) }
func (*ControllerPacketMetadata_Metadata) ProtoMessage()    {}
func (*ControllerPacketMetadata_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{15, 0}
}

func (m *ControllerPacketMetadata_Metadata) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ControllerPacketMetadata_Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ControllerPacketMetadata_Metadata) GetAnnotations() []string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ControllerPacketMetadata_Metadata) GetBitwidth() int32 {
	if m != nil {
		return m.Bitwidth
	}
	return 0
}

func init() {
	proto.RegisterType((*P4Info)(nil), "p4.config.P4Info")
	proto.RegisterType((*Preamble)(nil), "p4.config.Preamble")
	proto.RegisterType((*Extern)(nil), "p4.config.Extern")
	proto.RegisterType((*ExternInstance)(nil), "p4.config.ExternInstance")
	proto.RegisterType((*MatchField)(nil), "p4.config.MatchField")
	proto.RegisterType((*Table)(nil), "p4.config.Table")
	proto.RegisterType((*ActionRef)(nil), "p4.config.ActionRef")
	proto.RegisterType((*Action)(nil), "p4.config.Action")
	proto.RegisterType((*Action_Param)(nil), "p4.config.Action.Param")
	proto.RegisterType((*ActionProfile)(nil), "p4.config.ActionProfile")
	proto.RegisterType((*CounterSpec)(nil), "p4.config.CounterSpec")
	proto.RegisterType((*Counter)(nil), "p4.config.Counter")
	proto.RegisterType((*DirectCounter)(nil), "p4.config.DirectCounter")
	proto.RegisterType((*MeterSpec)(nil), "p4.config.MeterSpec")
	proto.RegisterType((*Meter)(nil), "p4.config.Meter")
	proto.RegisterType((*DirectMeter)(nil), "p4.config.DirectMeter")
	proto.RegisterType((*ControllerPacketMetadata)(nil), "p4.config.ControllerPacketMetadata")
	proto.RegisterType((*ControllerPacketMetadata_Metadata)(nil), "p4.config.ControllerPacketMetadata.Metadata")
	proto.RegisterEnum("p4.config.MatchField_MatchType", MatchField_MatchType_name, MatchField_MatchType_value)
	proto.RegisterEnum("p4.config.CounterSpec_Unit", CounterSpec_Unit_name, CounterSpec_Unit_value)
	proto.RegisterEnum("p4.config.MeterSpec_Unit", MeterSpec_Unit_name, MeterSpec_Unit_value)
	proto.RegisterEnum("p4.config.MeterSpec_Type", MeterSpec_Type_name, MeterSpec_Type_value)
}

func init() { proto.RegisterFile("proto/p4/config/p4info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5f, 0x6f, 0xdb, 0x36,
	0x10, 0xaf, 0x6c, 0xd9, 0x91, 0xce, 0xb1, 0xab, 0xb0, 0x59, 0xab, 0xa6, 0x03, 0x66, 0xa8, 0xc3,
	0x66, 0xb4, 0x9d, 0x0c, 0x64, 0x19, 0xb6, 0x61, 0xd8, 0x00, 0x37, 0x71, 0x57, 0x63, 0x4d, 0x62,
	0xd0, 0xce, 0xb6, 0x3e, 0x09, 0x8c, 0x44, 0x27, 0x42, 0xad, 0x3f, 0x90, 0xe8, 0x75, 0xee, 0x6b,
	0x5f, 0xf7, 0xb0, 0x01, 0xc3, 0x3e, 0xd2, 0x3e, 0xc1, 0xf6, 0x7d, 0x06, 0x92, 0x92, 0x22, 0x47,
	0x69, 0x11, 0x0f, 0xc8, 0x1b, 0x79, 0xf7, 0xbb, 0xe3, 0xf1, 0x7e, 0xc7, 0xe3, 0xc1, 0x87, 0x71,
	0x12, 0xb1, 0xa8, 0x1f, 0xef, 0xf5, 0xdd, 0x28, 0x9c, 0xf9, 0x67, 0xfd, 0x78, 0xcf, 0x0f, 0x67,
	0x91, 0x2d, 0xc4, 0x48, 0x8f, 0xf7, 0x6c, 0x29, 0xdf, 0xb9, 0x7f, 0x16, 0x45, 0x67, 0x73, 0xda,
	0x17, 0x8a, 0xd3, 0xc5, 0xac, 0x4f, 0xc2, 0xa5, 0x44, 0x59, 0xbf, 0xa9, 0xd0, 0x1c, 0xef, 0x8d,
	0xc2, 0x59, 0x84, 0x1e, 0xc3, 0x06, 0xfd, 0x95, 0xd1, 0x24, 0x4c, 0x4d, 0xa5, 0x5b, 0xef, 0xb5,
	0x76, 0xb7, 0xec, 0xc2, 0x85, 0x3d, 0x14, 0x1a, 0x9c, 0x23, 0x50, 0x0f, 0x9a, 0x8c, 0x9c, 0xce,
	0x69, 0x6a, 0xd6, 0x04, 0xd6, 0x28, 0x61, 0xa7, 0x5c, 0x81, 0x33, 0x3d, 0x77, 0x4b, 0x5c, 0xe6,
	0x47, 0x61, 0x6a, 0xd6, 0x2b, 0x6e, 0x07, 0x42, 0x83, 0x73, 0x04, 0x1a, 0xc0, 0x6d, 0xb9, 0x74,
	0xe2, 0x24, 0x9a, 0xf9, 0xdc, 0xbf, 0x2a, 0x8c, 0xcc, 0x8a, 0xd1, 0x58, 0x02, 0x70, 0x87, 0x94,
	0xb7, 0x29, 0xb2, 0x41, 0x73, 0xa3, 0x45, 0xc8, 0x68, 0x92, 0x9a, 0x0d, 0x61, 0x8b, 0x4a, 0xb6,
	0xfb, 0x52, 0x85, 0x0b, 0x0c, 0x3f, 0xd2, 0xf3, 0x13, 0xea, 0x32, 0xa7, 0x30, 0x6b, 0x56, 0x8e,
	0x3c, 0x10, 0x88, 0xdc, 0xb8, 0xe3, 0x95, 0xb7, 0x22, 0x19, 0x01, 0x15, 0x96, 0x1b, 0x95, 0x64,
	0x1c, 0x72, 0x05, 0xce, 0xf4, 0xe8, 0x1b, 0x68, 0x67, 0x87, 0x65, 0x06, 0x9a, 0x30, 0xb8, 0x5b,
	0x39, 0x4a, 0x9a, 0x6d, 0x7a, 0x17, 0x9b, 0x14, 0x11, 0xd8, 0x71, 0xa3, 0x90, 0x25, 0xd1, 0x7c,
	0x4e, 0x13, 0x27, 0x26, 0xee, 0x2b, 0x2a, 0xfc, 0x10, 0x8f, 0x30, 0x62, 0xea, 0xc2, 0xd3, 0xc3,
	0x95, 0xbb, 0xe6, 0xe0, 0xb1, 0xc0, 0x1e, 0x66, 0x50, 0x6c, 0xba, 0xef, 0xd0, 0x58, 0x33, 0xd0,
	0xc6, 0x09, 0x25, 0xc1, 0xe9, 0x9c, 0xa2, 0x0e, 0xd4, 0x7c, 0xcf, 0x54, 0xba, 0x4a, 0xaf, 0x8d,
	0x6b, 0xbe, 0x87, 0x10, 0xa8, 0x21, 0x09, 0xa8, 0x59, 0xeb, 0x2a, 0x3d, 0x1d, 0x8b, 0x35, 0xda,
	0x86, 0x06, 0x99, 0xfb, 0x84, 0x53, 0xcb, 0x85, 0x72, 0x83, 0xba, 0xd0, 0x22, 0x61, 0x18, 0x31,
	0x22, 0x69, 0xe7, 0x0c, 0xea, 0xb8, 0x2c, 0xb2, 0xfe, 0x50, 0xa0, 0x29, 0x4b, 0x0a, 0x7d, 0x0c,
	0x1d, 0x59, 0x54, 0x0e, 0x5b, 0xc6, 0xd4, 0x29, 0x8e, 0xdc, 0x94, 0xd2, 0xe9, 0x32, 0xa6, 0x23,
	0x0f, 0xf5, 0xc0, 0x28, 0xa3, 0x4a, 0x81, 0x74, 0x2e, 0x70, 0x47, 0x3c, 0xa4, 0x2f, 0x41, 0xf7,
	0xc3, 0x94, 0x91, 0xd0, 0xa5, 0x79, 0xc5, 0xdd, 0xaf, 0x14, 0xf2, 0x28, 0x43, 0xe0, 0x0b, 0xac,
	0xf5, 0x0a, 0x3a, 0xab, 0x4a, 0xd4, 0x07, 0x2d, 0xce, 0xb2, 0x21, 0x82, 0x6a, 0xed, 0xde, 0x29,
	0x79, 0xca, 0x13, 0x85, 0x0b, 0x10, 0xea, 0x81, 0xca, 0x5f, 0xa0, 0x88, 0xac, 0xb5, 0xbb, 0x6d,
	0xcb, 0x77, 0x67, 0xe7, 0xef, 0xce, 0x1e, 0x84, 0x4b, 0x2c, 0x10, 0xd6, 0xdb, 0x1a, 0xc0, 0x21,
	0x61, 0xee, 0xf9, 0x33, 0x9f, 0xce, 0xbd, 0x6b, 0xe5, 0xfa, 0x52, 0x56, 0xeb, 0x95, 0xac, 0xa2,
	0x1d, 0xd0, 0x4e, 0x7d, 0xf6, 0xda, 0xf7, 0xd8, 0xb9, 0xa9, 0x76, 0x95, 0x5e, 0x03, 0x17, 0x7b,
	0xf4, 0x1d, 0x40, 0xc0, 0xcf, 0x13, 0xf9, 0x33, 0x1b, 0x5d, 0xa5, 0xd7, 0xd9, 0xfd, 0xa8, 0x5c,
	0xa7, 0x45, 0x30, 0x72, 0xc9, 0xf3, 0x89, 0xf5, 0x20, 0x5f, 0x5a, 0x13, 0xd0, 0x0b, 0x39, 0xba,
	0x0d, 0xad, 0x93, 0xa3, 0xc9, 0x78, 0xb8, 0x3f, 0x7a, 0x36, 0x1a, 0x1e, 0x18, 0xb7, 0x90, 0x0e,
	0x8d, 0x1f, 0x07, 0x2f, 0x46, 0x07, 0x86, 0xc2, 0x97, 0xc3, 0x9f, 0x07, 0xfb, 0x53, 0xa3, 0x86,
	0x36, 0xa0, 0xfe, 0x62, 0x7c, 0x68, 0xd4, 0x51, 0x0b, 0x36, 0xa6, 0x43, 0x7c, 0x34, 0xc0, 0x2f,
	0x0d, 0x95, 0x03, 0xf0, 0xe0, 0xe8, 0xfb, 0xa1, 0xd1, 0xb0, 0xfe, 0xa9, 0x43, 0x43, 0x74, 0x8b,
	0xf5, 0x53, 0xfd, 0x15, 0x6c, 0xca, 0xfb, 0xcc, 0x78, 0xcc, 0x79, 0x1b, 0xfa, 0xe0, 0xca, 0x1b,
	0xe1, 0x56, 0x50, 0xac, 0x53, 0xf4, 0x05, 0xb4, 0xb2, 0x1e, 0x93, 0xd0, 0x59, 0x5e, 0x22, 0xdb,
	0xd5, 0xa6, 0x44, 0x67, 0x18, 0x48, 0xbe, 0xe4, 0x66, 0xf7, 0xdc, 0x28, 0x4c, 0x99, 0xe3, 0xd1,
	0x19, 0x59, 0xcc, 0x99, 0x93, 0x39, 0xf1, 0x3d, 0x91, 0xeb, 0x36, 0xde, 0x16, 0xea, 0x03, 0xa9,
	0x95, 0x3e, 0x46, 0x1e, 0x9a, 0xc0, 0xa7, 0x57, 0x9a, 0x9d, 0x93, 0xd4, 0x09, 0x16, 0xa2, 0x47,
	0x3a, 0x31, 0x49, 0x48, 0x90, 0x0a, 0x52, 0x34, 0x6c, 0x55, 0xdd, 0x3c, 0x27, 0xe9, 0xa1, 0x84,
	0x8e, 0x05, 0x12, 0x3d, 0x86, 0x2d, 0x3f, 0x88, 0xe7, 0x34, 0xa0, 0xa1, 0xe4, 0x9e, 0x47, 0xd1,
	0x14, 0x51, 0x18, 0xab, 0x8a, 0x91, 0x87, 0x6c, 0xb8, 0x93, 0xf5, 0x9c, 0x84, 0xa6, 0xd1, 0x22,
	0x71, 0xf9, 0x23, 0x93, 0xad, 0xaa, 0x8d, 0xb7, 0xa4, 0x0a, 0x67, 0x9a, 0x91, 0x97, 0xf2, 0xda,
	0x4b, 0xfd, 0x37, 0xd4, 0xd4, 0xba, 0x4a, 0xaf, 0x8e, 0xc5, 0x1a, 0x3d, 0x01, 0xf4, 0xda, 0x67,
	0xe7, 0x0e, 0x0d, 0x59, 0xb2, 0x74, 0x98, 0x1f, 0xd0, 0x68, 0xc1, 0x4c, 0x5d, 0x04, 0x6c, 0x70,
	0xcd, 0x90, 0x2b, 0xa6, 0x52, 0x6e, 0x7d, 0x0b, 0x7a, 0x91, 0xc3, 0x4a, 0x69, 0x5f, 0x2a, 0xe3,
	0x5a, 0xb5, 0x39, 0xfc, 0xab, 0x40, 0x53, 0xda, 0xaf, 0x5f, 0x16, 0x7d, 0x68, 0x66, 0xd9, 0x94,
	0x05, 0x71, 0xaf, 0xc2, 0xab, 0x2d, 0x72, 0x88, 0x33, 0xd8, 0x8e, 0x0f, 0x0d, 0x21, 0xb8, 0xf9,
	0x27, 0x68, 0xfd, 0xa5, 0x40, 0x7b, 0xe5, 0xef, 0x5a, 0xff, 0x7a, 0x0f, 0x40, 0x97, 0x25, 0xe3,
	0x67, 0x25, 0xdf, 0xc6, 0x9a, 0x10, 0x70, 0xe2, 0x1e, 0x42, 0x5b, 0x90, 0x94, 0xd2, 0x39, 0x75,
	0x59, 0x94, 0x88, 0xa6, 0xac, 0xe1, 0x4d, 0x2e, 0x9c, 0x64, 0xb2, 0x82, 0x5d, 0xf5, 0x82, 0x5d,
	0x6b, 0x09, 0xad, 0xec, 0x2f, 0x9b, 0xc4, 0xd4, 0x45, 0x7d, 0x50, 0x17, 0xa1, 0xcf, 0x44, 0x44,
	0x9d, 0xdd, 0x07, 0xd5, 0xdf, 0x93, 0xa3, 0xec, 0x93, 0xd0, 0x67, 0x58, 0x00, 0xad, 0xaf, 0x41,
	0xe5, 0xbb, 0x2b, 0xdb, 0xc2, 0xd3, 0x97, 0xd3, 0xe1, 0xc4, 0x50, 0x78, 0x0b, 0x18, 0x0f, 0xf6,
	0x7f, 0x18, 0x4e, 0x27, 0x46, 0x0d, 0x69, 0xa0, 0x3e, 0x3d, 0x9e, 0x3e, 0x37, 0xea, 0xd6, 0x1b,
	0xd8, 0xc8, 0x9c, 0xae, 0x9f, 0x8c, 0x47, 0xa0, 0xa6, 0x31, 0x75, 0xb3, 0x6e, 0x7b, 0xf7, 0xea,
	0x38, 0xb1, 0xc0, 0x14, 0xd7, 0xae, 0x97, 0xae, 0xfd, 0xa7, 0x02, 0xed, 0x95, 0x8f, 0xfd, 0x66,
	0x43, 0xf8, 0xa4, 0x18, 0x34, 0x72, 0x0a, 0x45, 0x34, 0x6d, 0x9c, 0x8d, 0x04, 0x53, 0xc9, 0xa3,
	0xf5, 0xb7, 0x02, 0xba, 0xf8, 0xf1, 0x05, 0x19, 0x9f, 0xad, 0x90, 0x71, 0xff, 0xf2, 0x64, 0x71,
	0x89, 0x0a, 0x0e, 0x17, 0x0d, 0xbe, 0xf6, 0x1e, 0xb8, 0x68, 0xed, 0x02, 0x66, 0xf5, 0xd7, 0x64,
	0xce, 0x7a, 0x04, 0xaa, 0xf8, 0x01, 0xb6, 0xa0, 0xbd, 0x7f, 0xfc, 0xe2, 0x18, 0x3b, 0x27, 0x47,
	0x83, 0x9f, 0x06, 0x78, 0x68, 0xdc, 0xe2, 0x3e, 0xa4, 0x48, 0x0a, 0x14, 0xeb, 0x17, 0x68, 0x88,
	0x43, 0xff, 0xd7, 0x3f, 0x5a, 0x4a, 0xeb, 0xf6, 0x55, 0xb7, 0x78, 0x0f, 0xaf, 0xbf, 0x2b, 0xd0,
	0x2a, 0x4d, 0x51, 0x37, 0x79, 0xfc, 0x75, 0x39, 0x7d, 0x5b, 0x03, 0xf3, 0x5d, 0xe3, 0xd8, 0xfa,
	0xf1, 0x3d, 0x07, 0xad, 0x18, 0xfb, 0x64, 0x9b, 0x7b, 0x72, 0x8d, 0xb1, 0xcf, 0x2e, 0xe6, 0xbf,
	0xc2, 0x7a, 0x67, 0x0e, 0x5a, 0x11, 0xc6, 0x8d, 0x37, 0xc0, 0xd3, 0xa6, 0x18, 0x84, 0x3e, 0xff,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x53, 0x77, 0xb7, 0x99, 0xb9, 0x0c, 0x00, 0x00,
}
