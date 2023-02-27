/*
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * The version of the OpenAPI document: 1.0.0-20230225
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.formance.formance.model;

import java.util.Objects;
import java.util.Arrays;
import com.formance.formance.model.Posting;
import com.formance.formance.model.WalletsVolume;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * WalletsTransaction
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class WalletsTransaction {
  public static final String SERIALIZED_NAME_LEDGER = "ledger";
  @SerializedName(SERIALIZED_NAME_LEDGER)
  private String ledger;

  public static final String SERIALIZED_NAME_TIMESTAMP = "timestamp";
  @SerializedName(SERIALIZED_NAME_TIMESTAMP)
  private OffsetDateTime timestamp;

  public static final String SERIALIZED_NAME_POSTINGS = "postings";
  @SerializedName(SERIALIZED_NAME_POSTINGS)
  private List<Posting> postings = new ArrayList<>();

  public static final String SERIALIZED_NAME_REFERENCE = "reference";
  @SerializedName(SERIALIZED_NAME_REFERENCE)
  private String reference;

  public static final String SERIALIZED_NAME_METADATA = "metadata";
  @SerializedName(SERIALIZED_NAME_METADATA)
  private Map<String, Object> metadata = new HashMap<>();

  public static final String SERIALIZED_NAME_TXID = "txid";
  @SerializedName(SERIALIZED_NAME_TXID)
  private Long txid;

  public static final String SERIALIZED_NAME_PRE_COMMIT_VOLUMES = "preCommitVolumes";
  @SerializedName(SERIALIZED_NAME_PRE_COMMIT_VOLUMES)
  private Map<String, Map<String, WalletsVolume>> preCommitVolumes = new HashMap<>();

  public static final String SERIALIZED_NAME_POST_COMMIT_VOLUMES = "postCommitVolumes";
  @SerializedName(SERIALIZED_NAME_POST_COMMIT_VOLUMES)
  private Map<String, Map<String, WalletsVolume>> postCommitVolumes = new HashMap<>();

  public WalletsTransaction() {
  }

  public WalletsTransaction ledger(String ledger) {
    
    this.ledger = ledger;
    return this;
  }

   /**
   * Get ledger
   * @return ledger
  **/
  @javax.annotation.Nullable

  public String getLedger() {
    return ledger;
  }


  public void setLedger(String ledger) {
    this.ledger = ledger;
  }


  public WalletsTransaction timestamp(OffsetDateTime timestamp) {
    
    this.timestamp = timestamp;
    return this;
  }

   /**
   * Get timestamp
   * @return timestamp
  **/
  @javax.annotation.Nonnull

  public OffsetDateTime getTimestamp() {
    return timestamp;
  }


  public void setTimestamp(OffsetDateTime timestamp) {
    this.timestamp = timestamp;
  }


  public WalletsTransaction postings(List<Posting> postings) {
    
    this.postings = postings;
    return this;
  }

  public WalletsTransaction addPostingsItem(Posting postingsItem) {
    this.postings.add(postingsItem);
    return this;
  }

   /**
   * Get postings
   * @return postings
  **/
  @javax.annotation.Nonnull

  public List<Posting> getPostings() {
    return postings;
  }


  public void setPostings(List<Posting> postings) {
    this.postings = postings;
  }


  public WalletsTransaction reference(String reference) {
    
    this.reference = reference;
    return this;
  }

   /**
   * Get reference
   * @return reference
  **/
  @javax.annotation.Nullable

  public String getReference() {
    return reference;
  }


  public void setReference(String reference) {
    this.reference = reference;
  }


  public WalletsTransaction metadata(Map<String, Object> metadata) {
    
    this.metadata = metadata;
    return this;
  }

  public WalletsTransaction putMetadataItem(String key, Object metadataItem) {
    if (this.metadata == null) {
      this.metadata = new HashMap<>();
    }
    this.metadata.put(key, metadataItem);
    return this;
  }

   /**
   * Metadata associated with the wallet.
   * @return metadata
  **/
  @javax.annotation.Nullable

  public Map<String, Object> getMetadata() {
    return metadata;
  }


  public void setMetadata(Map<String, Object> metadata) {
    this.metadata = metadata;
  }


  public WalletsTransaction txid(Long txid) {
    
    this.txid = txid;
    return this;
  }

   /**
   * Get txid
   * minimum: 0
   * @return txid
  **/
  @javax.annotation.Nonnull

  public Long getTxid() {
    return txid;
  }


  public void setTxid(Long txid) {
    this.txid = txid;
  }


  public WalletsTransaction preCommitVolumes(Map<String, Map<String, WalletsVolume>> preCommitVolumes) {
    
    this.preCommitVolumes = preCommitVolumes;
    return this;
  }

  public WalletsTransaction putPreCommitVolumesItem(String key, Map<String, WalletsVolume> preCommitVolumesItem) {
    if (this.preCommitVolumes == null) {
      this.preCommitVolumes = new HashMap<>();
    }
    this.preCommitVolumes.put(key, preCommitVolumesItem);
    return this;
  }

   /**
   * Get preCommitVolumes
   * @return preCommitVolumes
  **/
  @javax.annotation.Nullable

  public Map<String, Map<String, WalletsVolume>> getPreCommitVolumes() {
    return preCommitVolumes;
  }


  public void setPreCommitVolumes(Map<String, Map<String, WalletsVolume>> preCommitVolumes) {
    this.preCommitVolumes = preCommitVolumes;
  }


  public WalletsTransaction postCommitVolumes(Map<String, Map<String, WalletsVolume>> postCommitVolumes) {
    
    this.postCommitVolumes = postCommitVolumes;
    return this;
  }

  public WalletsTransaction putPostCommitVolumesItem(String key, Map<String, WalletsVolume> postCommitVolumesItem) {
    if (this.postCommitVolumes == null) {
      this.postCommitVolumes = new HashMap<>();
    }
    this.postCommitVolumes.put(key, postCommitVolumesItem);
    return this;
  }

   /**
   * Get postCommitVolumes
   * @return postCommitVolumes
  **/
  @javax.annotation.Nullable

  public Map<String, Map<String, WalletsVolume>> getPostCommitVolumes() {
    return postCommitVolumes;
  }


  public void setPostCommitVolumes(Map<String, Map<String, WalletsVolume>> postCommitVolumes) {
    this.postCommitVolumes = postCommitVolumes;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    WalletsTransaction walletsTransaction = (WalletsTransaction) o;
    return Objects.equals(this.ledger, walletsTransaction.ledger) &&
        Objects.equals(this.timestamp, walletsTransaction.timestamp) &&
        Objects.equals(this.postings, walletsTransaction.postings) &&
        Objects.equals(this.reference, walletsTransaction.reference) &&
        Objects.equals(this.metadata, walletsTransaction.metadata) &&
        Objects.equals(this.txid, walletsTransaction.txid) &&
        Objects.equals(this.preCommitVolumes, walletsTransaction.preCommitVolumes) &&
        Objects.equals(this.postCommitVolumes, walletsTransaction.postCommitVolumes);
  }

  @Override
  public int hashCode() {
    return Objects.hash(ledger, timestamp, postings, reference, metadata, txid, preCommitVolumes, postCommitVolumes);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class WalletsTransaction {\n");
    sb.append("    ledger: ").append(toIndentedString(ledger)).append("\n");
    sb.append("    timestamp: ").append(toIndentedString(timestamp)).append("\n");
    sb.append("    postings: ").append(toIndentedString(postings)).append("\n");
    sb.append("    reference: ").append(toIndentedString(reference)).append("\n");
    sb.append("    metadata: ").append(toIndentedString(metadata)).append("\n");
    sb.append("    txid: ").append(toIndentedString(txid)).append("\n");
    sb.append("    preCommitVolumes: ").append(toIndentedString(preCommitVolumes)).append("\n");
    sb.append("    postCommitVolumes: ").append(toIndentedString(postCommitVolumes)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

