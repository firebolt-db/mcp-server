# [](#bool_or)BOOL\_OR

Returns true if all non NULL input value are true, otherwise false. If all input values are NULL values, returns NULL.

## [](#syntax)Syntax

```
BOOL_AND(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The boolean expression used to calculate the result `BOOLEAN`

## [](#return-types)Return Types

`BOOLEAN`

## [](#example)Example

name totalprizedollars The Drift Championship 22,048 The Lost Track Showdown 5,336 The Acceleration Championship 19,274 The French Grand Prix 237 The Circuit Championship 9,739

We want to check if all tournaments have prize money

```
SELECT
	BOOL_ALL(totalprizeddollars IS NOT NULL AND totalprizeddollars > 0) as all_have_prizes 
FROM
	tournaments;
```

**Returns**

`true`