# [](#typeof)TYPEOF

Returns the type of a given expression.

## [](#syntax)Syntax

```
TYPEOF(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression to typeof. Any

## [](#return-types)Return Types

A text of the given expression data type.

## [](#example)Example

The following example returns the type of PI() function:

```
SELECT TYPEOF(RANDOM()) AS random_data_type;
```

**Returns:** `double precision`