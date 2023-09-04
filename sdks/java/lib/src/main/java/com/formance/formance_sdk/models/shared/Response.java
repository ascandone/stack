/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * Response - Success
 */

public class Response {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("cursor")
    public ResponseCursor cursor;

    public Response withCursor(ResponseCursor cursor) {
        this.cursor = cursor;
        return this;
    }
    
    /**
     * The payload
     */
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("data")
    public ResponseData data;

    public Response withData(ResponseData data) {
        this.data = data;
        return this;
    }
    
    public Response(){}
}