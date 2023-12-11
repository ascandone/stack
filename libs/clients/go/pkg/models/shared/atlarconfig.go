// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
)

type AtlarConfig struct {
	// The access key used by the connector for authorizing requests to the Atlar API.
	// You can obtain it along with the associated secret from the Atlar dashboard.
	//
	AccessKey string `json:"accessKey"`
	// The base URL the client uses for making requests towards the Atlar API.
	//
	BaseURL *string `default:"https://api.atlar.com" json:"baseUrl"`
	Name    string  `json:"name"`
	// Number of items to fetch when querying paginated APIs.
	//
	PageSize *int64 `default:"25" json:"pageSize"`
	// The frequency at which the connector tries to fetch new Transaction objects from the Atlar API.
	//
	PollingPeriod *string `default:"120s" json:"pollingPeriod"`
	// The secret used by the connector for authorizing requests to the Atlar API.
	// You can obtain it along with the associated access key from the Atlar dashboard.
	//
	Secret string `json:"secret"`
}

func (a AtlarConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AtlarConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AtlarConfig) GetAccessKey() string {
	if o == nil {
		return ""
	}
	return o.AccessKey
}

func (o *AtlarConfig) GetBaseURL() *string {
	if o == nil {
		return nil
	}
	return o.BaseURL
}

func (o *AtlarConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AtlarConfig) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *AtlarConfig) GetPollingPeriod() *string {
	if o == nil {
		return nil
	}
	return o.PollingPeriod
}

func (o *AtlarConfig) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}