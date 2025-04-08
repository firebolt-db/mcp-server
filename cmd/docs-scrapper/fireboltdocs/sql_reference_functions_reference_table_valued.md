## [](#table-valued-functions)Table-valued functions

A table-valued function (TVF) returns a set of rows. You can use a table-valued function anywhere you can use a table. For example, the `generate_series` TVF generates a range of numbers as rows.

The following code example generates a series of numbers from 1 to 100 and selects each number as x from the resulting series:

```
SELECT x FROM generate_series(1, 100) r(x);
```

You can use the `list_objects` TVF to explore files on Amazon S3. Use functions such as `read_parquet` to read data from Amazon S3.

* * *

- [GENERATE\_SERIES](/sql_reference/functions-reference/table-valued/generate-series.html)
- [LIST\_OBJECTS](/sql_reference/functions-reference/table-valued/list-objects.html)
- [READ\_CSV](/sql_reference/functions-reference/table-valued/read_csv.html)
- [READ\_PARQUET](/sql_reference/functions-reference/table-valued/read_parquet.html)