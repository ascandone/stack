# CreateTransactionRequest


## Fields

| Field                                                                                                                                                                                  | Type                                                                                                                                                                                   | Required                                                                                                                                                                               | Description                                                                                                                                                                            | Example                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `idempotencyKey`                                                                                                                                                                       | *String*                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                     | Use an idempotency key                                                                                                                                                                 |                                                                                                                                                                                        |
| `postTransaction`                                                                                                                                                                      | [com.formance.formance_sdk.models.shared.PostTransaction](../../models/shared/PostTransaction.md)                                                                                      | :heavy_check_mark:                                                                                                                                                                     | The request body must contain at least one of the following objects:<br/>  - `postings`: suitable for simple transactions<br/>  - `script`: enabling more complex transactions with Numscript<br/> |                                                                                                                                                                                        |
| `async`                                                                                                                                                                                | *Boolean*                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                     | Set async mode.                                                                                                                                                                        | true                                                                                                                                                                                   |
| `dryRun`                                                                                                                                                                               | *Boolean*                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                     | Set the dryRun mode. dry run mode doesn't add the logs to the database or publish a message to the message broker.                                                                     | true                                                                                                                                                                                   |
| `ledger`                                                                                                                                                                               | *String*                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                     | Name of the ledger.                                                                                                                                                                    | ledger001                                                                                                                                                                              |