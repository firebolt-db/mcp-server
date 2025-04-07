# [](#array_filter)ARRAY\_FILTER

Returns an array containing the elements from `<array>` for which the given Lambda function `<function>` returns something other than `0`.

The function can receive one or more arrays as its arguments. If more than one array is provided the following conditions should be met:

- The number of arguments of the Lambda function must be equal to the number of arrays provided. If the condition isn’t met, the query will not run and an error will be returned.
- All the provided arrays should be of the same length. If the condition isn’t met a runtime error will occur.

When multiple arrays are provided to the function, the function will evaluate the current elements from each array as its parameter. All of the elements at that index position must evaluate to true (or `1`) for this index to be included in the results. The elements that are returned are taken only from the first array provided.

## [](#syntax)Syntax

```
ARRAY_FILTER(<function>, <array> [, ...] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<function>` A [Lambda function](/Guides/loading-data/working-with-semi-structured-data/working-with-arrays.html#manipulating-arrays-with-lambda-functions) used to check elements in the array. Any Lambda function `<array> [, ...]` One or more arrays that will be evaluated by the Lambda function. Only the first array that is included will be filtered in the results. All the arrays must have exactly same number of elements. Arrays that each contain the same amount of elements

## [](#return-type)Return Type

`ARRAY` of the same type as the first input array

## [](#examples)Examples

In the example below, there is only one array.

```
SELECT
	ARRAY_FILTER(x -> x = 'a', [ 'a', 'b', 'c', 'd' ]);
```

**Returns**: `['a']`

In this example below, there are two arrays and two separate arguments for evaluation. The Lambda function searches the second array for all elements that are greater than 2. The elements in these positions are returned from the first array. The returned elements highlight the players who have completed above level 2.

```
SELECT
	ARRAY_FILTER(x, y -> y > 2, [ 'esimpson', 'sabrina21', 'kennethpark', 'rileyjon' ], [ 1, 2, 3, 9 ]) AS levels;
```

**Returns**: `['kennethpark', 'rileyjon']`

In this example below, there are three arrays, and Lambda function which meet the condition on two of them.

```
SELECT
	ARRAY_FILTER(x, y, z -> (y > 2 AND z = 'kennethpark'),
		[ 'Player A', 'Player B', 'Player C', 'Player D' ],
		[ 1, 2, 3, 9 ],
		[ 'kennethpark', 'rileyjon', 'kennethpark', 'rileyjon' ] )
	AS levels;
```

**Returns**: `['Player C']`