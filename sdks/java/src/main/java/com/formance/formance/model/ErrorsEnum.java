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
import com.google.gson.annotations.SerializedName;

import java.io.IOException;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;

/**
 * Gets or Sets ErrorsEnum
 */
@JsonAdapter(ErrorsEnum.Adapter.class)
public enum ErrorsEnum {
  
  INTERNAL("INTERNAL"),
  
  INSUFFICIENT_FUND("INSUFFICIENT_FUND"),
  
  VALIDATION("VALIDATION"),
  
  CONFLICT("CONFLICT"),
  
  NO_SCRIPT("NO_SCRIPT"),
  
  COMPILATION_FAILED("COMPILATION_FAILED"),
  
  METADATA_OVERRIDE("METADATA_OVERRIDE");

  private String value;

  ErrorsEnum(String value) {
    this.value = value;
  }

  public String getValue() {
    return value;
  }

  @Override
  public String toString() {
    return String.valueOf(value);
  }

  public static ErrorsEnum fromValue(String value) {
    for (ErrorsEnum b : ErrorsEnum.values()) {
      if (b.value.equals(value)) {
        return b;
      }
    }
    throw new IllegalArgumentException("Unexpected value '" + value + "'");
  }

  public static class Adapter extends TypeAdapter<ErrorsEnum> {
    @Override
    public void write(final JsonWriter jsonWriter, final ErrorsEnum enumeration) throws IOException {
      jsonWriter.value(enumeration.getValue());
    }

    @Override
    public ErrorsEnum read(final JsonReader jsonReader) throws IOException {
      String value = jsonReader.nextString();
      return ErrorsEnum.fromValue(value);
    }
  }
}

