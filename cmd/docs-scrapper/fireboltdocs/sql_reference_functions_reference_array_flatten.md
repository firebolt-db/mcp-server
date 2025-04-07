# [](#array_flatten)ARRAY\_FLATTEN

Converts an array of arrays into a flat array. For every element that is an array, this function extracts its elements into the new array. The resulting flattened array contains all the elements from all source arrays.

The function:

- Applies to one level of nested arrays.
- Does not accept arrays that are already flat.

## [](#syntax)Syntax

```
ARRAY_FLATTEN(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array of arrays to be flattened Any `ARRAY` of `ARRAY` types

## [](#return-type)Return Type

`ARRAY` of the same type as the input array

## [](#example)Example

The following example flattens multiple arrays of level IDs:

```
SELECT
	ARRAY_FLATTEN([ [ [ 1, 2 ] ], [ [ 2, 3 ], [ 3, 4 ] ] ])
```

**Returns**: `[ [ 1, 2 ], [ 2, 3 ], [ 3, 4 ] ]`