# [](#array_count_distinct)ARRAY\_COUNT\_DISTINCT

Returns the number of distinct (unique) elements in the array. As with `COUNT` and `COUNT(DISTINCT ...)` aggregations, `NULL` is not counted as a value if it occurs.

## [](#syntax)Syntax

```
ARRAY_COUNT_DISTINCT(<array>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<array>` The array of which to count the distinct elements. Any `ARRAY` type

## [](#return-type)Return Type

`INTEGER`

## [](#example)Example

```
SELECT ARRAY_COUNT_DISTINCT([ 1, 2, 4, 5, 2, NULL, 5, 1 ]) AS res;
```

**Returns**: `4`