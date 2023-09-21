# GetTransactionResponse


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `content_type`                                                                           | *str*                                                                                    | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `error_response`                                                                         | [Optional[shared.ErrorResponse]](../../models/shared/errorresponse.md)                   | :heavy_minus_sign:                                                                       | Error                                                                                    |
| `get_transaction_response`                                                               | [Optional[shared.GetTransactionResponse]](../../models/shared/gettransactionresponse.md) | :heavy_minus_sign:                                                                       | OK                                                                                       |
| `status_code`                                                                            | *int*                                                                                    | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `raw_response`                                                                           | [requests.Response](https://requests.readthedocs.io/en/latest/api/#requests.Response)    | :heavy_minus_sign:                                                                       | N/A                                                                                      |