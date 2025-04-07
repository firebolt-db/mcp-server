# [](#sum)SUM

Calculates the sum of an expression.

## [](#syntax)Syntax

```
SUM ([ DISTINCT ] <value>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` The expression used to calculate the sum. Any numeric type

Valid values for `<value>` include column names or expressions that evaluate to numeric values. When `DISTINCT` is being used, only the unique number of rows with no `NULL` values are summed.

## [](#return-types)Return Types

`NUMERIC`

## [](#precision-and-determinism)Precision and Determinism

Applying `SUM` to REAL and DOUBLE PRECISION is subject to [floating point arithmetic accuracy limitations](https://en.wikipedia.org/wiki/Floating-point_arithmetic#Accuracy_problems) and its resulting error. This error may add up when aggregating multiple values.

The order of operations while computing the aggregate is non-deterministic. This can lead to varying total floating point error when running a query multiple times. If this is not acceptable for your use-case, aggregate on [NUMERIC](/sql_reference/numeric-data-type) data instead.

## [](#examples)Examples

**Example**

This code example uses the following `tournaments` table:

name totalprizedollars The Drifting Thunderdome 24,768 The Lost Track Showdown 5,336 The Acceleration Championship 19,274 The Winter Wilderness Rally 21,560 The Circuit Championship 9,739 The Singapore Grand Prix 19,274

```
SELECT
	SUM(totalprizedollars)
FROM
	tournaments
```

**Returns**

`99,951`

**Example**

For the following code example, since both the Singapore Grand Prix and The Acceleration Championship have the same total prize of `19,274`, only one of these values in this sum in included:

```
SELECT
	SUM (DISTINCT totalprizedollars)
FROM
	tournaments
```

**Returns**

`80,677`