# [](#array_count)ARRAY\_COUNT

Counts the number of elements in an array for which `function(array[i])` evaluates to TRUE, if a function is provided. If `<function>` is not provided, counts the number of elements in the array that evaluate to TRUE, equivalent to using `x -> x`. To count the elements in an array without any conditions, use the [ARRAY\_LENGTH](/sql_reference/functions-reference/array/array-length.html) function instead.

## [](#syntax)Syntax

```
ARRAY_COUNT(<function>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<function>` Optional. A [Lambda function](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html#lambda-function-general-syntax) used to check elements in the array. If `<function>` is not provided, `x -> x` is used. Any Lambda function returning `BOOLEAN` `<array>` An array of elements Any `ARRAY` type if `<function>` is provided, else `ARRAY(BOOLEAN)`

## [](#return-type)Return Type

`INTEGER`

## [](#examples)Examples

The example below searches through the array for any elements that are greater than 3. Only one number that matches this criteria is found, so the function returns `1`

```
SELECT
	ARRAY_COUNT(x -> x > 3, [ 1, 2, 3, 9, NULL ]) AS levels;
```

**Returns**: `1`

In this example below, there is no `<function>` provided in the `ARRAY_COUNT` function. This means the function will count all of the elements in the array that evaluate to TRUE. Below, this is the case for all values except `FALSE` and `null`:

```
SELECT
	ARRAY_COUNT([TRUE, FALSE, 2::BOOLEAN, 3 is not null, null is null, null]) AS levels;
```

**Returns**: `4`