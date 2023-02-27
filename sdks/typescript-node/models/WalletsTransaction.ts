/**
 * Formance Stack API
 * Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 
 *
 * OpenAPI spec version: 1.0.20230227
 * Contact: support@formance.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { Posting } from '../models/Posting';
import { WalletsVolume } from '../models/WalletsVolume';
import { HttpFile } from '../http/http';

export class WalletsTransaction {
    'ledger'?: string;
    'timestamp': Date;
    'postings': Array<Posting>;
    'reference'?: string;
    /**
    * Metadata associated with the wallet.
    */
    'metadata'?: { [key: string]: any; };
    'txid': number;
    'preCommitVolumes'?: { [key: string]: { [key: string]: WalletsVolume; }; };
    'postCommitVolumes'?: { [key: string]: { [key: string]: WalletsVolume; }; };

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "ledger",
            "baseName": "ledger",
            "type": "string",
            "format": ""
        },
        {
            "name": "timestamp",
            "baseName": "timestamp",
            "type": "Date",
            "format": "date-time"
        },
        {
            "name": "postings",
            "baseName": "postings",
            "type": "Array<Posting>",
            "format": ""
        },
        {
            "name": "reference",
            "baseName": "reference",
            "type": "string",
            "format": ""
        },
        {
            "name": "metadata",
            "baseName": "metadata",
            "type": "{ [key: string]: any; }",
            "format": ""
        },
        {
            "name": "txid",
            "baseName": "txid",
            "type": "number",
            "format": "int64"
        },
        {
            "name": "preCommitVolumes",
            "baseName": "preCommitVolumes",
            "type": "{ [key: string]: { [key: string]: WalletsVolume; }; }",
            "format": ""
        },
        {
            "name": "postCommitVolumes",
            "baseName": "postCommitVolumes",
            "type": "{ [key: string]: { [key: string]: WalletsVolume; }; }",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return WalletsTransaction.attributeTypeMap;
    }

    public constructor() {
    }
}

