// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type StageStatus struct {
	Error        *string    `json:"error,omitempty"`
	InstanceID   string     `json:"instanceID"`
	Stage        float64    `json:"stage"`
	StartedAt    time.Time  `json:"startedAt"`
	TerminatedAt *time.Time `json:"terminatedAt,omitempty"`
}

func (o *StageStatus) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *StageStatus) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

func (o *StageStatus) GetStage() float64 {
	if o == nil {
		return 0.0
	}
	return o.Stage
}

func (o *StageStatus) GetStartedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartedAt
}

func (o *StageStatus) GetTerminatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TerminatedAt
}