// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/display_ad_format_setting.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

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

// Enumerates display ad format settings.
type DisplayAdFormatSettingEnum_DisplayAdFormatSetting int32

const (
	// Not specified.
	DisplayAdFormatSettingEnum_UNSPECIFIED DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 0
	// The value is unknown in this version.
	DisplayAdFormatSettingEnum_UNKNOWN DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 1
	// Text, image and native formats.
	DisplayAdFormatSettingEnum_ALL_FORMATS DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 2
	// Text and image formats.
	DisplayAdFormatSettingEnum_NON_NATIVE DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 3
	// Native format, i.e. the format rendering is controlled by the publisher
	// and not by Google.
	DisplayAdFormatSettingEnum_NATIVE DisplayAdFormatSettingEnum_DisplayAdFormatSetting = 4
)

var DisplayAdFormatSettingEnum_DisplayAdFormatSetting_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ALL_FORMATS",
	3: "NON_NATIVE",
	4: "NATIVE",
}
var DisplayAdFormatSettingEnum_DisplayAdFormatSetting_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ALL_FORMATS": 2,
	"NON_NATIVE":  3,
	"NATIVE":      4,
}

func (x DisplayAdFormatSettingEnum_DisplayAdFormatSetting) String() string {
	return proto.EnumName(DisplayAdFormatSettingEnum_DisplayAdFormatSetting_name, int32(x))
}
func (DisplayAdFormatSettingEnum_DisplayAdFormatSetting) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_display_ad_format_setting_4e0440badaec0ae7, []int{0, 0}
}

// Container for display ad format settings.
type DisplayAdFormatSettingEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisplayAdFormatSettingEnum) Reset()         { *m = DisplayAdFormatSettingEnum{} }
func (m *DisplayAdFormatSettingEnum) String() string { return proto.CompactTextString(m) }
func (*DisplayAdFormatSettingEnum) ProtoMessage()    {}
func (*DisplayAdFormatSettingEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_display_ad_format_setting_4e0440badaec0ae7, []int{0}
}
func (m *DisplayAdFormatSettingEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisplayAdFormatSettingEnum.Unmarshal(m, b)
}
func (m *DisplayAdFormatSettingEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisplayAdFormatSettingEnum.Marshal(b, m, deterministic)
}
func (dst *DisplayAdFormatSettingEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisplayAdFormatSettingEnum.Merge(dst, src)
}
func (m *DisplayAdFormatSettingEnum) XXX_Size() int {
	return xxx_messageInfo_DisplayAdFormatSettingEnum.Size(m)
}
func (m *DisplayAdFormatSettingEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DisplayAdFormatSettingEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DisplayAdFormatSettingEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DisplayAdFormatSettingEnum)(nil), "google.ads.googleads.v0.enums.DisplayAdFormatSettingEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.DisplayAdFormatSettingEnum_DisplayAdFormatSetting", DisplayAdFormatSettingEnum_DisplayAdFormatSetting_name, DisplayAdFormatSettingEnum_DisplayAdFormatSetting_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/display_ad_format_setting.proto", fileDescriptor_display_ad_format_setting_4e0440badaec0ae7)
}

var fileDescriptor_display_ad_format_setting_4e0440badaec0ae7 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x6d, 0x27, 0x13, 0x32, 0xd0, 0xd2, 0x83, 0x07, 0x65, 0x07, 0xf7, 0x00, 0x69, 0xc1,
	0xa3, 0x78, 0x48, 0x5d, 0x3b, 0x8a, 0x33, 0x2d, 0x76, 0xab, 0x20, 0x85, 0x12, 0x97, 0x1a, 0x0a,
	0x6d, 0x53, 0x9a, 0x6e, 0xe0, 0xd1, 0x57, 0xf1, 0xe8, 0xa3, 0x88, 0x0f, 0x25, 0x49, 0xb6, 0x9d,
	0xe6, 0x2e, 0xe5, 0x83, 0x1f, 0xbf, 0xaf, 0xff, 0x7c, 0xe0, 0x9e, 0x71, 0xce, 0xaa, 0xc2, 0x21,
	0x54, 0x38, 0x3a, 0xca, 0xb4, 0x71, 0x9d, 0xa2, 0x59, 0xd7, 0xc2, 0xa1, 0xa5, 0x68, 0x2b, 0xf2,
	0x91, 0x13, 0x9a, 0xbf, 0xf3, 0xae, 0x26, 0x7d, 0x2e, 0x8a, 0xbe, 0x2f, 0x1b, 0x06, 0xdb, 0x8e,
	0xf7, 0xdc, 0x1e, 0x6b, 0x07, 0x12, 0x2a, 0xe0, 0x5e, 0x87, 0x1b, 0x17, 0x2a, 0x7d, 0xf2, 0x69,
	0x80, 0xab, 0xa9, 0xae, 0x40, 0x34, 0x50, 0x05, 0x89, 0xf6, 0xfd, 0x66, 0x5d, 0x4f, 0x56, 0xe0,
	0xf2, 0x30, 0xb5, 0x2f, 0xc0, 0x68, 0x89, 0x93, 0xd8, 0x7f, 0x08, 0x83, 0xd0, 0x9f, 0x5a, 0x27,
	0xf6, 0x08, 0x9c, 0x2d, 0xf1, 0x23, 0x8e, 0x5e, 0xb0, 0x65, 0x48, 0x8a, 0xe6, 0xf3, 0x3c, 0x88,
	0x9e, 0x9f, 0xd0, 0x22, 0xb1, 0x4c, 0xfb, 0x1c, 0x00, 0x1c, 0xe1, 0x1c, 0xa3, 0x45, 0x98, 0xfa,
	0xd6, 0xc0, 0x06, 0x60, 0xb8, 0xcd, 0xa7, 0xde, 0xaf, 0x01, 0x6e, 0x56, 0xbc, 0x86, 0x47, 0x2f,
	0xf5, 0xae, 0x0f, 0x1f, 0x12, 0xcb, 0x57, 0xc6, 0xc6, 0xab, 0xb7, 0xb5, 0x19, 0xaf, 0x48, 0xc3,
	0x20, 0xef, 0x98, 0xc3, 0x8a, 0x46, 0x6d, 0xb0, 0x9b, 0xad, 0x2d, 0xc5, 0x3f, 0x2b, 0xde, 0xa9,
	0xef, 0x97, 0x39, 0x98, 0x21, 0xf4, 0x6d, 0x8e, 0x67, 0xba, 0x0a, 0x51, 0x01, 0x75, 0x94, 0x29,
	0x75, 0xa1, 0x9c, 0x44, 0xfc, 0xec, 0x78, 0x86, 0xa8, 0xc8, 0xf6, 0x3c, 0x4b, 0xdd, 0x4c, 0xf1,
	0xb7, 0xa1, 0xfa, 0xe9, 0xed, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x9a, 0x0f, 0xb8, 0xb9,
	0x01, 0x00, 0x00,
}
