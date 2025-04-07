# [](#min)MIN

Returns the minimum value within the requested window.

For more information on usage, please refer to [Window Functions](/sql_reference/functions-reference/window/).

## [](#syntax)Syntax

```
MIN( <expression> ) OVER ( [ PARTITION BY <partition_by> ] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` A value used for the `MIN` function Any `<partition_by>` An expression used for the `PARTITION BY` clause. Any

## [](#example)Example

The example below queries test scores for players in various grade levels. Unlike a regular `MIN()` aggregation, the window function highlights how each player individually compares to the lowest game score for their level.

```
SELECT
	nickname,
	level,
	current_score,
	MIN(current_score) OVER (PARTITION BY level) AS lowest_score
FROM
	players;
```

**Returns**:

nickname level current\_score lowest\_score kennethpark 9 76 2 sabrina21 7 90 15 burchdenise 5 79 4 ymatthews 6 85 9 rileyjon 8 80 20