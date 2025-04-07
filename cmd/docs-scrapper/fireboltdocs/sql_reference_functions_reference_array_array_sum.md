# [](#array_sum)ARRAY\_SUM

Returns the sum of elements of `<array>`.

## [](#syntax)Syntax

```
ARRAY_SUM(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array to be used to calculate the function. Any array of numeric types

## [](#return-type)Return Type

The return type is `BIGINT` if the element type of `<array>` is `INT` and `DOUBLE PRECISION` if the element type is `REAL`. Otherwise, it matches the element type.

## [](#example)Example

```
SELECT
	ARRAY_SUM([ 4, 1, 3, 2 ]) AS levels;
```

**Returns**: `10`