/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class ReadTriggerResponse {
    
    public String contentType;

    public ReadTriggerResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * General error
     */
    
    public com.formance.formance_sdk.models.shared.Error error;

    public ReadTriggerResponse withError(com.formance.formance_sdk.models.shared.Error error) {
        this.error = error;
        return this;
    }
    
    /**
     * A specific trigger
     */
    
    public com.formance.formance_sdk.models.shared.ReadTriggerResponse readTriggerResponse;

    public ReadTriggerResponse withReadTriggerResponse(com.formance.formance_sdk.models.shared.ReadTriggerResponse readTriggerResponse) {
        this.readTriggerResponse = readTriggerResponse;
        return this;
    }
    
    
    public Integer statusCode;

    public ReadTriggerResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;

    public ReadTriggerResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    public ReadTriggerResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
        this.contentType = contentType;
        this.statusCode = statusCode;
  }
}
