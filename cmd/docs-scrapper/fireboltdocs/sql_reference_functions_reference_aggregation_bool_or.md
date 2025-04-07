# [](#bool_or)BOOL\_OR

Returns true if any non NULL input value is true, otherwise false. If all input values are NULL values, returns NULL.

## [](#syntax)Syntax

```
BOOL_OR(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The boolean expression used to calculate the result `BOOLEAN`

## [](#return-types)Return Types

`BOOLEAN`

## [](#example)Example

name totalprizedollars The Drift Championship 22,048 The Lost Track Showdown 5,336 The Acceleration Championship 19,274 The French Grand Prix 237 The Circuit Championship 9,739

We want to see if any of the tournaments have prize value more than 20,000

```
SELECT
	BOOL_OR(totalprizeddollars > 20000) as has_big_prize
FROM
	tournaments;
```

**Returns**

`true`