# [](#is_infinite)IS\_INFINITE

Returns `TRUE` if the argument is infinite, and `FALSE` otherwise. Only `REAL` and `DOUBLE PRECISION` types can represent infinity in Firebolt, meaning that `IS_INFINITE` will always return `FALSE` for `NUMERIC` inputs.

## [](#syntax)Syntax

```
IS_FINITE(<value>);
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The input that will be checked to determine if it is an infinite number. `NUMERIC`, `DOUBLE PRECISION`, `REAL`

## [](#return-type)Return Type

`IS_INFINITE` returns a value of type `BOOLEAN`.

## [](#examples)Examples

The following code example checks whether the value inf, after being cast to a `DOUBLE PRECISION` data type, is an infinite number:

The following code example checks whether the value 10, after being cast to a `REAL` data type, is an infinite number: