# ExpandedTransaction


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `id`                                                                 | *int*                                                                | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `metadata`                                                           | dict[str, *str*]                                                     | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `post_commit_volumes`                                                | dict[str, dict[str, [Volume](../../models/shared/volume.md)]]        | :heavy_minus_sign:                                                   | N/A                                                                  |                                                                      |
| `postings`                                                           | list[[Posting](../../models/shared/posting.md)]                      | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `pre_commit_volumes`                                                 | dict[str, dict[str, [Volume](../../models/shared/volume.md)]]        | :heavy_minus_sign:                                                   | N/A                                                                  |                                                                      |
| `reference`                                                          | *Optional[str]*                                                      | :heavy_minus_sign:                                                   | N/A                                                                  | ref:001                                                              |
| `reverted`                                                           | *bool*                                                               | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `timestamp`                                                          | [date](https://docs.python.org/3/library/datetime.html#date-objects) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |