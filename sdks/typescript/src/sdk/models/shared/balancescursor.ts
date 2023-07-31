/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { AccountBalance } from "./accountbalance";
import { Expose, Type } from "class-transformer";

export class BalancesCursorCursor extends SpeakeasyBase {
  @SpeakeasyMetadata({ elemType: AccountBalance })
  @Expose({ name: "data" })
  @Type(() => AccountBalance)
  data: AccountBalance[];

  @SpeakeasyMetadata()
  @Expose({ name: "hasMore" })
  hasMore: boolean;

  @SpeakeasyMetadata()
  @Expose({ name: "next" })
  next?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "pageSize" })
  pageSize: number;

  @SpeakeasyMetadata()
  @Expose({ name: "previous" })
  previous?: string;
}

/**
 * OK
 */
export class BalancesCursor extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "cursor" })
  @Type(() => BalancesCursorCursor)
  cursor: BalancesCursorCursor;
}