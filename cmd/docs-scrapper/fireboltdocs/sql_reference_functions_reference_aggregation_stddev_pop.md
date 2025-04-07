# [](#stddev_pop)STDDEV\_POP

Computes the population standard deviation of all non-`NULL` numeric values produced by an expression. The population standard deviation shows how spread out the values in a population are, by measuring the average distance of each value from the population’s mean. For information about the sample standard deviation, which estimates the spread of values across sample rather than the full population, see [STDDEV\_SAMP](/sql_reference/functions-reference/aggregation/stddev-samp.html).

## [](#syntax)Syntax

```
STDDEV_POP(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression producing numeric values for which to calculate the population standard deviation. `REAL`, `DOUBLE PRECISION`

## [](#return-type)Return Type

`STDDEV_POP` returns a result of type `DOUBLE PRECISION`.

### [](#special-cases)Special cases

- If there are no non-`NULL` input values, the result is `NULL`.
- If the input contains an `Inf` or `NaN` value, the result will be `NaN`.

## [](#example)Example

The following code creates an `exams` table with a `grade` column of type `DOUBLE PRECISION`, and inserts five grade values into it:

```
CREATE TABLE exams (grade DOUBLE PRECISION);
INSERT INTO exams VALUES (4.0), (3.7), (3.3), (2.7), (2.7);
```

The following code calculates the population standard deviation of the grade values from the `exams` table, rounds the result to three decimal places, and returns it as `stddev`:

```
SELECT ROUND(STDDEV_POP(grade), 3) as stddev from exams;
```

**Returns** The previous code returns the following result:

stddev (DOUBLE PRECISION) 0.523