# [](#hll_count_estimate)HLL\_COUNT\_ESTIMATE

Extracts a cardinality estimate of a single HLL++ sketch that was previously built using the aggregate function [HLL\_COUNT\_BUILD](/sql_reference/functions-reference/aggregation/hll-count-build.html).

## [](#syntax)Syntax

```
HLL_COUNT_ESTIMATE(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An HLL++ sketch produced by the [HLL\_COUNT\_BUILD](/sql_reference/functions-reference/aggregation/hll-count-build.html) function. `BYTEA`

## [](#return-type)Return Type

`BIGINT`

## [](#example)Example

Following the [example](/sql_reference/functions-reference/aggregation/hll-count-build.html#example) in [HLL\_COUNT\_BUILD](/sql_reference/functions-reference/aggregation/hll-count-build.html):

```
SELECT hll_count_estimate(a) AS hll_estimate
FROM sketch_of_data_to_count
ORDER BY 1;
```

hll\_estimate (BIGINT) 3291008 4948957

```
SELECT hll_count_estimate(hll_count_merge(a)) AS hll_estimate
FROM sketch_of_data_to_count;
```

hll\_estimate (BIGINT) 6606880