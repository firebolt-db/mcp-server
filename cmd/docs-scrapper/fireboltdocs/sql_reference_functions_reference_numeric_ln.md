# [](#ln)LN

Returns natural (base e) logarithm of a numerical expression. The value for which `ln` is computed needs to be larger than 0, otherwise an error is returned. You can use the function [LOG](/sql_reference/functions-reference/numeric/log.html) if you want to provide a different base.

## [](#syntax)Syntax

```
LN(<value>);
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The value for which to compute the natural logarithm. `DOUBLE PRECISION`

## [](#return-type)Return Type

`DOUBLE PRECISION`

## [](#examples)Examples

The following example computes the natural logarithm of 1.0:

The following example returns the natural logarithm close to e:

The natural logarithm can only be computed for values that are larger than 0. All the following functions return an error: