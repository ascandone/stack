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

public class TaskCurrencyCloud {
    @JsonProperty("connectorId")
    public String connectorId;

    public TaskCurrencyCloud withConnectorId(String connectorId) {
        this.connectorId = connectorId;
        return this;
    }
    
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("createdAt")
    public OffsetDateTime createdAt;

    public TaskCurrencyCloud withCreatedAt(OffsetDateTime createdAt) {
        this.createdAt = createdAt;
        return this;
    }
    
    @JsonProperty("descriptor")
    public TaskCurrencyCloudDescriptor descriptor;

    public TaskCurrencyCloud withDescriptor(TaskCurrencyCloudDescriptor descriptor) {
        this.descriptor = descriptor;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("error")
    public String error;

    public TaskCurrencyCloud withError(String error) {
        this.error = error;
        return this;
    }
    
    @JsonProperty("id")
    public String id;

    public TaskCurrencyCloud withId(String id) {
        this.id = id;
        return this;
    }
    
    @JsonProperty("state")
    public java.util.Map<String, Object> state;

    public TaskCurrencyCloud withState(java.util.Map<String, Object> state) {
        this.state = state;
        return this;
    }
    
    @JsonProperty("status")
    public PaymentStatus status;

    public TaskCurrencyCloud withStatus(PaymentStatus status) {
        this.status = status;
        return this;
    }
    
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("updatedAt")
    public OffsetDateTime updatedAt;

    public TaskCurrencyCloud withUpdatedAt(OffsetDateTime updatedAt) {
        this.updatedAt = updatedAt;
        return this;
    }
    
    public TaskCurrencyCloud(@JsonProperty("connectorId") String connectorId, @JsonProperty("createdAt") OffsetDateTime createdAt, @JsonProperty("descriptor") TaskCurrencyCloudDescriptor descriptor, @JsonProperty("id") String id, @JsonProperty("state") java.util.Map<String, Object> state, @JsonProperty("status") PaymentStatus status, @JsonProperty("updatedAt") OffsetDateTime updatedAt) {
        this.connectorId = connectorId;
        this.createdAt = createdAt;
        this.descriptor = descriptor;
        this.id = id;
        this.state = state;
        this.status = status;
        this.updatedAt = updatedAt;
  }
}
