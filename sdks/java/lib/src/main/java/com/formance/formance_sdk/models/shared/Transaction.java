/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.formance.formance_sdk.utils.DateTimeDeserializer;
import com.formance.formance_sdk.utils.DateTimeSerializer;
import java.time.OffsetDateTime;


public class Transaction {
    @JsonProperty("metadata")
    public java.util.Map<String, String> metadata;

    public Transaction withMetadata(java.util.Map<String, String> metadata) {
        this.metadata = metadata;
        return this;
    }
    
    @JsonProperty("postings")
    public Posting[] postings;

    public Transaction withPostings(Posting[] postings) {
        this.postings = postings;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("reference")
    public String reference;

    public Transaction withReference(String reference) {
        this.reference = reference;
        return this;
    }
    
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("timestamp")
    public OffsetDateTime timestamp;

    public Transaction withTimestamp(OffsetDateTime timestamp) {
        this.timestamp = timestamp;
        return this;
    }
    
    @JsonProperty("txid")
    public Long txid;

    public Transaction withTxid(Long txid) {
        this.txid = txid;
        return this;
    }
    
    public Transaction(@JsonProperty("metadata") java.util.Map<String, String> metadata, @JsonProperty("postings") Posting[] postings, @JsonProperty("timestamp") OffsetDateTime timestamp, @JsonProperty("txid") Long txid) {
        this.metadata = metadata;
        this.postings = postings;
        this.timestamp = timestamp;
        this.txid = txid;
  }
}
