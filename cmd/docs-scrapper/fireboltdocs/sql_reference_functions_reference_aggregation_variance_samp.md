# [](#var_samp)VAR\_SAMP

Computes the sample variance of all non-`NULL` numeric values produced by an expression. The sample variance measures the average of the squared differences from the sample mean, indicating how spread out the values are within a sample. For information about the population variance, which measures how spread out the values are within the full population, see [VAR\_POP](/sql_reference/functions-reference/aggregation/variance-pop.html).

**Alias**: `VARIANCE`

## [](#syntax)Syntax

```
{ VARIANCE | VAR_SAMP }(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression producing numeric values for which to calculate the sample variance. `REAL`, `DOUBLE PRECISION`

## [](#return-type)Return Type

`VAR_SAMP` returns a result of type `DOUBLE PRECISION`.

### [](#special-cases)Special cases

- If there is at most one non-`NULL` input value, the result is `NULL`.
- If the input contains an `Inf` or `NaN` value, the result will be `NaN`.

## [](#example)Example

The following code creates an `exams` table with a `grade` column of type `DOUBLE PRECISION`, and inserts five grade values into it:

```
CREATE TABLE exams (grade DOUBLE PRECISION);
INSERT INTO exams VALUES (4.0), (3.7), (3.3), (2.7), (2.7);
```

The following code calculates the sample variance of the grade values from the `exams` table, rounds the result to three decimal places, and returns it as `variance`:

```
SELECT ROUND(VAR_SAMP(grade), 3) as variance from exams;
```

**Returns** The previous code returns the following result:

variance (DOUBLE PRECISION) 0.342