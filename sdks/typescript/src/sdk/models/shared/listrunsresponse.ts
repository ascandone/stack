/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { WorkflowInstance } from "./workflowinstance";
import { Expose, Type } from "class-transformer";

/**
 * List of workflow instances
 */
export class ListRunsResponse extends SpeakeasyBase {
    @SpeakeasyMetadata({ elemType: WorkflowInstance })
    @Expose({ name: "data" })
    @Type(() => WorkflowInstance)
    data: WorkflowInstance[];
}