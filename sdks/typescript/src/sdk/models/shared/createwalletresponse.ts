/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Wallet } from "./wallet";
import { Expose, Type } from "class-transformer";

/**
 * Wallet created
 */
export class CreateWalletResponse extends SpeakeasyBase {
    @SpeakeasyMetadata()
    @Expose({ name: "data" })
    @Type(() => Wallet)
    data: Wallet;
}