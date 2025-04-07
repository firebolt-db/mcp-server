# [](#array_sum)ARRAY\_SUM

Returns the sum of elements of `<array>`.

## [](#syntax)Syntax

```
ARRAY_SUM(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array to be used to calculate the function. Any array of numeric types

## [](#return-type)Return Type

The return type depends on the input type:

Array element type Return type `INTEGER` `BIGINT` `BIGINT` `NUMERIC(38, 0)` `NUMERIC(precision, scale)` `NUMERIC(precision, scale)` `REAL` `REAL` `DOUBLE` `DOUBLE`

## [](#example)Example

```
SELECT
	ARRAY_SUM([ 4, 1, 3, 2 ]) AS levels;
```

**Returns**: `10`