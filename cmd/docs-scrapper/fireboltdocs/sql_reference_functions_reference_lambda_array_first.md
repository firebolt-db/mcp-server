# [](#array_first)ARRAY\_FIRST

Returns the first element in the given array for which the given function returns `true`. The `<function>` parameter must be included.

## [](#syntax)Syntax

```
ARRAY_FIRST(<function>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<function>` A [Lambda function](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html#manipulating-arrays-with-lambda-functions) used to check elements in the array A Lambda function returning `BOOLEAN` `<array>` The array evaluated by the function Any array

## [](#return-type)Return Type

The element type of `<array>`

## [](#examples)Examples

The following example returns the first value in the `levels` array greater than 2:

```
SELECT
	ARRAY_FIRST(x -> x > 2, [ 1, 2, 4, 9 ]) AS levels;
```

**Returns**: `4`