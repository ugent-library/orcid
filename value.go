package orcid

import "time"

type StringValue struct {
	Value string `json:"value,omitempty"`
}

type TimeValue struct {
	// Value time.Time `json:"value,omitempty"`
	Value time.Time `json:"-"`
}

func String(v string) *StringValue {
	return &StringValue{Value: v}
}

func Time(v time.Time) *TimeValue {
	return &TimeValue{Value: v}
}
