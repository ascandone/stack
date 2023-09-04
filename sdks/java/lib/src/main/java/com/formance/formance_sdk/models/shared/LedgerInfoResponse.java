/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * LedgerInfoResponse - OK
 */

public class LedgerInfoResponse {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("data")
    public LedgerInfo data;

    public LedgerInfoResponse withData(LedgerInfo data) {
        this.data = data;
        return this;
    }
    
    public LedgerInfoResponse(){}
}