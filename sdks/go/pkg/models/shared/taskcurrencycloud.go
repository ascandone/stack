// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type TaskCurrencyCloudDescriptor struct {
	Name *string `json:"name,omitempty"`
}

func (o *TaskCurrencyCloudDescriptor) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type TaskCurrencyCloudState struct {
}

type TaskCurrencyCloud struct {
	ConnectorID string                      `json:"connectorId"`
	CreatedAt   time.Time                   `json:"createdAt"`
	Descriptor  TaskCurrencyCloudDescriptor `json:"descriptor"`
	Error       *string                     `json:"error,omitempty"`
	ID          string                      `json:"id"`
	State       TaskCurrencyCloudState      `json:"state"`
	Status      PaymentStatus               `json:"status"`
	UpdatedAt   time.Time                   `json:"updatedAt"`
}

func (o *TaskCurrencyCloud) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *TaskCurrencyCloud) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *TaskCurrencyCloud) GetDescriptor() TaskCurrencyCloudDescriptor {
	if o == nil {
		return TaskCurrencyCloudDescriptor{}
	}
	return o.Descriptor
}

func (o *TaskCurrencyCloud) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *TaskCurrencyCloud) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TaskCurrencyCloud) GetState() TaskCurrencyCloudState {
	if o == nil {
		return TaskCurrencyCloudState{}
	}
	return o.State
}

func (o *TaskCurrencyCloud) GetStatus() PaymentStatus {
	if o == nil {
		return PaymentStatus("")
	}
	return o.Status
}

func (o *TaskCurrencyCloud) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
