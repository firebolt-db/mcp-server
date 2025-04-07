# [](#array_max)ARRAY\_MAX

Returns the maximum element in an array.

## [](#syntax)Syntax

```
ARRAY_MAX(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array or array-type column to be checked `ARRAY`

## [](#return-type)Return Type

Same as the element type of the array.

## [](#example)Example

The following example calculates the maximum number in the `levels` array:

```
SELECT
	ARRAY_MAX([ 1, 2, 3, 4 ]) AS levels;
```

**Returns**: `4`