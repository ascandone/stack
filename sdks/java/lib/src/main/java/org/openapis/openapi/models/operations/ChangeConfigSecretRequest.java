/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.openapis.openapi.utils.SpeakeasyMetadata;

public class ChangeConfigSecretRequest {
    @SpeakeasyMetadata("request:mediaType=application/json")
    public org.openapis.openapi.models.shared.ConfigChangeSecret configChangeSecret;

    public ChangeConfigSecretRequest withConfigChangeSecret(org.openapis.openapi.models.shared.ConfigChangeSecret configChangeSecret) {
        this.configChangeSecret = configChangeSecret;
        return this;
    }
    
    /**
     * Config ID
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=id")
    public String id;

    public ChangeConfigSecretRequest withId(String id) {
        this.id = id;
        return this;
    }
    
    public ChangeConfigSecretRequest(@JsonProperty("id") String id) {
        this.id = id;
  }
}