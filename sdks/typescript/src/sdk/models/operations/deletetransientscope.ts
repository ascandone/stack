/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { AxiosResponse } from "axios";

export class DeleteTransientScopeRequest extends SpeakeasyBase {
  /**
   * Scope ID
   */
  @SpeakeasyMetadata({
    data: "pathParam, style=simple;explode=false;name=scopeId",
  })
  scopeId: string;

  /**
   * Transient scope ID
   */
  @SpeakeasyMetadata({
    data: "pathParam, style=simple;explode=false;name=transientScopeId",
  })
  transientScopeId: string;
}

export class DeleteTransientScopeResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  rawResponse?: AxiosResponse;
}
