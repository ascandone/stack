// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConfigChangeSecret struct {
	Secret string `json:"secret"`
}

func (o *ConfigChangeSecret) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}
