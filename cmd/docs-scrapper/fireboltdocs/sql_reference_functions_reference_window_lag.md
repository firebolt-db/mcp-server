# [](#lag)LAG

Returns the value of the input expression at the given offset before the current row within the requested window.

For more information on usage, please refer to [Window Functions](/sql_reference/functions-reference/window/).

## [](#syntax)Syntax

```
LAG ( <expression> [, <offset> [, <default> ]] )
    OVER ( [ PARTITION BY <partition_by> ] ORDER BY <order_by> [ { ASC | DESC } ] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<expression>` Any valid expression that will be returned based on the `<offset>.` Any `<offset>` The number of rows backward from the current row from which to obtain a value. A negative number will act as [LEAD()](/sql_reference/functions-reference/window/lead.html). The offset must be a literal `INTEGER`. `INTEGER` `<default>` The expression to return when the offset goes out of the bounds of the window. Must be a literal of the same type as `<expression>`. The default is `NULL`. Any `<partition_by>` The expression used for the `PARTITION BY` clause. Any `<order_by>` An expression used for the `ORDER BY` clause. Any

## [](#example)Example

In the example below, the `LAG` function is being used to find the players in each level who ranked before and after a certain player. In some cases, if the player has no one ranked before or after them, the `LAG` function returns `NULL`.

```
SELECT
	nickname,
	level,
	LAG(nickname, 1) OVER (PARTITION BY level ORDER BY nickname) AS player_before,
	LAG(first_name, -1) OVER (PARTITION BY level ORDER BY nickname) AS player_after
FROM
	players;
```

**Returns**:

nickname level player\_before player\_after kennethpark 9 NULL rileyjon rileyjon 9 kennethpark sabrina21 sabrina21 9 rileyjon ymatthews ymatthews 9 sabrina21 NULL