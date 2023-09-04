/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * CreateClientResponse - Client created
 */

public class CreateClientResponse {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("data")
    public Client data;

    public CreateClientResponse withData(Client data) {
        this.data = data;
        return this;
    }
    
    public CreateClientResponse(){}
}