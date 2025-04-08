# [](#is_finite)IS\_FINITE

Returns `TRUE` if the argument is finite, and `FALSE` otherwise. Only `REAL` and `DOUBLE PRECISION` types can represent infinity in Firebolt, meaning that `IS_FINITE` will always return `TRUE` for `NUMERIC` inputs.

## [](#syntax)Syntax

```
IS_FINITE(<value>);
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The input that will be checked to determine if it is a finite number. `NUMERIC`, `DOUBLE PRECISION`, `REAL`

## [](#return-type)Return Type

`IS_FINITE` returns a value of type `BOOLEAN`.

## [](#examples)Examples

The following example checks whether the value inf, after being cast to a `DOUBLE PRECISION` data type, is a finite number:

The following code example checks whether the value 10, after being cast to a `REAL` data type, is an infinite number: