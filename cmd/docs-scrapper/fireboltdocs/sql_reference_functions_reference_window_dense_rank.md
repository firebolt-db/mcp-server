# [](#dense_rank)DENSE\_RANK

Rank the current row within the requested window.

For more information on usage, please refer to [Window Functions](/sql_reference/functions-reference/window/).

## [](#syntax)Syntax

```
DENSE_RANK() OVER ([PARTITION BY <partition_by>] ORDER BY <order_by> [ASC|DESC] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<partition_by>` The expression used for the `PARTITION BY` clause. Any `<order_by>` The expression used in the `ORDER BY` clause. This parameter determines what value will be ranked. Any

## [](#return-types)Return Types

Same as input type

## [](#example)Example

In this example below, players are ranked based on their high scores for their game level.

```
SELECT
	nickname,
	level,
	highscore,
	DENSE_RANK() OVER (PARTITION BY level ORDER BY highscore DESC ) AS game_rank
FROM
	players;
```

**Returns**:

nickname level highscore game\_rank kennethpark 9 76 6 sabrina21 10 78 3 rileyjon 11 94 1 ymatthews 12 92 4