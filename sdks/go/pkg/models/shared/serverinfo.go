// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ServerInfo - Server information
type ServerInfo struct {
	Version string `json:"version"`
}

func (o *ServerInfo) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}