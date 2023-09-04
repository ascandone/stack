/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;


public class CreateTransactionRequest {
    /**
     * Use an idempotency key
     */
    @SpeakeasyMetadata("header:style=simple,explode=false,name=Idempotency-Key")
    public String idempotencyKey;

    public CreateTransactionRequest withIdempotencyKey(String idempotencyKey) {
        this.idempotencyKey = idempotencyKey;
        return this;
    }
    
    /**
     * The request body must contain at least one of the following objects:
     *   - `postings`: suitable for simple transactions
     *   - `script`: enabling more complex transactions with Numscript
     * 
     */
    @SpeakeasyMetadata("request:mediaType=application/json")
    public com.formance.formance_sdk.models.shared.PostTransaction postTransaction;

    public CreateTransactionRequest withPostTransaction(com.formance.formance_sdk.models.shared.PostTransaction postTransaction) {
        this.postTransaction = postTransaction;
        return this;
    }
    
    /**
     * Set async mode.
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=async")
    public Boolean async;

    public CreateTransactionRequest withAsync(Boolean async) {
        this.async = async;
        return this;
    }
    
    /**
     * Set the dryRun mode. dry run mode doesn't add the logs to the database or publish a message to the message broker.
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=dryRun")
    public Boolean dryRun;

    public CreateTransactionRequest withDryRun(Boolean dryRun) {
        this.dryRun = dryRun;
        return this;
    }
    
    /**
     * Name of the ledger.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=ledger")
    public String ledger;

    public CreateTransactionRequest withLedger(String ledger) {
        this.ledger = ledger;
        return this;
    }
    
    public CreateTransactionRequest(@JsonProperty("PostTransaction") com.formance.formance_sdk.models.shared.PostTransaction postTransaction, @JsonProperty("ledger") String ledger) {
        this.postTransaction = postTransaction;
        this.ledger = ledger;
  }
}
