# [](#array_reverse)ARRAY\_REVERSE

Returns an array of the same size and type as the original array, with the elements in reverse order. Nulls are retained.

## [](#syntax)Syntax

```
ARRAY_REVERSE(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array to be reversed `ARRAY` of any type

## [](#return-type)Return Type

`ARRAY` of the same type as the input array

## [](#example)Example

The following example returns the reverse of the input array:

```
SELECT
	ARRAY_REVERSE([ 1, 2, 3, 6 ]);
```

**Returns**: `[6,3,2,1]`

Only the outermost array is reversed for nested arrays:

```
SELECT
	ARRAY_REVERSE([[1,2,3], [4,5], NULL, [7], [8,9]]);
```

**Returns**: `[[8,9], [7], NULL, [4,5], [1,2,3]]`