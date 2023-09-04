// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type Attempt struct {
	Config         WebhooksConfig `json:"config"`
	CreatedAt      time.Time      `json:"createdAt"`
	ID             string         `json:"id"`
	NextRetryAfter *time.Time     `json:"nextRetryAfter,omitempty"`
	Payload        string         `json:"payload"`
	RetryAttempt   int64          `json:"retryAttempt"`
	Status         string         `json:"status"`
	StatusCode     int64          `json:"statusCode"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	WebhookID      string         `json:"webhookID"`
}

func (o *Attempt) GetConfig() WebhooksConfig {
	if o == nil {
		return WebhooksConfig{}
	}
	return o.Config
}

func (o *Attempt) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Attempt) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Attempt) GetNextRetryAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.NextRetryAfter
}

func (o *Attempt) GetPayload() string {
	if o == nil {
		return ""
	}
	return o.Payload
}

func (o *Attempt) GetRetryAttempt() int64 {
	if o == nil {
		return 0
	}
	return o.RetryAttempt
}

func (o *Attempt) GetStatus() string {
	if o == nil {
		return ""
	}
	return o.Status
}

func (o *Attempt) GetStatusCode() int64 {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Attempt) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Attempt) GetWebhookID() string {
	if o == nil {
		return ""
	}
	return o.WebhookID
}