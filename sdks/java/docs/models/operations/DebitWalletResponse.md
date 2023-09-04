# DebitWalletResponse


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `contentType`                                                                                                            | *String*                                                                                                                 | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `debitWalletResponse`                                                                                                    | [com.formance.formance_sdk.models.shared.DebitWalletResponse](../../models/shared/DebitWalletResponse.md)                | :heavy_minus_sign:                                                                                                       | Wallet successfully debited as a pending hold                                                                            |
| `statusCode`                                                                                                             | *Integer*                                                                                                                | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `rawResponse`                                                                                                            | [HttpResponse<byte[]>](https://docs.oracle.com/en/java/javase/11/docs/api/java.net.http/java/net/http/HttpResponse.html) | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `walletsErrorResponse`                                                                                                   | [com.formance.formance_sdk.models.shared.WalletsErrorResponse](../../models/shared/WalletsErrorResponse.md)              | :heavy_minus_sign:                                                                                                       | Error                                                                                                                    |