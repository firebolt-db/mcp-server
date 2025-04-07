# [](#array_agg)ARRAY\_AGG

Concatenates input values, including `NULL` values, into an array.

## [](#syntax)Syntax

```
ARRAY_AGG(<expression>)
```

## [](#parameters)Parameters

Parameter Description Supported input type `<expression>` Expression of any type to be accumulated into an array. Any

## [](#return-type)Return Type

`ARRAY` of the same type as the input data. If there is no input data, `ARRAY_AGG` returns `NULL`.

## [](#example)Example

For the following example, see the `player_information` table:

nickname playerid stephen70 1 burchdenise 7 sabrina21 NULL

The following code example code selects the columns in the `player_information` table and returns the values in two arrays, `nicknames` and `playerids`:

```
SELECT
  ARRAY_AGG(nickname) AS nicknames,
  ARRAY_AGG(playerid) AS playerids
FROM
	price_list;
```

**Returns**

`{'stephen70', 'burchdenise', 'sabrina21'}, {1, 7, NULL}`

The following code example shows that if a filter is added to the query which rejects all rows, `ARRAY_AGG` will return `NULL`:

```
SELECT
  ARRAY_AGG(nickname) AS nicknames,
  ARRAY_AGG(playerid) AS playerids
FROM
	price_list
WHERE
  playerid = 42;
```

**Returns**

`NULL, NULL`