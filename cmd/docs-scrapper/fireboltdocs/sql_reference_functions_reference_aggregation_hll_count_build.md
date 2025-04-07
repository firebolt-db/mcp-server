# [](#hll_count_build)HLL\_COUNT\_BUILD

Counts the approximate number of unique not NULL values, aggregating the values to HLL++ sketches represented as the [BYTEA data type](/sql_reference/bytea-data-type.html). Multiple sketches can be merged to a single sketch using the aggregate function [HLL\_COUNT\_MERGE](/sql_reference/functions-reference/aggregation/hll-count-merge.html). To estimate the final distinct count value, the scalar function [HLL\_COUNT\_ESTIMATE](/sql_reference/functions-reference/numeric/hll-count-estimate.html) can be used. `HLL_COUNT_BUILD` uses the HLL++ algorithm and allows you to control the set sketch size precision, similar to [HLL\_COUNT\_DISTINCT](/sql_reference/functions-reference/aggregation/hll-count-distinct.html).

`HLL_COUNT_BUILD` requires less memory than exact count distinct aggregation, but also introduces statistical uncertainty. The default precision is 12, with a maximum of 20 set optionally.

> Higher precision comes at a memory and performance cost.

## [](#syntax)Syntax

```
HLL_COUNT_BUILD(<expression> [, <precision> ])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` Any column name or function that return a column name. Any type `<precision>` Optional literal integer value to set precision. If not included, the default precision is 12. Precision range: 12-20. `INTEGER`, `BIGINT`

## [](#return-type)Return Type

`BYTEA`

## [](#example)Example

```
CREATE TABLE data_to_count AS
SELECT *
FROM generate_series(0, 10000000, 3) a;

SELECT count(distinct a) AS accurate_count
FROM data_to_count;
```

accurate\_count (BIGINT) 3333334

```
CREATE TABLE data_to_count2 AS
SELECT *
FROM generate_series(0, 10000000, 2) a;

SELECT count(distinct a) AS accurate_count
FROM data_to_count2;
```

accurate\_count (BIGINT) 5000001

```
CREATE TABLE sketch_of_data_to_count AS
SELECT hll_count_build(a) a
FROM data_to_count;

INSERT INTO sketch_of_data_to_count
SELECT hll_count_build(a)
FROM data_to_count2;

SELECT hll_count_estimate(a) AS hll_estimate, a AS sketch
FROM sketch_of_data_to_count
ORDER BY 1;
```

hll\_estimate (BIGINT) sketch (BYTEA) 3291008 \\x2f41676772656761746546…. 4948957 \\x2f41676772656761746546….

```
SELECT hll_count_estimate(hll_count_merge(a)) AS hll_estimate
FROM sketch_of_data_to_count;
```

hll\_estimate (BIGINT) 6606880