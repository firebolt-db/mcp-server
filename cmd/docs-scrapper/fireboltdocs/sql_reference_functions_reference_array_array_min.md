# [](#array_min)ARRAY\_MIN

Returns the minimum element in an array.

## [](#syntax)Syntax

```
ARRAY_MIN(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array or array-type column to be checked `ARRAY`

## [](#return-type)Return Type

Same as the element type of the array.

## [](#example)Example

The following example calculates the minimum number in the `levels` array:

```
SELECT
	ARRAY_MIN([ 1, 2, 3, 4 ]) AS levels;
```

**Returns**: `1`