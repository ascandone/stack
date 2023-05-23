/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * ServerInfo - Server information
 */
public class ServerInfo {
    @JsonProperty("version")
    public String version;

    public ServerInfo withVersion(String version) {
        this.version = version;
        return this;
    }
    
    public ServerInfo(@JsonProperty("version") String version) {
        this.version = version;
  }
}