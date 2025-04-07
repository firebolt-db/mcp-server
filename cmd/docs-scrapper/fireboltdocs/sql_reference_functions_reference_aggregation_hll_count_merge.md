# [](#hll_count_merge)HLL\_COUNT\_MERGE

Merges one or more HLL++ sketches that were previously built using the aggregate function [HLL\_COUNT\_BUILD](/sql_reference/functions-reference/aggregation/hll-count-build.html) into a new sketch.

Each sketch must be built on the same type and the same precision. Attempts to merge sketches for different types or precisions results in an error. For example, you cannot merge a sketch built from `INTEGER` data with one built from `TEXT` data, or a sketch built with precision 13 and a sketch built with precision 14.

## [](#syntax)Syntax

```
HLL_COUNT_MERGE(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` HLL++ sketch in a valid format, e.g. the output of the [HLL\_COUNT\_BUILD](/sql_reference/functions-reference/aggregation/hll-count-build.html) function. `BYTEA`

## [](#return-type)Return Type

`BYTEA`

## [](#example)Example

Following the [example](/sql_reference/functions-reference/aggregation/hll-count-build.html#example) in [HLL\_COUNT\_BUILD](/sql_reference/functions-reference/aggregation/hll-count-build.html):

```
SELECT hll_count_estimate(hll_count_merge(a)) AS hll_estimate, hll_count_merge(a) AS merged_sketch
FROM sketch_of_data_to_count;
```

hll\_estimate BIGINT merged\_sketch BYTEA 6606880 \\x2f4167677265676174654675….