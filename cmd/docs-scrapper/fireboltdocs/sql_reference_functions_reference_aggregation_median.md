# [](#median)MEDIAN

Returns the middle value in a given column. If number of values are even, `MEDIAN` returns the average of the two middle values.

## [](#syntax)Syntax

```
MEDIAN(<value>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The expression used to calculate the median value. `DOUBLE PRECISION`, `REAL`, `BIGINT`, `INT`

## [](#return-type)Return Type

`MEDIAN` returns a value of type `DOUBLE PRECISION`.

- This function ignores `NULL` values.
- If the input is empty, the function returns `NULL`.

## [](#examples)Examples

**Example**

The following query uses `UNNEST` to convert an array \[1,2,5] into a column, calculates the median value, and returns it as `result`:

```
SELECT MEDIAN(a) as result FROM UNNEST(array[1,2,5]) as a;
```

**Returns**

result (DOUBLE PRECISION) 2

The previous code example returns the middle element as the median because the number of elements in the array is odd.

**Example**

The following query uses `UNNEST` to convert an array \[100, NULL, 1, 2, 5] into a column, calculates the median value, and returns it as `result`:

```
SELECT MEDIAN(a) as result FROM UNNEST(array[100,NULL,1,2,5]) as a;
```

**Returns**

result (DOUBLE PRECISION) 3.5

The previous code example returns the average of the two middle elements after sorting the values, because the number of non-`NULL` elements is even after ignoring `NULL` values. The average of the middle elements, `2` and `5`, is calculated as : `((2+5)/2)`.