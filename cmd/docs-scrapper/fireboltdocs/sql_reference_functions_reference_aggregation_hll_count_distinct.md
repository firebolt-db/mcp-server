# [](#hll_count_distinct)HLL\_COUNT\_DISTINCT

Counts the approximate number of unique or not NULL values, to the precision specified. `HLL_COUNT_DISTINCT` uses the HLL++ algorithm and allows you to control the sketch size set precision.

`HLL_COUNT_DISTINCT` requires less memory than exact aggregation functions, but also introduces statistical uncertainty. The default precision is 12, with a maximum of 20.

> Higher precision comes at a memory and performance cost.

## [](#syntax)Syntax

```
HLL_COUNT_DISTINCT ( <expression> [, <precision> ] )
```

Parameter Description Supported input types `<expression>` Valid values for the expression include column names or functions that return a column name. Any type `<precision>` Optional integer value to set precision. If not included, the default precision is 12. Precision range: 12-20. `INTEGER`, `BIGINT`

## [](#return-type)Return Type

`BIGINT`

> `APPROX_COUNT_DISTINCT(expression)` and `HLL_COUNT_DISTINCT(expression, 17)` return the same results, as `APPROX_COUNT_DISTINCT` uses the HLL algorithm with the default parameter to control the sketch size set to 17.

## [](#return-type-1)Return Type

`NUMERIC`

## [](#example)Example

To understand the difference between `COUNT(DISTINCT pk)` with exact precision enabled, `APPROX_COUNT_DISTNCT(pk)`, and `HLL_COUNT_DISTINCT(pk, <precision>)`, consider a table, `count_test` with 8,388,608 unique `pk` values.

```
SELECT
	COUNT(DISTINCT pk) as count_distinct,
	APPROX_COUNT_DISTINCT(pk) as approx_count
	HLL_COUNT_DISTINCT(pk, 12) as hll12_count,
	HLL_COUNT_DISTINCT(pk, 20) as hll20_count
FROM
	count_test;
```

**Returns**

Assuming 8,388,608 unique pk values, the previous code example returns the following:

```
' +----------------+--------------+-------------+-------------+
' | count_distinct | approx_count | hll12_count | hll20_count |
' +----------------+--------------+-------------+-------------+
' |      8,388,608 |    8,427,387 |   8,667,274 |   8,377,014 |
' +----------------+--------------+-------------+-------------+
```

In the previous result, `approx_count` is using precision 17, `hll12_count` is using precision 12, and `hll20_count` is using precision 20, which is the most precise value.