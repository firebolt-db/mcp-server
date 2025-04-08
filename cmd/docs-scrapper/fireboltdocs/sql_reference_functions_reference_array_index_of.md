# [](#index_of)INDEX\_OF

Returns the index position of the first occurrence of the element in the array (or `0` if not found).

## [](#syntax)Syntax

```
INDEX_OF(<array>, <value>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array to be analyzed `ARRAY` `<value>` The element from the array that is to be matched Any integer that corresponds to an element in the array

## [](#return-type)Return Type

`INTEGER`

## [](#example)Example

The following example returns the index position of the 5 in the `levels` array:

```
SELECT
	INDEX_OF([ 1, 3, 4, 5, 7 ], 5) AS levels;
```

**Returns**: `4`