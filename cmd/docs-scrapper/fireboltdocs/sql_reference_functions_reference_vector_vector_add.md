## [](#vector_add)VECTOR\_ADD

Returns an array that is the sum of two input arrays.

## [](#syntax)Syntax

```
VECTOR_ADD(<array>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The first array in the addition calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric). `<array>` The second array in the addition calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric).

## [](#notes)Notes

Both input `ARRAY` arguments must have the same number of elements.

## [](#return-type)Return Type

`VECTOR_ADD` returns a result of type `ARRAY(BIGINT)` if the elements of `<array>` are of type `INT`, and returns a result of type `ARRAY(DOUBLE PRECISION)` if the elements are of type `REAL`. For other element types, `VECTOR_ADD` returns a result that matches the original element type, or follows Firebolt’s [type conversion](../../data-types.html#type-conversion) rules to convert them to compatible data types.

## [](#examples)Examples

**Example**

The following code example adds the two vectors, `[1, 2, 5]` and `[3, 4, -1]`, by computing `1+3`, `2+4`, and `5+(-1)`, which yields `[4, 6, 3]`:

```
select vector_add([1, 2, 5], [3, 4, -2]) as res
```

**Returns**

res (ARRAY(BIGINT)) {4,6,3}