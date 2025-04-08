# [](#log)LOG

Returns the common (base 10) logarithm of a numerical expression, or the logarithm to an arbitrary base if specified as the first argument. The value for which `log` is computed needs to be larger than 0, otherwise an error is returned. If a base is provided, it also needs to be larger than 0 and not equal to 1. You can use the function [LN](/sql_reference/functions-reference/numeric/ln.html) to compute the natural logarithm (base e).

**Alias:** `LOG10` (does not support a custom `base` argument)

## [](#syntax)Syntax

```
LOG([<base>,] <value>);
```

## [](#parameters)Parameters

Parameter Description Supported input types `<base>` Optional. The base for the logarithm. The default base is 10. `DOUBLE PRECISION` `<value>` The value for which to compute the logarithm. `DOUBLE PRECISION`

## [](#return-type)Return Type

`DOUBLE PRECISION`

## [](#example)Example

The following example returns the logarithm of 64.0 to base 2:

The following example returns the logarithm of 100.0 to the default base 10:

The logarithm can only be computed for values that are larger than 0. All the following functions return an error:

When a base is provided, it needs to be positive and not equal to zero. All the following functions return an error: