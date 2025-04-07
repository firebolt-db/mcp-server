# [](#least)LEAST

The LEAST function selects the smallest value from a list of any number of expressions. The expressions must all be convertible to a common data type, which will be the type of the result. NULL values in the argument list are ignored. The result will be NULL only if all the expressions evaluate to NULL.

## [](#syntax)Syntax

```
LEAST(<expression> [,...])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression(s) to select smallest value. Any

## [](#return-types)Return Types

Same as input type

## [](#example)Example

```
SELECT LEAST(NULL, 2^3, 3^2)
```

**Returns:** `8`