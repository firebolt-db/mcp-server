# [](#max)MAX

Returns the maximum value in its argument. NULL values are ignored. If all inputs are NULL, `MAX` returns NULL.

## [](#syntax)Syntax

```
MAX(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` The expression whose maximum to determine Any type

## [](#return-types)Return Types

Same as input type

## [](#examples)Examples

**Example**

This code example uses the following `tournaments` table:

name totalprizedollars The Drifting Thunderdome 24,768 The Lost Track Showdown 5,336 The Acceleration Championship 19,274 The Winter Wilderness Rally 21,560 The Circuit Championship 9,739

When used on the `totalprizedollars` column, `MAX` will return the highest value, as follows:

```
SELECT
	MAX(totalprizedollars) as maxprize
FROM
	tournaments;
```

**Returns**

`24,768`

**Example**

`MAX` can also work on text or array columns, in which case it returns the lexicographically largest value. In this example, the function assesses the `name` column in the `tournaments` table.

```
SELECT
	MAX(name) as maxtournament
FROM
	tournaments;
```

**Returns**

`The Winter Wilderness Rally`