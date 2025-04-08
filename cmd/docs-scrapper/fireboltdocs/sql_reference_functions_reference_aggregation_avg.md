# [](#avg)AVG

Calculates the average of an expression.

## [](#syntax)Syntax

```
AVG(<value>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The expression used to calculate the average Any numeric type

Valid values for the expression include column names or functions that return a column name (or columns) that contain numeric values.

The `AVG()` aggregation function ignores rows with `NULL` values. For example, an `AVG` from 3 rows containing `1`, `2`, and `NULL` returns `1.5` because the `NULL` row is not counted. To calculate an average that includes `NULL`, use `SUM(COLUMN)/COUNT(*)`.

## [](#return-types)Return Types

- `NUMERIC` if the input is type `INTEGER`, `BIGINT` or `NUMERIC`
- `DOUBLE PRECISION` if the input is type `REAL` or `DOUBLE PRECISION`

## [](#precision-and-determinism)Precision and Determinism

Applying `AVG` to REAL and DOUBLE PRECISION is subject to [floating point arithmetic accuracy limitations](https://en.wikipedia.org/wiki/Floating-point_arithmetic#Accuracy_problems) and its resulting error. This error may add up when aggregating multiple values.

The order of operations while computing the aggregate is non-deterministic. This can lead to varying total floating point error when running a query multiple times. If this is not acceptable for your use-case, aggregate on [NUMERIC](/sql_reference/numeric-data-type) data instead.

## [](#example)Example

The example below uses the following table `LevelPoints`. This table includes the maximum points a player can score at each level of the game:

levels maxpoints 1 50 2 100 3 150 4 200 5 250

Use the query below to find the average of the `maxpoints` value.

```
SELECT 
    AVG(maxpoints) AS AverageMaxPoints 
FROM levels;
```

**Returns**:

AverageMaxPoints 150