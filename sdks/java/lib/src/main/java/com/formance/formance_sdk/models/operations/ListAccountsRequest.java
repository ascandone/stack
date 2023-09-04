/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;


public class ListAccountsRequest {
    /**
     * Filter accounts by address pattern (regular expression placed between ^ and $).
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=address")
    public String address;

    public ListAccountsRequest withAddress(String address) {
        this.address = address;
        return this;
    }
    
    /**
     * Parameter used in pagination requests. Maximum page size is set to 15.
     * Set to the value of next for the next page of results.
     * Set to the value of previous for the previous page of results.
     * No other parameters can be set when this parameter is set.
     * 
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=cursor")
    public String cursor;

    public ListAccountsRequest withCursor(String cursor) {
        this.cursor = cursor;
        return this;
    }
    
    /**
     * Name of the ledger.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=ledger")
    public String ledger;

    public ListAccountsRequest withLedger(String ledger) {
        this.ledger = ledger;
        return this;
    }
    
    /**
     * Filter accounts by metadata key value pairs. Nested objects can be used like this -&gt; metadata[key]=value1&amp;metadata[a.nested.key]=value2
     */
    @SpeakeasyMetadata("queryParam:style=deepObject,explode=true,name=metadata")
    public java.util.Map<String, String> metadata;

    public ListAccountsRequest withMetadata(java.util.Map<String, String> metadata) {
        this.metadata = metadata;
        return this;
    }
    
    /**
     * The maximum number of results to return per page.
     * 
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=pageSize")
    public Long pageSize;

    public ListAccountsRequest withPageSize(Long pageSize) {
        this.pageSize = pageSize;
        return this;
    }
    
    public ListAccountsRequest(@JsonProperty("ledger") String ledger) {
        this.ledger = ledger;
  }
}