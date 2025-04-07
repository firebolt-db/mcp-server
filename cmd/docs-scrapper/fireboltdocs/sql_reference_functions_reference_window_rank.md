# [](#rank)RANK

Rank the current row within the requested window with gaps.

For more information on usage, please refer to [Window Functions](/sql_reference/functions-reference/window/).

## [](#syntax)Syntax

```
RANK() OVER ([PARTITION BY <partition_by>] ORDER BY <order_by> [ASC|DESC] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<partition_by>` The expression used for the `PARTITION BY` clause. Any `<order_by>` The expression used in the `ORDER BY` clause. This parameter determines what value will be ranked. Any

## [](#return-type)Return Type

`INTEGER`

## [](#example)Example

In this example below, players are ranked based on their test scores for their game level.

```
SELECT
	nickname,
	level,
	current_score,
	RANK() OVER (PARTITION BY level ORDER BY current_score DESC ) AS rank_in_game
FROM
	players;
```

**Returns**:

first\_name grade\_level test\_score rank\_in\_class kennethpark 9 76 6 burchdenise 12 89 5 ymatthews 11 75 3 sabrina21 10 78 3