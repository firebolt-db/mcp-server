# [](#avg)AVG

Returns the average value within the requested window.

For more information on usage, please refer to [Window Functions](/sql_reference/functions-reference/window/).

## [](#syntax)Syntax

```
AVG( <value> ) OVER ( [ PARTITION BY <partition_by> ] )
```

## [](#parameters)Parameters

Parameter Description Supported input types `<value>` A value used for the `AVG()` function Any numeric type `<partition_by` An expression used for the `PARTITION BY` clause Any

## [](#return-types)Return Types

- `NUMERIC` if the input is type `INTEGER`, `BIGINT` or `NUMERIC`
- `DOUBLE PRECISION` if the input is type `REAL` or `DOUBLE PRECISION`

## [](#example)Example

The example below is querying test scores for players in various game levels. Unlike a regular `AVG()` aggregation, the window function allows us to see how each student individually compares to the average test score for their game level.

```
SELECT
	nickname,
	level,
	currentscore,
	AVG(game_score) OVER (PARTITION BY level) AS score_average
FROM
	class_test;
```

**Returns**:

nickname level currentscore score\_average kennethpark 9 76 75.77777 sabrina21 7 90 81.33333 burchdenise 8 79 79.55555 ymatthews 6 85 93.88888 rileyjon 8 80 84.99999