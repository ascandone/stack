/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class ListTriggersOccurrencesResponse {
    
    public String contentType;

    public ListTriggersOccurrencesResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * General error
     */
    
    public com.formance.formance_sdk.models.shared.Error error;

    public ListTriggersOccurrencesResponse withError(com.formance.formance_sdk.models.shared.Error error) {
        this.error = error;
        return this;
    }
    
    /**
     * List of triggers occurrences
     */
    
    public com.formance.formance_sdk.models.shared.ListTriggersOccurrencesResponse listTriggersOccurrencesResponse;

    public ListTriggersOccurrencesResponse withListTriggersOccurrencesResponse(com.formance.formance_sdk.models.shared.ListTriggersOccurrencesResponse listTriggersOccurrencesResponse) {
        this.listTriggersOccurrencesResponse = listTriggersOccurrencesResponse;
        return this;
    }
    
    
    public Integer statusCode;

    public ListTriggersOccurrencesResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;

    public ListTriggersOccurrencesResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    public ListTriggersOccurrencesResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
        this.contentType = contentType;
        this.statusCode = statusCode;
  }
}
