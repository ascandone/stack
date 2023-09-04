/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;


public class UpdateScopeResponse {
    
    public String contentType;

    public UpdateScopeResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    
    public Integer statusCode;

    public UpdateScopeResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;

    public UpdateScopeResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    /**
     * Updated scope
     */
    
    public com.formance.formance_sdk.models.shared.UpdateScopeResponse updateScopeResponse;

    public UpdateScopeResponse withUpdateScopeResponse(com.formance.formance_sdk.models.shared.UpdateScopeResponse updateScopeResponse) {
        this.updateScopeResponse = updateScopeResponse;
        return this;
    }
    
    public UpdateScopeResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
        this.contentType = contentType;
        this.statusCode = statusCode;
  }
}