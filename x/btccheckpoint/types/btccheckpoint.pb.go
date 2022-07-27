// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btccheckpoint/btccheckpoint.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EpochStatus int32

const (
	// SUBMITTED Epoch is in submitted state when at least on checkpoint from it
	// is included in bitcoin
	Submitted EpochStatus = 0
	// CONFIRMED Epoch is in confirmed state when at least on checkpoint from it
	// is included in bitcoin and at least k-deep on main chain
	Confirmed EpochStatus = 1
	// FINALIZED Epoch is in confirmed state when at least on checkpoint from it
	// is included in bitcoin and at least w-deep on main chain
	Finalized EpochStatus = 2
	// SIGNED Epoch does not have any submissions or all of its submission are not on
	// main chain
	Signed EpochStatus = 3
)

var EpochStatus_name = map[int32]string{
	0: "EPOCH_STATUS_SUBMITTED",
	1: "EPOCH_STATUS_CONFIRMED",
	2: "EPOCH_STATUS_FINALIZED",
	3: "EPOCH_STATUS_SIGNED",
}

var EpochStatus_value = map[string]int32{
	"EPOCH_STATUS_SUBMITTED": 0,
	"EPOCH_STATUS_CONFIRMED": 1,
	"EPOCH_STATUS_FINALIZED": 2,
	"EPOCH_STATUS_SIGNED":    3,
}

func (x EpochStatus) String() string {
	return proto.EnumName(EpochStatus_name, int32(x))
}

func (EpochStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_da8b9af3dbd18a36, []int{0}
}

// Each provided OP_RETURN transaction can be idendtified by hash of block in
// which transaction was included and transaction index in the block
type TransactionKey struct {
	Index uint32                                                    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Hash  *github_com_babylonchain_babylon_types.BTCHeaderHashBytes `protobuf:"bytes,2,opt,name=hash,proto3,customtype=github.com/babylonchain/babylon/types.BTCHeaderHashBytes" json:"hash,omitempty"`
}

func (m *TransactionKey) Reset()         { *m = TransactionKey{} }
func (m *TransactionKey) String() string { return proto.CompactTextString(m) }
func (*TransactionKey) ProtoMessage()    {}
func (*TransactionKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_da8b9af3dbd18a36, []int{0}
}
func (m *TransactionKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionKey.Merge(m, src)
}
func (m *TransactionKey) XXX_Size() int {
	return m.Size()
}
func (m *TransactionKey) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionKey.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionKey proto.InternalMessageInfo

func (m *TransactionKey) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

// Checkpoint can be composed from multiple transactions, so to identify whole
// submission we need list of transaction keys.
// Each submission can generally be identified by this list of (txIdx, blockHash)
// tuples.
// Note: this could possibly be optimized as if transactions were in one block
// they would have the same block hash and different indexes, but each blockhash
// is only 33 (1  byte for prefix encoding and 32 byte hash), so there should
// be other strong arguments for this optimization
type SubmissionKey struct {
	Key []*TransactionKey `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
}

func (m *SubmissionKey) Reset()         { *m = SubmissionKey{} }
func (m *SubmissionKey) String() string { return proto.CompactTextString(m) }
func (*SubmissionKey) ProtoMessage()    {}
func (*SubmissionKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_da8b9af3dbd18a36, []int{1}
}
func (m *SubmissionKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmissionKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmissionKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmissionKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmissionKey.Merge(m, src)
}
func (m *SubmissionKey) XXX_Size() int {
	return m.Size()
}
func (m *SubmissionKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmissionKey.DiscardUnknown(m)
}

var xxx_messageInfo_SubmissionKey proto.InternalMessageInfo

func (m *SubmissionKey) GetKey() []*TransactionKey {
	if m != nil {
		return m.Key
	}
	return nil
}

// TODO: Determine if we should keep any block number or depth info.
// On one hand it may be usefull to determine if block is stable or not, on other
// depth/block number info, without context (i.e info about chain) is pretty useless
// and blockshash in enough to retrieve is from lightclient
type SubmissionData struct {
	// TODO: this could probably be better typed
	// Address of submitter of given checkpoint. Required to payup the reward to
	// submitter of given checkpoint
	Submitter []byte `protobuf:"bytes,1,opt,name=submitter,proto3" json:"submitter,omitempty"`
	// Required to recover address of sender of btc transction to payup the reward.
	// TODO: Maybe it is worth recovering senders while processing the InsertProof
	// message, and store only those. Another point is that it is not that simple
	// to recover sender of btc tx.
	Btctransaction [][]byte `protobuf:"bytes,2,rep,name=btctransaction,proto3" json:"btctransaction,omitempty"`
	Epoch          uint64   `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
}

func (m *SubmissionData) Reset()         { *m = SubmissionData{} }
func (m *SubmissionData) String() string { return proto.CompactTextString(m) }
func (*SubmissionData) ProtoMessage()    {}
func (*SubmissionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_da8b9af3dbd18a36, []int{2}
}
func (m *SubmissionData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubmissionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubmissionData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubmissionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmissionData.Merge(m, src)
}
func (m *SubmissionData) XXX_Size() int {
	return m.Size()
}
func (m *SubmissionData) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmissionData.DiscardUnknown(m)
}

var xxx_messageInfo_SubmissionData proto.InternalMessageInfo

func (m *SubmissionData) GetSubmitter() []byte {
	if m != nil {
		return m.Submitter
	}
	return nil
}

func (m *SubmissionData) GetBtctransaction() [][]byte {
	if m != nil {
		return m.Btctransaction
	}
	return nil
}

func (m *SubmissionData) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

// Data stored in db and indexed by epoch number
// TODO: Add btc blockheight at epooch end, when adding hadnling of epoching callbacks
type EpochData struct {
	// List of all received checkpoints during this epoch, sorted by order of
	// submission.
	Key []*SubmissionKey `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	// Current state of epoch.
	Status EpochStatus `protobuf:"varint,2,opt,name=status,proto3,enum=babylon.btccheckpoint.EpochStatus" json:"status,omitempty"`
	// Required to comunicate with checkpoint module about checkpoint status
	RawCheckpoint []byte `protobuf:"bytes,3,opt,name=raw_checkpoint,json=rawCheckpoint,proto3" json:"raw_checkpoint,omitempty"`
}

func (m *EpochData) Reset()         { *m = EpochData{} }
func (m *EpochData) String() string { return proto.CompactTextString(m) }
func (*EpochData) ProtoMessage()    {}
func (*EpochData) Descriptor() ([]byte, []int) {
	return fileDescriptor_da8b9af3dbd18a36, []int{3}
}
func (m *EpochData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochData.Merge(m, src)
}
func (m *EpochData) XXX_Size() int {
	return m.Size()
}
func (m *EpochData) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochData.DiscardUnknown(m)
}

var xxx_messageInfo_EpochData proto.InternalMessageInfo

func (m *EpochData) GetKey() []*SubmissionKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *EpochData) GetStatus() EpochStatus {
	if m != nil {
		return m.Status
	}
	return Submitted
}

func (m *EpochData) GetRawCheckpoint() []byte {
	if m != nil {
		return m.RawCheckpoint
	}
	return nil
}

func init() {
	proto.RegisterEnum("babylon.btccheckpoint.EpochStatus", EpochStatus_name, EpochStatus_value)
	proto.RegisterType((*TransactionKey)(nil), "babylon.btccheckpoint.TransactionKey")
	proto.RegisterType((*SubmissionKey)(nil), "babylon.btccheckpoint.SubmissionKey")
	proto.RegisterType((*SubmissionData)(nil), "babylon.btccheckpoint.SubmissionData")
	proto.RegisterType((*EpochData)(nil), "babylon.btccheckpoint.EpochData")
}

func init() {
	proto.RegisterFile("babylon/btccheckpoint/btccheckpoint.proto", fileDescriptor_da8b9af3dbd18a36)
}

var fileDescriptor_da8b9af3dbd18a36 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x49, 0x88, 0x94, 0x6d, 0x12, 0x45, 0xa6, 0xa0, 0x28, 0x42, 0xc6, 0x0a, 0x14,
	0xa5, 0x1c, 0x1c, 0xa9, 0x88, 0x3f, 0x42, 0x5c, 0x9a, 0xd8, 0x21, 0x16, 0x34, 0xa9, 0x6c, 0xf7,
	0xd2, 0x4b, 0xb4, 0xb6, 0x97, 0x78, 0xd5, 0xc4, 0x1b, 0x79, 0x37, 0x6a, 0xcc, 0x13, 0xa0, 0x9e,
	0x78, 0x81, 0x9e, 0x90, 0x78, 0x07, 0xde, 0x80, 0x63, 0x8f, 0x88, 0x03, 0x42, 0xc9, 0x8b, 0x20,
	0xaf, 0x0d, 0xa9, 0x4b, 0xab, 0xde, 0x76, 0xc6, 0xbf, 0x99, 0xef, 0x9b, 0xd1, 0x18, 0xee, 0x3a,
	0xc8, 0x89, 0xa6, 0x34, 0xe8, 0x38, 0xdc, 0x75, 0x7d, 0xec, 0x9e, 0xcc, 0x29, 0x09, 0x78, 0x36,
	0x52, 0xe7, 0x21, 0xe5, 0x54, 0xba, 0x97, 0xa2, 0x6a, 0xe6, 0x63, 0x73, 0x7b, 0x42, 0x27, 0x54,
	0x10, 0x9d, 0xf8, 0x95, 0xc0, 0xad, 0x25, 0xac, 0xd9, 0x21, 0x0a, 0x18, 0x72, 0x39, 0xa1, 0xc1,
	0x3b, 0x1c, 0x49, 0xdb, 0xf0, 0x0e, 0x09, 0x3c, 0xbc, 0x6c, 0x00, 0x05, 0xb4, 0xab, 0x66, 0x12,
	0x48, 0x87, 0xb0, 0xe8, 0x23, 0xe6, 0x37, 0xf2, 0x0a, 0x68, 0x57, 0xba, 0x6f, 0x7e, 0xfe, 0x7a,
	0xf8, 0x6a, 0x42, 0xb8, 0xbf, 0x70, 0x54, 0x97, 0xce, 0x3a, 0xa9, 0xa2, 0xeb, 0x23, 0x12, 0xfc,
	0x0d, 0x3a, 0x3c, 0x9a, 0x63, 0xa6, 0x76, 0xed, 0xde, 0x00, 0x23, 0x0f, 0x87, 0x03, 0xc4, 0xfc,
	0x6e, 0xc4, 0x31, 0x33, 0x45, 0xa7, 0xd6, 0x00, 0x56, 0xad, 0x85, 0x33, 0x23, 0x8c, 0xa5, 0xc2,
	0x2f, 0x61, 0xe1, 0x04, 0x47, 0x0d, 0xa0, 0x14, 0xda, 0x5b, 0x7b, 0x3b, 0xea, 0xb5, 0x53, 0xa8,
	0x59, 0xb3, 0x66, 0x5c, 0xd1, 0x9a, 0xc2, 0xda, 0xa6, 0x93, 0x86, 0x38, 0x92, 0x1e, 0xc0, 0x32,
	0x8b, 0x33, 0x9c, 0xe3, 0x50, 0xcc, 0x51, 0x31, 0x37, 0x09, 0xe9, 0x09, 0xac, 0x39, 0xdc, 0xe5,
	0x9b, 0x4e, 0x8d, 0xbc, 0x52, 0x68, 0x57, 0xcc, 0x2b, 0xd9, 0x78, 0x13, 0x78, 0x4e, 0x5d, 0xbf,
	0x51, 0x50, 0x40, 0xbb, 0x68, 0x26, 0x41, 0xeb, 0x2b, 0x80, 0x65, 0x3d, 0x7e, 0x09, 0xa5, 0x17,
	0x97, 0x4d, 0x3f, 0xbe, 0xc1, 0x74, 0x66, 0x4e, 0xe1, 0x59, 0x7a, 0x0d, 0x4b, 0x8c, 0x23, 0xbe,
	0x60, 0x62, 0xa3, 0xb5, 0xbd, 0xd6, 0x0d, 0xa5, 0x42, 0xc9, 0x12, 0xa4, 0x99, 0x56, 0x48, 0x3b,
	0xb0, 0x16, 0xa2, 0xd3, 0xf1, 0x86, 0x12, 0x06, 0x2b, 0x66, 0x35, 0x44, 0xa7, 0xbd, 0x7f, 0xc9,
	0xa7, 0xdf, 0x00, 0xdc, 0xba, 0x54, 0x2e, 0xed, 0xc2, 0xfb, 0xfa, 0xe1, 0xa8, 0x37, 0x18, 0x5b,
	0xf6, 0xbe, 0x7d, 0x64, 0x8d, 0xad, 0xa3, 0xee, 0x81, 0x61, 0xdb, 0xba, 0x56, 0xcf, 0x35, 0xab,
	0x67, 0xe7, 0x4a, 0xd9, 0x4a, 0x37, 0xe4, 0xfd, 0x87, 0xf6, 0x46, 0xc3, 0xbe, 0x61, 0x1e, 0xe8,
	0x5a, 0x1d, 0x24, 0x68, 0x8f, 0x06, 0x1f, 0x48, 0x38, 0xbb, 0x06, 0xed, 0x1b, 0xc3, 0xfd, 0xf7,
	0xc6, 0xb1, 0xae, 0xd5, 0xf3, 0x09, 0xda, 0x27, 0x01, 0x9a, 0x92, 0x8f, 0xd8, 0x93, 0x1e, 0xc1,
	0xbb, 0x59, 0x03, 0xc6, 0xdb, 0xa1, 0xae, 0xd5, 0x0b, 0x4d, 0x78, 0x76, 0xae, 0x94, 0x2c, 0x32,
	0x09, 0xb0, 0xd7, 0x2c, 0x7e, 0xfa, 0x22, 0xe7, 0xba, 0xa3, 0xef, 0x2b, 0x19, 0x5c, 0xac, 0x64,
	0xf0, 0x7b, 0x25, 0x83, 0xcf, 0x6b, 0x39, 0x77, 0xb1, 0x96, 0x73, 0x3f, 0xd6, 0x72, 0xee, 0xf8,
	0xf9, 0x6d, 0x67, 0xb7, 0xbc, 0xf2, 0x8b, 0x88, 0x33, 0x74, 0x4a, 0xe2, 0xdc, 0x9f, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0x7a, 0xe8, 0xcf, 0x48, 0x03, 0x00, 0x00,
}

func (m *TransactionKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Hash != nil {
		{
			size := m.Hash.Size()
			i -= size
			if _, err := m.Hash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBtccheckpoint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintBtccheckpoint(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SubmissionKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmissionKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmissionKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		for iNdEx := len(m.Key) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Key[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBtccheckpoint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SubmissionData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubmissionData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubmissionData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Epoch != 0 {
		i = encodeVarintBtccheckpoint(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Btctransaction) > 0 {
		for iNdEx := len(m.Btctransaction) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Btctransaction[iNdEx])
			copy(dAtA[i:], m.Btctransaction[iNdEx])
			i = encodeVarintBtccheckpoint(dAtA, i, uint64(len(m.Btctransaction[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Submitter) > 0 {
		i -= len(m.Submitter)
		copy(dAtA[i:], m.Submitter)
		i = encodeVarintBtccheckpoint(dAtA, i, uint64(len(m.Submitter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EpochData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RawCheckpoint) > 0 {
		i -= len(m.RawCheckpoint)
		copy(dAtA[i:], m.RawCheckpoint)
		i = encodeVarintBtccheckpoint(dAtA, i, uint64(len(m.RawCheckpoint)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Status != 0 {
		i = encodeVarintBtccheckpoint(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		for iNdEx := len(m.Key) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Key[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBtccheckpoint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBtccheckpoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovBtccheckpoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransactionKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovBtccheckpoint(uint64(m.Index))
	}
	if m.Hash != nil {
		l = m.Hash.Size()
		n += 1 + l + sovBtccheckpoint(uint64(l))
	}
	return n
}

func (m *SubmissionKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Key) > 0 {
		for _, e := range m.Key {
			l = e.Size()
			n += 1 + l + sovBtccheckpoint(uint64(l))
		}
	}
	return n
}

func (m *SubmissionData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Submitter)
	if l > 0 {
		n += 1 + l + sovBtccheckpoint(uint64(l))
	}
	if len(m.Btctransaction) > 0 {
		for _, b := range m.Btctransaction {
			l = len(b)
			n += 1 + l + sovBtccheckpoint(uint64(l))
		}
	}
	if m.Epoch != 0 {
		n += 1 + sovBtccheckpoint(uint64(m.Epoch))
	}
	return n
}

func (m *EpochData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Key) > 0 {
		for _, e := range m.Key {
			l = e.Size()
			n += 1 + l + sovBtccheckpoint(uint64(l))
		}
	}
	if m.Status != 0 {
		n += 1 + sovBtccheckpoint(uint64(m.Status))
	}
	l = len(m.RawCheckpoint)
	if l > 0 {
		n += 1 + l + sovBtccheckpoint(uint64(l))
	}
	return n
}

func sovBtccheckpoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBtccheckpoint(x uint64) (n int) {
	return sovBtccheckpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransactionKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBtccheckpoint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TransactionKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_types.BTCHeaderHashBytes
			m.Hash = &v
			if err := m.Hash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBtccheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SubmissionKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBtccheckpoint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubmissionKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmissionKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key, &TransactionKey{})
			if err := m.Key[len(m.Key)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBtccheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SubmissionData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBtccheckpoint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SubmissionData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubmissionData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submitter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Submitter = append(m.Submitter[:0], dAtA[iNdEx:postIndex]...)
			if m.Submitter == nil {
				m.Submitter = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Btctransaction", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Btctransaction = append(m.Btctransaction, make([]byte, postIndex-iNdEx))
			copy(m.Btctransaction[len(m.Btctransaction)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBtccheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EpochData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBtccheckpoint
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EpochData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key, &SubmissionKey{})
			if err := m.Key[len(m.Key)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= EpochStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RawCheckpoint", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RawCheckpoint = append(m.RawCheckpoint[:0], dAtA[iNdEx:postIndex]...)
			if m.RawCheckpoint == nil {
				m.RawCheckpoint = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBtccheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBtccheckpoint
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBtccheckpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBtccheckpoint
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBtccheckpoint
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBtccheckpoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBtccheckpoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBtccheckpoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBtccheckpoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBtccheckpoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBtccheckpoint = fmt.Errorf("proto: unexpected end of group")
)
