## [](#vector_squared_euclidean_distance)VECTOR\_SQUARED\_EUCLIDEAN\_DISTANCE

Returns the squared [Euclidean distance](/sql_reference/functions-reference/vector/vector-euclidean-distance.html), or squared [L2 distance](/sql_reference/functions-reference/vector/vector-euclidean-distance.html) between two vectors. The squared Euclidean distance measures how far apart two vectors based on the size of their differences, without considering direction. By squaring the difference, it emphasizes larger differences, which can help in finding outliers or large deviations.

## [](#syntax)Syntax

```
VECTOR_SQUARED_EUCLIDEAN_DISTANCE(<array>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The first array used in the distance calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric). `<array>` The second array used in the distance calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric).

## [](#notes)Notes

Both input `array` arguments must have the same number of elements.

## [](#return-type)Return Type

`DOUBLE`

## [](#examples)Examples

**Example**

The following code returns the squared Euclidean distance between two vectors:

```
SELECT VECTOR_SQUARED_EUCLIDEAN_DISTANCE([1, 2], [3, 4]) AS distance;
```

**Returns**

distance (DOUBLE PRECISION) 8

**Example**

The following code returns the squared Euclidean distance between two identical vectors:

```
SELECT VECTOR_SQUARED_EUCLIDEAN_DISTANCE([1, 1], [1, 1]) AS distance;
```

**Returns**

distance (DOUBLE PRECISION) 0

**Example**

The following code returns the squared Euclidean distance between two vectors that are very far apart:

```
SELECT VECTOR_SQUARED_EUCLIDEAN_DISTANCE([1, 1], [10, 10]) AS distance;
```

**Returns**

distance (DOUBLE PRECISION) 162