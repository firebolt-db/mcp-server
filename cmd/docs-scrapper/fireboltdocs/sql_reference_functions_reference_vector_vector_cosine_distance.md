## [](#vector_cosine_distance)VECTOR\_COSINE\_DISTANCE

Returns the cosine distance between two vectors, calculated based on the angle (θ) between them. Vector cosine distance emphasizes the directional difference between the vectors, rather than magnitude. It is calculated as `1 - cos(θ)`, where `cos(θ)` is the [cosine similarity](./vector-cosine-similarity) between the vectors. `VECTOR_COSINE_DISTANCE` returns a value in the range `[0, 2]`. A vector cosine distance of `0` means that the vectors are identical in direction. A distance of `1` means that they are orthogonal and have no correlation. A distance of `2` means that they point in opposite directions.

## [](#syntax)Syntax

```
VECTOR_COSINE_DISTANCE(<array>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The first array used in the distance calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric). `<array>` The second array used in the distance calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric).

## [](#notes)Notes

Both input `array` arguments must have the same number of elements.

## [](#return-type)Return Type

`DOUBLE`

## [](#examples)Examples

**Example**

The following code returns the cosine distance between two vectors:

```
SELECT VECTOR_COSINE_DISTANCE([1, 2], [3, 4]) AS distance;
```

**Returns**

distance (DOUBLE PRECISION) 0.01613008990009257