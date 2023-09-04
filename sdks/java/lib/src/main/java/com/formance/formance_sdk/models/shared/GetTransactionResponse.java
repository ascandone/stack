/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * GetTransactionResponse - OK
 */

public class GetTransactionResponse {
    @JsonProperty("data")
    public ExpandedTransaction data;

    public GetTransactionResponse withData(ExpandedTransaction data) {
        this.data = data;
        return this;
    }
    
    public GetTransactionResponse(@JsonProperty("data") ExpandedTransaction data) {
        this.data = data;
  }
}
