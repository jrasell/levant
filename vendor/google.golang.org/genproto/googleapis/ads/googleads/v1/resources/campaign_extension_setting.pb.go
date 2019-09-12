// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_extension_setting.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A campaign extension setting.
type CampaignExtensionSetting struct {
	// The resource name of the campaign extension setting.
	// CampaignExtensionSetting resource names have the form:
	//
	//
	// `customers/{customer_id}/campaignExtensionSettings/{campaign_id}~{extension_type}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The extension type of the customer extension setting.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// The resource name of the campaign. The linked extension feed items will
	// serve under this campaign.
	// Campaign resource names have the form:
	//
	// `customers/{customer_id}/campaigns/{campaign_id}`
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The resource names of the extension feed items to serve under the campaign.
	// ExtensionFeedItem resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ExtensionFeedItems []*wrappers.StringValue `protobuf:"bytes,4,rep,name=extension_feed_items,json=extensionFeedItems,proto3" json:"extension_feed_items,omitempty"`
	// The device for which the extensions will serve. Optional.
	Device               enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice `protobuf:"varint,5,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *CampaignExtensionSetting) Reset()         { *m = CampaignExtensionSetting{} }
func (m *CampaignExtensionSetting) String() string { return proto.CompactTextString(m) }
func (*CampaignExtensionSetting) ProtoMessage()    {}
func (*CampaignExtensionSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_campaign_extension_setting_fad2cef1ff7ae86a, []int{0}
}
func (m *CampaignExtensionSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignExtensionSetting.Unmarshal(m, b)
}
func (m *CampaignExtensionSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignExtensionSetting.Marshal(b, m, deterministic)
}
func (dst *CampaignExtensionSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignExtensionSetting.Merge(dst, src)
}
func (m *CampaignExtensionSetting) XXX_Size() int {
	return xxx_messageInfo_CampaignExtensionSetting.Size(m)
}
func (m *CampaignExtensionSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignExtensionSetting.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignExtensionSetting proto.InternalMessageInfo

func (m *CampaignExtensionSetting) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignExtensionSetting) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *CampaignExtensionSetting) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignExtensionSetting) GetExtensionFeedItems() []*wrappers.StringValue {
	if m != nil {
		return m.ExtensionFeedItems
	}
	return nil
}

func (m *CampaignExtensionSetting) GetDevice() enums.ExtensionSettingDeviceEnum_ExtensionSettingDevice {
	if m != nil {
		return m.Device
	}
	return enums.ExtensionSettingDeviceEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignExtensionSetting)(nil), "google.ads.googleads.v1.resources.CampaignExtensionSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_extension_setting.proto", fileDescriptor_campaign_extension_setting_fad2cef1ff7ae86a)
}

var fileDescriptor_campaign_extension_setting_fad2cef1ff7ae86a = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0x52, 0x98, 0x20, 0xb0, 0x3d, 0x44, 0x3c, 0x44, 0xd3, 0x40, 0x1d, 0x68, 0x52, 0x9f,
	0x6c, 0xa5, 0xbc, 0x20, 0x83, 0x90, 0x52, 0x18, 0x13, 0x3c, 0x4c, 0x55, 0x86, 0xfa, 0x80, 0x2a,
	0x45, 0x5e, 0x7d, 0x67, 0x2c, 0x35, 0xb6, 0x15, 0x3b, 0x85, 0xfd, 0x02, 0x9f, 0xc0, 0x23, 0x8f,
	0x7c, 0x0a, 0x9f, 0xc2, 0x57, 0xa0, 0xc6, 0xb1, 0xa1, 0x42, 0xdd, 0xfa, 0x76, 0xe2, 0x7b, 0xce,
	0xb9, 0xf7, 0xf8, 0x3a, 0xc9, 0x84, 0x2b, 0xc5, 0x97, 0x80, 0x29, 0x33, 0xd8, 0xc1, 0x35, 0x5a,
	0xe5, 0xb8, 0x01, 0xa3, 0xda, 0x66, 0x01, 0x06, 0x2f, 0x68, 0xad, 0xa9, 0xe0, 0xb2, 0x82, 0xaf,
	0x16, 0xa4, 0x11, 0x4a, 0x56, 0x06, 0xac, 0x15, 0x92, 0x23, 0xdd, 0x28, 0xab, 0xd2, 0x63, 0x27,
	0x44, 0x94, 0x19, 0x14, 0x3c, 0xd0, 0x2a, 0x47, 0xc1, 0xe3, 0xf0, 0xd5, 0xb6, 0x36, 0x20, 0xdb,
	0xda, 0xe0, 0xff, 0x9c, 0x2b, 0x06, 0x2b, 0xb1, 0x00, 0xd7, 0xe0, 0x70, 0xbc, 0xab, 0xda, 0x5e,
	0x6b, 0xaf, 0x79, 0xd2, 0x6b, 0xba, 0xaf, 0xcb, 0xf6, 0x0a, 0x7f, 0x69, 0xa8, 0xd6, 0xd0, 0x98,
	0xbe, 0x7e, 0xe4, 0x3d, 0xb5, 0xc0, 0x54, 0x4a, 0x65, 0xa9, 0x15, 0x4a, 0xf6, 0xd5, 0xa7, 0xdf,
	0x07, 0x49, 0xf6, 0xa6, 0xcf, 0x7d, 0xea, 0xed, 0x2f, 0xdc, 0x6c, 0xe9, 0xb3, 0x64, 0xdf, 0x27,
	0xab, 0x24, 0xad, 0x21, 0x8b, 0x86, 0xd1, 0xe8, 0x7e, 0xf9, 0xd0, 0x1f, 0x9e, 0xd3, 0x1a, 0x52,
	0x48, 0x0e, 0x36, 0xe7, 0xca, 0xe2, 0x61, 0x34, 0x3a, 0x18, 0xbf, 0x46, 0xdb, 0x6e, 0xab, 0x0b,
	0x83, 0x42, 0xb7, 0x8f, 0xd7, 0x1a, 0x4e, 0x65, 0x5b, 0x6f, 0x9e, 0x94, 0xfb, 0xf0, 0xef, 0x67,
	0xfa, 0x22, 0xb9, 0xe7, 0xf7, 0x93, 0x0d, 0x86, 0xd1, 0xe8, 0xc1, 0xf8, 0xc8, 0x37, 0xf0, 0xc9,
	0xd1, 0x85, 0x6d, 0x84, 0xe4, 0x33, 0xba, 0x6c, 0xa1, 0x0c, 0xec, 0xf4, 0x3c, 0x79, 0xf4, 0x77,
	0xc0, 0x2b, 0x00, 0x56, 0x09, 0x0b, 0xb5, 0xc9, 0xee, 0x0c, 0x07, 0xb7, 0xba, 0xa4, 0x41, 0xf9,
	0x0e, 0x80, 0xbd, 0x5f, 0xeb, 0xd2, 0xcf, 0xc9, 0x9e, 0x5b, 0x5a, 0x76, 0xb7, 0x0b, 0x3a, 0xdd,
	0x35, 0x68, 0x7f, 0xad, 0x6f, 0x3b, 0xf1, 0x66, 0xe2, 0x8d, 0x52, 0xd9, 0xfb, 0x4f, 0xbe, 0xc5,
	0xc9, 0xc9, 0x42, 0xd5, 0xe8, 0xd6, 0x67, 0x37, 0x79, 0xbc, 0x6d, 0x87, 0xd3, 0x75, 0xaa, 0x69,
	0xf4, 0xe9, 0x43, 0xef, 0xc1, 0xd5, 0x92, 0x4a, 0x8e, 0x54, 0xc3, 0x31, 0x07, 0xd9, 0x65, 0xf6,
	0x2f, 0x4d, 0x0b, 0x73, 0xc3, 0xdf, 0xf1, 0x32, 0xa0, 0x1f, 0xf1, 0xe0, 0xac, 0x28, 0x7e, 0xc6,
	0xc7, 0x67, 0xce, 0xb2, 0x60, 0x06, 0x39, 0xb8, 0x46, 0xb3, 0x1c, 0x95, 0x9e, 0xf9, 0xcb, 0x73,
	0xe6, 0x05, 0x33, 0xf3, 0xc0, 0x99, 0xcf, 0xf2, 0x79, 0xe0, 0xfc, 0x8e, 0x4f, 0x5c, 0x81, 0x90,
	0x82, 0x19, 0x42, 0x02, 0x8b, 0x90, 0x59, 0x4e, 0x48, 0xe0, 0x5d, 0xee, 0x75, 0xc3, 0x3e, 0xff,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x40, 0x38, 0xd0, 0xe6, 0xc9, 0x03, 0x00, 0x00,
}