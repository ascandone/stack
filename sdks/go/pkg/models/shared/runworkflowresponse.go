// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RunWorkflowResponse - The workflow instance
type RunWorkflowResponse struct {
	Data WorkflowInstance `json:"data"`
}

func (o *RunWorkflowResponse) GetData() WorkflowInstance {
	if o == nil {
		return WorkflowInstance{}
	}
	return o.Data
}
