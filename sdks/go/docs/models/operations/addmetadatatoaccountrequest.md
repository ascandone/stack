# AddMetadataToAccountRequest


## Fields

| Field                                                                                                               | Type                                                                                                                | Required                                                                                                            | Description                                                                                                         | Example                                                                                                             |
| ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `IdempotencyKey`                                                                                                    | **string*                                                                                                           | :heavy_minus_sign:                                                                                                  | Use an idempotency key                                                                                              |                                                                                                                     |
| `RequestBody`                                                                                                       | map[string]*string*                                                                                                 | :heavy_check_mark:                                                                                                  | metadata                                                                                                            |                                                                                                                     |
| `Address`                                                                                                           | *string*                                                                                                            | :heavy_check_mark:                                                                                                  | Exact address of the account. It must match the following regular expressions pattern:<br/>```<br/>^\w+(:\w+)*$<br/>```<br/> | users:001                                                                                                           |
| `DryRun`                                                                                                            | **bool*                                                                                                             | :heavy_minus_sign:                                                                                                  | Set the dry run mode. Dry run mode doesn't add the logs to the database or publish a message to the message broker. | true                                                                                                                |
| `Ledger`                                                                                                            | *string*                                                                                                            | :heavy_check_mark:                                                                                                  | Name of the ledger.                                                                                                 | ledger001                                                                                                           |