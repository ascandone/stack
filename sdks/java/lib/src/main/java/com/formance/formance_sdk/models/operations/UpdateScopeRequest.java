/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;


public class UpdateScopeRequest {
    @SpeakeasyMetadata("request:mediaType=application/json")
    public com.formance.formance_sdk.models.shared.UpdateScopeRequest updateScopeRequest;

    public UpdateScopeRequest withUpdateScopeRequest(com.formance.formance_sdk.models.shared.UpdateScopeRequest updateScopeRequest) {
        this.updateScopeRequest = updateScopeRequest;
        return this;
    }
    
    /**
     * Scope ID
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=scopeId")
    public String scopeId;

    public UpdateScopeRequest withScopeId(String scopeId) {
        this.scopeId = scopeId;
        return this;
    }
    
    public UpdateScopeRequest(@JsonProperty("scopeId") String scopeId) {
        this.scopeId = scopeId;
  }
}
