# [](#round)ROUND

Rounds a value to a specified number of decimal places.

## [](#syntax)Syntax

```
ROUND(<value> [, <decimal>])
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The value to be rounded   `<decimal>` Optional. An `INTEGER` constant that defines the decimal range of the returned value. By default, `ROUND` returns whole numbers. `INTEGER`

## [](#return-types)Return Types

`DOUBLE PRECISION`

## [](#example)Example

The following example returns the rounded value of 5.4. Since there is no specification of the decimal range, the functions returns a whole number:

The following example rounds the value 5.6930 to 1 decimal place: