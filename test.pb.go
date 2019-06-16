// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	SearchRequest
*/
package test

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

type Corpus int32

const (
	Corpus_UNIVERSAL Corpus = 0
	Corpus_WEB       Corpus = 1
	Corpus_IMAGES    Corpus = 2
	Corpus_LOCAL     Corpus = 3
	Corpus_NEWS      Corpus = 4
	Corpus_PRODUCTS  Corpus = 5
	Corpus_VIDEO     Corpus = 6
)

var Corpus_name = map[int32]string{
	0: "UNIVERSAL",
	1: "WEB",
	2: "IMAGES",
	3: "LOCAL",
	4: "NEWS",
	5: "PRODUCTS",
	6: "VIDEO",
}
var Corpus_value = map[string]int32{
	"UNIVERSAL": 0,
	"WEB":       1,
	"IMAGES":    2,
	"LOCAL":     3,
	"NEWS":      4,
	"PRODUCTS":  5,
	"VIDEO":     6,
}

func (x Corpus) String() string {
	return proto.EnumName(Corpus_name, int32(x))
}
func (Corpus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SearchRequest struct {
	Query         string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	PageNumber    int32  `protobuf:"varint,2,opt,name=page_number,json=pageNumber" json:"page_number,omitempty"`
	ResultPerPage int32  `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage" json:"result_per_page,omitempty"`
	Corpus        Corpus `protobuf:"varint,4,opt,name=corpus,enum=Corpus" json:"corpus,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func (m *SearchRequest) GetCorpus() Corpus {
	if m != nil {
		return m.Corpus
	}
	return Corpus_UNIVERSAL
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "SearchRequest")
	proto.RegisterEnum("Corpus", Corpus_name, Corpus_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0xcf, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0x71, 0xb7, 0x05, 0x5a, 0x46, 0xd1, 0xcd, 0xc4, 0x03, 0xb7, 0x12, 0x0f, 0x86, 0x78,
	0xe8, 0x41, 0x9f, 0x00, 0xe9, 0xc6, 0x90, 0x20, 0x90, 0x5d, 0xdb, 0x26, 0x5e, 0x08, 0x6d, 0x26,
	0xf5, 0xa0, 0x42, 0x97, 0xdd, 0x83, 0x4f, 0xe1, 0x2b, 0x1b, 0xc0, 0xe3, 0xf7, 0x9b, 0xff, 0x65,
	0x00, 0x0c, 0xf5, 0x66, 0xdd, 0xe9, 0xd6, 0xb4, 0x77, 0xbf, 0x0c, 0x02, 0x45, 0x8d, 0x3e, 0x7e,
	0x48, 0x3a, 0x5b, 0xea, 0x0d, 0xde, 0x82, 0x7b, 0xb6, 0xa4, 0x7f, 0x42, 0x16, 0xb1, 0xd8, 0x97,
	0xd3, 0xc0, 0x15, 0x5c, 0x76, 0xcd, 0x89, 0xea, 0x6f, 0xfb, 0x75, 0x20, 0x1d, 0xce, 0x22, 0x16,
	0xbb, 0x12, 0x06, 0x2a, 0x46, 0xc1, 0x7b, 0xb8, 0xd1, 0xd4, 0xdb, 0x4f, 0x53, 0x77, 0xa4, 0xeb,
	0xe1, 0x10, 0xce, 0xc7, 0x28, 0x98, 0xb8, 0x22, 0x5d, 0x35, 0x27, 0xc2, 0x15, 0x78, 0xc7, 0x56,
	0x77, 0xb6, 0x0f, 0x9d, 0x88, 0xc5, 0xd7, 0x8f, 0x8b, 0x75, 0x3a, 0x4e, 0xf9, 0xcf, 0x0f, 0xef,
	0xe0, 0x4d, 0x82, 0x01, 0xf8, 0xdb, 0x22, 0xdb, 0x09, 0xa9, 0x92, 0x9c, 0x5f, 0xe0, 0x02, 0xe6,
	0x7b, 0xf1, 0xcc, 0x19, 0x02, 0x78, 0xd9, 0x6b, 0xf2, 0x22, 0x14, 0x9f, 0xa1, 0x0f, 0x6e, 0x5e,
	0xa6, 0x49, 0xce, 0xe7, 0xb8, 0x04, 0xa7, 0x10, 0x7b, 0xc5, 0x1d, 0xbc, 0x82, 0x65, 0x25, 0xcb,
	0xcd, 0x36, 0x7d, 0x53, 0xdc, 0x1d, 0x92, 0x5d, 0xb6, 0x11, 0x25, 0xf7, 0x0e, 0xde, 0xf8, 0xf4,
	0xd3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x4c, 0x8b, 0x70, 0x02, 0x01, 0x00, 0x00,
}