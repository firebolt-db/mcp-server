# [](#date_add)DATE\_ADD

Computes a new TIMESTAMP or TIMESTAMPTZ value by adding or subtracting a specified number of time units from a DATE, TIMESTAMP, or TIMESTAMPTZ value. It’s similar to [arithmetic with intervals](/Reference/interval-arithmetic.html), but it also allows using a column reference for the `<quantity>`.

## [](#syntax)Syntax

```
DATE_ADD('<unit>', <quantity>, <expression>)
```

## [](#parameters)Parameters

Parameter Description `<unit>` A TEXT literal specifying the time unit. Must be one of `microsecond`,`millisecond`,`second`,`minute`,`hour`,`day`,`week`,`month`,`quarter`,`year`,`decade`,`century`, or `millennium`. `<quantity>` An INT or BIGINT specifying how often the time unit should be added or subtracted to . `<expression>` An expression evaluating to type DATE, TIMESTAMP, or TIMESTAMPTZ.

## [](#return-types)Return Types

TIMESTAMP if `<expression>` has type DATE or TIMESTAMP. TIMESTAMPTZ if `<expression>` has type TIMESTAMPTZ.

## [](#example)Example