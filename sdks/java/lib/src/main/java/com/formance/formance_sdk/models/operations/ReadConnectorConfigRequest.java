/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;

public class ReadConnectorConfigRequest {
    /**
     * The name of the connector.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=connector")
    public com.formance.formance_sdk.models.shared.Connector connector;

    public ReadConnectorConfigRequest withConnector(com.formance.formance_sdk.models.shared.Connector connector) {
        this.connector = connector;
        return this;
    }
    
    public ReadConnectorConfigRequest(@JsonProperty("connector") com.formance.formance_sdk.models.shared.Connector connector) {
        this.connector = connector;
  }
}