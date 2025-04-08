# [](#stddev_samp)STDDEV\_SAMP

Computes the sample standard deviation of all non-`NULL` numeric values produced by an expression. The sample standard deviation measures how spread out values are in a sample by calculating the square root of the average of squared deviations from the sample mean, using a correction for sample size. For information about the population standard deviation, which estimates the spread of values in the full population, see [STDDEV\_POP](/sql_reference/functions-reference/aggregation/stddev-pop.html).

**Alias**: `STDDEV`

## [](#syntax)Syntax

```
{ STDDEV | STDDEV_SAMP }(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression producing numeric values for which to calculate the sample standard deviation. `REAL`, `DOUBLE PRECISION`

## [](#return-type)Return Type

`STDDEV_SAMP` returns a result of type `DOUBLE PRECISION`.

### [](#special-cases)Special cases

- If there is at most one non-`NULL` input value, the result is `NULL`.
- If the input contains an `Inf` or `NaN` value, the result will be `NaN`.

## [](#example)Example

The following code creates an `exams` table with a `grade` column of type `DOUBLE PRECISION`, and inserts five grade values into it:

```
CREATE TABLE exams (grade DOUBLE PRECISION);
INSERT INTO exams VALUES (4.0), (3.7), (3.3), (2.7), (2.7);
```

The following code calculates the sample standard deviation of the grade values from the `exams` table, rounds the result to three decimal places, and returns it as `stddev`:

```
SELECT ROUND(STDDEV_SAMP(grade), 3) as stddev from exams;
```

**Returns** The previous code returns the following result:

stddev (DOUBLE PRECISION) 0.585