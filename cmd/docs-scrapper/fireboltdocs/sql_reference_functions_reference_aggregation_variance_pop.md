# [](#var_pop)VAR\_POP

Computes the population variance of all non-`NULL` numeric values produced by an expression. The population variance measures the average of the squared differences from the population mean, indicating how spread out the values are within the entire population. For information about the sample variance, which measures how spread out the values are within a sample, see [VAR\_SAMP](/sql_reference/functions-reference/aggregation/variance-samp.html).

## [](#syntax)Syntax

```
VAR_POP(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` An expression producing numeric values for which to calculate the population variance. `REAL`, `DOUBLE PRECISION`

## [](#return-type)Return Type

`VAR_POP` returns a result of type `DOUBLE PRECISION`.

### [](#special-cases)Special cases

- If there are no non-`NULL` input values, the result is `NULL`.
- If the input contains an `Inf` or `NaN` value, the result will be `NaN`.

## [](#example)Example

The following code creates an `exams` table with a `grade` column of type `DOUBLE PRECISION`, and inserts five grade values into it:

```
CREATE TABLE exams (grade DOUBLE PRECISION);
INSERT INTO exams VALUES (4.0), (3.7), (3.3), (2.7), (2.7);
```

The following code calculates the population variance of the grade values from the `exams` table, rounds the result to three decimal places, and returns it as `variance`:

```
SELECT ROUND(VAR_POP(grade), 3) as variance from exams;
```

**Returns** The previous code returns the following result:

variance (DOUBLE PRECISION) 0.274