## [](#vector_cosine_similarity)VECTOR\_COSINE\_SIMILARITY

Returns the cosine similarity between two vectors, calculated based on the angle (θ) between them. Vector cosine similarity measures how closely two vectors point in the same direction, and is calculated as `cos(θ)`. `VECTOR_COSINE_SIMILARITY` returns a value in the range `[-1, 1]`. A vector cosine similarity of `1` means that the vectors are identical in direction. A similarity of `0` means that they are orthogonal and have no correlation. A similarity of `-1` means that they point in opposite directions.

## [](#syntax)Syntax

```
VECTOR_COSINE_SIMILARITY(<array>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The first array used in the similarity calculation Any array of [numeric data types](/sql_reference/data-types.html#numeric). `<array>` The second array used in the similarity calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric).

## [](#notes)Notes

Both input `array` arguments must have the same number of elements.

## [](#return-type)Return Type

`DOUBLE`

## [](#examples)Examples

**Example**

The following code returns the cosine similarity between two vectors that point in very similar directions:

```
SELECT VECTOR_COSINE_SIMILARITY([1, 2], [3, 4]) AS similarity;
```

**Returns**

similarity (DOUBLE PRECISION) 0.9838699100999074

**Example**

The following code returns the cosine similarity between two vectors that point in very different directions:

```
SELECT VECTOR_COSINE_SIMILARITY([1, 2], [-3, -4]) AS similarity;
```

**Returns**

similarity (DOUBLE PRECISION) -0.9838699100999074

**Example**

The following code returns the cosine similarity between two vectors that are orthogonal to each other:

```
SELECT VECTOR_COSINE_SIMILARITY([2, 0], [0, 2]) AS similarity;
```

**Returns**

similarity (DOUBLE PRECISION) 0