// Code generated by protoc-gen-go.
// source: metrics.proto
// DO NOT EDIT!

/*
Package io_prometheus_client is a generated protocol buffer package.

It is generated from these files:
	metrics.proto

It has these top-level messages:
	LabelPair
	Gauge
	Counter
	Quantile
	Summary
	Custom
	Metric
	MetricFamily
*/
package io_prometheus_client

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type MetricType int32

const (
	MetricType_COUNTER MetricType = 0
	MetricType_GAUGE   MetricType = 1
	MetricType_SUMMARY MetricType = 2
	MetricType_CUSTOM  MetricType = 3
)

var MetricType_name = map[int32]string{
	0: "COUNTER",
	1: "GAUGE",
	2: "SUMMARY",
	3: "CUSTOM",
}
var MetricType_value = map[string]int32{
	"COUNTER": 0,
	"GAUGE":   1,
	"SUMMARY": 2,
	"CUSTOM":  3,
}

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}
func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}
func (x *MetricType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetricType_value, data, "MetricType")
	if err != nil {
		return err
	}
	*x = MetricType(value)
	return nil
}

type LabelPair struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LabelPair) Reset()         { *m = LabelPair{} }
func (m *LabelPair) String() string { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()    {}

func (m *LabelPair) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LabelPair) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Gauge struct {
	Value            *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}

func (m *Gauge) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Counter struct {
	Value            *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}

func (m *Counter) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Quantile struct {
	Quantile         *float64 `protobuf:"fixed64,1,opt,name=quantile" json:"quantile,omitempty"`
	Value            *float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Quantile) Reset()         { *m = Quantile{} }
func (m *Quantile) String() string { return proto.CompactTextString(m) }
func (*Quantile) ProtoMessage()    {}

func (m *Quantile) GetQuantile() float64 {
	if m != nil && m.Quantile != nil {
		return *m.Quantile
	}
	return 0
}

func (m *Quantile) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Summary struct {
	SampleCount      *uint64     `protobuf:"varint,1,opt,name=sample_count" json:"sample_count,omitempty"`
	SampleSum        *float64    `protobuf:"fixed64,2,opt,name=sample_sum" json:"sample_sum,omitempty"`
	Quantile         []*Quantile `protobuf:"bytes,3,rep,name=quantile" json:"quantile,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}

func (m *Summary) GetSampleCount() uint64 {
	if m != nil && m.SampleCount != nil {
		return *m.SampleCount
	}
	return 0
}

func (m *Summary) GetSampleSum() float64 {
	if m != nil && m.SampleSum != nil {
		return *m.SampleSum
	}
	return 0
}

func (m *Summary) GetQuantile() []*Quantile {
	if m != nil {
		return m.Quantile
	}
	return nil
}

type Custom struct {
	Value            *float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Custom) Reset()         { *m = Custom{} }
func (m *Custom) String() string { return proto.CompactTextString(m) }
func (*Custom) ProtoMessage()    {}

func (m *Custom) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Metric struct {
	Label            []*LabelPair `protobuf:"bytes,1,rep,name=label" json:"label,omitempty"`
	Gauge            *Gauge       `protobuf:"bytes,2,opt,name=gauge" json:"gauge,omitempty"`
	Counter          *Counter     `protobuf:"bytes,3,opt,name=counter" json:"counter,omitempty"`
	Summary          *Summary     `protobuf:"bytes,4,opt,name=summary" json:"summary,omitempty"`
	Custom           *Custom      `protobuf:"bytes,5,opt,name=custom" json:"custom,omitempty"`
	TimestampMs      *int64       `protobuf:"varint,6,opt,name=timestamp_ms" json:"timestamp_ms,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}

func (m *Metric) GetLabel() []*LabelPair {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *Metric) GetGauge() *Gauge {
	if m != nil {
		return m.Gauge
	}
	return nil
}

func (m *Metric) GetCounter() *Counter {
	if m != nil {
		return m.Counter
	}
	return nil
}

func (m *Metric) GetSummary() *Summary {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *Metric) GetCustom() *Custom {
	if m != nil {
		return m.Custom
	}
	return nil
}

func (m *Metric) GetTimestampMs() int64 {
	if m != nil && m.TimestampMs != nil {
		return *m.TimestampMs
	}
	return 0
}

type MetricFamily struct {
	Name             *string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Help             *string     `protobuf:"bytes,2,opt,name=help" json:"help,omitempty"`
	Type             *MetricType `protobuf:"varint,3,opt,name=type,enum=io.prometheus.client.MetricType" json:"type,omitempty"`
	Metric           []*Metric   `protobuf:"bytes,4,rep,name=metric" json:"metric,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *MetricFamily) Reset()         { *m = MetricFamily{} }
func (m *MetricFamily) String() string { return proto.CompactTextString(m) }
func (*MetricFamily) ProtoMessage()    {}

func (m *MetricFamily) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MetricFamily) GetHelp() string {
	if m != nil && m.Help != nil {
		return *m.Help
	}
	return ""
}

func (m *MetricFamily) GetType() MetricType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MetricType_COUNTER
}

func (m *MetricFamily) GetMetric() []*Metric {
	if m != nil {
		return m.Metric
	}
	return nil
}

func init() {
	proto.RegisterEnum("io.prometheus.client.MetricType", MetricType_name, MetricType_value)
}
