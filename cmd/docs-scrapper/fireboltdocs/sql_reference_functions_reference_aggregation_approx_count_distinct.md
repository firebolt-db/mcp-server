# [](#approx_count_distinct)APPROX\_COUNT\_DISTINCT

Counts the approximate number of unique values that are not NULL. Internally, `APPROX_COUNT_DISTINCT` uses HyperLogLog (HLL) sketches. Calling `APPROX_COUNT_DISTINCT` returns the same value as calling `HLL_COUNT_DISTINCT` with precision 17.

## [](#syntax)Syntax

```
APPROX_COUNT_DISTINCT(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` Expression on which to approximate the distinct count Any type

## [](#return-type)Return Type

`BIGINT`

## [](#example)Example

When aggregating on few distinct values, `APPROX_COUNT_DISTINCT` has no estimation error and returns exact results:

```
SELECT
    APPROX_COUNT_DISTINCT(number) as approximate,
    COUNT(DISTINCT number) as exact
FROM
    generate_series(1, 1000) r(number);
```

**Returns**:

approximate exact 1,000 1,000

`NULL` values do not change the result of `APPROX_COUNT_DISTINCT`:

```
SELECT
    APPROX_COUNT_DISTINCT(number) as approximate,
    COUNT(DISTINCT number) as exact
FROM
    (SELECT * FROM generate_series(1, 1000)
       UNION ALL
     SELECT NULL) r(number);
```

**Returns**:

approximate exact 1,000 1,000

As the number of distinct values grows, the result becomes an approximation:

```
SELECT
    APPROX_COUNT_DISTINCT(number) as approximate,
    COUNT(DISTINCT number) as exact
FROM
    generate_series(1, 50000) r(number);
```

**Returns**:

approximate exact 50,160 50,000

`APPROX_COUNT_DISTINCT` also works for compound types:

```
SELECT 
    APPROX_COUNT_DISTINCT(arr) as approximate,
    COUNT(DISTINCT arr) as exact
FROM 
    unnest([[1, 2], [3, 4], NULL, [NULL], [1, NULL]]) r(arr)
```

**Returns**:

approximate exact 4 4