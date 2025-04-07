# [](#count)COUNT

Count the number of values within the requested window.

For more information on usage, please refer to [Window Functions](/sql_reference/functions-reference/window/)

## [](#syntax)Syntax

```
COUNT( <value> ) OVER ( [ PARTITION BY <partition_by> ] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` A value used for the `COUNT()` function. Any numeric type `<partition_by>` An expression used for the `PARTITION BY` clause. Any

## [](#return-type)Return Type

`NUMERIC`

## [](#example)Example

The following example generates a count of how many video game players have registered on a specific day:

```
SELECT
	registeredon,
	COUNT(agecategory) OVER (PARTITION BY registeredon) AS count_of_players
FROM
	players;
```

**Returns**:

registeredon count\_of\_players 2020-11-15 12 2020-11-16 8 2020-11-17 4 2020-11-18 9