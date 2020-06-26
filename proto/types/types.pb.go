// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/types/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	merkle "github.com/tendermint/tendermint/proto/crypto/merkle"
	bits "github.com/tendermint/tendermint/proto/libs/bits"
	version "github.com/tendermint/tendermint/proto/version"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BlockIdFlag indicates which BlcokID the signature is for
type BlockIDFlag int32

const (
	BLOCKD_ID_FLAG_UNKNOWN BlockIDFlag = 0
	BlockIDFlagAbsent      BlockIDFlag = 1
	BlockIDFlagCommit      BlockIDFlag = 2
	BlockIDFlagNil         BlockIDFlag = 3
)

var BlockIDFlag_name = map[int32]string{
	0: "BLOCKD_ID_FLAG_UNKNOWN",
	1: "BLOCK_ID_FLAG_ABSENT",
	2: "BLOCK_ID_FLAG_COMMIT",
	3: "BLOCK_ID_FLAG_NIL",
}

var BlockIDFlag_value = map[string]int32{
	"BLOCKD_ID_FLAG_UNKNOWN": 0,
	"BLOCK_ID_FLAG_ABSENT":   1,
	"BLOCK_ID_FLAG_COMMIT":   2,
	"BLOCK_ID_FLAG_NIL":      3,
}

func (BlockIDFlag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{0}
}

// SignedMsgType is a type of signed message in the consensus.
type SignedMsgType int32

const (
	SIGNED_MSG_TYPE_UNKNOWN SignedMsgType = 0
	PrevoteType             SignedMsgType = 1
	PrecommitType           SignedMsgType = 2
	ProposalType            SignedMsgType = 3
)

var SignedMsgType_name = map[int32]string{
	0: "SIGNED_MSG_TYPE_UNKNOWN",
	1: "PREVOTE_TYPE",
	2: "PRECOMMIT_TYPE",
	3: "PROPOSAL_TYPE",
}

var SignedMsgType_value = map[string]int32{
	"SIGNED_MSG_TYPE_UNKNOWN": 0,
	"PREVOTE_TYPE":            1,
	"PRECOMMIT_TYPE":          2,
	"PROPOSAL_TYPE":           3,
}

func (SignedMsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{1}
}

// PartsetHeader
type PartSetHeader struct {
	Total                int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartSetHeader) Reset()         { *m = PartSetHeader{} }
func (m *PartSetHeader) String() string { return proto.CompactTextString(m) }
func (*PartSetHeader) ProtoMessage()    {}
func (*PartSetHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{0}
}
func (m *PartSetHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartSetHeader.Unmarshal(m, b)
}
func (m *PartSetHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartSetHeader.Marshal(b, m, deterministic)
}
func (m *PartSetHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartSetHeader.Merge(m, src)
}
func (m *PartSetHeader) XXX_Size() int {
	return xxx_messageInfo_PartSetHeader.Size(m)
}
func (m *PartSetHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PartSetHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PartSetHeader proto.InternalMessageInfo

func (m *PartSetHeader) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PartSetHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Part struct {
	Index                uint32             `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Bytes                []byte             `protobuf:"bytes,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Proof                merkle.SimpleProof `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Part) Reset()         { *m = Part{} }
func (m *Part) String() string { return proto.CompactTextString(m) }
func (*Part) ProtoMessage()    {}
func (*Part) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{1}
}
func (m *Part) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Part.Unmarshal(m, b)
}
func (m *Part) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Part.Marshal(b, m, deterministic)
}
func (m *Part) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Part.Merge(m, src)
}
func (m *Part) XXX_Size() int {
	return xxx_messageInfo_Part.Size(m)
}
func (m *Part) XXX_DiscardUnknown() {
	xxx_messageInfo_Part.DiscardUnknown(m)
}

var xxx_messageInfo_Part proto.InternalMessageInfo

func (m *Part) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Part) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *Part) GetProof() merkle.SimpleProof {
	if m != nil {
		return m.Proof
	}
	return merkle.SimpleProof{}
}

// BlockID
type BlockID struct {
	Hash                 []byte        `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PartsHeader          PartSetHeader `protobuf:"bytes,2,opt,name=parts_header,json=partsHeader,proto3" json:"parts_header"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{2}
}
func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockID.Unmarshal(m, b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return xxx_messageInfo_BlockID.Size(m)
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockID) GetPartsHeader() PartSetHeader {
	if m != nil {
		return m.PartsHeader
	}
	return PartSetHeader{}
}

// Header defines the structure of a Tendermint block header.
type Header struct {
	// basic block info
	Version version.Consensus `protobuf:"bytes,1,opt,name=version,proto3" json:"version"`
	ChainID string            `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Height  int64             `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Time    time.Time         `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
	// prev block info
	LastBlockID BlockID `protobuf:"bytes,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id"`
	// hashes of block data
	LastCommitHash []byte `protobuf:"bytes,6,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"`
	DataHash       []byte `protobuf:"bytes,7,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	// hashes from the app output from the prev block
	VotersHash      []byte `protobuf:"bytes,8,opt,name=voters_hash,json=votersHash,proto3" json:"voters_hash,omitempty"`
	NextVotersHash  []byte `protobuf:"bytes,9,opt,name=next_voters_hash,json=nextVotersHash,proto3" json:"next_voters_hash,omitempty"`
	ConsensusHash   []byte `protobuf:"bytes,10,opt,name=consensus_hash,json=consensusHash,proto3" json:"consensus_hash,omitempty"`
	AppHash         []byte `protobuf:"bytes,11,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	LastResultsHash []byte `protobuf:"bytes,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`
	// consensus info
	EvidenceHash         []byte   `protobuf:"bytes,13,opt,name=evidence_hash,json=evidenceHash,proto3" json:"evidence_hash,omitempty"`
	ProposerAddress      []byte   `protobuf:"bytes,14,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{3}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() version.Consensus {
	if m != nil {
		return m.Version
	}
	return version.Consensus{}
}

func (m *Header) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Header) GetLastBlockID() BlockID {
	if m != nil {
		return m.LastBlockID
	}
	return BlockID{}
}

func (m *Header) GetLastCommitHash() []byte {
	if m != nil {
		return m.LastCommitHash
	}
	return nil
}

func (m *Header) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (m *Header) GetVotersHash() []byte {
	if m != nil {
		return m.VotersHash
	}
	return nil
}

func (m *Header) GetNextVotersHash() []byte {
	if m != nil {
		return m.NextVotersHash
	}
	return nil
}

func (m *Header) GetConsensusHash() []byte {
	if m != nil {
		return m.ConsensusHash
	}
	return nil
}

func (m *Header) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *Header) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *Header) GetEvidenceHash() []byte {
	if m != nil {
		return m.EvidenceHash
	}
	return nil
}

func (m *Header) GetProposerAddress() []byte {
	if m != nil {
		return m.ProposerAddress
	}
	return nil
}

// Data contains the set of transactions included in the block
type Data struct {
	// Txs that will be applied by state @ block.Height+1.
	// NOTE: not all txs here are valid.  We're just agreeing on the order first.
	// This means that block.AppHash does not include these txs.
	Txs [][]byte `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	// Volatile
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{4}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetTxs() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *Data) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// Vote represents a prevote, precommit, or commit vote from validators for
// consensus.
type Vote struct {
	Type                 SignedMsgType `protobuf:"varint,1,opt,name=type,proto3,enum=tendermint.proto.types.SignedMsgType" json:"type,omitempty"`
	Height               int64         `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Round                int64         `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	BlockID              BlockID       `protobuf:"bytes,4,opt,name=block_id,json=blockId,proto3" json:"block_id"`
	Timestamp            time.Time     `protobuf:"bytes,5,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	ValidatorAddress     []byte        `protobuf:"bytes,6,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	ValidatorIndex       int64         `protobuf:"varint,7,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty"`
	Signature            []byte        `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{5}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vote.Unmarshal(m, b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return xxx_messageInfo_Vote.Size(m)
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetType() SignedMsgType {
	if m != nil {
		return m.Type
	}
	return SIGNED_MSG_TYPE_UNKNOWN
}

func (m *Vote) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Vote) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Vote) GetBlockID() BlockID {
	if m != nil {
		return m.BlockID
	}
	return BlockID{}
}

func (m *Vote) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Vote) GetValidatorAddress() []byte {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

func (m *Vote) GetValidatorIndex() int64 {
	if m != nil {
		return m.ValidatorIndex
	}
	return 0
}

func (m *Vote) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Commit contains the evidence that a block was committed by a set of validators.
type Commit struct {
	Height               int64          `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Round                int32          `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	BlockID              BlockID        `protobuf:"bytes,3,opt,name=block_id,json=blockId,proto3" json:"block_id"`
	Signatures           []CommitSig    `protobuf:"bytes,4,rep,name=signatures,proto3" json:"signatures"`
	Hash                 []byte         `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	BitArray             *bits.BitArray `protobuf:"bytes,6,opt,name=bit_array,json=bitArray,proto3" json:"bit_array,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{6}
}
func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Commit) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Commit) GetBlockID() BlockID {
	if m != nil {
		return m.BlockID
	}
	return BlockID{}
}

func (m *Commit) GetSignatures() []CommitSig {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Commit) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Commit) GetBitArray() *bits.BitArray {
	if m != nil {
		return m.BitArray
	}
	return nil
}

// CommitSig is a part of the Vote included in a Commit.
type CommitSig struct {
	BlockIdFlag          BlockIDFlag `protobuf:"varint,1,opt,name=block_id_flag,json=blockIdFlag,proto3,enum=tendermint.proto.types.BlockIDFlag" json:"block_id_flag,omitempty"`
	ValidatorAddress     []byte      `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Timestamp            time.Time   `protobuf:"bytes,3,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Signature            []byte      `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CommitSig) Reset()         { *m = CommitSig{} }
func (m *CommitSig) String() string { return proto.CompactTextString(m) }
func (*CommitSig) ProtoMessage()    {}
func (*CommitSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{7}
}
func (m *CommitSig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitSig.Unmarshal(m, b)
}
func (m *CommitSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitSig.Marshal(b, m, deterministic)
}
func (m *CommitSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitSig.Merge(m, src)
}
func (m *CommitSig) XXX_Size() int {
	return xxx_messageInfo_CommitSig.Size(m)
}
func (m *CommitSig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitSig.DiscardUnknown(m)
}

var xxx_messageInfo_CommitSig proto.InternalMessageInfo

func (m *CommitSig) GetBlockIdFlag() BlockIDFlag {
	if m != nil {
		return m.BlockIdFlag
	}
	return BLOCKD_ID_FLAG_UNKNOWN
}

func (m *CommitSig) GetValidatorAddress() []byte {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

func (m *CommitSig) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *CommitSig) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Proposal struct {
	Type                 SignedMsgType `protobuf:"varint,1,opt,name=type,proto3,enum=tendermint.proto.types.SignedMsgType" json:"type,omitempty"`
	Height               int64         `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Round                int32         `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	PolRound             int32         `protobuf:"varint,4,opt,name=pol_round,json=polRound,proto3" json:"pol_round,omitempty"`
	BlockID              BlockID       `protobuf:"bytes,5,opt,name=block_id,json=blockId,proto3" json:"block_id"`
	Timestamp            time.Time     `protobuf:"bytes,6,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	Signature            []byte        `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{8}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetType() SignedMsgType {
	if m != nil {
		return m.Type
	}
	return SIGNED_MSG_TYPE_UNKNOWN
}

func (m *Proposal) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Proposal) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Proposal) GetPolRound() int32 {
	if m != nil {
		return m.PolRound
	}
	return 0
}

func (m *Proposal) GetBlockID() BlockID {
	if m != nil {
		return m.BlockID
	}
	return BlockID{}
}

func (m *Proposal) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

func (m *Proposal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type SignedHeader struct {
	Header               *Header  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Commit               *Commit  `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedHeader) Reset()         { *m = SignedHeader{} }
func (m *SignedHeader) String() string { return proto.CompactTextString(m) }
func (*SignedHeader) ProtoMessage()    {}
func (*SignedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{9}
}
func (m *SignedHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedHeader.Unmarshal(m, b)
}
func (m *SignedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedHeader.Marshal(b, m, deterministic)
}
func (m *SignedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedHeader.Merge(m, src)
}
func (m *SignedHeader) XXX_Size() int {
	return xxx_messageInfo_SignedHeader.Size(m)
}
func (m *SignedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SignedHeader proto.InternalMessageInfo

func (m *SignedHeader) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *SignedHeader) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

type BlockMeta struct {
	BlockID              BlockID  `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id"`
	BlockSize            int64    `protobuf:"varint,2,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	Header               Header   `protobuf:"bytes,3,opt,name=header,proto3" json:"header"`
	NumTxs               int64    `protobuf:"varint,4,opt,name=num_txs,json=numTxs,proto3" json:"num_txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockMeta) Reset()         { *m = BlockMeta{} }
func (m *BlockMeta) String() string { return proto.CompactTextString(m) }
func (*BlockMeta) ProtoMessage()    {}
func (*BlockMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff06f8095857fb18, []int{10}
}
func (m *BlockMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMeta.Unmarshal(m, b)
}
func (m *BlockMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMeta.Marshal(b, m, deterministic)
}
func (m *BlockMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMeta.Merge(m, src)
}
func (m *BlockMeta) XXX_Size() int {
	return xxx_messageInfo_BlockMeta.Size(m)
}
func (m *BlockMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMeta.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMeta proto.InternalMessageInfo

func (m *BlockMeta) GetBlockID() BlockID {
	if m != nil {
		return m.BlockID
	}
	return BlockID{}
}

func (m *BlockMeta) GetBlockSize() int64 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *BlockMeta) GetHeader() Header {
	if m != nil {
		return m.Header
	}
	return Header{}
}

func (m *BlockMeta) GetNumTxs() int64 {
	if m != nil {
		return m.NumTxs
	}
	return 0
}

func init() {
	proto.RegisterEnum("tendermint.proto.types.BlockIDFlag", BlockIDFlag_name, BlockIDFlag_value)
	proto.RegisterEnum("tendermint.proto.types.SignedMsgType", SignedMsgType_name, SignedMsgType_value)
	proto.RegisterType((*PartSetHeader)(nil), "tendermint.proto.types.PartSetHeader")
	proto.RegisterType((*Part)(nil), "tendermint.proto.types.Part")
	proto.RegisterType((*BlockID)(nil), "tendermint.proto.types.BlockID")
	proto.RegisterType((*Header)(nil), "tendermint.proto.types.Header")
	proto.RegisterType((*Data)(nil), "tendermint.proto.types.Data")
	proto.RegisterType((*Vote)(nil), "tendermint.proto.types.Vote")
	proto.RegisterType((*Commit)(nil), "tendermint.proto.types.Commit")
	proto.RegisterType((*CommitSig)(nil), "tendermint.proto.types.CommitSig")
	proto.RegisterType((*Proposal)(nil), "tendermint.proto.types.Proposal")
	proto.RegisterType((*SignedHeader)(nil), "tendermint.proto.types.SignedHeader")
	proto.RegisterType((*BlockMeta)(nil), "tendermint.proto.types.BlockMeta")
}

func init() { proto.RegisterFile("proto/types/types.proto", fileDescriptor_ff06f8095857fb18) }

var fileDescriptor_ff06f8095857fb18 = []byte{
	// 1271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdd, 0x6e, 0x1a, 0x47,
	0x14, 0xf6, 0xc2, 0x62, 0xe0, 0x2c, 0x60, 0xbc, 0x4d, 0x13, 0x8a, 0x5b, 0x43, 0x70, 0x93, 0x92,
	0x1f, 0x2d, 0x92, 0x2b, 0x55, 0x8d, 0xd4, 0x1b, 0xb0, 0x1d, 0x07, 0xc5, 0xc6, 0x68, 0xa1, 0xe9,
	0xcf, 0xcd, 0x6a, 0x60, 0x27, 0xb0, 0xca, 0xb2, 0xbb, 0xda, 0x19, 0x2c, 0x3b, 0x95, 0x7a, 0x5d,
	0xf9, 0xaa, 0x2f, 0xe0, 0xab, 0xb4, 0x52, 0xdf, 0xa2, 0xbd, 0xec, 0x55, 0x1f, 0x21, 0x95, 0xd2,
	0x27, 0xa8, 0xd4, 0x07, 0xa8, 0xe6, 0x67, 0x17, 0x08, 0x76, 0x6b, 0x35, 0xe9, 0x8d, 0xbd, 0x73,
	0xce, 0x77, 0xce, 0xcc, 0xf9, 0xce, 0x77, 0x66, 0x6c, 0xb8, 0x11, 0x84, 0x3e, 0xf5, 0x1b, 0xf4,
	0x34, 0xc0, 0x44, 0xfc, 0x34, 0xb8, 0x45, 0xbf, 0x4e, 0xb1, 0x67, 0xe3, 0x70, 0xe2, 0x78, 0x54,
	0x58, 0x0c, 0xee, 0x2d, 0xdf, 0xa6, 0x63, 0x27, 0xb4, 0xad, 0x00, 0x85, 0xf4, 0xb4, 0x21, 0x82,
	0x47, 0xfe, 0xc8, 0x9f, 0x7d, 0x09, 0x74, 0xb9, 0x32, 0xf2, 0xfd, 0x91, 0x8b, 0x05, 0x64, 0x30,
	0x7d, 0xda, 0xa0, 0xce, 0x04, 0x13, 0x8a, 0x26, 0x81, 0x04, 0x6c, 0x88, 0x10, 0xd7, 0x19, 0x90,
	0xc6, 0xc0, 0xa1, 0x0b, 0xbb, 0x97, 0x2b, 0xc2, 0x39, 0x0c, 0x4f, 0x03, 0xea, 0x37, 0x26, 0x38,
	0x7c, 0xe6, 0xe2, 0x05, 0x80, 0x8c, 0x3e, 0xc6, 0x21, 0x71, 0x7c, 0x2f, 0xfa, 0x2d, 0x9c, 0xb5,
	0x07, 0x90, 0xef, 0xa2, 0x90, 0xf6, 0x30, 0x7d, 0x84, 0x91, 0x8d, 0x43, 0xfd, 0x1a, 0xa4, 0xa8,
	0x4f, 0x91, 0x5b, 0x52, 0xaa, 0x4a, 0x3d, 0x69, 0x8a, 0x85, 0xae, 0x83, 0x3a, 0x46, 0x64, 0x5c,
	0x4a, 0x54, 0x95, 0x7a, 0xce, 0xe4, 0xdf, 0xb5, 0x6f, 0x40, 0x65, 0xa1, 0x2c, 0xc2, 0xf1, 0x6c,
	0x7c, 0xc2, 0x23, 0xf2, 0xa6, 0x58, 0x30, 0xeb, 0xe0, 0x94, 0x62, 0x22, 0x43, 0xc4, 0x42, 0xdf,
	0x87, 0x54, 0x10, 0xfa, 0xfe, 0xd3, 0x52, 0xb2, 0xaa, 0xd4, 0xb5, 0xed, 0x7b, 0xc6, 0x12, 0x75,
	0xa2, 0x0e, 0x43, 0xd4, 0x61, 0xf4, 0x9c, 0x49, 0xe0, 0xe2, 0x2e, 0x0b, 0x69, 0xa9, 0xbf, 0xbe,
	0xac, 0xac, 0x98, 0x22, 0xbe, 0x36, 0x81, 0x74, 0xcb, 0xf5, 0x87, 0xcf, 0xda, 0xbb, 0xf1, 0xd9,
	0x94, 0xd9, 0xd9, 0xf4, 0x0e, 0xe4, 0x18, 0xed, 0xc4, 0x1a, 0xf3, 0xaa, 0xf8, 0x21, 0xb4, 0xed,
	0x5b, 0xc6, 0xc5, 0x9d, 0x32, 0x16, 0x28, 0x90, 0x1b, 0x69, 0x3c, 0x81, 0x30, 0xd5, 0xfe, 0x54,
	0x61, 0x55, 0x12, 0xb4, 0x03, 0x69, 0x49, 0x21, 0xdf, 0x51, 0xdb, 0xde, 0x5a, 0xce, 0x1a, 0x71,
	0xbc, 0xe3, 0x7b, 0x04, 0x7b, 0x64, 0x4a, 0x64, 0xce, 0x28, 0x52, 0xbf, 0x0d, 0x99, 0xe1, 0x18,
	0x39, 0x9e, 0xe5, 0xd8, 0xfc, 0x6c, 0xd9, 0x96, 0xf6, 0xea, 0x65, 0x25, 0xbd, 0xc3, 0x6c, 0xed,
	0x5d, 0x33, 0xcd, 0x9d, 0x6d, 0x5b, 0xbf, 0x0e, 0xab, 0x63, 0xec, 0x8c, 0xc6, 0x94, 0x13, 0x96,
	0x34, 0xe5, 0x4a, 0xff, 0x14, 0x54, 0x26, 0x92, 0x92, 0xca, 0x4f, 0x50, 0x36, 0x84, 0x82, 0x8c,
	0x48, 0x41, 0x46, 0x3f, 0x52, 0x50, 0x2b, 0xc3, 0x36, 0xfe, 0xfe, 0xf7, 0x8a, 0x62, 0xf2, 0x08,
	0xfd, 0x4b, 0xc8, 0xbb, 0x88, 0x50, 0x6b, 0xc0, 0xd8, 0x63, 0xdb, 0xa7, 0x78, 0x8a, 0xca, 0x65,
	0xd4, 0x48, 0x96, 0x5b, 0xef, 0xb0, 0x3c, 0xaf, 0x5e, 0x56, 0xb4, 0x03, 0x44, 0xa8, 0x34, 0x9a,
	0x9a, 0x1b, 0x2f, 0x6c, 0xbd, 0x0e, 0x45, 0x9e, 0x79, 0xe8, 0x4f, 0x26, 0x0e, 0xb5, 0x78, 0x4f,
	0x56, 0x79, 0x4f, 0x0a, 0xcc, 0xbe, 0xc3, 0xcd, 0x8f, 0x58, 0x77, 0x36, 0x20, 0x6b, 0x23, 0x8a,
	0x04, 0x24, 0xcd, 0x21, 0x19, 0x66, 0xe0, 0xce, 0x0a, 0x68, 0xc7, 0x3e, 0xc5, 0x21, 0x11, 0xee,
	0x0c, 0x77, 0x83, 0x30, 0x71, 0x40, 0x1d, 0x8a, 0x1e, 0x3e, 0xa1, 0xd6, 0x3c, 0x2a, 0x2b, 0xf6,
	0x61, 0xf6, 0x27, 0x33, 0xe4, 0x2d, 0x28, 0x0c, 0xa3, 0x0e, 0x08, 0x1c, 0x70, 0x5c, 0x3e, 0xb6,
	0x72, 0xd8, 0x7b, 0x90, 0x41, 0x41, 0x20, 0x00, 0x1a, 0x07, 0xa4, 0x51, 0x10, 0x70, 0xd7, 0x5d,
	0x58, 0xe7, 0x35, 0x85, 0x98, 0x4c, 0x5d, 0x2a, 0x93, 0xe4, 0x38, 0x66, 0x8d, 0x39, 0x4c, 0x61,
	0xe7, 0xd8, 0x2d, 0xc8, 0xe3, 0x63, 0xc7, 0xc6, 0xde, 0x10, 0x0b, 0x5c, 0x9e, 0xe3, 0x72, 0x91,
	0x91, 0x83, 0xee, 0x40, 0x31, 0x08, 0xfd, 0xc0, 0x27, 0x38, 0xb4, 0x90, 0x6d, 0x87, 0x98, 0x90,
	0x52, 0x41, 0xe4, 0x8b, 0xec, 0x4d, 0x61, 0xae, 0xdd, 0x07, 0x75, 0x17, 0x51, 0xa4, 0x17, 0x21,
	0x49, 0x4f, 0x48, 0x49, 0xa9, 0x26, 0xeb, 0x39, 0x93, 0x7d, 0x5e, 0x38, 0x8d, 0x7f, 0x25, 0x40,
	0x65, 0xa5, 0xeb, 0x0f, 0x40, 0x65, 0x9d, 0xe3, 0xe2, 0x2c, 0x5c, 0x2e, 0xf9, 0x9e, 0x33, 0xf2,
	0xb0, 0x7d, 0x48, 0x46, 0xfd, 0xd3, 0x00, 0x9b, 0x3c, 0x64, 0x4e, 0x6d, 0x89, 0x05, 0xb5, 0x5d,
	0x83, 0x54, 0xe8, 0x4f, 0x3d, 0x5b, 0x8a, 0x50, 0x2c, 0xf4, 0xc7, 0x90, 0x89, 0x45, 0xa4, 0x5e,
	0x4d, 0x44, 0x6b, 0x52, 0x44, 0xd1, 0xec, 0x9a, 0xe9, 0x81, 0x14, 0x4f, 0x0b, 0xb2, 0xf1, 0xad,
	0x27, 0x25, 0x79, 0x35, 0x55, 0xcf, 0xc2, 0xf4, 0x7b, 0xb0, 0x7e, 0x8c, 0x5c, 0xc7, 0x46, 0xd4,
	0x9f, 0x91, 0x2b, 0x14, 0x58, 0x8c, 0x1d, 0x92, 0x5d, 0xfd, 0x23, 0x58, 0x9b, 0x81, 0xc5, 0xfd,
	0x95, 0xe6, 0xd5, 0x15, 0x62, 0x73, 0x9b, 0x5f, 0x64, 0xef, 0x43, 0x96, 0x38, 0x23, 0x0f, 0xd1,
	0x69, 0x88, 0xa5, 0x1a, 0x67, 0x86, 0xda, 0x8b, 0x04, 0xac, 0x0a, 0x65, 0xcf, 0xb1, 0xa7, 0x5c,
	0xcc, 0x1e, 0x23, 0x35, 0x75, 0x11, 0x7b, 0xc9, 0x37, 0x65, 0x6f, 0x1f, 0x20, 0x3e, 0x12, 0x29,
	0xa9, 0xd5, 0x64, 0x5d, 0xdb, 0xbe, 0x79, 0x59, 0x3a, 0x71, 0xdc, 0x9e, 0x33, 0x92, 0x97, 0xd2,
	0x5c, 0x68, 0xac, 0xac, 0xd4, 0xdc, 0x5d, 0xda, 0x84, 0xec, 0xc0, 0xa1, 0x16, 0x0a, 0x43, 0x74,
	0xca, 0xe9, 0xd4, 0xb6, 0x3f, 0x5c, 0xce, 0xcd, 0x1e, 0x27, 0x83, 0x3d, 0x4e, 0x46, 0xcb, 0xa1,
	0x4d, 0x86, 0x35, 0x33, 0x03, 0xf9, 0x55, 0xfb, 0x43, 0x81, 0x6c, 0xbc, 0xad, 0xbe, 0x0f, 0xf9,
	0xa8, 0x74, 0xeb, 0xa9, 0x8b, 0x46, 0x52, 0xaa, 0x5b, 0xff, 0x52, 0xff, 0x43, 0x17, 0x8d, 0x4c,
	0x4d, 0x96, 0xcc, 0x16, 0x17, 0x37, 0x3c, 0x71, 0x49, 0xc3, 0x17, 0x14, 0x96, 0xfc, 0x6f, 0x0a,
	0x5b, 0xd0, 0x82, 0xfa, 0xba, 0x16, 0x7e, 0x4e, 0x40, 0xa6, 0xcb, 0x87, 0x18, 0xb9, 0xff, 0xfb,
	0x18, 0xc6, 0x42, 0xda, 0x80, 0x6c, 0xe0, 0xbb, 0x96, 0xf0, 0xa8, 0xdc, 0x93, 0x09, 0x7c, 0xd7,
	0x5c, 0x52, 0x59, 0xea, 0xad, 0xce, 0xe8, 0xea, 0x5b, 0x60, 0x30, 0xfd, 0x3a, 0x83, 0xdf, 0x42,
	0x4e, 0x10, 0x22, 0xdf, 0xda, 0x4f, 0x18, 0x13, 0xfc, 0x01, 0x17, 0x4f, 0xed, 0xe6, 0x65, 0x87,
	0x17, 0x78, 0x53, 0xa2, 0x59, 0x9c, 0x78, 0x85, 0xe4, 0xc3, 0xbf, 0xf9, 0xcf, 0xb3, 0x60, 0x4a,
	0x74, 0xed, 0x37, 0x05, 0xb2, 0xbc, 0xec, 0x43, 0x4c, 0xd1, 0x02, 0x79, 0xca, 0x9b, 0x92, 0xf7,
	0x01, 0x80, 0x48, 0x46, 0x9c, 0xe7, 0x58, 0x36, 0x36, 0xcb, 0x2d, 0x3d, 0xe7, 0x39, 0xd6, 0x3f,
	0x8b, 0x2b, 0x4d, 0x5e, 0xa5, 0x52, 0x39, 0xba, 0x51, 0xbd, 0x37, 0x20, 0xed, 0x4d, 0x27, 0x16,
	0x7b, 0x26, 0x54, 0x21, 0x19, 0x6f, 0x3a, 0xe9, 0x9f, 0x90, 0xbb, 0xbf, 0x28, 0xa0, 0xcd, 0x8d,
	0x8f, 0x5e, 0x86, 0xeb, 0xad, 0x83, 0xa3, 0x9d, 0xc7, 0xbb, 0x56, 0x7b, 0xd7, 0x7a, 0x78, 0xd0,
	0xdc, 0xb7, 0x3e, 0xef, 0x3c, 0xee, 0x1c, 0x7d, 0xd1, 0x29, 0xae, 0xe8, 0x0d, 0xb8, 0xc6, 0x7d,
	0xb1, 0xab, 0xd9, 0xea, 0xed, 0x75, 0xfa, 0x45, 0xa5, 0xfc, 0xee, 0xd9, 0x79, 0x75, 0x7d, 0x2e,
	0x4d, 0x73, 0x40, 0xb0, 0x47, 0x97, 0x03, 0x76, 0x8e, 0x0e, 0x0f, 0xdb, 0xfd, 0x62, 0x62, 0x29,
	0x40, 0xde, 0x90, 0x77, 0x60, 0x7d, 0x31, 0xa0, 0xd3, 0x3e, 0x28, 0x26, 0xcb, 0xfa, 0xd9, 0x79,
	0xb5, 0x30, 0x87, 0xee, 0x38, 0x6e, 0x39, 0xf3, 0xdd, 0x8b, 0xcd, 0x95, 0x9f, 0x7e, 0xd8, 0x5c,
	0xb9, 0xfb, 0xa3, 0x02, 0xf9, 0x85, 0x29, 0xd1, 0x37, 0xe0, 0x46, 0xaf, 0xbd, 0xdf, 0xd9, 0xdb,
	0xb5, 0x0e, 0x7b, 0xfb, 0x56, 0xff, 0xab, 0xee, 0xde, 0x5c, 0x15, 0x37, 0x21, 0xd7, 0x35, 0xf7,
	0x9e, 0x1c, 0xf5, 0xf7, 0xb8, 0xa7, 0xa8, 0x94, 0xd7, 0xce, 0xce, 0xab, 0x5a, 0x37, 0xc4, 0xec,
	0xef, 0x05, 0x1e, 0x7f, 0x0b, 0x0a, 0x5d, 0x73, 0x4f, 0x1c, 0x56, 0x80, 0x12, 0xe5, 0xf5, 0xb3,
	0xf3, 0x6a, 0xbe, 0x1b, 0x62, 0x21, 0x04, 0x0e, 0xdb, 0x82, 0x7c, 0xd7, 0x3c, 0xea, 0x1e, 0xf5,
	0x9a, 0x07, 0x02, 0x95, 0x2c, 0x17, 0xcf, 0xce, 0xab, 0xb9, 0x68, 0xc4, 0x19, 0x68, 0x76, 0xce,
	0x96, 0xf1, 0xf5, 0xfd, 0x91, 0x43, 0xc7, 0xd3, 0x81, 0x31, 0xf4, 0x27, 0x8d, 0x59, 0xf7, 0xe6,
	0x3f, 0xe7, 0xfe, 0x83, 0x18, 0xac, 0xf2, 0xc5, 0xc7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa8,
	0x32, 0x32, 0xb1, 0x57, 0x0c, 0x00, 0x00,
}
