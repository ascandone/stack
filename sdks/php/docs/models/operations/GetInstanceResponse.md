# GetInstanceResponse


## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `contentType`                                                                                                    | *string*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `error`                                                                                                          | [?\formance\stack\Models\Shared\Error](../../models/shared/Error.md)                                             | :heavy_minus_sign:                                                                                               | General error                                                                                                    |
| `getWorkflowInstanceResponse`                                                                                    | [?\formance\stack\Models\Shared\GetWorkflowInstanceResponse](../../models/shared/GetWorkflowInstanceResponse.md) | :heavy_minus_sign:                                                                                               | The workflow instance                                                                                            |
| `statusCode`                                                                                                     | *int*                                                                                                            | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `rawResponse`                                                                                                    | [\Psr\Http\Message\ResponseInterface](https://www.php-fig.org/psr/psr-7/#33-psrhttpmessageresponseinterface)     | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |