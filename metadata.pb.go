// Code generated by protoc-gen-go.
// source: btree_metatype.proto
// DO NOT EDIT!

package btree

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type BtreeMetaData struct {
	Root             *int32  `protobuf:"varint,1,opt,name=root" json:"root,omitempty"`
	LeafCount        *int32  `protobuf:"varint,2,opt,name=leaf_count" json:"leaf_count,omitempty"`
	NodeCount        *int32  `protobuf:"varint,3,opt,name=node_count" json:"node_count,omitempty"`
	LeafMax          *int32  `protobuf:"varint,4,opt,name=leaf_max" json:"leaf_max,omitempty"`
	NodeMax          *int32  `protobuf:"varint,5,opt,name=node_max" json:"node_max,omitempty"`
	FreeList         []int32 `protobuf:"varint,6,rep,name=free_list" json:"free_list,omitempty"`
	Size             *int32  `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	Version          *int32  `protobuf:"varint,8,opt,name=version" json:"version,omitempty"`
	IndexCursor      *int32  `protobuf:"varint,9,opt,name=index_cursor" json:"index_cursor,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *BtreeMetaData) Reset()         { *this = BtreeMetaData{} }
func (this *BtreeMetaData) String() string { return proto.CompactTextString(this) }
func (*BtreeMetaData) ProtoMessage()       {}

func (this *BtreeMetaData) GetRoot() int32 {
	if this != nil && this.Root != nil {
		return *this.Root
	}
	return 0
}

func (this *BtreeMetaData) GetLeafCount() int32 {
	if this != nil && this.LeafCount != nil {
		return *this.LeafCount
	}
	return 0
}

func (this *BtreeMetaData) GetNodeCount() int32 {
	if this != nil && this.NodeCount != nil {
		return *this.NodeCount
	}
	return 0
}

func (this *BtreeMetaData) GetLeafMax() int32 {
	if this != nil && this.LeafMax != nil {
		return *this.LeafMax
	}
	return 0
}

func (this *BtreeMetaData) GetNodeMax() int32 {
	if this != nil && this.NodeMax != nil {
		return *this.NodeMax
	}
	return 0
}

func (this *BtreeMetaData) GetSize() int32 {
	if this != nil && this.Size != nil {
		return *this.Size
	}
	return 0
}

func (this *BtreeMetaData) GetVersion() int32 {
	if this != nil && this.Version != nil {
		return *this.Version
	}
	return 0
}

func (this *BtreeMetaData) GetIndexCursor() int32 {
	if this != nil && this.IndexCursor != nil {
		return *this.IndexCursor
	}
	return 0
}

type IndexMetaData struct {
	Id               *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Version          *int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *IndexMetaData) Reset()         { *this = IndexMetaData{} }
func (this *IndexMetaData) String() string { return proto.CompactTextString(this) }
func (*IndexMetaData) ProtoMessage()       {}

func (this *IndexMetaData) GetId() int32 {
	if this != nil && this.Id != nil {
		return *this.Id
	}
	return 0
}

func (this *IndexMetaData) GetVersion() int32 {
	if this != nil && this.Version != nil {
		return *this.Version
	}
	return 0
}

type NodeRecordMetaData struct {
	Childrens        []int32  `protobuf:"varint,1,rep,name=childrens" json:"childrens,omitempty"`
	Keys             [][]byte `protobuf:"bytes,2,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *NodeRecordMetaData) Reset()         { *this = NodeRecordMetaData{} }
func (this *NodeRecordMetaData) String() string { return proto.CompactTextString(this) }
func (*NodeRecordMetaData) ProtoMessage()       {}

type LeafRecordMetaData struct {
	Keys             [][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty"`
	Values           [][]byte `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *LeafRecordMetaData) Reset()         { *this = LeafRecordMetaData{} }
func (this *LeafRecordMetaData) String() string { return proto.CompactTextString(this) }
func (*LeafRecordMetaData) ProtoMessage()       {}

func init() {
}