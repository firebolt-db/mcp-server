## [](#vector_subtract)VECTOR\_SUBTRACT

Returns an array that is the difference of two input arrays.

## [](#syntax)Syntax

```
VECTOR_SUBTRACT(<array>, <array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The first array in the difference calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric). `<array>` The second array in the difference calculation. Any array of [numeric data types](/sql_reference/data-types.html#numeric).

## [](#notes)Notes

Both input `array` arguments must have the same number of elements.

## [](#return-type)Return Type

`VECTOR_SUBTRACT` returns a result of type `ARRAY(BIGINT)` if the elements of `<array>` are of type `INT`, and returns a result of type `ARRAY(DOUBLE PRECISION)` if the elements are of type `REAL`. For other element types, `VECTOR_ADD` returns a result that matches the original element type, or follows Firebolt’s [type conversion](../../data-types.html#type-conversion) rules to convert them to compatible data types.

## [](#examples)Examples

**Example**

The following example subtracts the vector `[3, 4, -4]` from `[1, 5, 6]` by computing `1-3`, `5-4`, and `6-(-2)`, which yields `[-2, 1, 8]`:

```
select vector_subtract([1, 5, 6], [3, 4, -2]) as res
```

**Returns**

res (ARRAY(BIGINT)) {-2,1,8}