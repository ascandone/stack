/*
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * The version of the OpenAPI document: 1.0.20230227
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.formance.formance.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;

/**
 * StageSendSourceAccount
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class StageSendSourceAccount {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  public static final String SERIALIZED_NAME_LEDGER = "ledger";
  @SerializedName(SERIALIZED_NAME_LEDGER)
  private String ledger;

  public StageSendSourceAccount() {
  }

  public StageSendSourceAccount id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nonnull

  public String getId() {
    return id;
  }


  public void setId(String id) {
    this.id = id;
  }


  public StageSendSourceAccount ledger(String ledger) {
    
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


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    StageSendSourceAccount stageSendSourceAccount = (StageSendSourceAccount) o;
    return Objects.equals(this.id, stageSendSourceAccount.id) &&
        Objects.equals(this.ledger, stageSendSourceAccount.ledger);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, ledger);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class StageSendSourceAccount {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    ledger: ").append(toIndentedString(ledger)).append("\n");
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

