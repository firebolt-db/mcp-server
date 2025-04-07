# [](#max_by)MAX\_BY

Returns the value of the first argument for the row that contains the maximum of the second argument. If the maximum of the second argument is not unique, an arbitrary non-NULL value of the first argument is returned from the set of rows that maximize the second argument. If the first argument is NULL for all rows maximizing the second argument, NULL is returned.

## [](#syntax)Syntax

```
MAX_BY(<result>, <value>)
```

## [](#parameters)Parameters

Parameter Description Supported input types `<result>` The column from which the value is returned Any type `<value>` The column that is maximized Any type

## [](#return-types)Return Types

Same as input type of

## [](#examples)Examples

**Example**

This example uses the following table, `tournaments`:

name totalprizedollars The Drifting Thunderdome 24,768 The Lost Track Showdown 5,336 The Acceleration Championship 19,274 The French Grand Prix 237 The Circuit Championship 9,739

In the following code example, `MAX_BY` is used to find the name of the tournament with the highest total prize.

```
SELECT
	MAX_BY(name, totalprizedollars) as maxprizetournament
FROM
	tournaments
```

**Returns:** `The Drifting Thunderdome`

**Example**

When multiple rows maximize the second argument, an arbitrary one is chosen, preferring non-NULL values of the first argument:

```
SELECT MAX_BY(key, value)
FROM UNNEST(
    ['a', NULL, 'c', 'd', 'e', NULL],
    [10,  100,   1,  100, 100, NULL]
) t(key, value)
```

**Returns**

`'d'` or `'e'`, as rows 2, 4, and 5 maximize the second argument, but the first argument is NULL for row 2. Because non-NULL values of the first argument exist for the other rows, one of those values is returned. Which of them is non-deterministic, hence this query may return either `'d'` or `'e'`.

**Example**

However, if all rows maximizing the second argument are `NULL` in the first argument, `NULL` is returned:

```
SELECT MAX_BY(key, value)
FROM UNNEST(
    ['a', NULL, 'c', 'd', NULL, 'f'],
    [10,  100,   1,   2,  100, NULL]
) t(key, value)
```

**Returns**

`NULL`