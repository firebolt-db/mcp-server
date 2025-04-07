# [](#array_to_string)ARRAY\_TO\_STRING

Converts each array element to its text representation, and concatenates those using an optional delimiter. If no delimiter is provided, an empty string is used instead. `NULL` array elements are omitted.

**Alias:** `ARRAY_JOIN`

## [](#syntax)Syntax

```
ARRAY_TO_STRING(<array>[, <delimiter>])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` An array to be concatenated `ARRAY` `<delimiter>` The delimiter used for concatenating the array elements `TEXT`

## [](#return-type)Return Type

`TEXT`

## [](#example)Example

In the example below, the three elements are concatenated with no delimiter.

```
SELECT
	ARRAY_TO_STRING([ '1', '2', '3' ]) AS levels;
```

**Returns**: `123`

In this example below, the levels are concatenated separated by a comma.

```
SELECT
	ARRAY_TO_STRING([ '1', '2', '3' ], ',') AS levels;
```

**Returns**: `1,2,3`

In this example below, the elements of a nested array containing a `NULL` are concatenated.

```
SELECT
	ARRAY_TO_STRING([ [ 1, 2 ], [3, 4], [NULL, 5] ], ',') AS levels;
```

**Returns**: `1,2,3,4,5`