# [](#array_sort)ARRAY\_SORT

Returns the elements of the input array in ascending order.

If the argument `<function>` is provided, the sorting order is determined by the result of applying `<function>` on each element of the array.

## [](#syntax)Syntax

```
ARRAY_SORT([<function>,] <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input type `<function>` An optional function to be used to determine the sort order. Any lambda function that takes the elements of `<array>` as input `<array>` The array to be sorted. Any array

## [](#return-type)Return Type

`ARRAY` of the same type as the input array

## [](#example)Example

```
SELECT
	ARRAY_SORT([ 4, 1, 3, 2 ]);
```

**Returns**: `[1,2,3,4]`

In this example below, the modulus operator is used to calculate the remainder on any odd numbers. Therefore `ARRAY_SORT` puts the higher (odd) numbers last in the results.

```
SELECT
	ARRAY_SORT(x -> x % 2, [ 4, 1, 3, 2 ]);
```

**Returns**: `[4,2,1,3]`