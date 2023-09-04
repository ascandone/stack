/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";

export class GetInstanceStageHistoryRequest extends SpeakeasyBase {
    /**
     * The instance id
     */
    @SpeakeasyMetadata({ data: "pathParam, style=simple;explode=false;name=instanceID" })
    instanceID: string;

    /**
     * The stage number
     */
    @SpeakeasyMetadata({ data: "pathParam, style=simple;explode=false;name=number" })
    number: number;
}

export class GetInstanceStageHistoryResponse extends SpeakeasyBase {
    @SpeakeasyMetadata()
    contentType: string;

    /**
     * General error
     */
    @SpeakeasyMetadata()
    error?: shared.ErrorT;

    /**
     * The workflow instance stage history
     */
    @SpeakeasyMetadata()
    getWorkflowInstanceHistoryStageResponse?: shared.GetWorkflowInstanceHistoryStageResponse;

    @SpeakeasyMetadata()
    statusCode: number;

    @SpeakeasyMetadata()
    rawResponse?: AxiosResponse;
}
