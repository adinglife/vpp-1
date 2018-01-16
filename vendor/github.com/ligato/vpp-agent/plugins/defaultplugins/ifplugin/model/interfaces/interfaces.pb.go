// Code generated by protoc-gen-gogo.
// source: interfaces.proto
// DO NOT EDIT!

/*
Package interfaces is a generated protocol buffer package.

It is generated from these files:
	interfaces.proto

It has these top-level messages:
	Interfaces
	InterfacesState
	InterfaceErrors
*/
package interfaces

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type InterfaceType int32

const (
	InterfaceType_SOFTWARE_LOOPBACK   InterfaceType = 0
	InterfaceType_ETHERNET_CSMACD     InterfaceType = 1
	InterfaceType_MEMORY_INTERFACE    InterfaceType = 2
	InterfaceType_TAP_INTERFACE       InterfaceType = 3
	InterfaceType_AF_PACKET_INTERFACE InterfaceType = 4
	InterfaceType_VXLAN_TUNNEL        InterfaceType = 5
)

var InterfaceType_name = map[int32]string{
	0: "SOFTWARE_LOOPBACK",
	1: "ETHERNET_CSMACD",
	2: "MEMORY_INTERFACE",
	3: "TAP_INTERFACE",
	4: "AF_PACKET_INTERFACE",
	5: "VXLAN_TUNNEL",
}
var InterfaceType_value = map[string]int32{
	"SOFTWARE_LOOPBACK":   0,
	"ETHERNET_CSMACD":     1,
	"MEMORY_INTERFACE":    2,
	"TAP_INTERFACE":       3,
	"AF_PACKET_INTERFACE": 4,
	"VXLAN_TUNNEL":        5,
}

func (x InterfaceType) String() string {
	return proto.EnumName(InterfaceType_name, int32(x))
}

// from vpp/build-root/install-vpp-native/vpp/include/vnet/interface.h
type RxModeType int32

const (
	RxModeType_UNKNOWN   RxModeType = 0
	RxModeType_POLLING   RxModeType = 1
	RxModeType_INTERRUPT RxModeType = 2
	RxModeType_ADAPTIVE  RxModeType = 3
	RxModeType_DEFAULT   RxModeType = 4
)

var RxModeType_name = map[int32]string{
	0: "UNKNOWN",
	1: "POLLING",
	2: "INTERRUPT",
	3: "ADAPTIVE",
	4: "DEFAULT",
}
var RxModeType_value = map[string]int32{
	"UNKNOWN":   0,
	"POLLING":   1,
	"INTERRUPT": 2,
	"ADAPTIVE":  3,
	"DEFAULT":   4,
}

func (x RxModeType) String() string {
	return proto.EnumName(RxModeType_name, int32(x))
}

type Interfaces_Interface_Memif_MemifMode int32

const (
	Interfaces_Interface_Memif_ETHERNET    Interfaces_Interface_Memif_MemifMode = 0
	Interfaces_Interface_Memif_IP          Interfaces_Interface_Memif_MemifMode = 1
	Interfaces_Interface_Memif_PUNT_INJECT Interfaces_Interface_Memif_MemifMode = 2
)

var Interfaces_Interface_Memif_MemifMode_name = map[int32]string{
	0: "ETHERNET",
	1: "IP",
	2: "PUNT_INJECT",
}
var Interfaces_Interface_Memif_MemifMode_value = map[string]int32{
	"ETHERNET":    0,
	"IP":          1,
	"PUNT_INJECT": 2,
}

func (x Interfaces_Interface_Memif_MemifMode) String() string {
	return proto.EnumName(Interfaces_Interface_Memif_MemifMode_name, int32(x))
}

type InterfacesState_Interface_Status int32

const (
	InterfacesState_Interface_UNKNOWN_STATUS InterfacesState_Interface_Status = 0
	InterfacesState_Interface_UP             InterfacesState_Interface_Status = 1
	InterfacesState_Interface_DOWN           InterfacesState_Interface_Status = 2
	InterfacesState_Interface_DELETED        InterfacesState_Interface_Status = 3
)

var InterfacesState_Interface_Status_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "UP",
	2: "DOWN",
	3: "DELETED",
}
var InterfacesState_Interface_Status_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"UP":             1,
	"DOWN":           2,
	"DELETED":        3,
}

func (x InterfacesState_Interface_Status) String() string {
	return proto.EnumName(InterfacesState_Interface_Status_name, int32(x))
}

type InterfacesState_Interface_Duplex int32

const (
	InterfacesState_Interface_UNKNOWN_DUPLEX InterfacesState_Interface_Duplex = 0
	InterfacesState_Interface_HALF           InterfacesState_Interface_Duplex = 1
	InterfacesState_Interface_FULL           InterfacesState_Interface_Duplex = 2
)

var InterfacesState_Interface_Duplex_name = map[int32]string{
	0: "UNKNOWN_DUPLEX",
	1: "HALF",
	2: "FULL",
}
var InterfacesState_Interface_Duplex_value = map[string]int32{
	"UNKNOWN_DUPLEX": 0,
	"HALF":           1,
	"FULL":           2,
}

func (x InterfacesState_Interface_Duplex) String() string {
	return proto.EnumName(InterfacesState_Interface_Duplex_name, int32(x))
}

type Interfaces struct {
	Interface []*Interfaces_Interface `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *Interfaces) Reset()         { *m = Interfaces{} }
func (m *Interfaces) String() string { return proto.CompactTextString(m) }
func (*Interfaces) ProtoMessage()    {}

func (m *Interfaces) GetInterface() []*Interfaces_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type Interfaces_Interface struct {
	Name               string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description        string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type               InterfaceType `protobuf:"varint,3,opt,name=type,proto3,enum=interfaces.InterfaceType" json:"type,omitempty"`
	Enabled            bool          `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	PhysAddress        string        `protobuf:"bytes,5,opt,name=phys_address,proto3" json:"phys_address,omitempty"`
	Mtu                uint32        `protobuf:"varint,6,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Vrf                uint32        `protobuf:"varint,7,opt,name=vrf,proto3" json:"vrf,omitempty"`
	ContainerIpAddress string        `protobuf:"bytes,8,opt,name=container_ip_address,proto3" json:"container_ip_address,omitempty"`
	// Required format is "ipAddress/ipPrefix"
	IpAddresses    []string                             `protobuf:"bytes,10,rep,name=ip_addresses" json:"ip_addresses,omitempty"`
	Unnumbered     *Interfaces_Interface_Unnumbered     `protobuf:"bytes,11,opt,name=unnumbered" json:"unnumbered,omitempty"`
	RxModeSettings *Interfaces_Interface_RxModeSettings `protobuf:"bytes,12,opt,name=rxModeSettings" json:"rxModeSettings,omitempty"`
	Memif          *Interfaces_Interface_Memif          `protobuf:"bytes,101,opt,name=memif" json:"memif,omitempty"`
	Vxlan          *Interfaces_Interface_Vxlan          `protobuf:"bytes,102,opt,name=vxlan" json:"vxlan,omitempty"`
	Afpacket       *Interfaces_Interface_Afpacket       `protobuf:"bytes,103,opt,name=afpacket" json:"afpacket,omitempty"`
	Tap            *Interfaces_Interface_Tap            `protobuf:"bytes,104,opt,name=tap" json:"tap,omitempty"`
}

func (m *Interfaces_Interface) Reset()         { *m = Interfaces_Interface{} }
func (m *Interfaces_Interface) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface) ProtoMessage()    {}

func (m *Interfaces_Interface) GetUnnumbered() *Interfaces_Interface_Unnumbered {
	if m != nil {
		return m.Unnumbered
	}
	return nil
}

func (m *Interfaces_Interface) GetRxModeSettings() *Interfaces_Interface_RxModeSettings {
	if m != nil {
		return m.RxModeSettings
	}
	return nil
}

func (m *Interfaces_Interface) GetMemif() *Interfaces_Interface_Memif {
	if m != nil {
		return m.Memif
	}
	return nil
}

func (m *Interfaces_Interface) GetVxlan() *Interfaces_Interface_Vxlan {
	if m != nil {
		return m.Vxlan
	}
	return nil
}

func (m *Interfaces_Interface) GetAfpacket() *Interfaces_Interface_Afpacket {
	if m != nil {
		return m.Afpacket
	}
	return nil
}

func (m *Interfaces_Interface) GetTap() *Interfaces_Interface_Tap {
	if m != nil {
		return m.Tap
	}
	return nil
}

type Interfaces_Interface_Unnumbered struct {
	IsUnnumbered    bool   `protobuf:"varint,1,opt,name=isUnnumbered,proto3" json:"isUnnumbered,omitempty"`
	InterfaceWithIP string `protobuf:"bytes,2,opt,name=interfaceWithIP,proto3" json:"interfaceWithIP,omitempty"`
}

func (m *Interfaces_Interface_Unnumbered) Reset()         { *m = Interfaces_Interface_Unnumbered{} }
func (m *Interfaces_Interface_Unnumbered) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Unnumbered) ProtoMessage()    {}

type Interfaces_Interface_RxModeSettings struct {
	RxMode       RxModeType `protobuf:"varint,1,opt,name=rx_mode,proto3,enum=interfaces.RxModeType" json:"rx_mode,omitempty"`
	QueueID      uint32     `protobuf:"varint,2,opt,proto3" json:"QueueID,omitempty"`
	QueueIDValid uint32     `protobuf:"varint,3,opt,proto3" json:"QueueIDValid,omitempty"`
}

func (m *Interfaces_Interface_RxModeSettings) Reset()         { *m = Interfaces_Interface_RxModeSettings{} }
func (m *Interfaces_Interface_RxModeSettings) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_RxModeSettings) ProtoMessage()    {}

type Interfaces_Interface_Memif struct {
	Master         bool                                 `protobuf:"varint,1,opt,name=master,proto3" json:"master,omitempty"`
	Mode           Interfaces_Interface_Memif_MemifMode `protobuf:"varint,2,opt,name=mode,proto3,enum=interfaces.Interfaces_Interface_Memif_MemifMode" json:"mode,omitempty"`
	Id             uint32                               `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	SocketFilename string                               `protobuf:"bytes,4,opt,name=socket_filename,proto3" json:"socket_filename,omitempty"`
	Secret         string                               `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	RingSize       uint32                               `protobuf:"varint,6,opt,name=ring_size,proto3" json:"ring_size,omitempty"`
	BufferSize     uint32                               `protobuf:"varint,7,opt,name=buffer_size,proto3" json:"buffer_size,omitempty"`
	RxQueues       uint32                               `protobuf:"varint,8,opt,name=rx_queues,proto3" json:"rx_queues,omitempty"`
	TxQueues       uint32                               `protobuf:"varint,9,opt,name=tx_queues,proto3" json:"tx_queues,omitempty"`
}

func (m *Interfaces_Interface_Memif) Reset()         { *m = Interfaces_Interface_Memif{} }
func (m *Interfaces_Interface_Memif) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Memif) ProtoMessage()    {}

type Interfaces_Interface_Vxlan struct {
	SrcAddress string `protobuf:"bytes,1,opt,name=src_address,proto3" json:"src_address,omitempty"`
	DstAddress string `protobuf:"bytes,2,opt,name=dst_address,proto3" json:"dst_address,omitempty"`
	Vni        uint32 `protobuf:"varint,3,opt,name=vni,proto3" json:"vni,omitempty"`
}

func (m *Interfaces_Interface_Vxlan) Reset()         { *m = Interfaces_Interface_Vxlan{} }
func (m *Interfaces_Interface_Vxlan) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Vxlan) ProtoMessage()    {}

type Interfaces_Interface_Afpacket struct {
	HostIfName string `protobuf:"bytes,1,opt,name=host_if_name,proto3" json:"host_if_name,omitempty"`
}

func (m *Interfaces_Interface_Afpacket) Reset()         { *m = Interfaces_Interface_Afpacket{} }
func (m *Interfaces_Interface_Afpacket) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Afpacket) ProtoMessage()    {}

type Interfaces_Interface_Tap struct {
	Version    uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	HostIfName string `protobuf:"bytes,2,opt,name=host_if_name,proto3" json:"host_if_name,omitempty"`
	Namespace  string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	RxRingSize uint32 `protobuf:"varint,4,opt,name=rx_ring_size,proto3" json:"rx_ring_size,omitempty"`
	TxRingSize uint32 `protobuf:"varint,5,opt,name=tx_ring_size,proto3" json:"tx_ring_size,omitempty"`
}

func (m *Interfaces_Interface_Tap) Reset()         { *m = Interfaces_Interface_Tap{} }
func (m *Interfaces_Interface_Tap) String() string { return proto.CompactTextString(m) }
func (*Interfaces_Interface_Tap) ProtoMessage()    {}

type InterfacesState struct {
	Interface []*InterfacesState_Interface `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *InterfacesState) Reset()         { *m = InterfacesState{} }
func (m *InterfacesState) String() string { return proto.CompactTextString(m) }
func (*InterfacesState) ProtoMessage()    {}

func (m *InterfacesState) GetInterface() []*InterfacesState_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type InterfacesState_Interface struct {
	Name         string                                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InternalName string                                `protobuf:"bytes,2,opt,name=internal_name,proto3" json:"internal_name,omitempty"`
	Type         InterfaceType                         `protobuf:"varint,3,opt,name=type,proto3,enum=interfaces.InterfaceType" json:"type,omitempty"`
	IfIndex      uint32                                `protobuf:"varint,4,opt,name=if_index,proto3" json:"if_index,omitempty"`
	AdminStatus  InterfacesState_Interface_Status      `protobuf:"varint,5,opt,name=admin_status,proto3,enum=interfaces.InterfacesState_Interface_Status" json:"admin_status,omitempty"`
	OperStatus   InterfacesState_Interface_Status      `protobuf:"varint,6,opt,name=oper_status,proto3,enum=interfaces.InterfacesState_Interface_Status" json:"oper_status,omitempty"`
	LastChange   int64                                 `protobuf:"varint,7,opt,name=last_change,proto3" json:"last_change,omitempty"`
	PhysAddress  string                                `protobuf:"bytes,8,opt,name=phys_address,proto3" json:"phys_address,omitempty"`
	Speed        uint64                                `protobuf:"varint,9,opt,name=speed,proto3" json:"speed,omitempty"`
	Mtu          uint32                                `protobuf:"varint,10,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Duplex       InterfacesState_Interface_Duplex      `protobuf:"varint,11,opt,name=duplex,proto3,enum=interfaces.InterfacesState_Interface_Duplex" json:"duplex,omitempty"`
	Statistics   *InterfacesState_Interface_Statistics `protobuf:"bytes,100,opt,name=statistics" json:"statistics,omitempty"`
}

func (m *InterfacesState_Interface) Reset()         { *m = InterfacesState_Interface{} }
func (m *InterfacesState_Interface) String() string { return proto.CompactTextString(m) }
func (*InterfacesState_Interface) ProtoMessage()    {}

func (m *InterfacesState_Interface) GetStatistics() *InterfacesState_Interface_Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type InterfacesState_Interface_Statistics struct {
	InPackets       uint64 `protobuf:"varint,1,opt,name=in_packets,proto3" json:"in_packets,omitempty"`
	InBytes         uint64 `protobuf:"varint,2,opt,name=in_bytes,proto3" json:"in_bytes,omitempty"`
	OutPackets      uint64 `protobuf:"varint,3,opt,name=out_packets,proto3" json:"out_packets,omitempty"`
	OutBytes        uint64 `protobuf:"varint,4,opt,name=out_bytes,proto3" json:"out_bytes,omitempty"`
	DropPackets     uint64 `protobuf:"varint,5,opt,name=drop_packets,proto3" json:"drop_packets,omitempty"`
	PuntPackets     uint64 `protobuf:"varint,6,opt,name=punt_packets,proto3" json:"punt_packets,omitempty"`
	Ipv4Packets     uint64 `protobuf:"varint,7,opt,name=ipv4_packets,proto3" json:"ipv4_packets,omitempty"`
	Ipv6Packets     uint64 `protobuf:"varint,8,opt,name=ipv6_packets,proto3" json:"ipv6_packets,omitempty"`
	InNobufPackets  uint64 `protobuf:"varint,9,opt,name=in_nobuf_packets,proto3" json:"in_nobuf_packets,omitempty"`
	InMissPackets   uint64 `protobuf:"varint,10,opt,name=in_miss_packets,proto3" json:"in_miss_packets,omitempty"`
	InErrorPackets  uint64 `protobuf:"varint,11,opt,name=in_error_packets,proto3" json:"in_error_packets,omitempty"`
	OutErrorPackets uint64 `protobuf:"varint,12,opt,name=out_error_packets,proto3" json:"out_error_packets,omitempty"`
}

func (m *InterfacesState_Interface_Statistics) Reset()         { *m = InterfacesState_Interface_Statistics{} }
func (m *InterfacesState_Interface_Statistics) String() string { return proto.CompactTextString(m) }
func (*InterfacesState_Interface_Statistics) ProtoMessage()    {}

type InterfaceErrors struct {
	Interface []*InterfaceErrors_Interface `protobuf:"bytes,1,rep,name=interface" json:"interface,omitempty"`
}

func (m *InterfaceErrors) Reset()         { *m = InterfaceErrors{} }
func (m *InterfaceErrors) String() string { return proto.CompactTextString(m) }
func (*InterfaceErrors) ProtoMessage()    {}

func (m *InterfaceErrors) GetInterface() []*InterfaceErrors_Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type InterfaceErrors_Interface struct {
	InterfaceName string                                 `protobuf:"bytes,1,opt,name=interface_name,proto3" json:"interface_name,omitempty"`
	ErrorData     []*InterfaceErrors_Interface_ErrorData `protobuf:"bytes,2,rep,name=error_data" json:"error_data,omitempty"`
}

func (m *InterfaceErrors_Interface) Reset()         { *m = InterfaceErrors_Interface{} }
func (m *InterfaceErrors_Interface) String() string { return proto.CompactTextString(m) }
func (*InterfaceErrors_Interface) ProtoMessage()    {}

func (m *InterfaceErrors_Interface) GetErrorData() []*InterfaceErrors_Interface_ErrorData {
	if m != nil {
		return m.ErrorData
	}
	return nil
}

type InterfaceErrors_Interface_ErrorData struct {
	ChangeType   string `protobuf:"bytes,1,opt,name=change_type,proto3" json:"change_type,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,proto3" json:"error_message,omitempty"`
	LastChange   int64  `protobuf:"varint,3,opt,name=last_change,proto3" json:"last_change,omitempty"`
}

func (m *InterfaceErrors_Interface_ErrorData) Reset()         { *m = InterfaceErrors_Interface_ErrorData{} }
func (m *InterfaceErrors_Interface_ErrorData) String() string { return proto.CompactTextString(m) }
func (*InterfaceErrors_Interface_ErrorData) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("interfaces.InterfaceType", InterfaceType_name, InterfaceType_value)
	proto.RegisterEnum("interfaces.RxModeType", RxModeType_name, RxModeType_value)
	proto.RegisterEnum("interfaces.Interfaces_Interface_Memif_MemifMode", Interfaces_Interface_Memif_MemifMode_name, Interfaces_Interface_Memif_MemifMode_value)
	proto.RegisterEnum("interfaces.InterfacesState_Interface_Status", InterfacesState_Interface_Status_name, InterfacesState_Interface_Status_value)
	proto.RegisterEnum("interfaces.InterfacesState_Interface_Duplex", InterfacesState_Interface_Duplex_name, InterfacesState_Interface_Duplex_value)
}
