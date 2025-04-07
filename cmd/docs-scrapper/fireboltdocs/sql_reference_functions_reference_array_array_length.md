# [](#array_length)ARRAY\_LENGTH

Returns the length of the given array, i.e., how many elements it contains. Produces `NULL` for `NULL` arrays.

**Alias:** [LENGTH](/sql_reference/functions-reference/string/length.html) (when used with an array argument)

## [](#syntax)Syntax

```
ARRAY_LENGTH(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array whose length should be calculated `ARRAY`

## [](#return-type)Return Type

`INTEGER`

## [](#example)Example

```
SELECT
	ARRAY_LENGTH([ 1, 2, 3, 4 ]) AS levels;
```

**Returns**: `4`

```
SELECT
	ARRAY_LENGTH([ [ 1, 2, 3 ], [ 4, 5, 6, 7 ]) AS levels;
```

**Returns**: `2`