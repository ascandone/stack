/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export class WalletSubject extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "balance" })
  balance?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "identifier" })
  identifier: string;

  @SpeakeasyMetadata()
  @Expose({ name: "type" })
  type: string;
}
