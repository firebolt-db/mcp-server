## [](#vector_euclidean_distance)VECTOR\_EUCLIDEAN\_DISTANCE

Returns the Euclidean distance, or L2 distance, between two vectors. The Euclidean distance measures the shortest distance between two points in space along a straight line. It is calculated by summing the squared distances of a pair of vector elements, and then taking the square root.

## [](#syntax)Syntax

```
VECTOR_EUCLIDEAN_DISTANCE(<array>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The first array used in the distance calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric). `<array>` The second array used in the distance calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric).

## [](#notes)Notes

Both input `array` arguments must have the same number of elements.

## [](#return-type)Return Type

`DOUBLE`

## [](#examples)Examples

**Example**

The following code returns the Euclidean distance between two vectors:

```
SELECT VECTOR_EUCLIDEAN_DISTANCE([1, 2], [3, 4]) AS distance;
```

**Returns**

distance (DOUBLE PRECISION) 2.8284271247461903