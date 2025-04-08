## [](#vector_manhattan_distance)VECTOR\_MANHATTAN\_DISTANCE

Returns the Manhattan, or L1 distance, between two vectors. The Manhattan distance measures the total distance by moving strictly along orthogonal axes, similar to navigating streets in a city grid. It is calculated as the sum of absolute differences between corresponding elements of two vectors.

## [](#syntax)Syntax

```
VECTOR_MANHATTAN_DISTANCE(<array>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The first array used in the distance calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric). `<array>` The second array used in the distance calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric).

## [](#notes)Notes

Both input `array` arguments must have the same number of elements.

## [](#return-type)Return Type

`DOUBLE`

## [](#examples)Examples

**Example**

The following code returns the Manhattan distance between two vectors: The Manhattan distance is calculated as the sum of absolute differences: |3 - 1| + |2 - 4| = 4.

```
SELECT VECTOR_MANHATTAN_DISTANCE([1, 4], [3, 2]) AS distance;
```

**Returns**

distance (DOUBLE PRECISION) 4