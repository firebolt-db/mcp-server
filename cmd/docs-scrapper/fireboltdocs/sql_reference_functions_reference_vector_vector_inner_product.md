## [](#vector_inner_product)VECTOR\_INNER\_PRODUCT

Returns the inner product, also known as the dot or scalar product, between two vectors. The inner product measures how closely two vectors align with each other, both in magnitude and the cosine of the angle between them. It is calculated by multiplying the corresponding elements of the vectors and then summing the results.

If the vectors have similar directions and large magnitudes, the inner product has a large positive value. If they are orthogonal, the inner product is zero, regardless of magnitude. If they point in roughly opposite directions, and at least one has a small magnitude, the inner product is small and negative.

## [](#syntax)Syntax

```
VECTOR_INNER_PRODUCT(<array>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The first array used in the inner product calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric). `<array>` The second array used in the inner product calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric).

## [](#notes)Notes

Both input `array` arguments must have the same number of elements.

## [](#return-type)Return Type

`DOUBLE`

## [](#examples)Examples

**Example**

The following code returns the inner product of two vectors that have similar directions and magnitudes:

```
SELECT VECTOR_INNER_PRODUCT([1, 2], [3, 4]) AS product;
```

**Returns**

product (DOUBLE PRECISION) 11

**Example**

The following code returns the inner product of two vectors that are orthogonal to each other:

```
SELECT VECTOR_INNER_PRODUCT([3, 4], [-4, 3]) AS product;
```

**Returns**

product (DOUBLE PRECISION) 0

**Example**

The following code returns the inner product of two vectors that are pointing in very different directions that are not orthogonal:

```
SELECT VECTOR_INNER_PRODUCT([3, 4], [4, -2]) AS product;
```

**Returns**

product (DOUBLE PRECISION) 4